var geotag = geotag || {};

geotag.pointinpolygon = (function(){

    var _endpoint = null;
    
    var self = {

	'set_endpoint': function(url){
	    _endpoint = url;
	},

	'query': function(args, on_success, on_error) {

	    var enc_args = JSON.stringify(args);
	    
	    var url = _endpoint;

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
	    
	    req.open("POST", url, true);
	    req.send(args);
	},
			
    };
    
    return self;
    
})();
