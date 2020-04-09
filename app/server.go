package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-www-geotag/flags"
	"net/url"
)

func NewServer(ctx context.Context, fs *flag.FlagSet) (server.Server, error) {

	scheme, err := flags.StringVar(fs, "scheme")

	if err != nil {
		return nil, err
	}

	host, err := flags.StringVar(fs, "host")

	if err != nil {
		return nil, err
	}

	port, err := flags.IntVar(fs, "port")

	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s://%s:%d", scheme, host, port)

	u, err := url.Parse(address)

	if err != nil {
		return nil, err
	}

	return server.NewStaticServer(u.Scheme, u)
}
