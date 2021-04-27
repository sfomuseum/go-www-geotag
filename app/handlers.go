package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-crumb"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/jtacoma/uritemplates"
	"github.com/rs/cors"
	"github.com/sfomuseum/go-flags/lookup"
	"github.com/sfomuseum/go-http-leaflet-geotag"
	"github.com/sfomuseum/go-http-leaflet-layers"
	"github.com/sfomuseum/go-http-protomaps"
	tzhttp "github.com/sfomuseum/go-http-tilezen/http"
	"github.com/sfomuseum/go-www-geotag/api"
	"github.com/sfomuseum/go-www-geotag/geo"
	"github.com/sfomuseum/go-www-geotag/writer"
	"github.com/sfomuseum/go-www-geotag/www"
	"github.com/whosonfirst/go-cache"
	pip_api "github.com/whosonfirst/go-whosonfirst-spatial-pip/api"
	spatial_app "github.com/whosonfirst/go-whosonfirst-spatial/app"
	spatial_www "github.com/whosonfirst/go-whosonfirst-spatial-www/http"	
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var crumb_init sync.Once

var cr crumb.Crumb
var crumb_err error

func init() {
	// because the tangramjs or the protomaps stuff will add it
	geotag.INCLUDE_LEAFLET = false
	layers.INCLUDE_LEAFLET = false
}

func AppendAssetHandlers(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	map_renderer, err := lookup.StringVar(fs, "map-renderer")

	if err != nil {
		return err
	}

	// PROTOMAPS: this remains to be reconciled with the Tangram.js + Nextzen stuff
	// In the end what you see below might just be the simplest way to "reconcile"
	// things (20210423/thisisaaronland)

	switch map_renderer {
	case "protomaps":

		err := protomaps.AppendAssetHandlers(mux)

		if err != nil {
			return err
		}

	case "tangramjs":

		err := tangramjs.AppendAssetHandlers(mux)

		if err != nil {
			return err
		}

	default:
		// pass
	}

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		return err
	}

	err = geotag.AppendAssetHandlers(mux)

	if err != nil {
		return err
	}

	enable_map_layers, err := lookup.BoolVar(fs, "enable-map-layers")

	if err != nil {
		return err
	}

	if enable_map_layers {

		err = layers.AppendAssetHandlers(mux)

		if err != nil {
			return err
		}
	}

	err = www.AppendStaticAssetHandlers(mux)

	if err != nil {
		return err
	}

	return nil
}

func AppendPointInPolygonHandlerIfEnabled(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	enable_pip, err := lookup.BoolVar(fs, "enable-point-in-polygon")

	if err != nil {
		return err
	}

	if !enable_pip {
		return nil
	}

	return AppendPointInPolygonHandler(ctx, fs, mux)
}

func AppendPointInPolygonHandler(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	app, err := spatial_app.NewSpatialApplicationWithFlagSet(ctx, fs)

	if err != nil {
		return fmt.Errorf("Failed to create new spatial application, %v", err)
	}

	pip_opts := &pip_api.PointInPolygonHandlerOptions{
		// EnableGeoJSON: true,
	}

	pip_handler, err := pip_api.PointInPolygonHandler(app, pip_opts)

	if err != nil {
		return fmt.Errorf("Failed to create PointInPolygonHandler handler, %v", err)
	}

	path_pip, err := lookup.StringVar(fs, "path-point-in-polygon")

	if err != nil {
		return fmt.Errorf("Failed to create PointInPolygonHandler handler - unable to lookup -path-point-in-polygon flag, %v", err)
	}

	mux.Handle(path_pip, pip_handler)

	data_handler, err := spatial_www.NewDataHandler(app.SpatialDatabase)

	if err != nil {
		return fmt.Errorf("Failed to create DataHandler, %v", err)
	}

	path_pip_data, err := lookup.StringVar(fs, "path-point-in-polygon-data")

	if err != nil {
		return fmt.Errorf("Failed to create DataHandler handler - unable to lookup -path-point-in-polygon-data flag, %v", err)
	}

	mux.Handle(path_pip_data, data_handler)
	return nil
}

