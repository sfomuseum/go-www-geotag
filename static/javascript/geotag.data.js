var geotag = geotag || {};

geotag.data = (function(){

    var _endpoint = null;
    
    var self = {

	'set_endpoint': function(url){
	    _endpoint = url;
	},

	'fetch': function(id, on_success, on_error) {
	    
	    var url = _endpoint + id;

	    var req = new XMLHttpRequest();
	    
	    req.onload = function(){

		var rsp;
		
		try {
		    rsp = JSON.parse(this.responseText);
            	}
		
		catch (e){
		    on_error(e);
		    return false;
		}

		on_success(rsp);
       	    };

	    req.open("GET", url, true);
	    req.send();
	},
    };
    
    return self;
    
})();
