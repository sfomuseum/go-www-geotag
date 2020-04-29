var geotag = geotag || {};

geotag.writer = (function(){

    var self = {

	'write_geotag': function(id, f, on_success, on_error){

	    var enc_id = encodeURIComponent(id);
	    var endpoint = "/update?id=" + enc_id;	// FIX ME: read path from data attribute or... ?

	    var crumb = document.body.getAttribute("data-crumb");
	    
	    if (crumb){
		var enc_crumb = encodeURIComponent(crumb);
		endpoint = endpoint + "&crumb=" + enc_crumb;
	    }

	    geotag.api.execute_method('PUT', endpoint, f, on_success, on_error);
	}
    };
    
    return self;
    
})();
