package app

import (
	"context"
	"flag"
	_ "github.com/sfomuseum/go-flags/lookup"
	"github.com/sfomuseum/go-www-geotag/templates/html"
	"html/template"
)

func NewApplicationTemplates(ctx context.Context, fs *flag.FlagSet) (*template.Template, error) {

	t := template.New("geotag").Funcs(template.FuncMap{
		//
	})

	return t.ParseFS(html.FS, "*.html")
}
