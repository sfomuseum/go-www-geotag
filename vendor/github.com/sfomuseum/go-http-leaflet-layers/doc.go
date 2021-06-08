// `go-http-leaflet-layers` is an HTTP middleware package for including LeafletLayers.js assets in web applications. It exports two principal methods:
//
// * `layers.AppendAssetHandlers(*http.ServeMux)` which is used to append HTTP handlers to a `http.ServeMux` instance for serving LeafletLayers JavaScript files, and related assets.
// * `layers.AppendResourcesHandler(http.Handler, *LeafletLayersOptions)` which is used to rewrite any HTML produced by previous handler to include the necessary markup to load LeafletLayers JavaScript files and related assets.
package layers
