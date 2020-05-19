package api

import (
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-geojson-geotag"
	"github.com/sfomuseum/go-www-geotag/writer"
	_ "log"
	"net/http"
)

func WriterHandler(wr writer.Writer) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case "PUT":
			// pass
		default:
			http.Error(rsp, "Method not allowed.", http.StatusMethodNotAllowed)
			return
		}

		defer req.Body.Close()

		uid, err := sanitize.GetString(req, "id")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		geotag_f, err := geotag.NewGeotagFeatureWithReader(req.Body)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")

		ctx := req.Context()

		ctx, err = writer.SetIOWriterWithContext(ctx, rsp)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		err = wr.WriteFeature(ctx, uid, geotag_f)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		err = wr.Close(ctx)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
