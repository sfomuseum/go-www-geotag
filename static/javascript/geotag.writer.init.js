window.addEventListener("load", function load(event){

    var save = document.getElementById("writer-save");

    if (! save){
	console.log("Unable to load save button");
	return;
    }

    save.onclick = function(){

	var camera = geotag.camera.getCamera();

	if (! camera){
	    console.log("Unable to retrieve camera");
	    return false;
	}

	var id = "fixme";
	var fov = camera.getFieldOfView();

	var on_success = function(rsp){
	    console.log("WRITE OKAY", rsp);
	};

	var on_error = function(err){
	    console.log("WRITE ERROR", err);
	};

	geotag.writer.write_geotag(id, fov, on_success, on_error);
	return false;
    };
    
});
