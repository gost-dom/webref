package idl

import (
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"log/slog"
	"slices"
	"strings"

	. "github.com/gost-dom/webref/idl/legacy"
	"github.com/gost-dom/webref/internal/specs"
)

type RetType struct {
	TypeName string
	Nullable bool
}

func NewRetTypeUndefined() RetType { return RetType{TypeName: "undefined", Nullable: false} }

// Spec represents the information stored in Web IDL files.
type Spec struct {
	// ParsedIdlFile is a direct JSON deserialisation of the data.
	//
	// Note: This was the first implementation and is most complete in terms of
	// available data, but has a lower level of abstraction. Use the other
	// properties, When data is available on those, e.g., find an interface from
	// the Interfaces map.
	//
	// This property will eventually be removed
	ParsedIdlFile
	Interfaces   map[string]Interface
	Dictionaries map[string]Dictionary
}

func createInterfaceMember(m NameMember) InterfaceMember {
	return InterfaceMember{
		InternalSpec: m,
		Name:         m.Name,
		Stringifier:  m.Special == "stringifier",
	}
}

func (s *Spec) createDictionary(n Name) Dictionary {
	dictionary := s.IdlNames[n.Name]
	result := Dictionary{}
	for f := range dictionary.Fields() {
		entry := DictionaryEntry{
			Key:   f.Name,
			Value: convertType(*f.IdlType.IdlType),
		}
		result.Entries = append(result.Entries, entry)
	}
	return result
}

// interfaceGlobal extract the [Global] extended idl attribute from an interface
func interfaceGlobal(n Name) []string {
	for _, a := range n.ExtAttrs {
		if a.Type == "extended-attribute" &&
			a.Name == "Global" {
			switch a.Rhs.Type {
			case "identifier":
				return []string{a.Rhs.Value.ValueName}
			case "identifier-list":
				res := make([]string, len(a.Rhs.Value.Values))
				for i, v := range a.Rhs.Value.Values {
					res[i] = v.Value.ValueName
				}
				return res
			}
		}
	}
	return nil
}

func (s *Spec) createInterface(n Name) Interface {
	includedNames := s.IdlExtendedNames[n.Name].Includes()

	jsonAttributes := slices.Collect(n.Attributes())
	jsonOperations := slices.Collect(n.Operations())
	jsonConstructors := slices.Collect(n.Constructors())
	iterableTypes := n.IterableTypes()

	intf := Interface{
		BaseInterface: BaseInterface{
			Name:         n.Name,
			Attributes:   make([]Attribute, len(jsonAttributes)),
			Operations:   make([]Operation, len(jsonOperations)),
			Constructors: make([]Constructor, len(jsonConstructors)),
			InternalSpec: n,
			Mixin:        n.Type == "interface mixin",
			Partial:      n.Partial,
		},
		Inheritance: n.Inheritance,
		Includes:    make([]Interface, len(includedNames)),
		Global:      interfaceGlobal(n),
	}
	intf.IterableTypes = make([]Type, len(iterableTypes))
	for i, t := range iterableTypes {
		intf.IterableTypes[i] = convertType(t)
	}

	for i, n := range includedNames {
		intf.Includes[i] = s.createInterface(s.IdlNames[n])
	}
	for i, a := range jsonAttributes {
		intf.Attributes[i] = Attribute{
			InterfaceMember: createInterfaceMember(a),
			Type:            getAttributeType(a),
			Readonly:        a.Readonly,
		}
	}
	for i, a := range jsonOperations {
		intf.Operations[i] = Operation{
			InterfaceMember: createInterfaceMember(a),
			ReturnType:      getReturnType(a),
			Arguments:       createMethodArguments(a),
			Static:          a.Special == "static",
		}
	}
	for i, c := range jsonConstructors {
		intf.Constructors[i] = Constructor{
			Arguments: createMethodArguments(c),
		}
	}
	return intf
}

func getAttributeType(operation NameMember) Type {
	if t, ok := FindIdlTypeValue(operation.IdlType, "attribute-type"); ok {
		return convertType(t)
	}
	return Type{}
}

func getReturnType(operation NameMember) Type {
	if t, ok := FindIdlTypeValue(operation.IdlType, "return-type"); ok {
		return convertType(t)
	}
	return Type{}
}

func convertType(t IdlType) Type {
	if t.Generic == "sequence" {
		return convertSequence(t)
	}
	if t.Generic == "Promise" {
		return convertPromise(t)
	}
	if t.Union {
		res := Type{
			Kind:     KindUnion,
			Nullable: t.Nullable,
			Types:    make([]Type, len(t.IType.Types)),
		}
		for i, u := range t.IType.Types {
			res.Types[i] = convertType(u)
		}
		return res
	}
	return Type{Name: t.IType.TypeName, Nullable: t.Nullable}
}

func convertSequence(t IdlType) Type {
	innerIdl, _ := FindIdlTypeValue(t.IType, "return-type")
	inner := convertType(innerIdl)
	return Type{
		Kind:      KindSequence,
		Name:      t.IType.TypeName,
		Nullable:  t.Nullable,
		TypeParam: &inner,
	}
}

