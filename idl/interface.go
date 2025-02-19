package idl

import (
	"fmt"
	"iter"
)

type TypeKind int

const (
	// KindSimple represents most distinguishable types, but not "generic"
	// types, such as sequences
	KindSimple TypeKind = iota
	// The type is specified as a sequence<...>.
	//
	// When the kind is a sequence, the value is present on the TypeParam field.
	KindSequence
)

// Represents an [IDL type]
//
// [IDL type]: https://webidl.spec.whatwg.org/#idl-types
type Type struct {
	// Type name. This is a simplification so far.
	Kind     TypeKind
	Name     string
	Nullable bool
	// TypeParam contains the type parameter for generic types, e.g.
	// sequence<...>
	TypeParam *Type
}

type BaseInterface struct {
	Attributes []Attribute
	Name       string
	// Operations are the callable methods defined on the object. Note that the
	// IDL spec allows for overloads, which is represented by multiple entries
	// with the same name.
	Operations []Operation
	// Includes represent interfaces included using the includes IDL statement.
	//
	// See also: https://webidl.spec.whatwg.org/#includes-statement

	// Don't rely on this, it only exists during a refactoring process
	InternalSpec Name
	// Mixin tells if this is an [interface mixin]. A mixin doesn't result in a
	// class by itself, but a collection of methods implemented by multiple
	// classes. Eg., Node, Document, and DocumentFragment all serve as root nodes,
	// but don't inherit that from a common base class.
	//
	// [interface mixin]: https://webidl.spec.whatwg.org/#idl-interface-mixins
	Mixin bool
}

// Gets the specifications for operation with the specified name. If nothing is
// configured, default settings are returned.
func (i BaseInterface) GetOperation(name string) (Operation, bool) {
	for _, o := range i.Operations {
		if o.Name == name {
			return o, true
		}
	}
	return Operation{}, false
}

// Gets the specifications for operation with the specified name. If nothing is
// configured, default settings are returned.
func (i BaseInterface) GetAttribute(name string) (a Attribute, found bool) {
	for _, a := range i.Attributes {
		fmt.Println("Testing", a.Name)
		if a.Name == name {
			fmt.Println("Returning")
			return a, true
		}
	}
	return Attribute{}, false
}

// Interface represents an interface specification in the webref IDL files.
//
// For example, the following interface Animal is represented by an _interface_
//
//	[Exposed=Window]
//	interface Animal {
//		attribute DOMString name;
//	};
type Interface struct {
	BaseInterface
	Inheritance string
	Includes    []Interface
}

type InterfaceMember struct {
	// Don't rely on this, it only exists during a refactoring process
	InternalSpec NameMember
	Name         string
	// If a member is a stringifier, this means that this member is to be used
	// when a string representation is created of an object. Only one member can
	// be a stringifier. For an operation, an empty name means that it must be
	// called toString() in JavaScript
	Stringifier bool
}

// Represents an attribute on an IDL interface
type Attribute struct {
	InterfaceMember
	Type     Type
	Readonly bool
}

// Represents a method on an IDL interface
type Operation struct {
	InterfaceMember
	ReturnType Type
	Arguments  []Argument
	Static     bool
}

type Argument struct {
	Name     string
	Type     Type
	Variadic bool
	Optional bool
}

// // NOTE: This will be removed in favour of a slice on the type
// func (i Interface) Attributes() iter.Seq[NameMember] {
// 	return i.InternalSpec.Attributes()
// }

// Attributes iterates and return all attributes from the IDO interface i. If
// included is true, this will also iterate attributes from interfaces that i
// includes.
func (i Interface) AllAttributes(included bool) iter.Seq[Attribute] {
	return func(yield func(Attribute) bool) {
		for _, a := range i.Attributes {
			if !yield(a) {
				return
			}
		}
		if included {
			for _, ii := range i.Includes {
				for _, a := range ii.Attributes {
					if !yield(a) {
						return
					}
				}
			}
		}
	}
}
