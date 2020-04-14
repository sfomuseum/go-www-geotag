var geotag = geotag || {};

geotag.writer = (function(){

    var self = {

	'write_geotag': function(f, on_success, on_error){

	    var endpoint = "/update";	// FIX ME

	    var crumb = document.body.getAttribute("data-crumb");
	    
	    if (crumb){
		var enc_crumb = encodeURIComponent(crumb);
		endpoint = endpoint + "?crumb=" + enc_crumb;
	    }

	    geotag.api.execute_method('PUT', endpoint, f, on_success, on_error);
	}
    };
    
    return self;
    
})();
