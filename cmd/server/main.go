// server is an HTTP server that runs the go-www-geotag application.
package main

import (
	"context"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-www-geotag/app"
	_ "github.com/whosonfirst/go-whosonfirst-spatial-sqlite"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()

	fl, err := app.CommonFlags()

	if err != nil {
		log.Fatalf("Failed to instantiate common flags, %v", err)
	}

	flagset.Parse(fl)

	err = flagset.SetFlagsFromEnvVars(fl, "GEOTAG")

	if err != nil {
		log.Fatalf("Failed to set flags from env vars, %v", err)
	}

	mux := http.NewServeMux()

	err = app.AppendAssetHandlers(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append asset handlers, %v", err)
	}

	// The core "geotagging" application

	err = app.AppendEditorHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append editor handler, %v", err)
	}

	// Support for point-in-polygon lookups (for things that have been geotagged)

	err = app.AppendPointInPolygonHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append point-in-polygon handler, %v", err)
	}

	// Local Tilezen proxy caching layer

	err = app.AppendProxyTilesHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append proxy tiles handler, %v", err)
	}

	// Local Tilezen "tilepack" tile-serving layer

	err = app.AppendTilezenTilepackHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append proxy tiles handler, %v", err)
	}

	// Local Protomaps tile-serving layer

	err = app.AppendProtomapsTilesHandlerIfNecessary(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append protomaps tiles handler, %v", err)
	}

	// Abstract writer interface(s)

	err = app.AppendWriterHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append writer handler, %v", err)
	}

	// Public-facing server

	s, err := app.NewServer(ctx, fl)

	if err != nil {
		log.Fatalf("Failed to create application server, %v", err)
	}

	log.Printf("Listening on %s", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