func AppendEditorHandlerIfEnabled(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	enable_editor, err := lookup.BoolVar(fs, "enable-editor")

	if err != nil {
		return err
	}

	if !enable_editor {
		return nil
	}

	path, err := lookup.StringVar(fs, "path-editor")

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

	map_renderer, err := lookup.StringVar(fs, "map-renderer")

	if err != nil {
		return nil, err
	}

	nextzen_apikey, err := lookup.StringVar(fs, "nextzen-apikey")

	if err != nil {
		return nil, err
	}

	nextzen_style_url, err := lookup.StringVar(fs, "nextzen-style-url")

	if err != nil {
		return nil, err
	}

	nextzen_tile_url, err := lookup.StringVar(fs, "nextzen-tile-url")

	if err != nil {
		return nil, err
	}

	protomaps_tile_url, err := lookup.StringVar(fs, "protomaps-tile-url")

	if err != nil {
		return nil, err
	}

	initial_latitude, err := lookup.Float64Var(fs, "initial-latitude")

	if err != nil {
		return nil, err
	}

	if !geo.IsValidLatitude(initial_latitude) {
		return nil, errors.New("Invalid latitude")
	}

	initial_longitude, err := lookup.Float64Var(fs, "initial-longitude")

	if err != nil {
		return nil, err
	}

	if !geo.IsValidLongitude(initial_longitude) {
		return nil, errors.New("Invalid longitude")
	}

	initial_zoom, err := lookup.IntVar(fs, "initial-zoom")

	if err != nil {
		return nil, err
	}

	if initial_zoom < 1 || initial_zoom > 21 {
		return nil, errors.New("Invalid zoom")
	}

	enable_placeholder, err := lookup.BoolVar(fs, "enable-placeholder")

	if err != nil {
		return nil, err
	}

	placeholder_endpoint, err := lookup.StringVar(fs, "placeholder-endpoint")

	if err != nil {
		return nil, err
	}

	enable_pip, err := lookup.BoolVar(fs, "enable-point-in-polygon")

	if err != nil {
		return nil, err
	}

	path_pip, err := lookup.StringVar(fs, "path-point-in-polygon")

	if err != nil {
		return nil, err
	}

	path_pip_data, err := lookup.StringVar(fs, "path-point-in-polygon-data")

	if err != nil {
		return nil, err
	}
	
	enable_oembed, err := lookup.BoolVar(fs, "enable-oembed")

	if err != nil {
		return nil, err
	}

	oembed_endpoints, err := lookup.StringVar(fs, "oembed-endpoints")

	if err != nil {
		return nil, err
	}

	enable_proxy_tiles, err := lookup.BoolVar(fs, "enable-proxy-tiles")

	if err != nil {
		return nil, err
	}

	enable_writer, err := lookup.BoolVar(fs, "enable-writer")

	if err != nil {
		return nil, err
	}

	path_writer, err := lookup.StringVar(fs, "path-writer")

	if err != nil {
		return nil, err
	}

	enable_map_layers, err := lookup.BoolVar(fs, "enable-map-layers")

	if err != nil {
		return nil, err
	}

	// PROTOMAPS: this needs to stay in sync with code in AppendProtomapsTilesHandlerIfNecessary
	// This takes a singular protomaps-tile-url and if it is a file:// URI
	// its filename is joined with protomaps-tiles-path to create a new path and this
	// path is registered with the http.ServeMux (and handled by the http.Dir
	// instance). It would be nice if this could just be handled by / hidden in
	// sfomuseum/go-http-leaflet-protomaps but I am not sure if that actually
	// makes sense yet (20210423/thisisaaronland)

	if map_renderer == "protomaps" {

		u, err := url.Parse(protomaps_tile_url)

		if err != nil {
			return nil, err
		}

		if u.Scheme == "file" {

			protomaps_tiles_path, err := lookup.StringVar(fs, "protomaps-tiles-path")

			if err != nil {
				return nil, err
			}

			abs_path, err := filepath.Abs(u.Path)

			if err != nil {
				return nil, err
			}

			fname := filepath.Base(abs_path)

			protomaps_tile_url = filepath.Join(protomaps_tiles_path, fname)
		}
	}

	geotag_opts := geotag.DefaultLeafletGeotagOptions()

	editor_opts := &www.EditorHandlerOptions{
		Templates:        t,
		InitialLatitude:  initial_latitude,
		InitialLongitude: initial_longitude,
		InitialZoom:      initial_zoom,
		MapRenderer:      map_renderer,
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

	if enable_pip {

		editor_opts.EnablePointInPolygon = enable_pip
		editor_opts.PointInPolygonEndpoint = path_pip
		editor_opts.PointInPolygonDataEndpoint = path_pip_data
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

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()
	editor_handler = bootstrap.AppendResourcesHandler(editor_handler, bootstrap_opts)

	switch map_renderer {
	case "protomaps":

		pm_opts := protomaps.DefaultProtomapsOptions()
		pm_opts.TileURL = protomaps_tile_url

		editor_handler = protomaps.AppendResourcesHandler(editor_handler, pm_opts)

	case "tangramjs":

		if enable_proxy_tiles {

			path_proxy_tiles, err := lookup.StringVar(fs, "path-proxy-tiles")

			if err != nil {
				return nil, err
			}

			nextzen_tile_url = fmt.Sprintf("%s{z}/{x}/{y}.mvt", path_proxy_tiles)
		}

		tangramjs_opts := tangramjs.DefaultTangramJSOptions()
		tangramjs_opts.Nextzen.APIKey = nextzen_apikey
		tangramjs_opts.Nextzen.StyleURL = nextzen_style_url
		tangramjs_opts.Nextzen.TileURL = nextzen_tile_url

		editor_handler = tangramjs.AppendResourcesHandler(editor_handler, tangramjs_opts)

	default:
		return nil, fmt.Errorf("Invalid or unsupported -map-renderer flag ('%s')", map_renderer)
	}

	editor_handler = geotag.AppendResourcesHandler(editor_handler, geotag_opts)

	if enable_map_layers {
		layers_opts := layers.DefaultLeafletLayersOptions()
		editor_handler = layers.AppendResourcesHandler(editor_handler, layers_opts)
	}

	editor_handler, err = AppendCrumbHandler(ctx, fs, editor_handler)

	if err != nil {
		return nil, err
	}

	return editor_handler, nil
}

// PROTOMAPS: this needs to stay in sync with code in NewEditorHandler
// This takes a singular protomaps-tile-url and if it is a file:// URI
// creates an http.Dir instance for its parent directory. The filename (for protomaps-tile-url)
// is joined with protomaps-tiles-path to create a new path and this
// path is registered with the http.ServeMux (and handled by the http.Dir
// instance). It would be nice if this could just be handled by / hidden in
// sfomuseum/go-http-leaflet-protomaps but I am not sure if that actually
// makes sense yet (20210423/thisisaaronland)

func AppendProtomapsTilesHandlerIfNecessary(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	protomaps_tile_url, err := lookup.StringVar(fs, "protomaps-tile-url")

	if err != nil {
		return err
	}

	protomaps_tiles_path, err := lookup.StringVar(fs, "protomaps-tiles-path")

	if err != nil {
		return err
	}

	map_renderer, err := lookup.StringVar(fs, "map-renderer")

	if err != nil {
		return err
	}

	if map_renderer != "protomaps" {
		return nil
	}

	u, err := url.Parse(protomaps_tile_url)

	if err != nil {
		return err
	}

	if u.Scheme != "file" {
		return nil
	}

	mux_url, mux_handler, err := protomaps.FileHandlerFromPath(u.Path, protomaps_tiles_path)

	if err != nil {
		return err
	}

	mux.Handle(mux_url, mux_handler)
	return nil
}

func AppendWriterHandlerIfEnabled(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	enable_writer, err := lookup.BoolVar(fs, "enable-writer")

	if err != nil {
		return err
	}

	if !enable_writer {
		return nil
	}

	return AppendWriterHandler(ctx, fs, mux)
}

func AppendWriterHandler(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	path, err := lookup.StringVar(fs, "path-writer")

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

	writer_uri, err := lookup.StringVar(fs, "writer-uri")

	if err != nil {
		return nil, err
	}

	disable_writer_crumb, err := lookup.BoolVar(fs, "disable-writer-crumb")

	if err != nil {
		return nil, err
	}

	enable_writer_cors, err := lookup.BoolVar(fs, "enable-writer-cors")

	if err != nil {
		return nil, err
	}

	allowed_origins_str, err := lookup.StringVar(fs, "writer-cors-allowed-origins")

	if err != nil {
		return nil, err
	}

	wr, err := writer.NewWriter(ctx, writer_uri)

	if err != nil {
		return nil, err
	}

	handler, err := api.WriterHandler(wr)

	if err != nil {
		return nil, err
	}

	if !disable_writer_crumb {

		handler, err = AppendCrumbHandler(ctx, fs, handler)

		if err != nil {
			return nil, err
		}
	}

	if enable_writer_cors {

		allowed_origins := strings.Split(allowed_origins_str, ",")

		cors_handler := cors.New(cors.Options{
			AllowedOrigins: allowed_origins,
			AllowedMethods: []string{"PUT"},
		})

		handler = cors_handler.Handler(handler)
	}

	return handler, nil
}

func AppendProxyTilesHandlerIfEnabled(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	enable_proxy_tiles, err := lookup.BoolVar(fs, "enable-proxy-tiles")

	if err != nil {
		return err
	}

	if !enable_proxy_tiles {
		return nil
	}

	return AppendProxyTilesHandler(ctx, fs, mux)
}

func AppendProxyTilesHandler(ctx context.Context, fs *flag.FlagSet, mux *http.ServeMux) error {

	path_proxy_tiles, err := lookup.StringVar(fs, "path-proxy-tiles")

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

	proxy_tiles_cache_uri, err := lookup.StringVar(fs, "proxy-tiles-cache-uri")

	if err != nil {
		return nil, err
	}

	proxy_tiles_timeout, err := lookup.IntVar(fs, "proxy-tiles-timeout")

	if err != nil {
		return nil, err
	}

	test_proxy_tiles, err := lookup.BoolVar(fs, "proxy-tiles-test")

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

func AppendCrumbHandler(ctx context.Context, fs *flag.FlagSet, handler http.Handler) (http.Handler, error) {

	crumb_uri, err := lookup.StringVar(fs, "crumb-uri")

	if err != nil {
		return nil, err
	}

	if crumb_uri == "disabled" {
		log.Printf("[WARNING] -crumb-uri explicitly disabled. This is probably insecure.")
		return handler, nil
	}

	crumb_config, err := crumbConfigWithFlagSet(ctx, fs)

	if err != nil {
		return nil, err
	}

	handler = crumb.EnsureCrumbHandler(crumb_config, handler)
	return handler, nil
}

func crumbConfigWithFlagSet(ctx context.Context, fs *flag.FlagSet) (crumb.Crumb, error) {

	crumb_func := func() {

		crumb_uri, err := lookup.StringVar(fs, "crumb-uri")

		if err != nil {
			crumb_err = err
			return
		}

		if crumb_uri == "auto" {

			ttl := 3600
			key := "geotag"

			uri, err := crumb.NewRandomEncryptedCrumbURI(ctx, ttl, key)

			if err != nil {
				crumb_err = err
				return
			}

			crumb_uri = uri

		} else {

			u, err := url.Parse(crumb_uri)

			if err != nil {
				crumb_err = err
				return
			}

			q := u.Query()

			if q.Get("key") == "" {
				crumb_err = errors.New("Required key= property for crumb URI missing")
				return
			}
		}

		cr, crumb_err = crumb.NewCrumb(ctx, crumb_uri)
	}

	crumb_init.Do(crumb_func)

	if crumb_err != nil {
		return nil, crumb_err
	}

	if cr == nil {
		return nil, errors.New("Failed to generate new crumb")
	}

	return cr, nil
}
