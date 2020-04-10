package www

import (
	"encoding/json"
	"github.com/sfomuseum/go-www-geotag/geojson"
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

		var geotag_f *geojson.GeotagFeature

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&geotag_f)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := req.Context()

		uri := "fixme"
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