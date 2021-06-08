cli:
	go build -mod vendor -o bin/server cmd/server/main.go

docker:
	docker build -t geotag-www .	

debug:
	go run -mod vendor cmd/server/main.go -nextzen-apikey $(APIKEY) -enable-placeholder -placeholder-endpoint $(SEARCH) -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}&format=json' -enable-writer

debug-tilepack:
	go run -mod vendor cmd/server/main.go -enable-tilezen-tilepacks -tilezen-path-tilepack $(TILEPACK) -nextzen-tile-url '/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt' -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}&format=json' -enable-exif-writer -enable-oembed-cors-image -prefix /demo/geotag

debug-protomaps:
	go run -mod vendor cmd/server/main.go -map-renderer protomaps -protomaps-tile-url $(TILES) -enable-oembed -oembed-endpoints 'https://millsfield.sfomuseum.org/oembed/?url={url}&format=json' -enable-exif-writer -enable-oembed-cors-image

debug-docker:
	docker run -it -p 8080:8080 -e GEOTAG_ENABLE_POINT_IN_POLYGON=true -e GEOTAG_MAP_RENDERER=tangramjs -e GEOTAG_ENABLE_TILEZEN_TILEPACKS=true -e GEOTAG_TILEZEN_PATH_TILEPACK=/usr/local/data/sfo-tiles.db -e GEOTAG_SERVER_URI=http://0.0.0.0:8080 -e GEOTAG_ENABLE_EXIF_WRITER=true -e GEOTAG_SPATIAL_DATABASE_URI='sqlite://?dsn=/usr/local/data/sfomuseum-architecture.db' -e GEOTAG_NEXTZEN_TILE_URL='/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt' -e GEOTAG_PREFIX=/demo/geotag geotag-www 
