package legacy

import "iter"

// This type should be avoided, but may contain information missing on
// [Interface].
type Name struct {
	Type        string       `json:"type"`
	Name        string       `json:"name"`
	Members     []NameMember `json:"members"`
	Partial     bool         `json:"partial"`
	Href        string       `json:"href"`
	Inheritance string       `json:"Inheritance"`
	ExtAttrs    []ExtAttr    `json:"extAttrs"`
}

func (n Name) Attributes() iter.Seq[NameMember] {
	return func(yield func(NameMember) bool) {
		for _, m := range n.Members {
			if m.Type == "attribute" {
				if !yield(m) {
					return
				}
			}
		}
	}
}

func (n Name) membersOfType(t string) iter.Seq[NameMember] {
	return func(yield func(NameMember) bool) {
		for _, m := range n.Members {
			if m.Type == t {
				if !yield(m) {
					return
				}
			}
		}
	}
}

// Extract extended attributes that are simple strings or lists. E.g.,
// Global/Exposed in the SharedWorkerGlobalScope that targets two global names
// and is exposed to one.
//
//	[Global=(Worker,SharedWorker),Exposed=SharedWorker]
//	interface SharedWorkerGlobalScope : WorkerGlobalScope {
func (n Name) ExtendedAttributes(name string) []string {
	for _, a := range n.ExtAttrs {
		if a.Type == "extended-attribute" &&
			a.Name == name {
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

func (n Name) Operations() iter.Seq[NameMember] {
	return n.membersOfType("operation")
}

func (n Name) Constructors() iter.Seq[NameMember] {
	return n.membersOfType("constructor")
}
func (n Name) Fields() iter.Seq[NameMember] {
	return n.membersOfType("field")
}

func (n Name) IterableTypes() []IdlType {
	for _, m := range n.Members {
		if m.Type == "iterable" {
			types := m.IdlType
			if types.IdlType != nil {
				return []IdlType{*types.IdlType}
			}
			return types.Types
		}
	}
	return nil
}
