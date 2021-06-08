(function (L) {
    
    'use strict';

    L = 'default' in L ? L['default'] : L;

    // The source data for each of these SVG elements is in static/images.
    // We're encoding them as data URIs so that we don't need to worry about
    // paths for web applications that might be run as stand-alone services
    // but "fronted" by something like CloudFront and exposed as a "leaf" of
    // another domain. Basically, using paths or URIs to reference these elements
    // results in headaches so we just "bake" all our assets in to the JavaScript
    // proper. (20210608/thisisaaronland)
    
    var crosshair_icon_svg = 'data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4NCjwhLS0gR2VuZXJhdG9yOiBBZG9iZSBJbGx1c3RyYXRvciAxOS4xLjAsIFNWRyBFeHBvcnQgUGx1Zy1JbiAuIFNWRyBWZXJzaW9uOiA2LjAwIEJ1aWxkIDApICAtLT4NCjxzdmcgdmVyc2lvbj0iMS4xIiBpZD0iTGF5ZXJfMSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeD0iMHB4IiB5PSIwcHgiDQoJIHZpZXdCb3g9IjAgMCAyNTYgMjU2IiBzdHlsZT0iZW5hYmxlLWJhY2tncm91bmQ6bmV3IDAgMCAyNTYgMjU2OyIgeG1sOnNwYWNlPSJwcmVzZXJ2ZSI+DQo8c3R5bGUgdHlwZT0idGV4dC9jc3MiPg0KCS5zdDB7ZmlsbDojN0RBN0M2O2ZpbGwtb3BhY2l0eTowLjQ7c3Ryb2tlOiMwMDAwMDA7c3Ryb2tlLXdpZHRoOjE2O3N0cm9rZS1taXRlcmxpbWl0OjEwO30NCjwvc3R5bGU+DQo8Y2lyY2xlIGNsYXNzPSJzdDAiIGN4PSIxMjgiIGN5PSIxMjgiIHI9IjExNCIvPg0KPGNpcmNsZSBjeD0iMTI4IiBjeT0iMTI4IiByPSIyMi4zIi8+DQo8L3N2Zz4NCg==';
    
    var camera_icon_svg = 'data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4NCjwhLS0gR2VuZXJhdG9yOiBBZG9iZSBJbGx1c3RyYXRvciAxOS4xLjAsIFNWRyBFeHBvcnQgUGx1Zy1JbiAuIFNWRyBWZXJzaW9uOiA2LjAwIEJ1aWxkIDApICAtLT4NCjxzdmcgdmVyc2lvbj0iMS4xIiBpZD0iTGF5ZXJfMSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeD0iMHB4IiB5PSIwcHgiDQoJIHZpZXdCb3g9IjAgMCAyNTYgMjU2IiBzdHlsZT0iZW5hYmxlLWJhY2tncm91bmQ6bmV3IDAgMCAyNTYgMjU2OyIgeG1sOnNwYWNlPSJwcmVzZXJ2ZSI+DQo8c3R5bGUgdHlwZT0idGV4dC9jc3MiPg0KCS5zdDB7b3BhY2l0eTowLjQ7ZmlsbDojN0RBN0M2O3N0cm9rZTojODNDMUZGO3N0cm9rZS1taXRlcmxpbWl0OjEwO30NCgkuc3Qxe2ZpbGw6bm9uZTtzdHJva2U6IzZENkQ2RDtzdHJva2Utd2lkdGg6MjA7c3Ryb2tlLWxpbmVjYXA6cm91bmQ7c3Ryb2tlLWxpbmVqb2luOnJvdW5kO3N0cm9rZS1taXRlcmxpbWl0OjEwO30NCgkuc3Qye2ZpbGw6IzFBMUExQTt9DQo8L3N0eWxlPg0KPGc+DQoJPGc+DQoJCTxwb2x5Z29uIGNsYXNzPSJzdDAiIHBvaW50cz0iMjM5LDEyNS42IDY0LjYsMTkxLjQgMTMwLjQsMTcgCQkiLz4NCgkJPHBvbHlsaW5lIGNsYXNzPSJzdDEiIHBvaW50cz0iMjM5LDEyNS42IDY0LjYsMTkxLjQgMTMwLjQsMTcgCQkiLz4NCgk8L2c+DQoJPHJlY3QgeD0iMjAuNiIgeT0iMTU1LjkiIHRyYW5zZm9ybT0ibWF0cml4KDAuNzA3MSAwLjcwNzEgLTAuNzA3MSAwLjcwNzEgMTU0LjI2NzQgMTAuMzk1OCkiIGNsYXNzPSJzdDIiIHdpZHRoPSI4OCIgaGVpZ2h0PSI3MSIvPg0KPC9nPg0KPC9zdmc+DQo=';

    var crosshair_svg = 'data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4NCjwhLS0gR2VuZXJhdG9yOiBBZG9iZSBJbGx1c3RyYXRvciAxOS4xLjAsIFNWRyBFeHBvcnQgUGx1Zy1JbiAuIFNWRyBWZXJzaW9uOiA2LjAwIEJ1aWxkIDApICAtLT4NCjxzdmcgdmVyc2lvbj0iMS4xIiBpZD0iTGF5ZXJfMSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeD0iMHB4IiB5PSIwcHgiDQoJIHZpZXdCb3g9Ii0yNTUgMzQ3IDEwMCAxMDAiIHN0eWxlPSJlbmFibGUtYmFja2dyb3VuZDpuZXcgLTI1NSAzNDcgMTAwIDEwMDsiIHhtbDpzcGFjZT0icHJlc2VydmUiPg0KPHN0eWxlIHR5cGU9InRleHQvY3NzIj4NCgkuc3Qwe2ZpbGw6bm9uZTtzdHJva2U6IzAwMDAwMDtzdHJva2Utd2lkdGg6MztzdHJva2UtbGluZWNhcDpyb3VuZDtzdHJva2UtbGluZWpvaW46cm91bmQ7c3Ryb2tlLW1pdGVybGltaXQ6MTA7fQ0KPC9zdHlsZT4NCjxjaXJjbGUgY2xhc3M9InN0MCIgY3g9Ii0yMDUiIGN5PSIzOTciIHI9IjQ3Ii8+DQo8Y2lyY2xlIGN4PSItMjA1IiBjeT0iMzk3IiByPSI0LjgiLz4NCjwvc3ZnPg0K';

    var camera_svg = 'data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4NCjwhLS0gR2VuZXJhdG9yOiBBZG9iZSBJbGx1c3RyYXRvciAxOS4xLjAsIFNWRyBFeHBvcnQgUGx1Zy1JbiAuIFNWRyBWZXJzaW9uOiA2LjAwIEJ1aWxkIDApICAtLT4NCjxzdmcgdmVyc2lvbj0iMS4xIiBpZD0iTGF5ZXJfMSINCgkgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgeD0iMHB4IiB5PSIwcHgiIHZpZXdCb3g9Ii0yNDUgMzY1IDY0IDY0Ig0KCSBzdHlsZT0iZW5hYmxlLWJhY2tncm91bmQ6bmV3IC0yNDUgMzY1IDY0IDY0OyIgeG1sOnNwYWNlPSJwcmVzZXJ2ZSI+DQo8c3R5bGUgdHlwZT0idGV4dC9jc3MiPg0KCS5zdDB7ZmlsbDojNEE0QTRBO30NCjwvc3R5bGU+DQo8cGF0aCBpZD0iUmVjdGFuZ2xlLTIiIGNsYXNzPSJzdDAiIGQ9Ik0tMjI3LjUsMzg1LjZoLTE0LjV2MjkuMWg1OC4xdi0yOS4xaC0xNC41di03LjNoLTI5LjFWMzg1LjZ6Ii8+DQo8L3N2Zz4NCg==';    

    var marker_svg = 'data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4NCjwhLS0gR2VuZXJhdG9yOiBBZG9iZSBJbGx1c3RyYXRvciAxOS4xLjAsIFNWRyBFeHBvcnQgUGx1Zy1JbiAuIFNWRyBWZXJzaW9uOiA2LjAwIEJ1aWxkIDApICAtLT4NCjxzdmcgdmVyc2lvbj0iMS4xIiBpZD0iTGF5ZXJfMSIgeG1sbnM6c2tldGNoPSJodHRwOi8vd3d3LmJvaGVtaWFuY29kaW5nLmNvbS9za2V0Y2gvbnMiDQoJIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHg9IjBweCIgeT0iMHB4IiB2aWV3Qm94PSItMjExLjIgMzcyLjkgNDguMiA0OC4yIg0KCSBzdHlsZT0iZW5hYmxlLWJhY2tncm91bmQ6bmV3IC0yMTEuMiAzNzIuOSA0OC4yIDQ4LjI7IiB4bWw6c3BhY2U9InByZXNlcnZlIj4NCjxzdHlsZSB0eXBlPSJ0ZXh0L2NzcyI+DQoJLnN0MHtmaWxsOiM0QTRBNEE7fQ0KCS5zdDF7b3BhY2l0eTowLjI7fQ0KPC9zdHlsZT4NCjxjaXJjbGUgY2xhc3M9InN0MCIgY3g9Ii0xODcuMSIgY3k9IjM5NyIgcj0iOC4zIi8+DQo8Y2lyY2xlIGNsYXNzPSJzdDEiIGN4PSItMTg3LjEiIGN5PSIzOTciIHI9IjIyIi8+DQo8L3N2Zz4NCg==';

    
var GeotagPhotoCrosshair = L.Evented.extend({
  options: {      
      controlCrosshairImg: crosshair_icon_svg,
      crosshairHTML: '<img alt="Center of the map; crosshair location" title="Crosshair" src="' + crosshair_svg + '" width="100px" />'
  },

  initialize: function initialize(options) {
    L.setOptions(this, options);
  },

  addTo: function addTo(map) {
    this._map = map;
    var container = map.getContainer();
    this._element = L.DomUtil.create('div', 'leaflet-geotag-photo-crosshair', container);
    this._element.innerHTML = this.options.crosshairHTML;

    this._boundOnInput = this._onInput.bind(this);
    this._boundOnChange = this._onChange.bind(this);

    this._map.on('move', this._boundOnInput);
    this._map.on('moveend', this._boundOnChange);

    return this;
  },

  removeFrom: function removeFrom(map) {
    if (this._map && this._boundOnInput && this._boundOnChange) {
      this._map.off('move', this._boundOnInput);
      this._map.off('moveend', this._boundOnChange);
    }

    if (this._element) {
      L.DomUtil.remove(this._element);
    }

    return this;
  },

  _onInput: function _onInput() {
    this.fire('input');
  },

  _onChange: function _onChange() {
    this.fire('change');
  },

  getCrosshairLatLng: function getCrosshairLatLng() {
    return this._map.getCenter();
  },

  getCrosshairPoint: function getCrosshairPoint() {
    if (this._map) {
      var center = this.getCrosshairLatLng();
      return {
        type: 'Point',
        coordinates: [center.lng, center.lat]
      };
    }
  }

});

/**
 * Unwrap a coordinate from a Point Feature, Geometry or a single coordinate.
 *
 * @param {Array<any>|Geometry|Feature<Point>} obj any value
 * @returns {Array<number>} coordinates
 */
function getCoord$1(obj) {
    if (!obj) throw new Error('No obj passed');

    var coordinates = getCoords(obj);

    // getCoord() must contain at least two numbers (Point)
    if (coordinates.length > 1 &&
        typeof coordinates[0] === 'number' &&
        typeof coordinates[1] === 'number') {
        return coordinates;
    } else {
        throw new Error('Coordinate is not a valid Point');
    }
}

/**
 * Unwrap coordinates from a Feature, Geometry Object or an Array of numbers
 *
 * @param {Array<any>|Geometry|Feature<any>} obj any value
 * @returns {Array<any>} coordinates
 */
function getCoords(obj) {
    if (!obj) throw new Error('No obj passed');
    var coordinates;

    // Array of numbers
    if (obj.length) {
        coordinates = obj;

    // Geometry Object
    } else if (obj.coordinates) {
        coordinates = obj.coordinates;

    // Feature
    } else if (obj.geometry && obj.geometry.coordinates) {
        coordinates = obj.geometry.coordinates;
    }
    // Checks if coordinates contains a number
    if (coordinates) {
        containsNumber(coordinates);
        return coordinates;
    }
    throw new Error('No valid coordinates');
}

/**
 * Checks if coordinates contains a number
 *
 * @private
 * @param {Array<any>} coordinates GeoJSON Coordinates
 * @returns {boolean} true if Array contains a number
 */
function containsNumber(coordinates) {
    if (coordinates.length > 1 &&
        typeof coordinates[0] === 'number' &&
        typeof coordinates[1] === 'number') {
        return true;
    }
    if (coordinates[0].length) {
        return containsNumber(coordinates[0]);
    }
    throw new Error('coordinates must only contain numbers');
}

/**
 * Enforce expectations about types of GeoJSON objects for Turf.
 *
 * @alias geojsonType
 * @param {GeoJSON} value any GeoJSON object
 * @param {string} type expected GeoJSON type
 * @param {string} name name of calling function
 * @throws {Error} if value is not the expected type.
 */
function geojsonType(value, type, name) {
    if (!type || !name) throw new Error('type and name required');

    if (!value || value.type !== type) {
        throw new Error('Invalid input to ' + name + ': must be a ' + type + ', given ' + value.type);
    }
}

/**
 * Enforce expectations about types of {@link Feature} inputs for Turf.
 * Internally this uses {@link geojsonType} to judge geometry types.
 *
 * @alias featureOf
 * @param {Feature} feature a feature with an expected geometry type
 * @param {string} type expected GeoJSON type
 * @param {string} name name of calling function
 * @throws {Error} error if value is not the expected type.
 */
function featureOf(feature, type, name) {
    if (!feature) throw new Error('No feature passed');
    if (!name) throw new Error('.featureOf() requires a name');
    if (!feature || feature.type !== 'Feature' || !feature.geometry) {
        throw new Error('Invalid input to ' + name + ', Feature with geometry required');
    }
    if (!feature.geometry || feature.geometry.type !== type) {
        throw new Error('Invalid input to ' + name + ': must be a ' + type + ', given ' + feature.geometry.type);
    }
}

/**
 * Enforce expectations about types of {@link FeatureCollection} inputs for Turf.
 * Internally this uses {@link geojsonType} to judge geometry types.
 *
 * @alias collectionOf
 * @param {FeatureCollection} featureCollection a FeatureCollection for which features will be judged
 * @param {string} type expected GeoJSON type
 * @param {string} name name of calling function
 * @throws {Error} if value is not the expected type.
 */
function collectionOf(featureCollection, type, name) {
    if (!featureCollection) throw new Error('No featureCollection passed');
    if (!name) throw new Error('.collectionOf() requires a name');
    if (!featureCollection || featureCollection.type !== 'FeatureCollection') {
        throw new Error('Invalid input to ' + name + ', FeatureCollection required');
    }
    for (var i = 0; i < featureCollection.features.length; i++) {
        var feature = featureCollection.features[i];
        if (!feature || feature.type !== 'Feature' || !feature.geometry) {
            throw new Error('Invalid input to ' + name + ', Feature with geometry required');
        }
        if (!feature.geometry || feature.geometry.type !== type) {
            throw new Error('Invalid input to ' + name + ': must be a ' + type + ', given ' + feature.geometry.type);
        }
    }
}

var geojsonType_1 = geojsonType;
var collectionOf_1 = collectionOf;
var featureOf_1 = featureOf;
var getCoord_1 = getCoord$1;
var getCoords_1 = getCoords;

var index$1 = {
	geojsonType: geojsonType_1,
	collectionOf: collectionOf_1,
	featureOf: featureOf_1,
	getCoord: getCoord_1,
	getCoords: getCoords_1
};

/**
 * Wraps a GeoJSON {@link Geometry} in a GeoJSON {@link Feature}.
 *
 * @name feature
 * @param {Geometry} geometry input geometry
 * @param {Object} properties properties
 * @returns {FeatureCollection} a FeatureCollection of input features
 * @example
 * var geometry = {
 *      "type": "Point",
 *      "coordinates": [
 *        67.5,
 *        32.84267363195431
 *      ]
 *    }
 *
 * var feature = turf.feature(geometry);
 *
 * //=feature
 */
function feature(geometry, properties) {
    if (!geometry) throw new Error('No geometry passed');

    return {
        type: 'Feature',
        properties: properties || {},
        geometry: geometry
    };
}
var feature_1 = feature;

/**
 * Takes coordinates and properties (optional) and returns a new {@link Point} feature.
 *
 * @name point
 * @param {Array<number>} coordinates longitude, latitude position (each in decimal degrees)
 * @param {Object=} properties an Object that is used as the {@link Feature}'s
 * properties
 * @returns {Feature<Point>} a Point feature
 * @example
 * var pt1 = turf.point([-75.343, 39.984]);
 *
 * //=pt1
 */
var point$1 = function (coordinates, properties) {
    if (!coordinates) throw new Error('No coordinates passed');
    if (coordinates.length === undefined) throw new Error('Coordinates must be an array');
    if (coordinates.length < 2) throw new Error('Coordinates must be at least 2 numbers long');
    if (typeof coordinates[0] !== 'number' || typeof coordinates[1] !== 'number') throw new Error('Coordinates must numbers');

    return feature({
        type: 'Point',
        coordinates: coordinates
    }, properties);
};

/**
 * Takes an array of LinearRings and optionally an {@link Object} with properties and returns a {@link Polygon} feature.
 *
 * @name polygon
 * @param {Array<Array<Array<number>>>} coordinates an array of LinearRings
 * @param {Object=} properties a properties object
 * @returns {Feature<Polygon>} a Polygon feature
 * @throws {Error} throw an error if a LinearRing of the polygon has too few positions
 * or if a LinearRing of the Polygon does not have matching Positions at the
 * beginning & end.
 * @example
 * var polygon = turf.polygon([[
 *  [-2.275543, 53.464547],
 *  [-2.275543, 53.489271],
 *  [-2.215118, 53.489271],
 *  [-2.215118, 53.464547],
 *  [-2.275543, 53.464547]
 * ]], { name: 'poly1', population: 400});
 *
 * //=polygon
 */
var polygon = function (coordinates, properties) {
    if (!coordinates) throw new Error('No coordinates passed');

    for (var i = 0; i < coordinates.length; i++) {
        var ring = coordinates[i];
        if (ring.length < 4) {
            throw new Error('Each LinearRing of a Polygon must have 4 or more Positions.');
        }
        for (var j = 0; j < ring[ring.length - 1].length; j++) {
            if (ring[ring.length - 1][j] !== ring[0][j]) {
                throw new Error('First and last Position are not equivalent.');
            }
        }
    }

    return feature({
        type: 'Polygon',
        coordinates: coordinates
    }, properties);
};

/**
 * Creates a {@link LineString} based on a
 * coordinate array. Properties can be added optionally.
 *
 * @name lineString
 * @param {Array<Array<number>>} coordinates an array of Positions
 * @param {Object=} properties an Object of key-value pairs to add as properties
 * @returns {Feature<LineString>} a LineString feature
 * @throws {Error} if no coordinates are passed
 * @example
 * var linestring1 = turf.lineString([
 *   [-21.964416, 64.148203],
 *   [-21.956176, 64.141316],
 *   [-21.93901, 64.135924],
 *   [-21.927337, 64.136673]
 * ]);
 * var linestring2 = turf.lineString([
 *   [-21.929054, 64.127985],
 *   [-21.912918, 64.134726],
 *   [-21.916007, 64.141016],
 *   [-21.930084, 64.14446]
 * ], {name: 'line 1', distance: 145});
 *
 * //=linestring1
 *
 * //=linestring2
 */
var lineString = function (coordinates, properties) {
    if (!coordinates) throw new Error('No coordinates passed');

    return feature({
        type: 'LineString',
        coordinates: coordinates
    }, properties);
};

/**
 * Takes one or more {@link Feature|Features} and creates a {@link FeatureCollection}.
 *
 * @name featureCollection
 * @param {Feature[]} features input features
 * @returns {FeatureCollection} a FeatureCollection of input features
 * @example
 * var features = [
 *  turf.point([-75.343, 39.984], {name: 'Location A'}),
 *  turf.point([-75.833, 39.284], {name: 'Location B'}),
 *  turf.point([-75.534, 39.123], {name: 'Location C'})
 * ];
 *
 * var fc = turf.featureCollection(features);
 *
 * //=fc
 */
var featureCollection = function (features) {
    if (!features) throw new Error('No features passed');

    return {
        type: 'FeatureCollection',
        features: features
    };
};

/**
 * Creates a {@link Feature<MultiLineString>} based on a
 * coordinate array. Properties can be added optionally.
 *
 * @name multiLineString
 * @param {Array<Array<Array<number>>>} coordinates an array of LineStrings
 * @param {Object=} properties an Object of key-value pairs to add as properties
 * @returns {Feature<MultiLineString>} a MultiLineString feature
 * @throws {Error} if no coordinates are passed
 * @example
 * var multiLine = turf.multiLineString([[[0,0],[10,10]]]);
 *
 * //=multiLine
 *
 */
var multiLineString = function (coordinates, properties) {
    if (!coordinates) throw new Error('No coordinates passed');

    return feature({
        type: 'MultiLineString',
        coordinates: coordinates
    }, properties);
};

/**
 * Creates a {@link Feature<MultiPoint>} based on a
 * coordinate array. Properties can be added optionally.
 *
 * @name multiPoint
 * @param {Array<Array<number>>} coordinates an array of Positions
 * @param {Object=} properties an Object of key-value pairs to add as properties
 * @returns {Feature<MultiPoint>} a MultiPoint feature
 * @throws {Error} if no coordinates are passed
 * @example
 * var multiPt = turf.multiPoint([[0,0],[10,10]]);
 *
 * //=multiPt
 *
 */
var multiPoint = function (coordinates, properties) {
    if (!coordinates) throw new Error('No coordinates passed');

    return feature({
        type: 'MultiPoint',
        coordinates: coordinates
    }, properties);
};


/**
 * Creates a {@link Feature<MultiPolygon>} based on a
 * coordinate array. Properties can be added optionally.
 *
 * @name multiPolygon
 * @param {Array<Array<Array<Array<number>>>>} coordinates an array of Polygons
 * @param {Object=} properties an Object of key-value pairs to add as properties
 * @returns {Feature<MultiPolygon>} a multipolygon feature
 * @throws {Error} if no coordinates are passed
 * @example
 * var multiPoly = turf.multiPolygon([[[[0,0],[0,10],[10,10],[10,0],[0,0]]]]);
 *
 * //=multiPoly
 *
 */
var multiPolygon = function (coordinates, properties) {
    if (!coordinates) throw new Error('No coordinates passed');

    return feature({
        type: 'MultiPolygon',
        coordinates: coordinates
    }, properties);
};

/**
 * Creates a {@link Feature<GeometryCollection>} based on a
 * coordinate array. Properties can be added optionally.
 *
 * @name geometryCollection
 * @param {Array<{Geometry}>} geometries an array of GeoJSON Geometries
 * @param {Object=} properties an Object of key-value pairs to add as properties
 * @returns {Feature<GeometryCollection>} a GeoJSON GeometryCollection Feature
 * @example
 * var pt = {
 *     "type": "Point",
 *       "coordinates": [100, 0]
 *     };
 * var line = {
 *     "type": "LineString",
 *     "coordinates": [ [101, 0], [102, 1] ]
 *   };
 * var collection = turf.geometryCollection([pt, line]);
 *
 * //=collection
 */
var geometryCollection = function (geometries, properties) {
    if (!geometries) throw new Error('No geometries passed');

    return feature({
        type: 'GeometryCollection',
        geometries: geometries
    }, properties);
};

var factors = {
    miles: 3960,
    nauticalmiles: 3441.145,
    degrees: 57.2957795,
    radians: 1,
    inches: 250905600,
    yards: 6969600,
    meters: 6373000,
    metres: 6373000,
    kilometers: 6373,
    kilometres: 6373,
    feet: 20908792.65
};

/*
 * Convert a distance measurement from radians to a more friendly unit.
 *
 * @name radiansToDistance
 * @param {number} distance in radians across the sphere
 * @param {string} [units=kilometers] can be degrees, radians, miles, or kilometers
 * inches, yards, metres, meters, kilometres, kilometers.
 * @returns {number} distance
 */
var radiansToDistance = function (radians, units) {
    var factor = factors[units || 'kilometers'];
    if (factor === undefined) throw new Error('Invalid unit');

    return radians * factor;
};

/*
 * Convert a distance measurement from a real-world unit into radians
 *
 * @name distanceToRadians
 * @param {number} distance in real units
 * @param {string} [units=kilometers] can be degrees, radians, miles, or kilometers
 * inches, yards, metres, meters, kilometres, kilometers.
 * @returns {number} radians
 */
var distanceToRadians$1 = function (distance, units) {
    var factor = factors[units || 'kilometers'];
    if (factor === undefined) throw new Error('Invalid unit');

    return distance / factor;
};

/*
 * Convert a distance measurement from a real-world unit into degrees
 *
 * @name distanceToRadians
 * @param {number} distance in real units
 * @param {string} [units=kilometers] can be degrees, radians, miles, or kilometers
 * inches, yards, metres, meters, kilometres, kilometers.
 * @returns {number} degrees
 */
var distanceToDegrees = function (distance, units) {
    var factor = factors[units || 'kilometers'];
    if (factor === undefined) throw new Error('Invalid unit');

    return (distance / factor) * 57.2958;
};

var index$3 = {
	feature: feature_1,
	point: point$1,
	polygon: polygon,
	lineString: lineString,
	featureCollection: featureCollection,
	multiLineString: multiLineString,
	multiPoint: multiPoint,
	multiPolygon: multiPolygon,
	geometryCollection: geometryCollection,
	radiansToDistance: radiansToDistance,
	distanceToRadians: distanceToRadians$1,
	distanceToDegrees: distanceToDegrees
};

//http://en.wikipedia.org/wiki/Haversine_formula
//http://www.movable-type.co.uk/scripts/latlong.html
var getCoord = index$1.getCoord;
var helpers = index$3;
var point = helpers.point;
var distanceToRadians = helpers.distanceToRadians;

/**
 * Takes a {@link Point} and calculates the location of a destination point given a distance in degrees, radians, miles, or kilometers; and bearing in degrees. This uses the [Haversine formula](http://en.wikipedia.org/wiki/Haversine_formula) to account for global curvature.
 *
 * @name destination
 * @param {Feature<Point>} from starting point
 * @param {number} distance distance from the starting point
 * @param {number} bearing ranging from -180 to 180
 * @param {string} [units=kilometers] miles, kilometers, degrees, or radians
 * @returns {Feature<Point>} destination point
 * @example
 * var point = {
 *   "type": "Feature",
 *   "properties": {
 *     "marker-color": "#0f0"
 *   },
 *   "geometry": {
 *     "type": "Point",
 *     "coordinates": [-75.343, 39.984]
 *   }
 * };
 * var distance = 50;
 * var bearing = 90;
 * var units = 'miles';
 *
 * var destination = turf.destination(point, distance, bearing, units);
 * destination.properties['marker-color'] = '#f00';
 *
 * var result = {
 *   "type": "FeatureCollection",
 *   "features": [point, destination]
 * };
 *
 * //=result
 */
var index = function (from, distance, bearing, units) {
    var degrees2radians = Math.PI / 180;
    var radians2degrees = 180 / Math.PI;
    var coordinates1 = getCoord(from);
    var longitude1 = degrees2radians * coordinates1[0];
    var latitude1 = degrees2radians * coordinates1[1];
    var bearing_rad = degrees2radians * bearing;

    var radians = distanceToRadians(distance, units);

    var latitude2 = Math.asin(Math.sin(latitude1) * Math.cos(radians) +
        Math.cos(latitude1) * Math.sin(radians) * Math.cos(bearing_rad));
    var longitude2 = longitude1 + Math.atan2(Math.sin(bearing_rad) *
        Math.sin(radians) * Math.cos(latitude1),
        Math.cos(radians) - Math.sin(latitude1) * Math.sin(latitude2));

    return point([radians2degrees * longitude2, radians2degrees * latitude2]);
};

/**
 * Callback for coordEach
 *
 * @private
 * @callback coordEachCallback
 * @param {[number, number]} currentCoords The current coordinates being processed.
 * @param {number} currentIndex The index of the current element being processed in the
 * array.Starts at index 0, if an initialValue is provided, and at index 1 otherwise.
 */

/**
 * Iterate over coordinates in any GeoJSON object, similar to Array.forEach()
 *
 * @name coordEach
 * @param {Object} layer any GeoJSON object
 * @param {Function} callback a method that takes (currentCoords, currentIndex)
 * @param {boolean} [excludeWrapCoord=false] whether or not to include
 * the final coordinate of LinearRings that wraps the ring in its iteration.
 * @example
 * var features = {
 *   "type": "FeatureCollection",
 *   "features": [
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [26, 37]
 *       }
 *     },
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [36, 53]
 *       }
 *     }
 *   ]
 * };
 * turf.coordEach(features, function (currentCoords, currentIndex) {
 *   //=currentCoords
 *   //=currentIndex
 * });
 */
function coordEach(layer, callback, excludeWrapCoord) {
    var i, j, k, g, l, geometry, stopG, coords,
        geometryMaybeCollection,
        wrapShrink = 0,
        currentIndex = 0,
        isGeometryCollection,
        isFeatureCollection = layer.type === 'FeatureCollection',
        isFeature = layer.type === 'Feature',
        stop = isFeatureCollection ? layer.features.length : 1;

  // This logic may look a little weird. The reason why it is that way
  // is because it's trying to be fast. GeoJSON supports multiple kinds
  // of objects at its root: FeatureCollection, Features, Geometries.
  // This function has the responsibility of handling all of them, and that
  // means that some of the `for` loops you see below actually just don't apply
  // to certain inputs. For instance, if you give this just a
  // Point geometry, then both loops are short-circuited and all we do
  // is gradually rename the input until it's called 'geometry'.
  //
  // This also aims to allocate as few resources as possible: just a
  // few numbers and booleans, rather than any temporary arrays as would
  // be required with the normalization approach.
    for (i = 0; i < stop; i++) {

        geometryMaybeCollection = (isFeatureCollection ? layer.features[i].geometry :
        (isFeature ? layer.geometry : layer));
        isGeometryCollection = geometryMaybeCollection.type === 'GeometryCollection';
        stopG = isGeometryCollection ? geometryMaybeCollection.geometries.length : 1;

        for (g = 0; g < stopG; g++) {
            geometry = isGeometryCollection ?
            geometryMaybeCollection.geometries[g] : geometryMaybeCollection;
            coords = geometry.coordinates;

            wrapShrink = (excludeWrapCoord &&
                (geometry.type === 'Polygon' || geometry.type === 'MultiPolygon')) ?
                1 : 0;

            if (geometry.type === 'Point') {
                callback(coords, currentIndex);
                currentIndex++;
            } else if (geometry.type === 'LineString' || geometry.type === 'MultiPoint') {
                for (j = 0; j < coords.length; j++) {
                    callback(coords[j], currentIndex);
                    currentIndex++;
                }
            } else if (geometry.type === 'Polygon' || geometry.type === 'MultiLineString') {
                for (j = 0; j < coords.length; j++)
                    for (k = 0; k < coords[j].length - wrapShrink; k++) {
                        callback(coords[j][k], currentIndex);
                        currentIndex++;
                    }
            } else if (geometry.type === 'MultiPolygon') {
                for (j = 0; j < coords.length; j++)
                    for (k = 0; k < coords[j].length; k++)
                        for (l = 0; l < coords[j][k].length - wrapShrink; l++) {
                            callback(coords[j][k][l], currentIndex);
                            currentIndex++;
                        }
            } else if (geometry.type === 'GeometryCollection') {
                for (j = 0; j < geometry.geometries.length; j++)
                    coordEach(geometry.geometries[j], callback, excludeWrapCoord);
            } else {
                throw new Error('Unknown Geometry Type');
            }
        }
    }
}
var coordEach_1 = coordEach;

/**
 * Callback for coordReduce
 *
 * The first time the callback function is called, the values provided as arguments depend
 * on whether the reduce method has an initialValue argument.
 *
 * If an initialValue is provided to the reduce method:
 *  - The previousValue argument is initialValue.
 *  - The currentValue argument is the value of the first element present in the array.
 *
 * If an initialValue is not provided:
 *  - The previousValue argument is the value of the first element present in the array.
 *  - The currentValue argument is the value of the second element present in the array.
 *
 * @private
 * @callback coordReduceCallback
 * @param {*} previousValue The accumulated value previously returned in the last invocation
 * of the callback, or initialValue, if supplied.
 * @param {[number, number]} currentCoords The current coordinate being processed.
 * @param {number} currentIndex The index of the current element being processed in the
 * array.Starts at index 0, if an initialValue is provided, and at index 1 otherwise.
 */

/**
 * Reduce coordinates in any GeoJSON object, similar to Array.reduce()
 *
 * @name coordReduce
 * @param {Object} layer any GeoJSON object
 * @param {Function} callback a method that takes (previousValue, currentCoords, currentIndex)
 * @param {*} [initialValue] Value to use as the first argument to the first call of the callback.
 * @param {boolean} [excludeWrapCoord=false] whether or not to include
 * the final coordinate of LinearRings that wraps the ring in its iteration.
 * @returns {*} The value that results from the reduction.
 * @example
 * var features = {
 *   "type": "FeatureCollection",
 *   "features": [
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [26, 37]
 *       }
 *     },
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [36, 53]
 *       }
 *     }
 *   ]
 * };
 * turf.coordReduce(features, function (previousValue, currentCoords, currentIndex) {
 *   //=previousValue
 *   //=currentCoords
 *   //=currentIndex
 *   return currentCoords;
 * });
 */
function coordReduce(layer, callback, initialValue, excludeWrapCoord) {
    var previousValue = initialValue;
    coordEach(layer, function (currentCoords, currentIndex) {
        if (currentIndex === 0 && initialValue === undefined) {
            previousValue = currentCoords;
        } else {
            previousValue = callback(previousValue, currentCoords, currentIndex);
        }
    }, excludeWrapCoord);
    return previousValue;
}
var coordReduce_1 = coordReduce;

/**
 * Callback for propEach
 *
 * @private
 * @callback propEachCallback
 * @param {*} currentProperties The current properties being processed.
 * @param {number} currentIndex The index of the current element being processed in the
 * array.Starts at index 0, if an initialValue is provided, and at index 1 otherwise.
 */

/**
 * Iterate over properties in any GeoJSON object, similar to Array.forEach()
 *
 * @name propEach
 * @param {Object} layer any GeoJSON object
 * @param {Function} callback a method that takes (currentProperties, currentIndex)
 * @example
 * var features = {
 *   "type": "FeatureCollection",
 *   "features": [
 *     {
 *       "type": "Feature",
 *       "properties": {"foo": "bar"},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [26, 37]
 *       }
 *     },
 *     {
 *       "type": "Feature",
 *       "properties": {"hello": "world"},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [36, 53]
 *       }
 *     }
 *   ]
 * };
 * turf.propEach(features, function (currentProperties, currentIndex) {
 *   //=currentProperties
 *   //=currentIndex
 * });
 */
function propEach(layer, callback) {
    var i;
    switch (layer.type) {
    case 'FeatureCollection':
        for (i = 0; i < layer.features.length; i++) {
            callback(layer.features[i].properties, i);
        }
        break;
    case 'Feature':
        callback(layer.properties, 0);
        break;
    }
}
var propEach_1 = propEach;


/**
 * Callback for propReduce
 *
 * The first time the callback function is called, the values provided as arguments depend
 * on whether the reduce method has an initialValue argument.
 *
 * If an initialValue is provided to the reduce method:
 *  - The previousValue argument is initialValue.
 *  - The currentValue argument is the value of the first element present in the array.
 *
 * If an initialValue is not provided:
 *  - The previousValue argument is the value of the first element present in the array.
 *  - The currentValue argument is the value of the second element present in the array.
 *
 * @private
 * @callback propReduceCallback
 * @param {*} previousValue The accumulated value previously returned in the last invocation
 * of the callback, or initialValue, if supplied.
 * @param {*} currentProperties The current properties being processed.
 * @param {number} currentIndex The index of the current element being processed in the
 * array.Starts at index 0, if an initialValue is provided, and at index 1 otherwise.
 */

/**
 * Reduce properties in any GeoJSON object into a single value,
 * similar to how Array.reduce works. However, in this case we lazily run
 * the reduction, so an array of all properties is unnecessary.
 *
 * @name propReduce
 * @param {Object} layer any GeoJSON object
 * @param {Function} callback a method that takes (previousValue, currentProperties, currentIndex)
 * @param {*} [initialValue] Value to use as the first argument to the first call of the callback.
 * @returns {*} The value that results from the reduction.
 * @example
 * var features = {
 *   "type": "FeatureCollection",
 *   "features": [
 *     {
 *       "type": "Feature",
 *       "properties": {"foo": "bar"},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [26, 37]
 *       }
 *     },
 *     {
 *       "type": "Feature",
 *       "properties": {"hello": "world"},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [36, 53]
 *       }
 *     }
 *   ]
 * };
 * turf.propReduce(features, function (previousValue, currentProperties, currentIndex) {
 *   //=previousValue
 *   //=currentProperties
 *   //=currentIndex
 *   return currentProperties
 * });
 */
function propReduce(layer, callback, initialValue) {
    var previousValue = initialValue;
    propEach(layer, function (currentProperties, currentIndex) {
        if (currentIndex === 0 && initialValue === undefined) {
            previousValue = currentProperties;
        } else {
            previousValue = callback(previousValue, currentProperties, currentIndex);
        }
    });
    return previousValue;
}
var propReduce_1 = propReduce;

/**
 * Callback for featureEach
 *
 * @private
 * @callback featureEachCallback
 * @param {Feature<any>} currentFeature The current feature being processed.
 * @param {number} currentIndex The index of the current element being processed in the
 * array.Starts at index 0, if an initialValue is provided, and at index 1 otherwise.
 */

/**
 * Iterate over features in any GeoJSON object, similar to
 * Array.forEach.
 *
 * @name featureEach
 * @param {Object} layer any GeoJSON object
 * @param {Function} callback a method that takes (currentFeature, currentIndex)
 * @example
 * var features = {
 *   "type": "FeatureCollection",
 *   "features": [
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [26, 37]
 *       }
 *     },
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [36, 53]
 *       }
 *     }
 *   ]
 * };
 * turf.featureEach(features, function (currentFeature, currentIndex) {
 *   //=currentFeature
 *   //=currentIndex
 * });
 */
function featureEach(layer, callback) {
    if (layer.type === 'Feature') {
        callback(layer, 0);
    } else if (layer.type === 'FeatureCollection') {
        for (var i = 0; i < layer.features.length; i++) {
            callback(layer.features[i], i);
        }
    }
}
var featureEach_1 = featureEach;

/**
 * Callback for featureReduce
 *
 * The first time the callback function is called, the values provided as arguments depend
 * on whether the reduce method has an initialValue argument.
 *
 * If an initialValue is provided to the reduce method:
 *  - The previousValue argument is initialValue.
 *  - The currentValue argument is the value of the first element present in the array.
 *
 * If an initialValue is not provided:
 *  - The previousValue argument is the value of the first element present in the array.
 *  - The currentValue argument is the value of the second element present in the array.
 *
 * @private
 * @callback featureReduceCallback
 * @param {*} previousValue The accumulated value previously returned in the last invocation
 * of the callback, or initialValue, if supplied.
 * @param {Feature<any>} currentFeature The current Feature being processed.
 * @param {number} currentIndex The index of the current element being processed in the
 * array.Starts at index 0, if an initialValue is provided, and at index 1 otherwise.
 */

/**
 * Reduce features in any GeoJSON object, similar to Array.reduce().
 *
 * @name featureReduce
 * @param {Object} layer any GeoJSON object
 * @param {Function} callback a method that takes (previousValue, currentFeature, currentIndex)
 * @param {*} [initialValue] Value to use as the first argument to the first call of the callback.
 * @returns {*} The value that results from the reduction.
 * @example
 * var features = {
 *   "type": "FeatureCollection",
 *   "features": [
 *     {
 *       "type": "Feature",
 *       "properties": {"foo": "bar"},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [26, 37]
 *       }
 *     },
 *     {
 *       "type": "Feature",
 *       "properties": {"hello": "world"},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [36, 53]
 *       }
 *     }
 *   ]
 * };
 * turf.featureReduce(features, function (previousValue, currentFeature, currentIndex) {
 *   //=previousValue
 *   //=currentFeature
 *   //=currentIndex
 *   return currentFeature
 * });
 */
function featureReduce(layer, callback, initialValue) {
    var previousValue = initialValue;
    featureEach(layer, function (currentFeature, currentIndex) {
        if (currentIndex === 0 && initialValue === undefined) {
            previousValue = currentFeature;
        } else {
            previousValue = callback(previousValue, currentFeature, currentIndex);
        }
    });
    return previousValue;
}
var featureReduce_1 = featureReduce;

/**
 * Get all coordinates from any GeoJSON object.
 *
 * @name coordAll
 * @param {Object} layer any GeoJSON object
 * @returns {Array<Array<number>>} coordinate position array
 * @example
 * var features = {
 *   "type": "FeatureCollection",
 *   "features": [
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [26, 37]
 *       }
 *     },
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [36, 53]
 *       }
 *     }
 *   ]
 * };
 * var coords = turf.coordAll(features);
 * //=coords
 */
function coordAll(layer) {
    var coords = [];
    coordEach(layer, function (coord) {
        coords.push(coord);
    });
    return coords;
}
var coordAll_1 = coordAll;

/**
 * Iterate over each geometry in any GeoJSON object, similar to Array.forEach()
 *
 * @name geomEach
 * @param {Object} layer any GeoJSON object
 * @param {Function} callback a method that takes (currentGeometry, currentIndex)
 * @example
 * var features = {
 *   "type": "FeatureCollection",
 *   "features": [
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [26, 37]
 *       }
 *     },
 *     {
 *       "type": "Feature",
 *       "properties": {},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [36, 53]
 *       }
 *     }
 *   ]
 * };
 * turf.geomEach(features, function (currentGeometry, currentIndex) {
 *   //=currentGeometry
 *   //=currentIndex
 * });
 */
function geomEach(layer, callback) {
    var i, j, g, geometry, stopG,
        geometryMaybeCollection,
        isGeometryCollection,
        currentIndex = 0,
        isFeatureCollection = layer.type === 'FeatureCollection',
        isFeature = layer.type === 'Feature',
        stop = isFeatureCollection ? layer.features.length : 1;

  // This logic may look a little weird. The reason why it is that way
  // is because it's trying to be fast. GeoJSON supports multiple kinds
  // of objects at its root: FeatureCollection, Features, Geometries.
  // This function has the responsibility of handling all of them, and that
  // means that some of the `for` loops you see below actually just don't apply
  // to certain inputs. For instance, if you give this just a
  // Point geometry, then both loops are short-circuited and all we do
  // is gradually rename the input until it's called 'geometry'.
  //
  // This also aims to allocate as few resources as possible: just a
  // few numbers and booleans, rather than any temporary arrays as would
  // be required with the normalization approach.
    for (i = 0; i < stop; i++) {

        geometryMaybeCollection = (isFeatureCollection ? layer.features[i].geometry :
        (isFeature ? layer.geometry : layer));
        isGeometryCollection = geometryMaybeCollection.type === 'GeometryCollection';
        stopG = isGeometryCollection ? geometryMaybeCollection.geometries.length : 1;

        for (g = 0; g < stopG; g++) {
            geometry = isGeometryCollection ?
            geometryMaybeCollection.geometries[g] : geometryMaybeCollection;

            if (geometry.type === 'Point' ||
                geometry.type === 'LineString' ||
                geometry.type === 'MultiPoint' ||
                geometry.type === 'Polygon' ||
                geometry.type === 'MultiLineString' ||
                geometry.type === 'MultiPolygon') {
                callback(geometry, currentIndex);
                currentIndex++;
            } else if (geometry.type === 'GeometryCollection') {
                for (j = 0; j < geometry.geometries.length; j++) {
                    callback(geometry.geometries[j], currentIndex);
                    currentIndex++;
                }
            } else {
                throw new Error('Unknown Geometry Type');
            }
        }
    }
}
var geomEach_1 = geomEach;

/**
 * Callback for geomReduce
 *
 * The first time the callback function is called, the values provided as arguments depend
 * on whether the reduce method has an initialValue argument.
 *
 * If an initialValue is provided to the reduce method:
 *  - The previousValue argument is initialValue.
 *  - The currentValue argument is the value of the first element present in the array.
 *
 * If an initialValue is not provided:
 *  - The previousValue argument is the value of the first element present in the array.
 *  - The currentValue argument is the value of the second element present in the array.
 *
 * @private
 * @callback geomReduceCallback
 * @param {*} previousValue The accumulated value previously returned in the last invocation
 * of the callback, or initialValue, if supplied.
 * @param {*} currentGeometry The current Feature being processed.
 * @param {number} currentIndex The index of the current element being processed in the
 * array.Starts at index 0, if an initialValue is provided, and at index 1 otherwise.
 */

/**
 * Reduce geometry in any GeoJSON object, similar to Array.reduce().
 *
 * @name geomReduce
 * @param {Object} layer any GeoJSON object
 * @param {Function} callback a method that takes (previousValue, currentGeometry, currentIndex)
 * @param {*} [initialValue] Value to use as the first argument to the first call of the callback.
 * @returns {*} The value that results from the reduction.
 * @example
 * var features = {
 *   "type": "FeatureCollection",
 *   "features": [
 *     {
 *       "type": "Feature",
 *       "properties": {"foo": "bar"},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [26, 37]
 *       }
 *     },
 *     {
 *       "type": "Feature",
 *       "properties": {"hello": "world"},
 *       "geometry": {
 *         "type": "Point",
 *         "coordinates": [36, 53]
 *       }
 *     }
 *   ]
 * };
 * turf.geomReduce(features, function (previousValue, currentGeometry, currentIndex) {
 *   //=previousValue
 *   //=currentGeometry
 *   //=currentIndex
 *   return currentGeometry
 * });
 */
function geomReduce(layer, callback, initialValue) {
    var previousValue = initialValue;
    geomEach(layer, function (currentGeometry, currentIndex) {
        if (currentIndex === 0 && initialValue === undefined) {
            previousValue = currentGeometry;
        } else {
            previousValue = callback(previousValue, currentGeometry, currentIndex);
        }
    });
    return previousValue;
}
var geomReduce_1 = geomReduce;

var index$6 = {
	coordEach: coordEach_1,
	coordReduce: coordReduce_1,
	propEach: propEach_1,
	propReduce: propReduce_1,
	featureEach: featureEach_1,
	featureReduce: featureReduce_1,
	coordAll: coordAll_1,
	geomEach: geomEach_1,
	geomReduce: geomReduce_1
};

var each = index$6.coordEach;
var point$2 = index$3.point;

/**
 * Takes one or more features and calculates the centroid using
 * the mean of all vertices.
 * This lessens the effect of small islands and artifacts when calculating
 * the centroid of a set of polygons.
 *
 * @name centroid
 * @param {(Feature|FeatureCollection)} features input features
 * @returns {Feature<Point>} the centroid of the input features
 * @example
 * var poly = {
 *   "type": "Feature",
 *   "properties": {},
 *   "geometry": {
 *     "type": "Polygon",
 *     "coordinates": [[
 *       [105.818939,21.004714],
 *       [105.818939,21.061754],
 *       [105.890007,21.061754],
 *       [105.890007,21.004714],
 *       [105.818939,21.004714]
 *     ]]
 *   }
 * };
 *
 * var centroidPt = turf.centroid(poly);
 *
 * var result = {
 *   "type": "FeatureCollection",
 *   "features": [poly, centroidPt]
 * };
 *
 * //=result
 */
var index$5 = function (features) {
    var xSum = 0, ySum = 0, len = 0;
    each(features, function (coord) {
        xSum += coord[0];
        ySum += coord[1];
        len++;
    }, true);
    return point$2([xSum / len, ySum / len]);
};

var getCoord$2 = index$1.getCoord;
//http://en.wikipedia.org/wiki/Haversine_formula
//http://www.movable-type.co.uk/scripts/latlong.html

/**
 * Takes two {@link Point|points} and finds the geographic bearing between them.
 *
 * @name bearing
 * @param {Feature<Point>} start starting Point
 * @param {Feature<Point>} end ending Point
 * @param {boolean} [final=false] calculates the final bearing if true
 * @returns {number} bearing in decimal degrees
 * @addToMap point1, point2
 * @example
 * var point1 = {
 *   "type": "Feature",
 *   "properties": {
 *     "marker-color": '#f00'
 *   },
 *   "geometry": {
 *     "type": "Point",
 *     "coordinates": [-75.343, 39.984]
 *   }
 * };
 * var point2 = {
 *   "type": "Feature",
 *   "properties": {
 *     "marker-color": '#0f0'
 *   },
 *   "geometry": {
 *     "type": "Point",
 *     "coordinates": [-75.534, 39.123]
 *   }
 * };
 *
 * var bearing = turf.bearing(point1, point2);
 * point1.properties.bearing = bearing
 * //=bearing
 */
function bearing(start, end, final) {
    if (final === true) return calculateFinalBearing(start, end);

    var degrees2radians = Math.PI / 180;
    var radians2degrees = 180 / Math.PI;
    var coordinates1 = getCoord$2(start);
    var coordinates2 = getCoord$2(end);

    var lon1 = degrees2radians * coordinates1[0];
    var lon2 = degrees2radians * coordinates2[0];
    var lat1 = degrees2radians * coordinates1[1];
    var lat2 = degrees2radians * coordinates2[1];
    var a = Math.sin(lon2 - lon1) * Math.cos(lat2);
    var b = Math.cos(lat1) * Math.sin(lat2) -
        Math.sin(lat1) * Math.cos(lat2) * Math.cos(lon2 - lon1);

    var bear = radians2degrees * Math.atan2(a, b);

    return bear;
}

/**
 * Calculates Final Bearing
 * @private
 * @param {Feature<Point>} start starting Point
 * @param {Feature<Point>} end ending Point
 * @returns {number} bearing
 */
function calculateFinalBearing(start, end) {
    // Swap start & end
    var bear = bearing(end, start);
    bear = (bear + 180) % 360;
    return bear;
}

var index$8 = bearing;

var getCoord$3 = index$1.getCoord;
var radiansToDistance$1 = index$3.radiansToDistance;
//http://en.wikipedia.org/wiki/Haversine_formula
//http://www.movable-type.co.uk/scripts/latlong.html

/**
 * Calculates the distance between two {@link Point|points} in degrees, radians,
 * miles, or kilometers. This uses the
 * [Haversine formula](http://en.wikipedia.org/wiki/Haversine_formula)
 * to account for global curvature.
 *
 * @name distance
 * @param {Feature<Point>} from origin point
 * @param {Feature<Point>} to destination point
 * @param {string} [units=kilometers] can be degrees, radians, miles, or kilometers
 * @returns {number} distance between the two points
 * @example
 * var from = {
 *   "type": "Feature",
 *   "properties": {},
 *   "geometry": {
 *     "type": "Point",
 *     "coordinates": [-75.343, 39.984]
 *   }
 * };
 * var to = {
 *   "type": "Feature",
 *   "properties": {},
 *   "geometry": {
 *     "type": "Point",
 *     "coordinates": [-75.534, 39.123]
 *   }
 * };
 * var units = "miles";
 *
 * var points = {
 *   "type": "FeatureCollection",
 *   "features": [from, to]
 * };
 *
 * //=points
 *
 * var distance = turf.distance(from, to, units);
 *
 * //=distance
 */
var index$9 = function (from, to, units) {
    var degrees2radians = Math.PI / 180;
    var coordinates1 = getCoord$3(from);
    var coordinates2 = getCoord$3(to);
    var dLat = degrees2radians * (coordinates2[1] - coordinates1[1]);
    var dLon = degrees2radians * (coordinates2[0] - coordinates1[0]);
    var lat1 = degrees2radians * coordinates1[1];
    var lat2 = degrees2radians * coordinates2[1];

    var a = Math.pow(Math.sin(dLat / 2), 2) +
          Math.pow(Math.sin(dLon / 2), 2) * Math.cos(lat1) * Math.cos(lat2);

    return radiansToDistance$1(2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a)), units);
};

