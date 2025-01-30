package specs

import (
	"embed"
	"io/fs"
	"path"
)

//go:embed curated/idlparsed/*.json
//go:embed curated/elements/*.json
var WebRef embed.FS

func Open(name string) (fs.File, error) {
	filename := path.Join("curated", name)
	return WebRef.Open(filename)
}
