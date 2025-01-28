// Package elements contains functionality to tell which HTML tags correspond to
// which IDL interface names. E.g., the <a> tag corresponds to the
// HTMLAnchorElement IDL interface specification.
package elements

import (
	"encoding/json"
	"fmt"

	"github.com/gost-dom/browser/code-gen/webref"
)

type ElementJSON struct {
	Name      string `json:"name"`
	Interface string `json:"interface"`
}

type ElementsJSON struct {
	Elements []ElementJSON `json:"elements"`
}

type Elements ElementsJSON

// GetTagNameForInterface finds the tagname for an element that is represented
// by interface i in the DOM. For example, the anchor tag, <a>, is represented
// by an HTMLAnchorElement in code, so the return value for
// GetTagNameForInterface("HTMLAnchorElement") is "a"
func (n Elements) GetTagNameForInterface(i string) (string, bool) {
	for _, e := range n.Elements {
		if e.Interface == i {
			return e.Name, true
		}
	}
	return "", false
}

// GetTagNameForInterfaceError is like [Elements.GetTagNameForInterface], but
// returns an error instead of a boolean if the element is not found.
func (n Elements) GetTagNameForInterfaceError(i string) (res string, err error) {
	// TODO: Take into account is multiple tag names can result in the same
	// elements. Then the caller needs to specify the tag.
	var ok bool
	if res, ok = n.GetTagNameForInterface(i); !ok {
		err = fmt.Errorf("Could not find the tag name corresponding to IDL interface: %s", i)
	}
	return
}

func Load() (Elements, error) {
	output := ElementsJSON{}
	err := json.Unmarshal(webref.Html_defs, &output)
	return Elements(output), err
}
