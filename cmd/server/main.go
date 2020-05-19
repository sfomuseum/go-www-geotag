package main

import (
	"context"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-www-geotag/app"
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

	err = app.AppendEditorHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append editor handler, %v", err)
	}

	err = app.AppendProxyTilesHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append proxy tiles handler, %v", err)
	}

	err = app.AppendWriterHandlerIfEnabled(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append writer handler, %v", err)
	}

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
