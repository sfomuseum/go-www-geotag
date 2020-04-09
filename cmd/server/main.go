package main

import (
	"context"
	"fmt"
	"github.com/pkg/browser"
	"github.com/sfomuseum/go-www-geotag/app"
	"github.com/sfomuseum/go-www-geotag/flags"
	"log"
	"net/http"
	"net/url"
	"time"
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

	err := app.AppendApplicationHandler(ctx, fl, mux, "/")

	if err != nil {
		log.Fatalf("Failed to append application handlers, %v", err)
	}

	address := fmt.Sprintf("%s://%s:%d", *scheme, *host, *port)

	u, err := url.Parse(address)

	if err != nil {
		log.Fatalf("Failed to parse address '%s', %v", address, err)
	}

	s, err := server.NewStaticServer(u.Scheme, u)

	if err != nil {
		log.Fatalf("Failed to create new server for '%s', %v", u, err)
	}

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {

		log.Printf("Listening on %s", s.Address())

		err = s.ListenAndServe(mux)

		if err != nil {
			log.Fatalf("Failed to start server, %v", err)
		}
	}()

	if *open_browser {

		ticker := time.NewTicker(200 * time.Millisecond)

		available := false
		attempts := 0

		for {
			select {
			case <-ticker.C:

				rsp, err := http.Get(s.Address())

				if err == nil && rsp.StatusCode == 200 {
					available = true
					break
				}

				attempts += 1

			default:
				// pass
			}

			if available == true {
				break
			}

			if attempts >= 10 {
				log.Fatalf("Failed to determine whether %s is available", s.Address())
			}
		}

		err := browser.OpenURL(s.Address())

		if err != nil {
			log.Fatalf("Failed to open URL %s, %v", s.Address(), err)
		}
	}

	<-ctx.Done()
}
