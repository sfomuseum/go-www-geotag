package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/jtacoma/uritemplates"
	"github.com/sfomuseum/go-http-leaflet-geotag"
	tzhttp "github.com/sfomuseum/go-http-tilezen/http"
	"github.com/sfomuseum/go-www-geotag/flags"
	"github.com/sfomuseum/go-www-geotag/geo"
	"github.com/sfomuseum/go-www-geotag/writer"
	"github.com/sfomuseum/go-www-geotag/www"
	"github.com/whosonfirst/go-cache"
	_ "log"
	"net/http"
	"net/url"
	"strings"
	"time"
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

func AppendEditorHandler(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	path, err := flags.StringVar(fs, "path-editor")

	if err != nil {
		return err
	}

	handler, err := NewEditorHandler(ctx, fs)

	if err != nil {
		return err
	}

	mux.Handle(path, handler)
	return nil
}

func NewEditorHandler(ctx context.Context, fs *flag.FlagSet) (http.Handler, error) {

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

	enable_proxy_tiles, err := flags.BoolVar(fs, "enable-proxy-tiles")

	if err != nil {
		return nil, err
	}

	enable_writer, err := flags.BoolVar(fs, "enable-writer")

	if err != nil {
		return nil, err
	}

	path_writer, err := flags.StringVar(fs, "path-writer")

	if err != nil {
		return nil, err
	}

	if enable_proxy_tiles {

		path_proxy_tiles, err := flags.StringVar(fs, "path-proxy-tiles")

		if err != nil {
			return nil, err
		}

		nextzen_tile_url = fmt.Sprintf("%s{z}/{x}/{y}.mvt", path_proxy_tiles)
	}

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	tangramjs_opts := tangramjs.DefaultTangramJSOptions()
	tangramjs_opts.Nextzen.APIKey = nextzen_apikey
	tangramjs_opts.Nextzen.StyleURL = nextzen_style_url
	tangramjs_opts.Nextzen.TileURL = nextzen_tile_url

	geotag_opts := geotag.DefaultLeafletGeotagOptions()

	editor_opts := &www.EditorHandlerOptions{
		Templates:        t,
		InitialLatitude:  initial_latitude,
		InitialLongitude: initial_longitude,
		InitialZoom:      initial_zoom,
	}

	if enable_writer {
		editor_opts.EnableWriter = enable_writer
		editor_opts.WriterPath = path_writer
	}

	if enable_placeholder {

		_, err := url.Parse(placeholder_endpoint)

		if err != nil {
			return nil, err
		}

		editor_opts.EnablePlaceholder = enable_placeholder
		editor_opts.PlaceholderEndpoint = placeholder_endpoint
	}

	if enable_oembed {

		editor_opts.EnableOEmbed = enable_oembed

		urls := strings.Split(oembed_endpoints, ",")

		for _, oembed_url := range urls {

			_, err := url.Parse(oembed_url)

			if err != nil {
				return nil, err
			}

			t, err := uritemplates.Parse(oembed_url)

			if err != nil {
				return nil, err
			}

			has_url := false

			for _, n := range t.Names() {

				if n == "url" {
					has_url = true
					break
				}
			}

			if !has_url {
				return nil, errors.New("Invalid oEmbed endpoint URL template")
			}
		}

		editor_opts.OEmbedEndpoints = urls
	}

	editor_handler, err := www.EditorHandler(editor_opts)

	if err != nil {
		return nil, err
	}

	editor_handler = bootstrap.AppendResourcesHandler(editor_handler, bootstrap_opts)
	editor_handler = tangramjs.AppendResourcesHandler(editor_handler, tangramjs_opts)
	editor_handler = geotag.AppendResourcesHandler(editor_handler, geotag_opts)

	return editor_handler, nil
}

func AppendWriterHandlerIfEnabled(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	enable_writer, err := flags.BoolVar(fs, "enable-writer")

	if err != nil {
		return err
	}

	if !enable_writer {
		return nil
	}

	return AppendWriterHandler(ctx, fs, mux)
}

func AppendWriterHandler(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	path, err := flags.StringVar(fs, "path-writer")

	if err != nil {
		return err
	}

	handler, err := NewWriterHandler(ctx, fs)

	if err != nil {
		return err
	}

	mux.Handle(path, handler)
	return nil
}

func NewWriterHandler(ctx context.Context, fs *flag.FlagSet) (http.Handler, error) {

	writer_uri, err := flags.StringVar(fs, "writer-uri")

	if err != nil {
		return nil, err
	}

	wr, err := writer.NewWriter(ctx, writer_uri)

	if err != nil {
		return nil, err
	}

	return www.WriterHandler(wr)
}

func AppendProxyTilesHandlerIfEnabled(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	enable_proxy_tiles, err := flags.BoolVar(fs, "enable-proxy-tiles")

	if err != nil {
		return err
	}

	if !enable_proxy_tiles {
		return nil
	}

	return AppendProxyTilesHandler(ctx, fs, mux)
}

func AppendProxyTilesHandler(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	path_proxy_tiles, err := flags.StringVar(fs, "path-proxy-tiles")

	if err != nil {
		return err
	}

	handler, err := NewProxyTilesHandler(ctx, fs)

	if err != nil {
		return err
	}

	mux.Handle(path_proxy_tiles, handler)
	return nil
}

func NewProxyTilesHandler(ctx context.Context, fs *flag.FlagSet) (http.Handler, error) {

	proxy_tiles_cache_uri, err := flags.StringVar(fs, "proxy-tiles-cache-uri")

	if err != nil {
		return nil, err
	}

	proxy_tiles_timeout, err := flags.IntVar(fs, "proxy-tiles-timeout")

	if err != nil {
		return nil, err
	}

	test_proxy_tiles, err := flags.BoolVar(fs, "proxy-tiles-test")

	if err != nil {
		return nil, err
	}

	tile_cache, err := cache.NewCache(ctx, proxy_tiles_cache_uri)

	if err != nil {
		return nil, err
	}

	timeout := time.Duration(proxy_tiles_timeout) * time.Second

	proxy_opts := &tzhttp.TilezenProxyHandlerOptions{
		Cache:   tile_cache,
		Timeout: timeout,
	}

	proxy_handler, err := tzhttp.TilezenProxyHandler(proxy_opts)

	if err != nil {
		return nil, err
	}

	if test_proxy_tiles {

		test_ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		req, err := http.NewRequest("GET", "tile.nextzen.org", nil)

		if err != nil {
			return nil, err
		}

		cl := new(http.Client)

		_, err = cl.Do(req.WithContext(test_ctx))

		if err != nil {
			return nil, err
		}

	}

	return proxy_handler, nil
}
