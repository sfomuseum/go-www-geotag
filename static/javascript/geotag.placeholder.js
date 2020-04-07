var geotag = geotag || {};

geotag.placeholder = (function(){

    var _endpoint = null;
    
    var self = {

	'set_endpoint': function(url){
	    _endpoint = url;
	},

	'search': function(text, on_success, on_error) {

	    var q = {
		'text': text,
	    };

	    var q_str = self.build_query_string(q);
	    var url = _endpoint + 'search/?' + q_str;

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
	    
	    req.open("get", url, true);
	    req.send();	    	    
	},
	
	'build_query_string': function(args){

	    var pairs = [];

	    for (var k in args){

		var v = args[k];

		var enc_k = encodeURIComponent(k);
		var enc_v = encodeURIComponent(v);
		
		var pair = enc_k + "=" + enc_v;
		pairs.push(pair);
	    }

	    return pairs.join("&");
	},
		
    };
    
    return self;
    
})();
