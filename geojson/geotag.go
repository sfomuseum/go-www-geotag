package geojson

import (
	"encoding/json"
)

type GeotagProperties struct {
	Angle    float64 `json:"angle"`
	Bearing  float64 `json:"bearing"`
	Distance float64 `json:"distance"`
}

type GeotagCoordinate [2]float64

type GeotagPoint struct {
	Type        string           `json:"type"`
	Coordinates GeotagCoordinate `json:"coordinates"`
}

type GeotagLineString struct {
	Type        string              `json:"type"`
	Coordinates [2]GeotagCoordinate `json:"coordinates"`
}

type GeotagGeometryCollection struct {
	Type       string         `json:"type"`
	Geometries [2]interface{} `json:"geometries"`
}

type GeotagFeature struct {
	Type       string                   `json:"type"`
	Geometry   GeotagGeometryCollection `json:"geometry"`
	Properties GeotagProperties         `json:"properties"`
}

func (f *GeotagFeature) PointOfView() (*GeotagPoint, error) {

	// if there is a better way to do this I wish I
	// knew what it was... (20200410/thisisaaronland)
	
	raw := f.Geometry.Geometries[0]
	enc, err := json.Marshal(raw)

	if err != nil {
		return nil, err
	}

	var pov *GeotagPoint
	err = json.Unmarshal(enc, &pov)

	if err != nil {
		return nil, err
	}
	
	return pov, nil
}
