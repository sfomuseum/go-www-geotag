var geotag = geotag || {};

geotag.writer = (function(){

    var self = {

	'write_geotag': function(id, f, on_success, on_error){

	    var endpoint = "/update";	// FIX ME
	    geotag.api.execute_method('PUT', endpoint, f, on_success, on_error);
	}
    };
    
    return self;
    
})();
