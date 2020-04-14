package www

import (
	"errors"
	"html/template"
	"net/http"
	"strings"
)

type EditorHandlerOptions struct {
	Templates           *template.Template
	InitialLatitude     float64
	InitialLongitude    float64
	InitialZoom         int
	EnablePlaceholder   bool
	PlaceholderEndpoint string
	EnableOEmbed        bool
	OEmbedEndpoints     []string
	EnableWriter        bool
	WriterPath          string
}

type EditorHandlerVars struct {
	PageContext         string
	InitialLatitude     float64
	InitialLongitude    float64
	InitialZoom         int
	EnablePlaceholder   bool
	PlaceholderEndpoint string
	EnableOEmbed        bool
	OEmbedEndpoints     string
	EnableWriter        bool
	WriterPath          string
}

func EditorHandler(opts *EditorHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup("editor")

	if t == nil {
		return nil, errors.New("Missing 'editor' template")
	}

	oembed_endpoints := strings.Join(opts.OEmbedEndpoints, ",")

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		vars := EditorHandlerVars{
			PageContext:         "Editor",
			InitialLatitude:     opts.InitialLatitude,
			InitialLongitude:    opts.InitialLongitude,
			InitialZoom:         opts.InitialZoom,
			EnablePlaceholder:   opts.EnablePlaceholder,
			PlaceholderEndpoint: opts.PlaceholderEndpoint,
			EnableOEmbed:        opts.EnableOEmbed,
			OEmbedEndpoints:     oembed_endpoints,
			EnableWriter:        opts.EnableWriter,
			WriterPath:          opts.WriterPath,
		}

		err := t.Execute(rsp, vars)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
