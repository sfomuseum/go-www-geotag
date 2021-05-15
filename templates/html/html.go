// package html provides an embed.FS filesystem of HTML templates used by the go-www-geotag application.
package html

import (
	"embed"
)

//go:embed *.html
var FS embed.FS
