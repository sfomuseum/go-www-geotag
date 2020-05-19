package writer

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sfomuseum/go-geojson-geotag"
	"io"
	"os"
)

func init() {
	ctx := context.Background()
	RegisterWriter(ctx, "stdout", NewStdoutWriter)
}

type StdoutWriter struct {
	Writer
}

func NewStdoutWriter(ctx context.Context, uri string) (Writer, error) {

	wr := &StdoutWriter{}
	return wr, nil
}

func (wr *StdoutWriter) WriteFeature(ctx context.Context, uri string, f *geotag.GeotagFeature) error {

	body, err := json.Marshal(f)

	if err != nil {
		return err
	}

	br := bytes.NewReader(body)
	_, err = io.Copy(os.Stdout, br)

	if err != nil {
		return err
	}

	return nil
}

func (wr *StdoutWriter) Close(ctx context.Context) error {
	return nil
}
