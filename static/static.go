package static

import (
	"embed"
)

//go:embed css/* javascript/* wasm/*
var FS embed.FS
