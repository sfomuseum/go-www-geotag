package writer

import (
	"context"
	"github.com/aaronland/go-roster"
	"github.com/sfomuseum/go-www-geotag/geojson"
	"net/url"
)

type Writer interface {
	WriteFeature(context.Context, string, *geojson.GeotagFeature) error
}

type WriterInitializeFunc func(ctx context.Context, uri string) (Writer, error)

var writers roster.Roster

func ensureRoster() error {

	if writers == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		writers = r
	}

	return nil
}

func RegisterWriter(ctx context.Context, scheme string, f WriterInitializeFunc) error {

	err := ensureRoster()

	if err != nil {
		return err
	}

	return writers.Register(ctx, scheme, f)
}

func NewWriter(ctx context.Context, uri string) (Writer, error) {

	err := ensureRoster()

	if err != nil {
		return nil, err
	}

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := writers.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	f := i.(WriterInitializeFunc)
	return f(ctx, uri)
}
