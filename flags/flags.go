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

func ValidateCommonFlags(fs *flag.FlagSet) error {
	return nil
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

	fs.Bool("enable-search", false, "...")
	fs.String("search-endpoint", "", "...")

	fs.Bool("enable-oembed", false, "...")
	fs.String("oembed-endpoints", "", "...")

	fs.Bool("open-browser", false, "...")

	err := AppendNextzenFlags(fs)

	if err != nil {
		return nil, err
	}

	err = AppendLeafletFlags(fs)

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

// because eventually we're going to put this in a go-http-leaflet specific package
// (20200408/thisisaaronland)

func AppendLeafletFlags(fs *flag.FlagSet) error {

	fs.Float64("initial-latitude", 37.61799, "...")
	fs.Float64("initial-longitude", -122.370943, "...")
	fs.Int("initial-zoom", 14, "...")

	return nil
}
