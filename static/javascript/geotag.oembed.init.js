window.addEventListener("load", function load(event){

    var endpoints_str = document.body.getAttribute("data-oembed-endpoints");    
    var endpoints_map = geotag.oembed.endpoints_map_from_string(endpoints_str);
    
    if (! endpoints_map){
	console.log("Unabled to build endpoints map");
	return;
    }

    geotag.oembed.set_endpoints_map(endpoints_map);
    
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

    var m = document.getElementById("oembed-meta");

    if (! i){
	console.log("Missing oembed-meta element.");
	return;
    }
    
    var f = document.getElementById("oembed-fetch");

    if (! f){
	console.log("Missing oembed-fetch element.");
	return;
    }

    var handle_geotag_props = function(props){
	
	var camera = geotag.camera.getCamera();
	
	var camera_lat = props["geotag:camera_latitude"];
	var camera_lon = props["geotag:longitude"];	    
	
	camera_lat = parseFloat(camera_lat);
	camera_lon = parseFloat(camera_lon);	    
	
	if ((camera_lat) && (camera_lon)){		    
	    camera.setCameraLatLng([camera_lat, camera_lon]);
	}
	
	var target_lat = props["geom:target_latitude"];
	var target_lon = props["geom:target_longitude"];	    
	
	camera_lat = parseFloat(camera_lat);
	camera_lon = parseFloat(camera_lon);	    
	
	if ((target_lat) && (target_lon)){		    
	    target.setTargetLatLng([target_lat, target_lon]);
	}
	
	var angle = props["geotag:angle"];
	angle = parseFloat(angle);
	
	if (angle){
	    camera.setAngle(angle);
	}
    };
    
    f.onclick = function(e){

	var url = q.value;
	
	var on_success_geojson = function(feature){
	    var props = feature["properties"];
	    handle_geotag_props(props);
	};

	var on_error_geojson = function(err){
	    console.log("SAD GEOJSON", err);
	};
	
	var on_success = function(rsp){

	    i.innerHTML = "";
	    m.innerHTML = "";	
	    
	    var img = document.createElement("img");
	    img.setAttribute("height", rsp["height"]);
	    img.setAttribute("width", rsp["width"]);	    
	    img.setAttribute("src", rsp["url"]);
	    img.setAttribute("class", "card-img-top");

	    i.appendChild(img);
	    i.style.display = "block";
	    
	    var title = rsp["title"];

	    var title = document.createElement("a");
	    title.setAttribute("href", rsp["author_url"]);
	    title.appendChild(document.createTextNode(rsp["title"]));

	    m.appendChild(title);
	    m.style.display = "block";

	    if (rsp["geotag:geojson_url"]){
		whosonfirst.net.fetch(rsp["geotag:geojson_url"], on_success_geojson, on_error_geojson);
		return;
	    }

	    handle_geotag_props(rsp);
	};

	var on_error = function(err){
	    console.log("ERROR", err);
	};

	geotag.oembed.fetch(url, on_success, on_error);

	i.style.display = "none";
	m.style.display = "none";

	return false;
    };
    
});
