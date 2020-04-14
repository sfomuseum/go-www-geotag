package api

import (
	"github.com/sfomuseum/go-geojson-geotag"
	"github.com/sfomuseum/go-www-geotag/writer"
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

		geotag_f, err := geotag.NewGeotagFeatureWithReader(req.Body)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := req.Context()

		uri := geotag_f.Id

		if uri == "" {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		err = wr.WriteFeature(ctx, uri, geotag_f)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
