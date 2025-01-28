package webref

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed specs/curated/elements/html.json
var Html_defs []byte

//go:embed specs/curated/idlparsed/*.json
var WebRef embed.FS

func OpenIdlParsed(name string) (fs.File, error) {
	filename := fmt.Sprintf("specs/curated/idlparsed/%s.json", name)
	return WebRef.Open(filename)
}
