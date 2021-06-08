# go-http-leaflet-geotag

![](docs/images/camera_sfo.png)

Go HTTP middleware for the Leaflet.GeotagPhoto plugin.

## Important

Documentation is incomplete.

## Example

```
$> go run -mod vendor cmd/example/main.go 
2020/04/06 11:28:37 Listening for requests on localhost:8080
```

## An abbreviated code example

This is an _abbreviated example_ of code to demonstrate appending `Leaflet.GeotagPhoto` related handlers to an `http.ServeMux` instance and updating a user-defined `http.Handler` to append `Leaflet.GeotagPhoto` Javascript and CSS links to its output.

_Error handling has been removed for the sake of brevity._

```
import (
	"github.com/sfomuseum/go-http-leaflet-geotag"
	"net/http"
)

func main() {

	// code...
     	
	geotag_opts := geotag.DefaultLeafletGeotagOptions()
	
	mux := http.NewServeMux()

	geotag.AppendAssetHandlers(mux)

	camera_handler, _ := PageHandler(t, "camera")

	camera_handler = geotag.AppendResourcesHandler(camera_handler, geotag_opts)

	mux.Handle("/camera/", camera_handler)

	// code...
}	
```

For a complete example please consult [cmd/example/main.go](cmd/example/main.go). The (Javascript) code for the web application itself is contained in the HTML files in the [templates/html](templates/html) directory.

## Changes

### Leaflet.GeotagPhoto.js

This package uses a modified version of the `Leaflet.GeotagPhoto.js` file that is part of the original [Leaflet.GeotagPhoto](https://github.com/sfomuseum/Leaflet.GeotagPhoto). It encodes all the static assets (icons and markers) as base64-encoded data URIs rather than linking to the resources themselves. This is to account for situations where an application may be run inside or behind one or more load balancers or proxies and where keeping track of URL rewriting rules becomes a burden.

## See also

* https://github.com/sfomuseum/Leaflet.GeotagPhoto
* https://github.com/aaronland/go-http-leaflet