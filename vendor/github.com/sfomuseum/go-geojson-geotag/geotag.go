package geotag

import (
	"bytes"
	"encoding/json"
	"io"
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

type GeotagPolygon struct {
	Type        string             `json:"type"`
	Coordinates [][]GeotagCoordinate `json:"coordinates"`
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

type FieldOfViewFeature struct {	// this is mostly for testing
	Type       string                   `json:"type"`
	Geometry   *GeotagPolygon `json:"geometry"`
	Properties GeotagProperties         `json:"properties"`
}

func NewGeotagFeature(body []byte) (*GeotagFeature, error) {

	br := bytes.NewReader(body)
	return NewGeotagFeatureWithReader(br)
}

func NewGeotagFeatureWithReader(fh io.Reader) (*GeotagFeature, error) {

	var f *GeotagFeature

	decoder := json.NewDecoder(fh)
	err := decoder.Decode(&f)

	if err != nil {
		return nil, err
	}

	return f, err
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

func (f *GeotagFeature) HorizonLine() (*GeotagLineString, error) {

	raw := f.Geometry.Geometries[1]
	enc, err := json.Marshal(raw)

	if err != nil {
		return nil, err
	}

	var ls *GeotagLineString
	err = json.Unmarshal(enc, &ls)

	if err != nil {
		return nil, err
	}

	return ls, nil
}

func (f *GeotagFeature) FieldOfView() (*GeotagPolygon, error) {

	pov, err := f.PointOfView()

	if err != nil {
		return nil, err
	}

	hl, err := f.HorizonLine()

	if err != nil {
		return nil, err
	}

	coords := make([]GeotagCoordinate, 0)

	coords = append(coords, pov.Coordinates)
	coords = append(coords, hl.Coordinates[1])	// right-hand rule
	coords = append(coords, hl.Coordinates[0])
	coords = append(coords, pov.Coordinates)

	rings := [][]GeotagCoordinate{ coords }
	
	poly := &GeotagPolygon{
		Type:        "Polygon",
		Coordinates: rings,
	}

	return poly, nil
}
