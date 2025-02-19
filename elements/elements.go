// Package elements is part of the gost-dom project and contains html element type lookup.
//
// Load a specification using function [Load].
//
// When a browser loads the following HTML.
//
//	<a href="/example">An example</a>
//
// The element will be represented by the HTMLAnchorElement IDL interface. This
// package contains the mapping from the tag name, "a", to the IDL interface
// name, "HTMLAnchorElement" for all HTML elements.
package elements

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gost-dom/webref/internal/specs"
)

type ElementJSON struct {
	Name      string `json:"name"`
	Interface string `json:"interface"`
}

type ElementsJSON struct {
	Elements []ElementJSON `json:"elements"`
}

type Elements = ElementsJSON

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

// Load loads the sources for a specific standard. The names correspond to the
// files in the [ed/elements] folder in the curated branch of the [webref]
// repository
//
// [ed/elements]: https://github.com/w3c/webref/tree/curated/ed/elements
// [webref]: https://github.com/w3c/webref
func Load(standard string) (res Elements, err error) {
	var (
		b []byte
		r io.Reader
	)
	if r, err = specs.Open(fmt.Sprintf("elements/%s.json", standard)); err == nil {
		if b, err = io.ReadAll(r); err == nil {
			err = json.Unmarshal(b, &res)
		}
	}
	return
}
