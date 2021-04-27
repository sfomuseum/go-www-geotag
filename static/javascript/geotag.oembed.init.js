window.addEventListener("load", function load(event){

    var endpoints_str = document.body.getAttribute("data-oembed-endpoints");    
    var endpoints_map = geotag.oembed.endpoints_map_from_string(endpoints_str);
    
    if (! endpoints_map){
	console.log("Unabled to build endpoints map");
	return;
    }

    geotag.oembed.set_endpoints_map(endpoints_map);
    
    var q_el = document.getElementById("oembed-url");

    if (! q_el){
	console.log("Missing oembed-url element.");
	return;
    }

    var i_el = document.getElementById("oembed-image");

    if (! i_el){
	console.log("Missing oembed-image element.");
	return;
    }

    var m_el = document.getElementById("oembed-meta");

    if (! m_el){
	console.log("Missing oembed-meta element.");
	return;
    }
    
    var f_el = document.getElementById("oembed-fetch");

    if (! f_el){
	console.log("Missing oembed-fetch element.");
	return;
    }

    var handle_geotag_props = function(props){

	var camera = geotag.camera.getCamera();
	
	var camera_lat = props["geotag:camera_latitude"];
	var camera_lon = props["geotag:camera_longitude"];	    

	camera_lat = parseFloat(camera_lat);
	camera_lon = parseFloat(camera_lon);	    

	if ((camera_lat) && (camera_lon)){
	    camera.setCameraLatLng([camera_lat, camera_lon]);
	}
	
	var target_lat = props["geotag:target_latitude"];
	var target_lon = props["geotag:target_longitude"];	    
	
	target_lat = parseFloat(target_lat);
	target_lon = parseFloat(target_lon);	    

	if ((target_lat) && (target_lon)){
	    camera.setTargetLatLng([target_lat, target_lon]);
	}
	
	var angle = props["geotag:angle"];
	angle = parseFloat(angle);
	
	if (angle){
	    camera.setAngle(angle);
	}
    };
    
    f_el.onclick = function(e){

	var url = q_el.value;
	
	var on_success_geojson = function(feature){
	    var props = feature["properties"];
	    handle_geotag_props(props);
	};

	var on_error_geojson = function(err){
	    console.log("SAD GEOJSON", err);
	};
	
	var on_success = function(rsp){

	    i_el.innerHTML = "";
	    m_el.innerHTML = "";	
	    
	    var img = document.createElement("img");
	    img.setAttribute("src", rsp["url"]);

	    i_el.appendChild(img);
	    i_el.style.display = "block";
	    
	    var title = rsp["title"];

	    var title = document.createElement("a");
	    title.setAttribute("href", rsp["author_url"]);
	    title.appendChild(document.createTextNode(rsp["title"]));

	    m_el.appendChild(title);
	    m_el.style.display = "block";

	    var col = document.getElementById("col-oembed");
	    col.style.display = "block";
	    
	    // 

	    var uri = rsp["geotag:uri"];
	    document.body.setAttribute("data-geotag-uri", uri);
	    
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

	i_el.style.display = "none";
	m_el.style.display = "none";

	return false;
    };

    var query_str = location.search;

    if (query_str.indexOf("?") === 0) {
	query_str = query_str.substr(1);
    }

    var query_pairs = query_str.split("&");
    var count_pairs = query_pairs.length;

    var query = {};
    
    for (var i=0; i < count_pairs; i++){
	
	var str_pair = query_pairs[i];
	var pair = str_pair.split("=");

	var key = pair[0];
	var value = pair[1];

	var values = query[key];

	if (! values){
	    values = [];
	}

	values.push(value);
	query[key] = values;
    }

    if (query["oembed-url"]){

	var str_url = query["oembed-url"][0];
	var url = new URL(str_url);

	if (url){
	    q_el.value = url.toString();
	    f_el.click();
	} else {
	    console.log("Invalid oembed-url parameter");
	}
    }   

});
