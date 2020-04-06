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
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg geotag -o assets.go static static/javascript static/css static/images

debug:
	go run -mod vendor cmd/example/main.go -templates 'templates/html/*.html'

js:
	curl -s -o static/javascript/Leaflet.GeotagPhoto.Camera.js https://raw.githubusercontent.com/sfomuseum/Leaflet.GeotagPhoto/master/src/Leaflet.GeotagPhoto.Camera.js
	curl -s -o static/javascript/Leaflet.GeotagPhoto.CameraControl.js https://raw.githubusercontent.com/sfomuseum/Leaflet.GeotagPhoto/master/src/Leaflet.GeotagPhoto.CameraControl.js
	curl -s -o static/javascript/Leaflet.GeotagPhoto.Crosshair.js https://raw.githubusercontent.com/sfomuseum/Leaflet.GeotagPhoto/master/src/Leaflet.GeotagPhoto.Crosshair.js

css:
	curl -s -o static/css/Leaflet.GeotagPhoto.css https://raw.githubusercontent.com/sfomuseum/Leaflet.GeotagPhoto/master/css/Leaflet.GeotagPhoto.css

images:
	curl -s -o static/images/camera-icon.svg https://raw.githubusercontent.com/sfomuseum/Leaflet.GeotagPhoto/master/images/camera-icon.svg
	curl -s -o static/images/camera.svg https://raw.githubusercontent.com/sfomuseum/Leaflet.GeotagPhoto/master/images/camera.svg
	curl -s -o static/images/crosshair-icon.svg https://raw.githubusercontent.com/sfomuseum/Leaflet.GeotagPhoto/master/images/crosshair-icon.svg
	curl -s -o static/images/crosshair.svg https://raw.githubusercontent.com/sfomuseum/Leaflet.GeotagPhoto/master/images/crosshair.svg
	curl -s -o static/images/marker.svg https://raw.githubusercontent.com/sfomuseum/Leaflet.GeotagPhoto/master/images/marker.svg
