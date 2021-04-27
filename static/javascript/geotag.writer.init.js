window.addEventListener("load", function load(event){

    var save = document.getElementById("writer-save");

    if (! save){
	console.log("Unable to load save button");
	return;
    }

    save.onclick = function(){

	var wrapper_el = document.getElementById("feature");	
	var body_el = document.getElementById("feature-body");
	
	var str_f = body_el.innerText;
	var f = JSON.parse(str_f);

	var uri = document.body.getAttribute("data-geotag-uri", uri);

	if (! uri){
	    console.log("Missing data-geotag-uri attribute");
	    // return false;
	}
	
	var on_success = function(rsp){
	    console.log("WRITE OKAY", rsp);
	};

	var on_error = function(err){
	    console.log("WRITE ERROR", err);
	};

	geotag.writer.write_geotag(uri, f, on_success, on_error);
	return false;
    };
    
});
