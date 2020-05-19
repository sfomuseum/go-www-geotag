var geotag = geotag || {};

geotag.camera = (function(){

    var _camera = null;

    var self = {

	'init': function(lat, lon){

	    lat = parseFloat(lat);
	    lon = parseFloat(lon);
	    
	    if (! lat){
		lat = 0.0;
	    }

	    if (! lon){
		lon = 0.0;
	    }

	    self.initCamera(lat, lon);
	},
	
	'getCamera': function(){
	    return _camera;
	},

	'initCamera': function(lat, lon){

	    var cameraPoint = [lon, lat]
	    var targetPoint = [lon, lat]
	    
	    var points = {
		type: 'Feature',
		properties: {
		    angle: 20
		},
		geometry: {
		    type: 'GeometryCollection',
		    geometries: [
			{
			    type: 'Point',
			    coordinates: cameraPoint
			},
			{
			    type: 'Point',
			    coordinates: targetPoint
			}
		    ]
		}
	    };
	    
	    _camera = L.geotagPhoto.camera(points, {
		minAngle: 10
	    });
	},
	
	'setLatLon': function(lat, lon){

	    console.log("SET", lat, lon);
	    
	    _camera.setCameraLatLng([lat, lon]);
	    _camera.setTargetLatLng([lat, lon]);

	    console.log("WTF", _camera.getCameraLatLng())
	    return _camera;
	},
	
    };

    return self;
    
})();
