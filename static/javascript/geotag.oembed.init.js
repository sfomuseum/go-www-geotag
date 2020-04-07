window.addEventListener("load", function load(event){
    
    var q = document.getElementById("oembed-url");

    if (! q){
	console.log("Missing oembed-url element.");
	return;
    }

    var i = document.getElementById("oembed-image");

    if (! i){
	console.log("Missing oembed-image element.");
	return;
    }

    var f = document.getElementById("oembed-fetch");

    if (! f){
	console.log("Missing oembed-fetch element.");
	return;
    }
    
    f.onclick = function(e){

	var el = e.target;
	var url = el.value;

	var on_success = function(rsp){
	    console.log("OKAY", rsp);
	};

	var on_error = function(err){
	    console.log("ERROR", err);
	};

	var endpoint = "https://millsfield.sfomuseum.org/oembed/";
	
	geotag.oembed.fetch(endpoint, url, on_success, on_error);
	return false;
    };
    
});