function fromFeature (feature, options) {
  options = options || {};
  feature = checkFeatures(feature, options);
  return processFeature(feature, options)
}

var units = 'meters';

function tanDeg (deg) {
  var rad = deg * Math.PI / 180;
  return Math.tan(rad)
}

function cosDeg (deg) {
  var rad = deg * Math.PI / 180;
  return Math.cos(rad)
}

function getNested (feature, options) {
  var properties = feature.properties || {};
  if (options.nested) {
    if (properties[options.nested]) {
      properties = properties[options.nested];
    } else {
      properties = {};
    }
  }
  return properties
}

function checkFeatures (feature, options) {
  var properties = getNested(feature, options);
  var angle = properties.angle || options.angle;

  var geometryType = feature.geometry.type;

  if (geometryType === 'GeometryCollection') {
    if (feature.geometry.geometries.length === 3 &&
        feature.geometry.geometries[0].type === 'Point' &&
        feature.geometry.geometries[1].type === 'Point' &&
        feature.geometry.geometries[2].type === 'Point') {
      return feature
    }
  }

  if (angle === undefined) {
    throw new Error('feature must include angle property, or global angle option must be set')
  }

  if (geometryType === 'LineString') {
    if (feature.geometry.coordinates.length === 2) {
      return feature
    } else {
      throw new Error('only accepts only accepts LineStrings with two points')
    }
  } else if (geometryType === 'GeometryCollection') {
    if (feature.geometry.geometries.length === 2 &&
        feature.geometry.geometries[0].type === 'Point' &&
        feature.geometry.geometries[1].type === 'Point') {
      return feature
    } else {
      throw new Error('only accepts GeometryCollections containing two Points')
    }
  } else if (geometryType === 'Point') {
    if (properties.bearing !== undefined && properties.distance !== undefined) {
      return feature
    } else {
      throw new Error('only accepts single Points with distance and bearing properties')
    }
  } else {
    throw new Error('only accepts LineStrings with two points, GeometryCollections \n' +
      'containing two Points, or single Points with distance and bearing properties\n' +
      ' - without the angle property set, GeometryCollections with three Points are accepted')
  }
}

