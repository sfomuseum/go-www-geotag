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

	// PROTOMAPS: this remains to be reconciled with the Tangram.js + Nextzen stuff	
	// These flags are assigned to the <div id="map"> element in templates/html/editor.html
	// and are handled by the Javascript code static/javascript/geotag.maps.js
	
	MapRenderer         string
	ProtomapsTileURL    string
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
	// PROTOMAPS: See notes above
	MapRenderer         string
	ProtomapsTileURL    string
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
			MapRenderer:         opts.MapRenderer,
			ProtomapsTileURL:    opts.ProtomapsTileURL,
		}

		rsp.Header().Set("Content-Type", "text/html; charset=utf-8")

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