func convertPromise(t IdlType) Type {
	innerIdl, _ := FindIdlTypeValue(t.IType, "return-type")
	inner := convertType(innerIdl)
	return Type{
		Kind:      KindPromise,
		Name:      t.IType.TypeName,
		Nullable:  t.Nullable,
		TypeParam: &inner,
	}
}

func createMethodArguments(n NameMember) []Argument {
	result := make([]Argument, len(n.Arguments))
	for i, a := range n.Arguments {
		result[i] = Argument{
			Name:     a.Name,
			Type:     convertType(*a.IdlType.IdlType),
			Variadic: a.Variadic,
			Optional: a.Optional,
		}
		if a.Default != nil {
			result[i].Default = &DefaultValue{
				Type:  DefaultValueType(a.Default.Type),
				Value: a.Default.Value,
			}
		}
	}
	return result
}

// initialize fills out the high-level representations from the low level parsed
// JSON data.
func (s *Spec) initialize() {
	s.Interfaces = make(map[string]Interface)
	s.Dictionaries = make(map[string]Dictionary)
	for name, spec := range s.IdlNames {
		switch spec.Type {
		case "interface", "interface mixin":
			s.Interfaces[name] = s.createInterface(spec)
		case "dictionary":
			s.Dictionaries[name] = s.createDictionary(spec)
		}
	}
}

func (s *Spec) Partials(name string) iter.Seq[Interface] {
	return func(yield func(Interface) bool) {
		for _, spec := range s.IdlExtendedNames[name] {
			switch spec.Type {
			case "interface", "interface mixin":
				if !yield(s.createInterface(spec.Name)) {
					return
				}
			}
		}
	}
}

// Load loads the IDL specs for a specific web API. The names correspond to the
// files in the [ed/idlparsed] folder in the curated branch of the
// [webref] repository.
//
// [ed/idlparsed]: https://github.com/w3c/webref/tree/curated/ed/idlparsed
// [webref]: https://github.com/w3c/webref
func Load(apiName string) (Spec, error) {
	file, err := specs.Open(fmt.Sprintf("idlparsed/%s.json", apiName))
	if err != nil {
		return Spec{}, err
	}
	defer file.Close()

	return ParseIdlJsonReader(file)
}

type TypeSpec struct {
	Spec         *Spec
	IdlInterface Interface
}

type MemberSpec struct{ NameMember }
type AttributeSpec struct{ NameMember }

func (t *TypeSpec) Members() []NameMember {
	return t.IdlInterface.InternalSpec.Members
}

func (t *TypeSpec) Constructor() (res MemberSpec, found bool) {
	idx := slices.IndexFunc(t.IdlInterface.InternalSpec.Members, func(n NameMember) bool {
		return n.Type == "constructor"
	})
	found = idx >= 0
	if found {
		res = MemberSpec{t.IdlInterface.InternalSpec.Members[idx]}
	}
	return
}

func (t *TypeSpec) Inheritance() string {
	return t.IdlInterface.InternalSpec.Inheritance
}

func (t *TypeSpec) InstanceMethods() iter.Seq[MemberSpec] {
	return func(yield func(v MemberSpec) bool) {
		for i, member := range t.IdlInterface.InternalSpec.Members {
			if member.Special == "static" {
				continue
			}
			if member.Type == "operation" && member.Name != "" {
				// Empty name seems to indicate a named property getter. Not sure yet.
				firstIndex := slices.IndexFunc(
					t.IdlInterface.InternalSpec.Members,
					func(m NameMember) bool {
						return m.Name == member.Name
					},
				)
				if firstIndex < i {
					slog.Warn("Function overloads", "Name", member.Name)
					continue
				} else {
					if !yield(MemberSpec{member}) {
						return
					}
				}
			}
		}
	}
}

func (t *TypeSpec) Attributes() iter.Seq[AttributeSpec] {
	return func(yield func(v AttributeSpec) bool) {
		for _, member := range t.IdlInterface.InternalSpec.Members {
			if member.IsAttribute() {
				if !yield(AttributeSpec{member}) {
					return
				}
			}
		}
	}
}

func ParseIdlJsonReader(reader io.Reader) (Spec, error) {
	spec := Spec{}
	b, err := io.ReadAll(reader)
	if err == nil {
		err = json.Unmarshal(b, &spec)
	}
	spec.initialize()
	return spec, err
}

func (s *Spec) GetType(name string) (TypeSpec, bool) {
	result, ok := s.Interfaces[name]
	return TypeSpec{s, result}, ok
}

func (s AttributeSpec) AttributeType() RetType {
	r, n := FindMemberAttributeType(s.NameMember)
	return RetType{r, n}
}

func (s MemberSpec) ReturnType() RetType {
	r, n := FindMemberReturnType(s.NameMember)
	return RetType{r, n}
}

func (t RetType) IsUndefined() bool { return t.TypeName == "undefined" }
func (t RetType) IsDefined() bool   { return !t.IsUndefined() }

func (t RetType) IsNode() bool {
	loweredName := strings.ToLower(t.TypeName)
	switch loweredName {
	case "node":
		return true
	case "document":
		return true
	case "documentfragment":
		return true
	}
	if strings.HasSuffix(loweredName, "element") {
		return true
	}
	return false
}
