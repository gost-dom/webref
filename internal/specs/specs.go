package specs

import (
	"embed"
	"io/fs"
	"path"
)

//go:embed curated/**/*.json
var WebRef embed.FS

func Open(name string) (fs.File, error) {
	filename := path.Join("curated", name)
	return WebRef.Open(filename)
}
