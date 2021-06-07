package www

import (
	"errors"
	"html/template"
	"net/http"
	"strings"
)

type EditorHandlerOptions struct {
	Templates             *template.Template
	InitialLatitude       float64
	InitialLongitude      float64
	InitialZoom           int
	EnablePlaceholder     bool
	PlaceholderEndpoint   string
	EnableOEmbed          bool
	EnableOEmbedCORSImage bool
	OEmbedEndpoints       []string
	EnableWriter          bool
	// The local URI for invoking the application's write API
	WriterPath string
	// The whosonfirst/go-writer.Writer URI for publishing geotag information
	WriterURI                  string
	EnableExifWriter           bool
	MapRenderer                string
	EnablePointInPolygon       bool
	PointInPolygonEndpoint     string
	PointInPolygonDataEndpoint string
}

type EditorHandlerVars struct {
	PageContext                string
	InitialLatitude            float64
	InitialLongitude           float64
	InitialZoom                int
	EnablePlaceholder          bool
	PlaceholderEndpoint        string
	EnableOEmbed               bool
	EnableOEmbedCORSImage      bool
	OEmbedEndpoints            string
	EnableWriter               bool
	WriterPath                 string
	WriterURI                  string
	EnableExifWriter           bool
	MapRenderer                string
	EnablePointInPolygon       bool
	PointInPolygonEndpoint     string
	PointInPolygonDataEndpoint string
}

func EditorHandler(opts *EditorHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup("editor")

	if t == nil {
		return nil, errors.New("Missing 'editor' template")
	}

	oembed_endpoints := strings.Join(opts.OEmbedEndpoints, ",")

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		vars := EditorHandlerVars{
			PageContext:                "Editor",
			InitialLatitude:            opts.InitialLatitude,
			InitialLongitude:           opts.InitialLongitude,
			InitialZoom:                opts.InitialZoom,
			EnablePlaceholder:          opts.EnablePlaceholder,
			PlaceholderEndpoint:        opts.PlaceholderEndpoint,
			EnableOEmbed:               opts.EnableOEmbed,
			EnableOEmbedCORSImage:      opts.EnableOEmbedCORSImage,
			OEmbedEndpoints:            oembed_endpoints,
			EnablePointInPolygon:       opts.EnablePointInPolygon,
			PointInPolygonEndpoint:     opts.PointInPolygonEndpoint,
			PointInPolygonDataEndpoint: opts.PointInPolygonDataEndpoint,
			EnableWriter:               opts.EnableWriter,
			WriterPath:                 opts.WriterPath,
			WriterURI:                  opts.WriterURI,
			EnableExifWriter:           opts.EnableExifWriter,
			MapRenderer:                opts.MapRenderer,
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
