package static

import (
	"embed"
)

//go:embed javascript/* css/*
var FS embed.FS
