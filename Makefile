CWD=$(shell pwd)

go-bindata:
	mkdir -p cmd/go-bindata
	mkdir -p cmd/go-bindata-assetfs
	curl -s -o cmd/go-bindata/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata/master/cmd/go-bindata/main.go
	curl -s -o cmd/go-bindata-assetfs/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata-assetfs/master/cmd/go-bindata-assetfs/main.go

bake:
	@make bake-templates
	@make bake-assets

bake-templates:
	go build -o bin/go-bindata cmd/go-bindata/main.go
	rm -rf templates/html/*~
	bin/go-bindata -pkg templates -o assets/templates/html.go templates/html

bake-assets:	
	go build -o bin/go-bindata cmd/go-bindata/main.go
	go build -o bin/go-bindata-assetfs cmd/go-bindata-assetfs/main.go
	rm -f static/*~ static/css/*~ static/images/*~ static/javascript/*~
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg www -o www/assetsfs.go static static/javascript static/css static/images

debug:
	@make bake
	go run -mod vendor cmd/server/main.go -nextzen-apikey $(APIKEY) -enable-search -search-endpoint $(SEARCH) -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/' -port 8081
