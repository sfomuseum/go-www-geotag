// https://github.com/perliedman/leaflet-control-geocoder/
// I am uncertain about nearly everything below because the docs for custom
// plugins are all geared towards modern/compiled JavaScript so it's a bit
// unclear how best to do things by hand (20210428/thisisaaronland)
var Placeholder = function () {
    
    function Placeholder(options) {
	this.options = {
            serviceUrl: 'http://localhost:3000'
	};
	L.Util.setOptions(this, options);
    }
    
    var _proto = Placeholder.prototype;
    
    _proto.geocode = function geocode(query, cb, context) {
	
	var _this = this;

	var hierarchy = [
	    "neighbourhood",
	    "locality",
	    "county",
	    "region",
	    "country",
	];
	
	var params = {
            text: query
	};
	
	this.getJSON(this.options.serviceUrl + '/parser/search', params, function (data){

	    var results = [];

		var count = data.length;
		
		for (var i=0; i < count; i++){
		    
		    var pl = data[i];
		    var result = {};
		    
		    var bounds = pl.geom.bbox.split(",");
		    
		    var swlon = parseFloat(bounds[0]);		
		    var swlat = parseFloat(bounds[1]);
		    var nelon = parseFloat(bounds[2]);		
		    var nelat = parseFloat(bounds[3]);
		    
		    bounds = L.latLngBounds([swlat, swlon], [nelat, nelon]);

		    var center = L.latLng( pl.geom.lat, pl.geom.lon );
		   
		    var placetype = pl.placetype;
		    var lineage = pl.lineage[0];
		    
		    var names = []
		    var ok = false;
		    
		    for (var j in hierarchy){
			
			var pt = hierarchy[j];
			
			if (pt == placetype){
			    ok = true;
			}
			
			if (!ok){
			    continue
			}

			if ((pt == "county") && (names.length)){
			    continue;
			}
			
			if (! lineage[pt]){
			    continue;
			}
			
			if ((names.length > 0) && (pt == "region") && (lineage[pt].abbr)){
			    names.push(lineage[pt].abbr);
			    continue;
			}
			
			names.push(lineage[pt].name);
		    }
		    
		    var name = names.join(", ");
		    
		    if (name == ""){
			name = pl.name;
		    }
		    
		    name = name + " (" + placetype + ")";
		    result.name = name;

		    result.center = center;
		    result.bbox = bounds;
		    
		    result.placetype = placetype;
		    result.id = pl.id;
		    
		    results.push(result);		
		}
	    
	    cb.call(context, results);
	});
    };

    _proto.suggest = function suggest(query, cb, context) {
	return this.geocode(query, cb, context);
    };

    _proto.reverse = function reverse(location, scale, cb, context) {
	cb.call(context, []);
    };
	 
    _proto.getJSON = function(url, params, callback) {
	
	var xmlHttp = new XMLHttpRequest();
	
	xmlHttp.onreadystatechange = function () {
	    if (xmlHttp.readyState !== 4) {
		return;
	    }
	    
	    var message;
	    
	    if (xmlHttp.status !== 200 && xmlHttp.status !== 304) {
		message = '';
	    } else if (typeof xmlHttp.response === 'string') {
		// IE doesn't parse JSON responses even with responseType: 'json'.
		    try {
			message = JSON.parse(xmlHttp.response);
		    } catch (e) {
			// Not a JSON response
			message = xmlHttp.response;
		    }
	    } else {
		message = xmlHttp.response;
	    }
	    
	    callback(message);
	};
	
	xmlHttp.open('GET', url + this.getParamString(params), true);
	xmlHttp.responseType = 'json';
	xmlHttp.setRequestHeader('Accept', 'application/json');
	xmlHttp.send(null);
    };

    _proto.getParamString = function(obj, existingUrl, uppercase) {
	  var params = [];
	
	for (var i in obj) {
	    var key = encodeURIComponent(uppercase ? i.toUpperCase() : i);
	    var value = obj[i];
	    
	    if (!Array.isArray(value)) {
		params.push(key + '=' + encodeURIComponent(String(value)));
	    } else {
		for (var j = 0; j < value.length; j++) {
		    params.push(key + '=' + encodeURIComponent(value[j]));
		}
	    }
	}
	
	return (!existingUrl || existingUrl.indexOf('?') === -1 ? '?' : '&') + params.join('&');
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
