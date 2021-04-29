# go-www-geotag

![](docs/images/geotag-three-columns.png)

A web application, written in Go, for geotagging images.

## Important

This is work in progress, documentation is incomplete.

## Background

Have a look at the [Geotagging at SFO Museum](https://millsfield.sfomuseum.org/blog/tags/geotagging) series of blog posts.
 
## Tools

```
$> make cli
go build -mod vendor -o bin/server cmd/server/main.go
```

### server

```
$> ./bin/server -h
  -crumb-uri string
    	A valid aaronland/go-http-crumb.Crumb URI for generating (CSRF) crumbs. If the value is 'auto' then a random crumb URI will be generated. (default "auto")
  -custom-placetypes string
    	A JSON-encoded string containing custom placetypes defined using the syntax described in the whosonfirst/go-whosonfirst-placetypes repository.
  -disable-writer-crumb
    	Do not require a valid CSRF crumb for all writes.
  -enable-custom-placetypes
    	Enable wof:placetype values that are not explicitly defined in the whosonfirst/go-whosonfirst-placetypes repository.
  -enable-editor
    	Enable the geotagging editor interface. (default true)
  -enable-map-layers
    	Enable use of the leaflet-layers-control Leaflet control element for custom custom map overlays.
  -enable-oembed
    	Enable oEmbed lookups for images.
  -enable-placeholder
    	Enable use of the Placeholder API for location searches.
  -enable-point-in-polygon
    	Enable point-in-polygon lookups for results.
  -enable-proxy-tiles
    	Enable the use of a local tile proxy for Nextzen map tiles.
  -enable-writer
    	Enable output of the leaflet-geotag plugin to be written to a go-www-geotag/writer.Writer instance.
  -enable-writer-cors
    	Enable CORS support for the writer endpoint.
  -initial-latitude float
    	A valid latitude for the map's initial view. (default 37.61799)
  -initial-longitude float
    	A valid longitude for the map's initial view. (default -122.370943)
  -initial-zoom int
    	A valid zoom level for the map's initial view. (default 15)
  -is-wof
    	Input data is WOF-flavoured GeoJSON. (Pass a value of '0' or 'false' if you need to index non-WOF documents. (default true)
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/emitter URI. Supported schemes are: directory://, featurecollection://, file://, filelist://, geojsonl://, repo://. (default "repo://")
  -map-renderer string
    	Valid options are: protomaps, tangramjs (default "tangramjs")
  -nextzen-apikey string
    	A valid Nextzen API key
  -nextzen-style-url string
    	A valid URL for loading a Tangram.js style bundle. (default "/tangram/refill-style.zip")
  -nextzen-tile-url string
    	A valid Nextzen tile URL template for loading map tiles. (default "https://{s}.tile.nextzen.org/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt")
  -oembed-endpoints string
    	A comma-separated list of valid oEmbed endpoints to query.
  -path-editor string
    	A relative path for the geotag editor application. (default "/")
  -path-point-in-polygon string
    	The URI for point-in-polygon API requests. (default "/point-in-polygon/")
  -path-point-in-polygon-data string
    	The URI for point-in-polygon data requests. (default "/point-in-polygon/data/")
  -path-proxy-tiles string
    	The URL (a relative path) for proxied tiles. (default "/tiles/")
  -path-writer string
    	A relative path for sending write updates. (default "/update")
  -placeholder-endpoint string
    	A valid Placeholder API endpoint to query.
  -properties-reader-uri string
    	A valid whosonfirst/go-reader.Reader URI. Available options are: [file:// fs:// null://]
  -protomaps-tile-url string
    	A valid Protomaps tile URL for loading map tiles.
  -protomaps-tiles-path string
    	The leading prefix for Protomap tile URLs. (default "/pmtiles/")
  -proxy-tiles-cache-uri string
    	A valid tile proxy DSN string. (default "gocache://")
  -proxy-tiles-test
    	Ensure outbound network connectivity for proxy tiles
  -proxy-tiles-timeout int
    	The maximum number of seconds to allow for fetching a tile from the proxy. (default 30)
  -server-uri string
    	A valid aaronland/go-http-server.Server URI for creating an application server. (default "http://localhost:8080")
  -spatial-database-uri string
    	A valid whosonfirst/go-whosonfirst-spatial/data.SpatialDatabase URI. options are: [sqlite://]
  -verbose
    	Be chatty.
  -writer-cors-allowed-origins string
    	A comma-separated list of origins to allow for CORS support. (default "*")
  -writer-uri string
    	A valid go-www-geotag/writer.Writer URI for creating a writer.Writer instance. (default "stdout://")
```

For example:

```
$> bin/server \
	-map-renderer protomaps \
	-protomaps-tile-url file:///usr/local/data/pmtiles/sfo.pmtiles \
	-enable-oembed \
	-oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}&format=json' \
	-enable-point-in-polygon \
	-spatial-database-uri 'sqlite://?dsn=/usr/local/data/sfomuseum-architecture.db'
	
2021/04/28 13:14:41 Listening on http://localhost:8080
```

If you visit `http://localhost:8080` in your web browser you'll see something like this:

![](docs/images/geotag-three-columns.png)

In this configuration the `server` application a simple three-pane interface for geotagging images.

* The left-hand panel contains a map a `Leaflet.GeotagPhoto` camera for defining a focal point and field of view.

* The middle panel contains controls for loading images and for filtering and selecting reverse-geocoding results (that are updated as the camera's focal point changes).

* The right-hand panels contains GeoJSON `Feature` data containing information about the camera's focal point and field of view as well as any reverse-geocoding data, if present, encoded in the `wof:parent_id` and `wof:hierarchy` properties. 

This example application does not have any way to publish data outside of copy-pasting the raw GeoJSON data in to another document. That will be addressed in subsequent examples.

_Related: [Geotagging at SFO Museum, Part 3 – What Is the Simplest Thing?](https://millsfield.sfomuseum.org/blog/2020/04/27/geotagging-simple/)_

Watching through the parameters step-by-step the first being defined is the toolchain for rendering map tiles. In this example we're using [Protomaps.js](https://github.com/protomaps/protomaps.js) which renders maps tiles using a single "PMTiles" database file, specified in the `-protomaps-tile-url` flag, which can be hosted locally or on a remote file-server.

```
	-map-renderer protomaps \
	-protomaps-tile-url file:///usr/local/data/pmtiles/sfo.pmtiles \
```

_Related: [Protomaps: A new way to make maps with OpenStreetMap](https://protomaps.com/blog/new-way-to-make-maps/). To make your Protomaps database files try the [Protomaps bundle/download tool](https://protomaps.com/bundles)._

The second set of options enable controls in the user interface for loading images from a remote OEmbed endpoint. The URI in the `-oembed-endpoints` is where the application will resolve image requests.

```
	-enable-oembed \
	-oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}&format=json' \
```

_Related: [Geotagging at SFO Museum, Part 5 – Images](https://millsfield.sfomuseum.org/blog/2020/04/29/geotagging-images/)_

The final set of flags enabled "reverse geocoding" (sometimes called "point-in-polygon") lookups whenever you move the camera's focal point. These lookups are performed using the [go-whosonfirst-spatial](https://github.com/whosonfirst/go-whosonfirst-spatial) packages reading data stored in a SQLite database, specified by the `-spatial-database-uri`. flag

```
	-enable-point-in-polygon \
	-spatial-database-uri 'sqlite://?dsn=/usr/local/data/sfomuseum-architecture.db'
```

The SQLite databases are created using the `wof-sqlite-index-features` tool in the [go-whosonfirst-sqlite-features-index](https://github.com/whosonfirst/go-whosonfirst-sqlite-features-index) package, reading data from one or more Who's On First style repositories of data. For examples, this is how the `sfomuseum-architecture.db` database would be created reading data from the [sfomuseum-data-architecture](https://github.com/sfomuseum-data/sfomuseum-data-architecture) repository:

```
$> ./bin/wof-sqlite-index-features \
	-all \
	-dsn /usr/local/data/sfomuseum-architecture.db \
	/usr/local/data/sfomuseum-data-architecture/
	
2021/04/22 12:49:41 time to index paths (1) 2.702893585s
```

_Related: [Reverse-Geocoding in Time at SFO Museum](https://millsfield.sfomuseum.org/blog/2021/03/26/spatial/)_

...

```
$> bin/server \
	-map-renderer tangramjs \
	-nextzen-apikey {NEXTZEN_APIKEY} \
	-enable-oembed \
	-oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}&format=json' \
	-enable-point-in-polygon \
	-spatial-database-uri 'sqlite://?dsn=/usr/local/data/sfomuseum-architecture.db' \
	-enable-placeholder \
	-placeholder-endpoint http://localhost:3000
```

![](docs/images/geotag-three-columns-tangram.png)

```
	-map-renderer tangramjs \
	-nextzen-apikey {NEXTZEN_APIKEY} \
```

![](docs/images/geotag-three-columns-placeholder.png)

```
	-enable-placeholder \
	-placeholder-endpoint http://localhost:3000
```

![](docs/images/geotag-three-columns-gowanus.png)

## See also

* https://github.com/nypl-spacetime/Leaflet.GeotagPhoto
* https://github.com/sfomuseum/go-http-leaflet-geotag
* https://github.com/aaronland/go-http-leaflet
* https://github.com/aaronland/go-http-tangramjs
* https://github.com/sfomuseum/go-http-protomaps
* https://github.com/aaronland/go-http-bootstrap