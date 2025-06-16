package legacy

import (
	"encoding/json"
	"fmt"
)

// This type should be avoided, but may contain information missing on
// [Interface].
type ParsedIdlFile struct {
	Parsed `json:"idlParsed"`
}

// This type should be avoided, but may contain information missing on
// [Interface].
type Parsed struct {
	IdlNames         map[string]Name
	IdlExtendedNames map[string]ExtendedNames
}

type ArgumentType struct {
	Stuff
	Default  *ArgumentDefault `json:"default"`
	Optional bool             `json:"optional"`
	Variadic bool             `json:"variadic"`
}

type ArgumentDefault struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

// This type should be avoided, but may contain information missing on
// [Interface].
type ExtendedName struct {
	Fragment string
	Type     string
	ExtAttrs []ExtAttr `json:"extAttrs"`
	Target   string
	Includes string
}

// This type should be avoided, but may contain information missing on
// [Interface].
type ExtendedNames []ExtendedName

func (nn ExtendedNames) Includes() []string {
	res := make([]string, 0)
	for _, n := range nn {
		if n.Type == "includes" {
			res = append(res, n.Includes)
		}
	}
	return res
}

type ExtAttr struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Rhs  struct {
		Type  string     `json:"type"`
		Value ValueTypes `json:"value"`
	} `json:"rhs"`
}

type IdlType struct {
	Type     string    `json:"type"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	Generic  string    `json:"generic"`
	Nullable bool      `json:"nullable"`
	Union    bool      `json:"union"`
	IType    IdlTypes  `json:"idlType"`
}

func (t *IdlTypes) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &t.Types)
	if err != nil {
		typ := new(IdlType)
		err = json.Unmarshal(bytes, &typ)
		if err == nil {
			t.IdlType = typ
		}
	}
	if err != nil {
		err = json.Unmarshal(bytes, &t.TypeName)
	}
	return err

}

type IdlTypes struct {
	Types    []IdlType
	IdlType  *IdlType
	TypeName string
}

func (i IdlTypes) String() string {
	if len(i.Types) > 0 {
		return fmt.Sprintf("%v", i.Types)
	}
	if i.IdlType != nil {
		return fmt.Sprintf("%v", *i.IdlType)
	}
	return i.TypeName
}

func (t *ValueTypes) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &t.Values)
	if err != nil {
		val := new(ValueType)
		err = json.Unmarshal(bytes, val)
		if err == nil {
			t.Value = val
		}
	}
	if err != nil {
		err = json.Unmarshal(bytes, &t.ValueName)
	}
	return err
}

func FindIdlTypeValue(idl IdlTypes, expectedType string) (IdlType, bool) {
	types := idl.Types
	if len(types) == 0 && idl.IdlType != nil {
		types = []IdlType{*idl.IdlType}
	}
	for _, t := range types {
		if t.Type == expectedType {
			return t, true
		}
	}
	return IdlType{}, false
}

func FindIdlType(idl IdlTypes, expectedType string) (string, bool) {
	if t, ok := FindIdlTypeValue(idl, expectedType); ok {
		return t.IType.TypeName, t.Nullable
	}
	return "", false
}

type ValueType struct {
	Value ValueTypes `json:"value"`
}

type ValueTypes struct {
	Values    []ValueType
	Value     *ValueType
	ValueName string
}

// IdlNameType represent the value of "type" of an exported name in an IDL
// files, and affects how the data is to be interpreted. Corresponds to the
// values in the json path:
//
//	idlParsed.idlNames[name].type
type IdlNameType string

const (
	IdlNameInterface IdlNameType = "interface"
)
