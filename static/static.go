// package static provides an embed.FS instance containing static CSS and JavaScript assets used by the go-www-geotag application.
package static

import (
	"embed"
)

//go:embed css/* javascript/* wasm/*
var FS embed.FS
