package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-www-geotag/flags"
	"net/url"
)

func NewServerWithFlagSet(ctx context.Context, fs *flag.FlagSet) (server.Server, error) {

	scheme, _ := flags.StringVar(fs, "scheme")
	host, _ := flags.StringVar(fs, "host")
	port, _ := flags.IntVar(fs, "port")

	address := fmt.Sprintf("%s://%s:%d", scheme, host, port)

	u, err := url.Parse(address)

	if err != nil {
		return nil, err
	}

	return server.NewStaticServer(u.Scheme, u)
}
