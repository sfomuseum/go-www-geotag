package www

import (
	"errors"
	"html/template"
	"net/http"
	"strings"
)

type IndexHandlerOptions struct {
	Templates        *template.Template
	InitialLatitude  float64
	InitialLongitude float64
	InitialZoom      int
	EnableSearch     bool
	SearchEndpoint   string
	EnableOEmbed     bool
	OEmbedEndpoints  []string
}

type IndexHandlerVars struct {
	InitialLatitude  float64
	InitialLongitude float64
	InitialZoom      int
	EnableSearch     bool
	SearchEndpoint   string
	EnableOEmbed     bool
	OEmbedEndpoints  string
}

func IndexHandler(opts *IndexHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup("index")

	if t == nil {
		return nil, errors.New("Missing 'index' template")
	}

	oembed_endpoints := strings.Join(opts.OEmbedEndpoints, ",")

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		vars := IndexHandlerVars{
			InitialLatitude:  opts.InitialLatitude,
			InitialLongitude: opts.InitialLongitude,
			InitialZoom:      opts.InitialZoom,
			EnableSearch:     opts.EnableSearch,
			SearchEndpoint:   opts.SearchEndpoint,
			EnableOEmbed:     opts.EnableOEmbed,
			OEmbedEndpoints:  oembed_endpoints,
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
