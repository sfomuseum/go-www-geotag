package app

import (
	"context"
	"flag"
	"github.com/sfomuseum/go-flags"
	"github.com/sfomuseum/go-www-geotag/assets/templates"
	"html/template"
)

func NewApplicationTemplates(ctx context.Context, fs *flag.FlagSet) (*template.Template, error) {

	path_templates, err := flags.StringVar(fs, "path-templates")

	if err != nil {
		return nil, err
	}

	t := template.New("geotag").Funcs(template.FuncMap{
		//
	})

	if path_templates != "" {

		parsed, err := t.ParseGlob(path_templates)

		if err != nil {
			return nil, err
		}

		t = parsed

	} else {

		for _, name := range templates.AssetNames() {

			body, err := templates.Asset(name)

			if err != nil {
				return nil, err
			}

			t, err = t.Parse(string(body))

			if err != nil {
				return nil, err
			}
		}
	}

	return t, nil
}
