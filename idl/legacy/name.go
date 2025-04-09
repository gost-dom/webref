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

func (n Name) Operations() iter.Seq[NameMember] {
	return func(yield func(NameMember) bool) {
		for _, m := range n.Members {
			if m.Type == "operation" {
				if !yield(m) {
					return
				}
			}
		}
	}
}

func (n Name) Constructors() iter.Seq[NameMember] {
	return func(yield func(NameMember) bool) {
		for _, m := range n.Members {
			if m.Type == "constructor" {
				if !yield(m) {
					return
				}
			}
		}
	}
}
