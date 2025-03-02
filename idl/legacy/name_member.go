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
