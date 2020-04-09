package app

import (
	"context"
	"flag"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/sfomuseum/go-http-leaflet-geotag"
	"github.com/sfomuseum/go-www-geotag/flags"
	"github.com/sfomuseum/go-www-geotag/www"
	"net/http"
	"net/url"
	"strings"
)

func init() {
	geotag.INCLUDE_LEAFLET = false // because the tangramjs stuff will add it
}

func AppendAssetHandlersWithFlagSet(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	err := tangramjs.AppendAssetHandlers(mux)

	if err != nil {
		return err
	}

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		return err
	}

	err = geotag.AppendAssetHandlers(mux)

	if err != nil {
		return err
	}

	err = www.AppendStaticAssetHandlers(mux)

	if err != nil {
		return err
	}

	return nil
}

func AppendApplicationHandler(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux, path string) error {

	handler, err := NewApplicationHandler(ctx, fs)

	if err != nil {
		return err
	}

	mux.Handle(path, handler)
	return nil
}

func NewApplicationHandler(ctx context.Context, fs *flag.FlagSet) (http.Handler, error) {

	t, err := NewApplicationTemplatesWithFlagSet(ctx, fs)

	if err != nil {
		return nil, err
	}

	nextzen_apikey, _ := flags.StringVar(fs, "nextzen-apikey")
	nextzen_style_url, _ := flags.StringVar(fs, "nextzen-style-url")
	nextzen_tile_url, _ := flags.StringVar(fs, "nextzen-tile-url")

	initial_latitude, _ := flags.Float64Var(fs, "initial-latitude")
	initial_longitude, _ := flags.Float64Var(fs, "initial-longitudex")
	initial_zoom, _ := flags.IntVar(fs, "initial-zoom")

	enable_placeholder, _ := flags.BoolVar(fs, "enable-placeholder")
	placeholder_endpoint, _ := flags.StringVar(fs, "placeholder-endpoint")

	enable_oembed, _ := flags.BoolVar(fs, "enable-oembed")
	oembed_endpoints, _ := flags.StringVar(fs, "oembed-endpoints")

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	tangramjs_opts := tangramjs.DefaultTangramJSOptions()
	tangramjs_opts.Nextzen.APIKey = nextzen_apikey
	tangramjs_opts.Nextzen.StyleURL = nextzen_style_url
	tangramjs_opts.Nextzen.TileURL = nextzen_tile_url

	geotag_opts := geotag.DefaultLeafletGeotagOptions()

	index_opts := &www.IndexHandlerOptions{
		Templates:           t,
		InitialLatitude:     initial_latitude,
		InitialLongitude:    initial_longitude,
		InitialZoom:         initial_zoom,
		EnablePlaceholder:   enable_placeholder,
		PlaceholderEndpoint: placeholder_endpoint,
	}

	if enable_oembed {

		index_opts.EnableOEmbed = enable_oembed

		urls := strings.Split(oembed_endpoints, ",")

		for _, oembed_url := range urls {

			_, err := url.Parse(oembed_url)

			if err != nil {
				return nil, err
			}
		}

		index_opts.OEmbedEndpoints = urls
	}

	index_handler, err := www.IndexHandler(index_opts)

	if err != nil {
		return nil, err
	}

	index_handler = bootstrap.AppendResourcesHandler(index_handler, bootstrap_opts)
	index_handler = tangramjs.AppendResourcesHandler(index_handler, tangramjs_opts)
	index_handler = geotag.AppendResourcesHandler(index_handler, geotag_opts)

	return index_handler, nil
}
