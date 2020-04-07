var geotag = geotag || {};

geotag.camera = (function(){

    var _camera = null;

    var self = {

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

	    return _camera;
	},
	
	'setLatLon': function(lat, lon){

	    if (! _camera){
		return self.initCamera(lat, lon);
	    }

	    _camera.setCameraLatLng([lat, lon]);
	    return _camera;
	},
	
    };

    return self;
    
})();