function processFeature (feature, options) {
  var geometryType = feature.geometry.type;
  if (geometryType === 'Point') {
    return processPoint(feature, options)
  } else if (geometryType === 'LineString') {
    return processLineString(feature, options)
  } else if (geometryType === 'GeometryCollection') {
    return processGeometryCollection(feature, options)
  }
}

function processPoint (feature, options) {
  var properties = getNested(feature, options);

  var distance = properties.distance;
  var angle = properties.angle || options.angle;

  var centroid = index(feature, distance, properties.bearing, units);

  var distCentroid = tanDeg(angle / 2) * distance;

  var points = [
    index(centroid, distCentroid, properties.bearing + 90, units),
    index(centroid, -distCentroid, properties.bearing + 90, units)
  ];

  return {
    type: 'Feature',
    properties: feature.properties,
    geometry: {
      type: 'GeometryCollection',
      geometries: [
        feature.geometry,
        {
          type: 'LineString',
          coordinates: [
            points[0].geometry.coordinates,
            points[1].geometry.coordinates
          ]
        }
      ]
    }
  }
}

function processLineString (feature, options) {
  var properties = getNested(feature, options);
  var angle = properties.angle || options.angle;

  var centroid = index$5(feature);

  var points = feature.geometry.coordinates.map(function (coordinate) {
    return {
      type: 'Feature',
      geometry: {
        type: 'Point',
        coordinates: coordinate
      }
    }
  });

  var distCentroid = index$9(points[0], centroid, units);
  var bearing = index$8(points[0], points[1]);

  var distCamera = distCentroid / tanDeg(angle / 2);
  var camera = index(centroid, distCamera, bearing + 90, units);

  return {
    type: 'Feature',
    properties: feature.properties,
    geometry: {
      type: 'GeometryCollection',
      geometries: [
        camera.geometry,
        feature.geometry
      ]
    }
  }
}

