package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/sfomuseum/go-http-leaflet-geotag"
	"github.com/sfomuseum/go-www-geotag/assets/templates"
	"github.com/sfomuseum/go-www-geotag/www"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func main() {

	scheme := flag.String("scheme", "http", "...")
	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")

	path_templates := flag.String("path-templates", "", "...")

	nextzen_apikey := flag.String("nextzen-apikey", "", "A valid Nextzen API key")
	nextzen_style_url := flag.String("nextzen-style-url", "/tangram/refill-style.zip", "...")
	nextzen_tile_url := flag.String("nextzen-tile-url", tangramjs.NEXTZEN_MVT_ENDPOINT, "...")

	initial_latitude := flag.Float64("initial-latitude", 37.61799, "...")
	initial_longitude := flag.Float64("initial-longitude", -122.370943, "...")
	initial_zoom := flag.Int("initial-zoom", 14, "...")

	enable_search := flag.Bool("enable-search", false, "...")
	search_endpoint := flag.String("search-endpoint", "", "...")

	enable_oembed := flag.Bool("enable-oembed", false, "...")

	// valid oembed endpoints here...
	
	flag.Parse()

	t := template.New("geotag").Funcs(template.FuncMap{
		//
	})

	if *path_templates != "" {

		parsed, err := t.ParseGlob(*path_templates)

		if err != nil {
			log.Fatalf("Unable to parse templates, %v\n", err)
		}

		t = parsed

	} else {

		for _, name := range templates.AssetNames() {

			body, err := templates.Asset(name)

			if err != nil {
				log.Fatalf("Unable to load template '%s', %v\n", name, err)
			}

			t, err = t.Parse(string(body))

			if err != nil {
				log.Fatalf("Unable to parse template '%s', %v\n", name, err)
			}
		}
	}

	mux := http.NewServeMux()

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	tangramjs_opts := tangramjs.DefaultTangramJSOptions()
	tangramjs_opts.Nextzen.APIKey = *nextzen_apikey
	tangramjs_opts.Nextzen.StyleURL = *nextzen_style_url
	tangramjs_opts.Nextzen.TileURL = *nextzen_tile_url

	geotag_opts := geotag.DefaultLeafletGeotagOptions()
	geotag.INCLUDE_LEAFLET = false // because the tangramjs stuff will add it

	err := tangramjs.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append tangram.js assets, %v", err)
	}

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append bootstrap assets, %v", err)
	}

	err = geotag.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append leaflet-geotag assets, %v", err)
	}

	err = www.AppendStaticAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append static assets, %v", err)
	}

	index_opts := &www.IndexHandlerOptions{
		Templates:        t,
		InitialLatitude:  *initial_latitude,
		InitialLongitude: *initial_longitude,
		InitialZoom:      *initial_zoom,
		EnableSearch: *enable_search,
		SearchEndpoint: *search_endpoint,
		EnableOEmbed: *enable_oembed,		
	}

	index_handler, err := www.IndexHandler(index_opts)

	if err != nil {
		log.Fatalf("failed to create index handler because %s", err)
	}

	index_handler = bootstrap.AppendResourcesHandler(index_handler, bootstrap_opts)
	index_handler = tangramjs.AppendResourcesHandler(index_handler, tangramjs_opts)
	index_handler = geotag.AppendResourcesHandler(index_handler, geotag_opts)

	mux.Handle("/", index_handler)

	address := fmt.Sprintf("%s://%s:%d", *scheme, *host, *port)

	u, err := url.Parse(address)

	if err != nil {
		log.Fatalf("Failed to parse address '%s', %v", address, err)
	}

	s, err := server.NewStaticServer(u.Scheme, u)

	if err != nil {
		log.Fatalf("Failed to create new server for '%s', %v", u, err)
	}

	log.Printf("Listening on %s", s.Address())

	err = s.ListenAndServe(mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

}
