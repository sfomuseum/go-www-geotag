package main

import (
	"context"
	"github.com/sfomuseum/go-www-geotag/app"
	"github.com/sfomuseum/go-www-geotag/flags"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()

	fl, err := flags.CommonFlags()

	if err != nil {
		log.Fatalf("Failed to instantiate common flags, %v", err)
	}

	flags.Parse(fl)

	err = flags.ValidateCommonFlags(fl)

	if err != nil {
		log.Fatalf("Failed to validate common flags, %v", err)
	}

	mux := http.NewServeMux()

	err = app.AppendAssetHandlersWithFlagSet(ctx, fl, mux)

	if err != nil {
		log.Fatalf("Failed to append asset handlers, %v", err)
	}

	err = app.AppendApplicationHandler(ctx, fl, mux, "/")

	if err != nil {
		log.Fatalf("Failed to append application handlers, %v", err)
	}

	s, err := app.NewServerWithFlagSet(ctx, fl)

	if err != nil {
		log.Fatalf("Failed to create application server, %v", err)
	}

	log.Printf("Listening on %s", s.Address())

	err = s.ListenAndServe(mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