function processGeometryCollection (feature, options) {
  var points = feature.geometry.geometries;

  var camera = points[0];
  var target = points[1];
  var targetBearing = index$8(camera, target);

  var angle;

  if (points.length === 2) {
    var properties = getNested(feature, options);
    angle = properties.angle || options.angle;
  } else {
    var angleBearing = index$8(camera, points[2]);
    var bearingDiff = (angleBearing - targetBearing + 360) % 180;

    if (bearingDiff < 90) {
      angle = bearingDiff * 2;
    } else {
      angle = (180 - bearingDiff) * 2;
    }
  }

  var distance = index$9(camera, target, units);
  var distFieldOfViewCorner = distance / cosDeg(angle / 2);

  var fieldOfViewPoint1 = index(camera, distFieldOfViewCorner, targetBearing - angle / 2, units);
  var fieldOfViewPoint2 = index(camera, distFieldOfViewCorner, targetBearing + angle / 2, units);

  return {
    type: 'Feature',
    properties: Object.assign({}, feature.properties, {
      angle: angle,
      bearing: targetBearing,
      distance: distance
    }),
    geometry: {
      type: 'GeometryCollection',
      geometries: [
        camera,
        {
          type: 'LineString',
          coordinates: [
            fieldOfViewPoint1.geometry.coordinates,
            fieldOfViewPoint2.geometry.coordinates
          ]
        }
      ]
    }
  }
}

