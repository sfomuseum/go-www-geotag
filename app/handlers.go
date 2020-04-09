package app

import (
	"context"
	"errors"
	"flag"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/sfomuseum/go-http-leaflet-geotag"
	"github.com/sfomuseum/go-www-geotag/flags"
	"github.com/sfomuseum/go-www-geotag/geo"
	"github.com/sfomuseum/go-www-geotag/www"
	"net/http"
	"net/url"
	"strings"
)

func init() {
	geotag.INCLUDE_LEAFLET = false // because the tangramjs stuff will add it
}

func AppendAssetHandlers(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

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

	t, err := NewApplicationTemplates(ctx, fs)

	if err != nil {
		return nil, err
	}

	nextzen_apikey, err := flags.StringVar(fs, "nextzen-apikey")

	if err != nil {
		return nil, err
	}

	nextzen_style_url, err := flags.StringVar(fs, "nextzen-style-url")

	if err != nil {
		return nil, err
	}

	nextzen_tile_url, err := flags.StringVar(fs, "nextzen-tile-url")

	if err != nil {
		return nil, err
	}

	initial_latitude, err := flags.Float64Var(fs, "initial-latitude")

	if err != nil {
		return nil, err
	}

	if !geo.IsValidLatitude(initial_latitude) {
		return nil, errors.New("Invalid latitude")
	}

	initial_longitude, err := flags.Float64Var(fs, "initial-longitude")

	if err != nil {
		return nil, err
	}

	if !geo.IsValidLongitude(initial_longitude) {
		return nil, errors.New("Invalid longitude")
	}

	initial_zoom, err := flags.IntVar(fs, "initial-zoom")

	if err != nil {
		return nil, err
	}

	if initial_zoom < 1 || initial_zoom > 21 {
		return nil, errors.New("Invalid zoom")
	}

	enable_placeholder, err := flags.BoolVar(fs, "enable-placeholder")

	if err != nil {
		return nil, err
	}

	placeholder_endpoint, err := flags.StringVar(fs, "placeholder-endpoint")

	if err != nil {
		return nil, err
	}

	enable_oembed, err := flags.BoolVar(fs, "enable-oembed")

	if err != nil {
		return nil, err
	}

	oembed_endpoints, err := flags.StringVar(fs, "oembed-endpoints")

	if err != nil {
		return nil, err
	}

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	tangramjs_opts := tangramjs.DefaultTangramJSOptions()
	tangramjs_opts.Nextzen.APIKey = nextzen_apikey
	tangramjs_opts.Nextzen.StyleURL = nextzen_style_url
	tangramjs_opts.Nextzen.TileURL = nextzen_tile_url

	geotag_opts := geotag.DefaultLeafletGeotagOptions()

	index_opts := &www.IndexHandlerOptions{
		Templates:        t,
		InitialLatitude:  initial_latitude,
		InitialLongitude: initial_longitude,
		InitialZoom:      initial_zoom,
	}

	if enable_placeholder {

		_, err := url.Parse(placeholder_endpoint)

		if err != nil {
			return nil, err
		}

		index_opts.EnablePlaceholder = enable_placeholder
		index_opts.PlaceholderEndpoint = placeholder_endpoint
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
