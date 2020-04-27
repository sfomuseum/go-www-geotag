var geotag = geotag || {};

geotag.oembed = (function(){

    var _endpoints = {};
    
    var self = {

	'set_endpoints_map': function(endpoints_map){
	    _endpoints = endpoints_map;
	},
	
	'endpoints_map_from_string': function(endpoints_str){

	    var endpoints_map = {};
    
	    var endpoints = endpoints_str.split(",");
	    var count = endpoints.length;

	    for (var i=0; i < count; i++){

		var endpoint = endpoints[i];
		var url = new URL(endpoint);
		
		if (! url){
		    console.log("Invalid oEmbed endpoint", endpoint);
		    return;
		}

		var hostname = url.hostname;
		endpoints_map[hostname] = endpoint;
	    }

	    return endpoints_map;
	},

	'derive_endpoint': function(url){

	    var parsed = new URL(url);

	    if (! parsed){
		return null;
	    }

	    var hostname = parsed.hostname;
	    return _endpoints[hostname];
	},
	
	'fetch': function(url, on_success, on_error) {

	    // URLs with #hash compotents trigger CORS errors
	    // I don't know why... (20200427/thisisaaronland)
	    
	    var u = new URL(url);

	    if (! u){
		on_error("Invalid URL");
		return;
	    }
	    
	    url = u.href.replace(u.hash, "")

	    var endpoint_template = self.derive_endpoint(url);

	    if (! endpoint_template){
		on_error("Unable able to derive oEmbed endpoint");
		return;
	    }

	    var enc_url = encodeURIComponent(url);
	    var oembed_url = endpoint_template.replace("{url}", enc_url);

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
	    
	    req.open("get", oembed_url, true);
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
