package app

import (
	"flag"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/sfomuseum/go-flags/flagset"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
)

func CommonFlags() (*flag.FlagSet, error) {

	fs := flagset.NewFlagSet("geotag")

	fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server.Server URI for creating an application server.")

	fs.Bool("enable-editor", true, "Enable the geotagging editor interface.")

	fs.Bool("enable-placeholder", false, "Enable use of the Placeholder API for location searches.")
	fs.String("placeholder-endpoint", "", "A valid Placeholder API endpoint to query.")

	fs.Bool("enable-oembed", false, "Enable oEmbed lookups for images.")
	fs.String("oembed-endpoints", "", "A comma-separated list of valid oEmbed endpoints to query.")

	fs.Bool("enable-writer", false, "Enable output of the leaflet-geotag plugin to be written to a go-www-geotag/writer.Writer instance.")
	fs.String("writer-uri", "stdout://", "A valid go-www-geotag/writer.Writer URI for creating a writer.Writer instance OR 'exif://' which will enable embedding focal point but not field-of-view geotagging information focal point but not field-of-view in a downloadable JPEG representation of an image.")

	fs.Bool("disable-writer-crumb", false, "Do not require a valid CSRF crumb for all writes.")

	fs.Bool("enable-writer-cors", false, "Enable CORS support for the writer endpoint.")
	fs.String("writer-cors-allowed-origins", "*", "A comma-separated list of origins to allow for CORS support.")

	fs.Bool("enable-map-layers", false, "Enable use of the leaflet-layers-control Leaflet control element for custom custom map overlays.")

	fs.String("path-editor", "/", "A relative path for the geotag editor application.")
	fs.String("path-writer", "/update", "A relative path for sending write updates.")

	fs.String("crumb-uri", "auto", "A valid aaronland/go-http-crumb.Crumb URI for generating (CSRF) crumbs. If the value is 'auto' then a random crumb URI will be generated.")

	err := AppendMapFlags(fs)

	if err != nil {
		return nil, err
	}

	err = AppendPointInPolygonFlags(fs)

	if err != nil {
		return nil, err
	}

	return fs, nil
}

func AppendPointInPolygonFlags(fs *flag.FlagSet) error {

	fs.Bool("enable-point-in-polygon", false, "Enable point-in-polygon lookups for results.")

	fs.String("path-point-in-polygon", "/point-in-polygon/", "The URI for point-in-polygon API requests.")
	fs.String("path-point-in-polygon-data", "/point-in-polygon/data/", "The URI for point-in-polygon data requests.")

	err := spatial_flags.AppendCommonFlags(fs)

	if err != nil {
		return err
	}

	err = spatial_flags.AppendIndexingFlags(fs)

	if err != nil {
		return err
	}

	return nil
}

func AppendMapFlags(fs *flag.FlagSet) error {

	fs.String("map-renderer", "tangramjs", "Valid options are: protomaps, tangramjs")

	err := AppendLeafletFlags(fs)

	if err != nil {
		return err
	}

	err = AppendProtomapsFlags(fs)

	if err != nil {
		return err
	}

	err = AppendNextzenFlags(fs)

	if err != nil {
		return err
	}

	err = AppendTilezenFlags(fs)

	if err != nil {
		return err
	}

	return nil
}

func AppendProtomapsFlags(fs *flag.FlagSet) error {

	fs.String("protomaps-tile-url", "", "A valid Protomaps tile URL for loading map tiles.")
	fs.String("protomaps-tiles-path", "/pmtiles/", "The leading prefix for Protomap tile URLs.")

	return nil
}

// because eventually we're going to put this in a nextzen specific package
// (20200408/thisisaaronland)

func AppendNextzenFlags(fs *flag.FlagSet) error {

	fs.String("nextzen-apikey", "", "A valid Nextzen API key")
	fs.String("nextzen-style-url", "/tangram/refill-style.zip", "A valid URL for loading a Tangram.js style bundle.")
	fs.String("nextzen-tile-url", tangramjs.NEXTZEN_MVT_ENDPOINT, "A valid Nextzen tile URL template for loading map tiles.")

	return nil
}

// because eventually we're going to put this in a sfomuseum/go-http-tilezen specific package
// (20200410/thisisaaronland)

func AppendTilezenFlags(fs *flag.FlagSet) error {

	// https://github.com/tilezen/go-tilepacks

	fs.Bool("enable-tilezen-tilepacks", false, "Enable to use of local Tilezen MBTiles database.")
	fs.String("tilezen-path-tilepack", "", "A valid path for a local Tilezen MBTiles database.")
	fs.String("tilezen-url-tiles", "/tilezen", "The URL (a relative path) for serving local Tilezen MBTiles database.")

	// https://github.com/sfomuseum/go-http-tilezen

	fs.Bool("enable-proxy-tiles", false, "Enable the use of a local tile proxy for Nextzen map tiles.")

	fs.String("path-proxy-tiles", "/tiles/", "The URL (a relative path) for proxied tiles.")
	fs.String("proxy-tiles-cache-uri", "gocache://", "A valid tile proxy DSN string.")
	fs.Int("proxy-tiles-timeout", 30, "The maximum number of seconds to allow for fetching a tile from the proxy.")
	fs.Bool("proxy-tiles-test", false, "Ensure outbound network connectivity for proxy tiles")

	return nil
}

// because eventually we're going to put this in a go-http-leaflet specific package
// (20200408/thisisaaronland)

func AppendLeafletFlags(fs *flag.FlagSet) error {

	fs.Float64("initial-latitude", 37.61799, "A valid latitude for the map's initial view.")
	fs.Float64("initial-longitude", -122.370943, "A valid longitude for the map's initial view.")
	fs.Int("initial-zoom", 15, "A valid zoom level for the map's initial view.")

	return nil
}
