var geotag = geotag || {};
geotag.writer = geotag.writer || {};

geotag.writer.api = (function(){

    var ed = document.getElementById("editor");
    var wr_path = ed.getAttribute("data-writer-path");
    
    var self = {

	'write_geotag': function(id, f, on_success, on_error){
	    
	    var enc_id = encodeURIComponent(id);
	    var endpoint = wr_path + "?id=" + enc_id;

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
