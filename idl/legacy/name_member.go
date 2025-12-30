package legacy

import "strings"

type Stuff struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	IdlType  IdlTypes  `json:"idlType"`
}

type NameMember struct {
	Stuff
	Arguments []ArgumentType `json:"arguments"`
	Special   string         `json:"special"`
	Readonly  bool           `json:"readOnly"`
	Href      string         `json:"href"`
}

func (n NameMember) HasExtendedAttributes(name string) bool {
	for _, a := range n.ExtAttrs {
		if a.Type == "extended-attribute" &&
			a.Name == name {
			return true
		}
	}
	return false
}

func (n NameMember) ExtendedAttributes(name string) []string {
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

func FindMemberReturnType(member NameMember) (string, bool) {
	return FindIdlType(member.IdlType, "return-type")
}

func FindMemberAttributeType(member NameMember) (string, bool) {
	return FindIdlType(member.IdlType, "attribute-type")
}

func (member NameMember) IsAttribute() bool {
	if member.Type != "attribute" {
		return false
	}
	t, ok := FindIdlTypeValue(member.IdlType, "attribute-type")
	return ok && !strings.HasSuffix(t.IType.TypeName, "EventHandler")
}
