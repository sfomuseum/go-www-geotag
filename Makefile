cli:
	go build -mod vendor -o bin/server cmd/server/main.go

debug:
	go run -mod vendor cmd/server/main.go -nextzen-apikey $(APIKEY) -enable-placeholder -placeholder-endpoint $(SEARCH) -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}&format=json' -enable-writer

debug-protomaps:
	go run -mod vendor cmd/server/main.go -map-renderer protomaps -protomaps-tile-url $(TILES) -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}&format=json' -enable-writer -writer-uri exif://
