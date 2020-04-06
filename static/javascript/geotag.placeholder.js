var geotag = geotag || {};

geotag.placeholder = (function(){

    var _endpoint = null;
    
    var self = {

	'set_endpoint': function(url){
	    _endpoint = url;
	},

	'query': function(text, on_success, on_error) {

	    var q = {
		'text': text,
	    };

	    var q_str = self.query_string(q);
	    var url = _endpoint + 'query/?' + q_str;

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
	
	'query_string': function(args){

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
	
	'render_query_results': function(rsp){

	    var wrapper = document.createElement("div");

	    var count = rsp.length;

	    for (var i=0; i < count; i++){

		var value = rsp[i];
		
		var row = document.createElement("div");
		row.appendChild(document.createTextNode(value));
		wrapper.appendChild(row);
	    }

	    return wrapper;
	},
	
    };
    
    return self;
    
})();
