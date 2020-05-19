package writer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/sfomuseum/go-geojson-geotag"
	"io"
)

const IOWRITER_TARGET_KEY string = "github.com/sfomuseum/go-geojson-geotag#io_writer"

type IOWriter struct {
	Writer
}

func init() {

	ctx := context.Background()
	err := RegisterWriter(ctx, "io", NewIOWriter)

	if err != nil {
		panic(err)
	}
}

func NewIOWriter(ctx context.Context, uri string) (Writer, error) {

	wr := &IOWriter{}
	return wr, nil
}

func (wr *IOWriter) WriteFeature(ctx context.Context, uri string, f *geotag.GeotagFeature) error {

	target, err := GetIOWriterFromContext(ctx)

	if err != nil {
		return err
	}

	body, err := json.Marshal(f)

	if err != nil {
		return err
	}

	br := bytes.NewReader(body)

	_, err = io.Copy(target, br)
	return err
}

func (wr *IOWriter) Close(ctx context.Context) error {
	return nil
}

func SetIOWriterWithContext(ctx context.Context, wr io.Writer) (context.Context, error) {

	ctx = context.WithValue(ctx, IOWRITER_TARGET_KEY, wr)
	return ctx, nil
}

func GetIOWriterFromContext(ctx context.Context) (io.Writer, error) {

	v := ctx.Value(IOWRITER_TARGET_KEY)

	if v == nil {
		return nil, errors.New("Missing writer")
	}

	var target io.Writer

	switch v.(type) {
	case io.Writer:
		target = v.(io.Writer)
	default:
		return nil, errors.New("Invalid writer")
	}

	return target, nil
}
