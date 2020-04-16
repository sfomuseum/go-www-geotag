package app

import (
	"context"
	"flag"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-flags"
)

func NewServer(ctx context.Context, fs *flag.FlagSet) (server.Server, error) {

	server_uri, err := flags.StringVar(fs, "server-uri")

	if err != nil {
		return nil, err
	}

	return server.NewServer(ctx, server_uri)
}