var GeotagPhotoCameraControl = L.Control.extend({
  options: {
    position: 'topleft'
  },

  initialize: function initialize(geotagPhotoCamera, options) {
    this._geotagPhotoCamera = geotagPhotoCamera;
    L.setOptions(this, options);
  },

  onAdd: function onAdd(map) {
    this._map = map;

    var controlName = 'leaflet-control-geotag-photo-';
    var container = L.DomUtil.create('div', controlName + ' leaflet-bar');

    var cameraImg = '<img role="none" src="' + this.options.cameraImg + '" />';
    var crosshairImg = '<img role="none" src="' + this.options.crosshairImg + '" />';

    this._cameraButton = this._createButton(cameraImg, 'Move camera back to map (C)', controlName + 'camera', container, this._centerCamera);

    this._crosshairButton = this._createButton(crosshairImg, 'Move map back to camera (M)', controlName + 'crosshair', container, this._centerMap);

    this._boundMapKeyPress = this._mapKeyPress.bind(this);
    this._map.on('keypress', this._boundMapKeyPress);

    return container;
  },

  _createButton: function _createButton(html, title, className, container, fn) {
    var link = L.DomUtil.create('a', className, container);
    link.innerHTML = html;
    link.href = '#';
    link.title = title;

    /*
     * Will force screen readers like VoiceOver to read this as "Zoom in - button"
     */
    link.setAttribute('role', 'button');
    link.setAttribute('aria-label', title);

    L.DomEvent.on(link, 'mousedown dblclick', L.DomEvent.stopPropagation).on(link, 'click', L.DomEvent.stop).on(link, 'click', fn, this).on(link, 'click', this._refocusOnMap, this);

    return link;
  },

  onRemove: function onRemove(map) {
    L.DomUtil.remove(this._element);
    map.off('keypress', this._boundMapKeyPress);
  },

  _mapKeyPress: function _mapKeyPress(evt) {
    if (evt.originalEvent.charCode === 99) {
      // C key
      this._centerCamera();
    } else if (evt.originalEvent.charCode === 109) {
      // M key
      this._centerMap();
    }
  },

  _centerCamera: function _centerCamera() {
    if (this._map && this._geotagPhotoCamera) {
      this._geotagPhotoCamera.centerBounds(this._map.getBounds());
    }
  },

  _centerMap: function _centerMap() {
    if (this._map && this._geotagPhotoCamera) {
      this._map.fitBounds(this._geotagPhotoCamera.getBounds());
    }
  }

});

