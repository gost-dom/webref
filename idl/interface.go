package idl

import "iter"

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
func (i BaseInterface) GetOperation(name string) Operation {
	for _, o := range i.Operations {
		if o.Name == name {
			return o
		}
	}
	return Operation{}
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

// Represents an attribute on an IDL interface
type Attribute struct {
	// Don't rely on this, it only exists during a refactoring process
	InternalSpec NameMember
	Name         string
	Type         Type
	Readonly     bool
}

// Represents a method on an IDL interface
type Operation struct {
	Name       string
	ReturnType Type
	Arguments  []Argument
	Static     bool
	// Don't rely on this, it only exists during a refactoring process
	InternalSpec NameMember
}

type Argument struct {
	Name     string
	Type     Type
	Variadic bool
}

type Type struct {
	// Type name. This is a simplification so far.
	Name     string
	Nullable bool
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
