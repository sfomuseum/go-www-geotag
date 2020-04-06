package www

import (
	"errors"
	"html/template"
	"net/http"
)

type IndexHandlerOptions struct {
	Templates        *template.Template
	InitialLatitude  float64
	InitialLongitude float64
	InitialZoom      int
}

type IndexHandlerVars struct {
	InitialLatitude  float64
	InitialLongitude float64
	InitialZoom      int
}

func IndexHandler(opts *IndexHandlerOptions) (http.Handler, error) {

	t := opts.Templates.Lookup("index_2")

	if t == nil {
		return nil, errors.New("Missing 'index' template")
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		vars := IndexHandlerVars{
			InitialLatitude:  opts.InitialLatitude,
			InitialLongitude: opts.InitialLongitude,
			InitialZoom:      opts.InitialZoom,
		}

		err := t.Execute(rsp, vars)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