L.geotagPhotoCameraControl = function (geotagPhotoCamera, options) {
  return new GeotagPhotoCameraControl(geotagPhotoCamera, options);
};

var GeotagPhotoCamera = L.FeatureGroup.extend({

  options: {
    // Whether the camera is draggable with mouse/touch or not
    draggable: true,

    // Whether to show camera control buttons
    control: true,

    // Whether the angle of the field-of-view can be changed with a draggable marker
    angleMarker: true,

    minAngle: 5,
    maxAngle: 120,

    // Control button images
    controlCameraImg: camera_icon_svg,
    controlCrosshairImg: crosshair_icon_svg,

    cameraIcon: L.icon({
      iconUrl: camera_svg,
      iconSize: [38, 38],
      iconAnchor: [19, 19]
    }),

    targetIcon: L.icon({
      iconUrl: marker_svg,
      iconSize: [32, 32],
      iconAnchor: [16, 16]
    }),

    angleIcon: L.icon({
      iconUrl: marker_svg,	
      iconSize: [32, 32],
      iconAnchor: [16, 16]
    }),

    outlineStyle: {
      color: 'black',
      opacity: 0.5,
      weight: 2,
      dashArray: '5, 7',
      lineCap: 'round',
      lineJoin: 'round'
    },

    fillStyle: {
      weight: 0,
      fillOpacity: 0.2,
      fillColor: '#3388ff'
    }
  },

  initialize: function initialize(feature, options) {
    L.setOptions(this, options);

    this.options.minAngle = Math.max(this.options.minAngle, 1);
    this.options.maxAngle = Math.min(this.options.maxAngle, 179);

    this._fieldOfView = fromFeature(feature);
    this._angle = this._fieldOfView.properties.angle;

    var layers = this._createLayers();
    L.LayerGroup.prototype.initialize.call(this, layers);

    this.setDraggable(this.options.draggable);
  },

  _createLayers: function _createLayers() {
    this._cameraIcon = this.options.cameraIcon;
    this._targetIcon = this.options.targetIcon;
    this._angleIcon = this.options.angleIcon;

    var pointList = this._getPointList(this._fieldOfView);
    var cameraLatLng = this._getCameraFromPointList(pointList);
    var targetLatLng = this._getTargetFromPointList(pointList);
    var angleLatLng = this._getAngleFromPointList(pointList);

    this._polyline = L.polyline(pointList, this.options.outlineStyle);
    this._polygon = L.polygon(pointList, Object.assign(this.options.fillStyle, {
      className: 'field-of-view'
    }));

    this._control = L.geotagPhotoCameraControl(this, {
      cameraImg: this.options.controlCameraImg,
      crosshairImg: this.options.controlCrosshairImg
    });

    this._cameraMarker = L.marker(cameraLatLng, {
      icon: this._cameraIcon,
      draggable: this.options.draggable,
      zIndexOffset: 600,
      title: 'Camera',
      alt: 'Location of the camera'
    }).on('drag', this._onMarkerDrag, this).on('dragend', this._onMarkerDragEnd, this);

    this._targetMarker = L.marker(targetLatLng, {
      icon: this._targetIcon,
      draggable: this.options.draggable,
      zIndexOffset: 200,
      title: 'Target',
      alt: 'Location of the target'
    }).on('drag', this._onMarkerDrag, this).on('dragend', this._onMarkerDragEnd, this);

    this._angleMarker = L.marker(angleLatLng, {
      icon: this._angleIcon,
      draggable: this.options.draggable,
      zIndexOffset: 400,
      title: 'Angle',
      alt: 'Field of view angle'
    }).on('drag', this._onAngleMarkerDrag, this).on('dragend', this._onMarkerDragEnd, this);

    var boundUpdateMarkerBearings = this._updateMarkerBearings.bind(this);
    var markerSetPos = function markerSetPos(pos) {
      var protoMarkerSetPos = L.Marker.prototype._setPos;
      protoMarkerSetPos.call(this, pos);
      boundUpdateMarkerBearings();
    };

    this._cameraMarker._setPos = this._targetMarker._setPos = markerSetPos;

    return [this._polygon, this._polyline, this._targetMarker, this._angleMarker, this._cameraMarker];
  },

  addTo: function addTo(map) {
    this._map = map;

    L.FeatureGroup.prototype.addTo.call(this, map);

    if (this.options.control) {
      this._control.addTo(map);
    }

    this._boundOnDocumentKeyDown = this._onDocumentKeyDown.bind(this);
    document.addEventListener('keydown', this._boundOnDocumentKeyDown);

    this.setDraggable(this.options.draggable);
    this._updateMarkerBearings(this._fieldOfView);

    return this;
  },

  removeFrom: function removeFrom(map) {
    this._map = undefined;

    L.FeatureGroup.prototype.removeFrom.call(this, map);

    if (this._boundOnDocumentKeyDown) {
      document.removeEventListener('keydown', this._boundOnDocumentKeyDown);
    }

    return this;
  },

  _getPointList: function _getPointList(fieldOfView) {
    return [[fieldOfView.geometry.geometries[1].coordinates[0][1], fieldOfView.geometry.geometries[1].coordinates[0][0]], [fieldOfView.geometry.geometries[0].coordinates[1], fieldOfView.geometry.geometries[0].coordinates[0]], [fieldOfView.geometry.geometries[1].coordinates[1][1], fieldOfView.geometry.geometries[1].coordinates[1][0]]];
  },

  _getCameraFromPointList: function _getCameraFromPointList(pointList) {
    return pointList[1];
  },

  _getTargetFromPointList: function _getTargetFromPointList(pointList) {
    return [(pointList[0][0] + pointList[2][0]) / 2, (pointList[0][1] + pointList[2][1]) / 2];
  },

  _getAngleFromPointList: function _getAngleFromPointList(pointList) {
    return pointList[2];
  },

  _addRotateTransform: function _addRotateTransform(element, rotation) {
    if (!element) {
      return;
    }

    var transform = element.style[L.DomUtil.TRANSFORM];
    var rotate = 'rotate(' + rotation + ')';

    element.style.transformOrigin = 'center center';

    if (transform.indexOf('rotate') !== -1) {
      element.style[L.DomUtil.TRANSFORM] = transform.replace(/rotate\(.*?\)/, rotate);
    } else {
      element.style[L.DomUtil.TRANSFORM] = transform + ' ' + rotate;
    }
  },

  _updateMarkerBearings: function _updateMarkerBearings(fieldOfView) {
    fieldOfView = fieldOfView || this._fieldOfView;

    var bearing = fieldOfView.properties.bearing;
    var angle = fieldOfView.properties.angle;
    this._addRotateTransform(this._cameraMarker._icon, bearing + 'deg');
    this._addRotateTransform(this._targetMarker._icon, bearing + 'deg');
    this._addRotateTransform(this._angleMarker._icon, bearing + angle / 2 + 'deg');
  },

  _drawFieldOfView: function _drawFieldOfView(fieldOfView) {
    fieldOfView = fieldOfView || this._fieldOfView;

    var pointList = this._getPointList(fieldOfView);
    this._polyline.setLatLngs(pointList);
    this._polygon.setLatLngs(pointList);
  },

  _updateFieldOfView: function _updateFieldOfView() {
    var angle = this._angle;
    var cameraLatLng = this._cameraMarker.getLatLng();
    var targetLatLng = this._targetMarker.getLatLng();

    var cameraTarget = {
      type: 'Feature',
      properties: {
        angle: angle
      },
      geometry: {
        type: 'GeometryCollection',
        geometries: [this._geoJsonPoint(cameraLatLng), this._geoJsonPoint(targetLatLng)]
      }
    };

    this._fieldOfView = fromFeature(cameraTarget);

    var angleLatLng = this._getAngleFromPointList(this._getPointList(this._fieldOfView));
    this._angleMarker.setLatLng(angleLatLng);

    this._updateMarkerBearings(this._fieldOfView);
    this._drawFieldOfView(this._fieldOfView);
  },

  _onAngleMarkerDrag: function _onAngleMarkerDrag(evt) {
    var cameraLatLng = this._cameraMarker.getLatLng();
    var targetLatLng = this._targetMarker.getLatLng();
    var angleLatLng = this._angleMarker.getLatLng();

    var points = {
      type: 'Feature',
      geometry: {
        type: 'GeometryCollection',
        geometries: [this._geoJsonPoint(cameraLatLng), this._geoJsonPoint(targetLatLng), this._geoJsonPoint(angleLatLng)]
      }
    };

    this._fieldOfView = fromFeature(points);
    this.setAngle(this._fieldOfView.properties.angle);
  },

  _onMarkerDrag: function _onMarkerDrag(evt) {
    this._updateFieldOfView();
    this.fire('input');
  },

  _onMarkerDragEnd: function _onMarkerDragEnd(evt) {
    this.fire('change');
  },

  _moveMarker: function _moveMarker(marker, offset) {
    var point = this._map.latLngToContainerPoint(marker.getLatLng());
    point = point.add(offset);
    var latLng = this._map.containerPointToLatLng(point);
    marker.setLatLng(latLng);

    this._updateFieldOfView();
    this.fire('change');
  },

  _onMarkerKeyDown: function _onMarkerKeyDown(marker, evt) {
    var moveDelta = 20;
    if (evt.shiftKey) {
      moveDelta = moveDelta * 4;
    }

    if (evt.keyCode === 37) {
      // left
      this._moveMarker(marker, L.point(-moveDelta, 0));
    } else if (evt.keyCode === 38) {
      // up
      this._moveMarker(marker, L.point(0, -moveDelta));
    } else if (evt.keyCode === 39) {
      // right
      this._moveMarker(marker, L.point(moveDelta, 0));
    } else if (evt.keyCode === 40) {
      // down
      this._moveMarker(marker, L.point(0, moveDelta));
    }
  },

  _onAngleMarkerKeyDown: function _onAngleMarkerKeyDown(evt) {
    var angleDelta = 2.5;
    if (evt.shiftKey) {
      angleDelta = angleDelta * 4;
    }

    if (evt.keyCode === 37) {
      // left
      this.setAngle(this._angle - angleDelta);
    } else if (evt.keyCode === 39) {
      // right
      this.setAngle(this._angle + angleDelta);
    }
  },

  _onDocumentKeyDown: function _onDocumentKeyDown(evt) {
    if (document.activeElement === this._cameraMarker._icon) {
      this._onMarkerKeyDown(this._cameraMarker, evt);
    } else if (document.activeElement === this._targetMarker._icon) {
      this._onMarkerKeyDown(this._targetMarker, evt);
    } else if (document.activeElement === this._angleMarker._icon) {
      this._onAngleMarkerKeyDown(evt);
    }
  },

  _setMarkerVisible: function _setMarkerVisible(marker, visible) {
    marker._icon.style.display = visible ? 'inherit' : 'none';
  },

  _geoJsonPoint: function _geoJsonPoint(latLng) {
    return {
      type: 'Point',
      coordinates: [latLng.lng, latLng.lat]
    };
  },

  getFieldOfView: function getFieldOfView() {
    return this._fieldOfView;
  },

  getCameraLatLng: function getCameraLatLng() {
    return this._cameraMarker.getLatLng();
  },

  getTargetLatLng: function getTargetLatLng() {
    return this._targetMarker.getLatLng();
  },

  getCameraPoint: function getCameraPoint() {
    return this._geoJsonPoint(this.getCameraLatLng());
  },

  getTargetPoint: function getTargetPoint() {
    return this._geoJsonPoint(this.getTargetLatLng());
  },

  getCenter: function getCenter() {
    if (!this._map) {
      return;
    }

    return L.latLngBounds([this.getCameraLatLng(), this.getTargetLatLng()]).getCenter();
  },

  centerBounds: function centerBounds(bounds) {
    var cameraBounds = this.getBounds();

    if (!bounds.contains(cameraBounds)) {
      var center = this.getCenter();
      var cameraLatLng = this.getCameraLatLng();
      var targetLatLng = this.getTargetLatLng();

      var boundsCenter = bounds.getCenter();

      var newCameraLatLng = [boundsCenter.lat - (center.lat - cameraLatLng.lat), boundsCenter.lng - (center.lng - cameraLatLng.lng)];

      var newTargetLatLng = [boundsCenter.lat - (center.lat - targetLatLng.lat), boundsCenter.lng - (center.lng - targetLatLng.lng)];

      this.setCameraAndTargetLatLng(newCameraLatLng, newTargetLatLng);
    }
  },

  setCameraLatLng: function setCameraLatLng(latLng) {
    if (!this._map) {
      return;
    }

    this._cameraMarker.setLatLng(latLng);
    this._updateFieldOfView();
    this.fire('change');
  },

  setTargetLatLng: function setTargetLatLng(latLng) {
    if (!this._map) {
      return;
    }

    this._targetMarker.setLatLng(latLng);
    this._updateFieldOfView();
    this.fire('change');
  },

  setCameraAndTargetLatLng: function setCameraAndTargetLatLng(cameraLatLng, targetLatLng) {
    if (!this._map) {
      return;
    }

    this._cameraMarker.setLatLng(cameraLatLng);
    this._targetMarker.setLatLng(targetLatLng);
    this._updateFieldOfView();
    this.fire('change');
  },

  getBounds: function getBounds() {
    if (!this._fieldOfView) {
      return;
    }

    var pointList = this._getPointList(this._fieldOfView);
    return L.latLngBounds(pointList);
  },

  setAngle: function setAngle(angle) {
    this._angle = Math.max(Math.min(angle, this.options.maxAngle), this.options.minAngle);
    this._updateFieldOfView();
    this.fire('input');
  },

  setDraggable: function setDraggable(draggable) {
    if (!this._map) {
      return;
    }

    if (draggable) {
      this._cameraMarker.dragging.enable();
      this._targetMarker.dragging.enable();
      this._angleMarker.dragging.enable();
    } else {
      this._cameraMarker.dragging.disable();
      this._targetMarker.dragging.disable();
      this._angleMarker.dragging.disable();
    }

    this._setMarkerVisible(this._targetMarker, draggable);
    this._setMarkerVisible(this._angleMarker, draggable && this.options.angleMarker);
  }

});

// Object.assign polyfill, for IE<=11. From:
//   https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/assign#Polyfill
// TODO: I'm sure Babel can add this polyfill, too.
if (typeof Object.assign !== 'function') {
  Object.assign = function (target, varArgs) {
    if (target == null) {
      throw new TypeError('Cannot convert undefined or null to object');
    }

    var to = Object(target);

    for (var index = 1; index < arguments.length; index++) {
      var nextSource = arguments[index];

      if (nextSource != null) {
        for (var nextKey in nextSource) {
          if (Object.prototype.hasOwnProperty.call(nextSource, nextKey)) {
            to[nextKey] = nextSource[nextKey];
          }
        }
      }
    }
    return to;
  };
}

L.geotagPhoto = {
  crosshair: function crosshair(options) {
    return new GeotagPhotoCrosshair(options);
  },
  camera: function camera(feature, options) {
    return new GeotagPhotoCamera(feature, options);
  }
};

L.GeotagPhoto = {
  Crosshair: GeotagPhotoCrosshair,
  Camera: GeotagPhotoCamera
};

}(L));
