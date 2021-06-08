package layers

import (
	"fmt"
	"github.com/aaronland/go-http-leaflet"
	"github.com/aaronland/go-http-rewrite"
	"github.com/sfomuseum/go-http-leaflet-layers/static"
	"io/fs"
	_ "log"
	"net/http"
	"path/filepath"
	"strings"
)

// By default the go-http-protomaps package will also include and reference Leaflet.js resources using the aaronland/go-http-leaflet package. If you want or need to disable this behaviour set this variable to false.
var APPEND_LEAFLET_RESOURCES = true

// By default the go-http-protomaps package will also include and reference Leaflet.js assets using the aaronland/go-http-leaflet package. If you want or need to disable this behaviour set this variable to false.
var APPEND_LEAFLET_ASSETS = true

// LeafletLayersOptions provides a list of JavaScript and CSS link to include with HTML output.
type LeafletLayersOptions struct {
	JS  []string
	CSS []string
}

// Return a *LeafletLayersOptions struct with default paths and URIs.
func DefaultLeafletLayersOptions() *LeafletLayersOptions {

	opts := &LeafletLayersOptions{
		CSS: []string{
			"/css/leaflet.layers.control.css",
		},
		JS: []string{
			"/javascript/leaflet.layers.control.js",
		},
	}

	return opts
}

// AppendResourcesHandler will rewrite any HTML produced by previous handler to include the necessary markup to load LeafletLayers JavaScript files and related assets.
func AppendResourcesHandler(next http.Handler, opts *LeafletLayersOptions) http.Handler {

	if APPEND_LEAFLET_RESOURCES {
		leaflet_opts := leaflet.DefaultLeafletOptions()
		next = leaflet.AppendResourcesHandler(next, leaflet_opts)
	}

	return AppendResourcesHandlerWithPrefix(next, opts, "")
}

// AppendResourcesHandlerWithPrefix will rewrite any HTML produced by previous handler to include the necessary markup to load LeafletLayers JavaScript files and related assets ensuring that all URIs are prepended with a prefix.
func AppendResourcesHandlerWithPrefix(next http.Handler, opts *LeafletLayersOptions, prefix string) http.Handler {

	if APPEND_LEAFLET_RESOURCES {
		leaflet_opts := leaflet.DefaultLeafletOptions()
		next = leaflet.AppendResourcesHandlerWithPrefix(next, leaflet_opts, prefix)
	}

	js := opts.JS
	css := opts.CSS

	if prefix != "" {

		for i, path := range js {
			js[i] = appendPrefix(prefix, path)
		}

		for i, path := range css {
			css[i] = appendPrefix(prefix, path)
		}
	}

	ext_opts := &rewrite.AppendResourcesOptions{
		JavaScript:  js,
		Stylesheets: css,
	}

	return rewrite.AppendResourcesHandler(next, ext_opts)
}

// AssetsHandler returns a net/http FS instance containing the embedded LeafletLayers assets that are included with this package.
func AssetsHandler() (http.Handler, error) {
	http_fs := http.FS(static.FS)
	return http.FileServer(http_fs), nil
}

// AssetsHandler returns a net/http FS instance containing the embedded LeafletLayers assets that are included with this package ensuring that all URLs are stripped of prefix.
func AssetsHandlerWithPrefix(prefix string) (http.Handler, error) {

	fs_handler, err := AssetsHandler()

	if err != nil {
		return nil, err
	}

	prefix = strings.TrimRight(prefix, "/")

	if prefix == "" {
		return fs_handler, nil
	}

	rewrite_func := func(req *http.Request) (*http.Request, error) {
		req.URL.Path = strings.Replace(req.URL.Path, prefix, "", 1)
		return req, nil
	}

	rewrite_handler := rewrite.RewriteRequestHandler(fs_handler, rewrite_func)
	return rewrite_handler, nil
}

// Append all the files in the net/http FS instance containing the embedded LeafletLayers assets to an *http.ServeMux instance.
func AppendAssetHandlers(mux *http.ServeMux) error {
	return AppendAssetHandlersWithPrefix(mux, "")
}

// Append all the files in the net/http FS instance containing the embedded LeafletLayers assets to an *http.ServeMux instance ensuring that all URLs are prepended with prefix.
func AppendAssetHandlersWithPrefix(mux *http.ServeMux, prefix string) error {

	if APPEND_LEAFLET_ASSETS {

		err := leaflet.AppendAssetHandlersWithPrefix(mux, prefix)

		if err != nil {
			return err
		}
	}

	asset_handler, err := AssetsHandlerWithPrefix(prefix)

	if err != nil {
		return nil
	}

	walk_func := func(path string, info fs.DirEntry, err error) error {

		if path == "." {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if prefix != "" {
			path = appendPrefix(prefix, path)
		}

		if !strings.HasPrefix(path, "/") {
			path = fmt.Sprintf("/%s", path)
		}

		// log.Println("APPEND", path)

		mux.Handle(path, asset_handler)
		return nil
	}

	return fs.WalkDir(static.FS, ".", walk_func)
}

func appendPrefix(prefix string, path string) string {

	prefix = strings.TrimRight(prefix, "/")

	if prefix != "" {
		path = strings.TrimLeft(path, "/")
		path = filepath.Join(prefix, path)
	}

	return path
}
