CWD=$(shell pwd)

go-bindata:
	mkdir -p cmd/go-bindata
	mkdir -p cmd/go-bindata-assetfs
	curl -s -o cmd/go-bindata/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata/master/cmd/go-bindata/main.go
	curl -s -o cmd/go-bindata-assetfs/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata-assetfs/master/cmd/go-bindata-assetfs/main.go

src:
	curl -s -o static/javascript/leaflet.layers.control.js https://raw.githubusercontent.com/sfomuseum/leaflet-layers-control/master/src/leaflet.layers.control.js
	curl -s -o static/css/leaflet.layers.control.css https://raw.githubusercontent.com/sfomuseum/leaflet-layers-control/master/src/leaflet.layers.control.css

bake:
	# @make bake-templates
	@make bake-assets

bake-templates:
	go build -o bin/go-bindata cmd/go-bindata/main.go
	rm -rf templates/html/*~
	bin/go-bindata -pkg templates -o assets/templates/html.go templates/html

bake-assets:	
	go build -o bin/go-bindata cmd/go-bindata/main.go
	go build -o bin/go-bindata-assetfs cmd/go-bindata-assetfs/main.go
	rm -f static/*~ static/css/*~ static/javascript/*~
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg layers -o assets.go static static/javascript static/css

debug:
	go run -mod vendor cmd/example/main.go -templates 'templates/html/*.html'
