window.addEventListener("load", function load(event){

    var endpoint = document.body.getAttribute("data-placeholder-endpoint");

    if (! endpoint){
	return;
    }
    
    var q = document.getElementById("placeholder-query");

    if (! q){
	console.log("Missing placeholder-query element.");
	return;
    }

    var r = document.getElementById("placeholder-results");

    if (! r){
	console.log("Missing placeholder-results element.");
	return;
    }

    geotag.placeholder.set_endpoint(endpoint);
   
    q.onkeyup = function(e){

	var el = e.target;
	var text = el.value;
	
	if (text.length < 3){
	    return;
	}

	var render_search_results = function(rsp){	   

	    var count = rsp.length;

	    if (count == 0){
		return;
	    }

	    var wrapper = document.createElement("div");
	    
	    for (var i=0; i < count; i++){

		var row = rsp[i];
		var id = row["id"];
		var pt = row["placetype"];
		var name = row["name"];

		var geom = row["geom"];
		var lat = geom["lat"];
		var lon = geom["lon"];
		
		var row = document.createElement("div");
		row.setAttribute("class", "placeholder-row");
		row.setAttribute("data-whosonfirst-id", id);
		row.setAttribute("data-latitude", lat);
		row.setAttribute("data-longitude", lon);

		var value = name + " (" + id + ") " + pt;

		row.appendChild(document.createTextNode(value));

		row.onclick = function(e){
		    var el = e.target;		    
		    var id = el.getAttribute("data-whosonfirst-id");

		    var lat = el.getAttribute("data-latitude");
		    var lon = el.getAttribute("data-longitude");		    

		    var map = geotag.maps.getMapById("map");

		    if (map){
			map.setView([lat, lon], 13);
		    }

		    var camera = geotag.camera.getCamera();
		    
		    camera.setCameraLatLng([lat, lon]);
		    camera.setTargetLatLng([lat, lon]);    
		    
		    r.style.display = "none";
		    r.innerHTML = "";
		};
		
		wrapper.appendChild(row);
	    }

	    return wrapper;
	};
	
	var on_success = function(rsp){

	    var el = render_search_results(rsp);
	    
	    r.appendChild(el);
	    r.style.display = "block";	    
	};

	var on_error = function(err){
	    console.log("SAD", err);
	};

	geotag.placeholder.search(text, on_success, on_error);
	
	r.style.display = "none";
	r.innerHTML = "";
    };
});
