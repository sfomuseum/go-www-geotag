  var Placeholder = /*#__PURE__*/function () {
    function Placeholder(options) {
      this.options = {
        serviceUrl: 'http://localhost:3000'
      };
      this._lastSuggest = 0;
      L.Util.setOptions(this, options);
    }

    var _proto = Placeholder.prototype;

    _proto.geocode = function geocode(query, cb, context) {
      var _this = this;

      var params = geocodingParams(this.options, {
        text: query
      });
      getJSON(this.options.serviceUrl + '/search', params, function (data) {
        cb.call(context, _this._parseResults(data, 'bbox'));
      });
    };

    _proto.suggest = function suggest(query, cb, context) {
      var _this2 = this;

      var params = geocodingParams(this.options, {
        api_key: this.options.apiKey,
        text: query
      });
      getJSON(this.options.serviceUrl + '/autocomplete', params, function (data) {
        if (data.geocoding.timestamp > _this2._lastSuggest) {
          _this2._lastSuggest = data.geocoding.timestamp;
          cb.call(context, _this2._parseResults(data, 'bbox'));
        }
      });
    };

    _proto.reverse = function reverse(location, scale, cb, context) {
      var _this3 = this;

      var params = reverseParams(this.options, {
        api_key: this.options.apiKey,
        'point.lat': location.lat,
        'point.lon': location.lng
      });
      getJSON(this.options.serviceUrl + '/reverse', params, function (data) {
        cb.call(context, _this3._parseResults(data, 'bounds'));
      });
    };

    _proto._parseResults = function _parseResults(data, bboxname) {
      var results = [];
      L.geoJSON(data, {
        pointToLayer: function pointToLayer(feature, latlng) {
          return L.circleMarker(latlng);
        },
        onEachFeature: function onEachFeature(feature, layer) {
          var result = {};
          var bbox;
          var center;

          if (layer.getBounds) {
            bbox = layer.getBounds();
            center = bbox.getCenter();
          } else if (layer.feature.bbox) {
            center = layer.getLatLng();
            bbox = L.latLngBounds(L.GeoJSON.coordsToLatLng(layer.feature.bbox.slice(0, 2)), L.GeoJSON.coordsToLatLng(layer.feature.bbox.slice(2, 4)));
          } else {
            center = layer.getLatLng();
            bbox = L.latLngBounds(center, center);
          }

          result.name = layer.feature.properties.label;
          result.center = center;
          result[bboxname] = bbox;
          result.properties = layer.feature.properties;
          results.push(result);
        }
      });
      return results;
    };

    return Placeholder;
  }();
  /**
   * [Class factory method](https://leafletjs.com/reference.html#class-class-factories) for {@link Placeholder}
   * @param options the options
   */

  function placeholder(options) {
    return new Placeholder(options);
  }
