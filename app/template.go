package app

import (
	"context"
	"flag"
	"github.com/sfomuseum/go-flags/lookup"
	"github.com/sfomuseum/go-www-geotag/templates/html"
	"html/template"
)

func NewApplicationTemplates(ctx context.Context, fs *flag.FlagSet) (*template.Template, error) {

	prefix, err := lookup.StringVar(fs, "prefix")

	if err != nil {
		return nil, err
	}

	t := template.New("geotag").Funcs(template.FuncMap{
		"EnsureRoot": func(p string) string {
			return EnsureRoot(p, prefix)
		},
	})

	return t.ParseFS(html.FS, "*.html")
}
