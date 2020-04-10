package flags

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-tangramjs"
	"log"
	"os"
	"strings"
)

func Parse(fs *flag.FlagSet) {

	args := os.Args[1:]

	if len(args) > 0 && args[0] == "-h" {
		fs.Usage()
		os.Exit(0)
	}

	if len(args) > 0 && args[0] == "-setenv" {
		SetFromEnv(fs)
	}

	fs.Parse(args)
}

func SetFromEnv(fs *flag.FlagSet) {

	fs.VisitAll(func(fl *flag.Flag) {

		name := fl.Name
		env := name

		env = strings.ToUpper(env)
		env = strings.Replace(env, "-", "_", -1)
		env = fmt.Sprintf("GEOTAG_%s", env)

		val, ok := os.LookupEnv(env)

		if ok {
			log.Printf("set -%s flag (%s) from %s environment variable\n", name, val, env)
			fs.Set(name, val)
		}

	})
}

func NewFlagSet(name string) *flag.FlagSet {

	fs := flag.NewFlagSet(name, flag.ExitOnError)

	fs.Usage = func() {
		fs.PrintDefaults()
	}

	return fs
}

func CommonFlags() (*flag.FlagSet, error) {

	fs := NewFlagSet("common")

	fs.String("scheme", "http", "...")
	fs.String("host", "localhost", "...")
	fs.Int("port", 8080, "...")

	fs.String("path-templates", "", "...")

	fs.Bool("enable-placeholder", false, "...")
	fs.String("placeholder-endpoint", "", "...")

	fs.Bool("enable-oembed", false, "...")
	fs.String("oembed-endpoints", "", "...")

	fs.Bool("enable-writer", false, "...")
	fs.String("writer-uri", "stdout://", "...")

	fs.String("path-editor", "/", "...")
	fs.String("path-writer", "/update", "...")

	err := AppendLeafletFlags(fs)

	if err != nil {
		return nil, err
	}

	err = AppendNextzenFlags(fs)

	if err != nil {
		return nil, err
	}

	err = AppendTilezenFlags(fs)

	if err != nil {
		return nil, err
	}

	return fs, nil
}

// because eventually we're going to put this in a nextzen specific package
// (20200408/thisisaaronland)

func AppendNextzenFlags(fs *flag.FlagSet) error {

	fs.String("nextzen-apikey", "", "A valid Nextzen API key")
	fs.String("nextzen-style-url", "/tangram/refill-style.zip", "...")
	fs.String("nextzen-tile-url", tangramjs.NEXTZEN_MVT_ENDPOINT, "...")

	return nil
}

// because eventually we're going to put this in a sfomuseum/go-http-tilezen specific package
// (20200410/thisisaaronland)

func AppendTilezenFlags(fs *flag.FlagSet) error {

	fs.Bool("enable-proxy-tiles", false, "...")

	fs.String("path-proxy-tiles", "/tiles/", "The URL (a relative path) for proxied tiles.")
	fs.String("proxy-tiles-cache-uri", "gocache://", "A valid tile proxy DSN string.")
	fs.Int("proxy-tiles-timeout", 30, "The maximum number of seconds to allow for fetching a tile from the proxy.")
	fs.Bool("proxy-tiles-test", false, "Ensure outbound network connectivity for proxy tiles")

	return nil
}

// because eventually we're going to put this in a go-http-leaflet specific package
// (20200408/thisisaaronland)

func AppendLeafletFlags(fs *flag.FlagSet) error {

	fs.Float64("initial-latitude", 37.61799, "...")
	fs.Float64("initial-longitude", -122.370943, "...")
	fs.Int("initial-zoom", 14, "...")

	return nil
}
