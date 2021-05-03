window.addEventListener("load", function load(event){

    // PROTOMAPS: TBD is whether or not we want to be looking for or validating
    // map-related properties here rather than in geotag.maps.js - the reason
    // being that we are introducing support for multiple map renderers: Tangram.js
    // + Nextzen or Protomaps. Support for these renderers is handled by their
    // respective go-http-leaflet-{RENDERER} middleware packages. As of this writing
    // the Protomaps handlers don't append properties to data.body yet. Regardless
    // the flag to signal which renderer to use is specific to this application
    // and currently being assigned to the <div id="map"> element and processed in
    // geotag.maps.js (20210423/thisisaaronland)
    
    var api_key = document.body.getAttribute("data-nextzen-api-key");
    var style_url = document.body.getAttribute("data-nextzen-style-url");
    var tile_url = document.body.getAttribute("data-nextzen-tile-url");    

    /*
    if (! api_key){
	console.log("Missing API key");
	return;
    }
    
    if (! style_url){
	console.log("Missing style URL");
	return;
    }
    
    if (! tile_url){
	console.log("Missing tile URL");
	return;
    }
     */
    
    var init_lat = document.body.getAttribute("data-initial-latitude");

    if (! init_lat){
	console.log("Missing initial latitude");
	return;
    }
    
    var init_lon = document.body.getAttribute("data-initial-longitude");

    if (! init_lon){
	console.log("Missing initial longitude");
	return;
    }
    
    var init_zoom = document.body.getAttribute("data-initial-zoom");    

    if (! init_zoom){
	console.log("Missing initial zoom");
	return;
    }
    
    var map_el = document.getElementById("map");

    if (! map_el){
	console.log("Missing map element");	
	return;
    }

    var map_args = {
	"api_key": api_key,
	"style_url": style_url,
	"tile_url": tile_url,
    };

    // we need to do this _before_ Tangram starts trying to draw things
    // map_el.style.display = "block";
    
    var map = geotag.maps.getMap(map_el, map_args);

    if (! map){
	console.log("Unable to instantiate map");
	return;
    }

    // START OF geocoder/placeholder stuff
    
    var ph_endpoint = document.body.getAttribute("data-placeholder-endpoint");
    
    if (ph_endpoint){

	var ph_opts = {
	    serviceUrl: ph_endpoint
	};
	
	var ph = new Placeholder(ph_opts);
	
	var geocoder_opts = {
	    geocoder: ph,
	    defaultMarkGeocode: false,
	};
	
	var geocoder = L.Control.geocoder(geocoder_opts);
	
	geocoder.on('markgeocode', function(e){
	    
	    var bbox = e.geocode.bbox;		    
	    map.fitBounds(bbox);
	    
	    var camera = geotag.camera.getCamera();

	    if (camera){		
		var center = e.geocode.center;
		camera.setCameraLatLng(center);
		camera.setTargetLatLng(center);		
	    }
	});
	
	geocoder.addTo(map);
    }
    
    // END OF geocoder/placeholder stuff
    
    var hash = new L.Hash(map);

    var hash_str = location.hash;

    if (hash_str){

	var parsed = geotag.maps.parseHash(hash_str);

	if (parsed){
	    init_lat = parsed['latitude'];
	    init_lon = parsed['longitude'];
	    init_zoom = parsed['zoom'];
	}
    }
    
    map.setView([init_lat, init_lon], init_zoom);

    // START OF point in polygon
    
    var ed = document.getElementById("editor");

    var pip_enabled = false;

    if (ed.getAttribute("data-point-in-polygon") == "enabled"){
	pip_enabled = true;
    }

    var pip_layer = "";
    
    var pip_update = function(){

	var f = camera.getFieldOfView();
	
	var geoms = f.geometry.geometries;
	var pos = geoms[0];
	
	var pip_q = {
	    "longitude": pos.coordinates[0],	    
	    "latitude": pos.coordinates[1],
	};
	
	if (pip_inception.value != ""){
	    pip_q.inception = pip_inception.value;
	}
	
	if (pip_cessation.value != ""){
	    pip_q.cessation = pip_cessation.value;
	}
	
	geotag.pointinpolygon.query(pip_q, pip_onsuccess, pip_onerror);
	pip_reset_candidates();
    };
    
    var pip_candidates = document.getElementById("point-in-polygon-candidates");
    
    var pip_inception = document.getElementById("point-in-polygon-inception");
    var pip_cessation = document.getElementById("point-in-polygon-cessation");
    
    if (pip_inception){
	pip_inception.onchange = pip_update;
    }

    if (pip_cessation){
	pip_cessation.onchange = pip_update;
    }
    
    var pip_onsuccess = function(rsp){

	// pip_candidates.style.display = "none";
	// pip_candidates.innerHTML = "";
	
	var places = rsp.places;
	var count = places.length;

	var options = {};
	var labels = [];

	for (var i=0; i < count; i++){

	    var pl = places[i];

	    var id = pl["wof:id"];
	    var name = pl["wof:name"];
	    var placetype = pl["wof:placetype"];
	    var inception = pl["edtf:inception"];
	    var cessation = pl["edtf:cessation"];

	    var label = name + " (" + inception + "/" + cessation + ")";

	    labels.push(label);
	    options[label] = id;
	}

	labels = labels.sort();
	
	var sel = document.createElement("select");
	sel.setAttribute("id", "point-in-polygon-parent-id");

	var opt = document.createElement("option");
	sel.appendChild(opt);
	
	for (var i=0; i < count; i++){

	    var label = labels[i];
	    var id = options[label];
	    
	    var opt = document.createElement("option");
	    opt.setAttribute("value", id);

	    opt.appendChild(document.createTextNode(label));
	    sel.appendChild(opt);
	}

	sel.onchange = function(e){
	    
	    var el = e.target;
	    var str_id = el.value;

	    // Remove any existing wof: properties
	    // ... is a special flag to remove this property

	    if (str_id == ""){
		
		var props = {
		    "wof:parent_id":"...",
		    "wof:hierarchy":"...",
		};

		update_feature_properties(props);
		return;
	    }
	    
	    var parent_id = parseInt(str_id);

	    var props = {
		    "wof:parent_id": parent_id,		    
	    };

	    update_feature_properties(props);

            // fetch the hierarchy for this feature asychronously		
	    geotag.data.fetch(parent_id, data_onsuccess, data_onerror);
	};

	var label = document.createElement("label");
	label.setAttribute("for", "point-in-polygon-parent-id");
	label.appendChild(document.createTextNode("Candidates"));

	pip_candidates.innerHTML = "";
	pip_candidates.appendChild(label);
	pip_candidates.appendChild(sel);

	pip_candidates.style.display = "block";
    };
    
    var pip_onerror  = function(err){	
	console.log("PIP ERROR", err);
	pip_reset_candidates();
    };

    var data_onsuccess = function(f){

	if (pip_layer){
	    map.removeLayer(pip_layer);
	}

	pip_layer = L.geoJSON(f);
	pip_layer.addTo(map);

	pip_layer.bringToBack();
	
	var props = {
	    "wof:hierarchy": f.properties["wof:hierarchy"],  
	};
	
	update_feature_properties(props);
    };

    var data_onerror = function(err){
	console.log("ERROR", err);
    };

    var pip_reset_candidates = function(){

	var el = document.getElementById("point-in-polygon-parent-id");

	if (el){
	    el.innerHTML = "";
	}
    };
    
    // END OF point in polygon    
    
    var camera = geotag.camera.getCamera();
    camera.addTo(map);
    
    camera.setCameraLatLng([init_lat, init_lon]);
    camera.setTargetLatLng([init_lat, init_lon]);    

    var last_pos;
    
    var on_update = function(e){
	
	var f = camera.getFieldOfView();

	// START OF point in polygon

	if (pip_enabled){

	    var this_pos = JSON.stringify(f.geometry.geometries[0]);

	    if ((! last_pos) || (last_pos != this_pos)){
		
		if (pip_layer){
		    map.removeLayer(pip_layer);
		}
	    	
		pip_update();
	    }

	    last_pos = this_pos;
	}
	
	// END OF point in polygon	
	
	var el = document.getElementById("feature");

	if (el){
	    el.innerHTML = "";
	    el.appendChild(render_feature(f));
	}

	var out = document.getElementById("output");

	if (out){
	    out.style.display = "block";
	}

	var save = document.getElementById("writer-save");

	if (save){
	    save.style.display = "block";
	}
	
    };

    var render_feature = function(f){

	var enc = JSON.stringify(f, null, 2);
	var pre = document.createElement("pre");
	pre.setAttribute("id", "feature-body");
	pre.appendChild(document.createTextNode(enc));

	return pre;
    };

    var update_feature_properties = function(props){

	var wrapper_el = document.getElementById("feature");	
	var body_el = document.getElementById("feature-body");
	
	var str_f = body_el.innerText;
	var f = JSON.parse(str_f);

	for (var k in props){

	    var v = props[k];

	    // ... is a special flag to remove this property
	    
	    if (v == "..."){
		delete f.properties[k];
		continue;
	    }

	    f.properties[k] = props[k];
	}

	body_el = render_feature(f);

	wrapper_el.innerHTML = "";
	wrapper_el.appendChild(body_el);
    };
    
    camera.on('change', on_update);
    camera.on('input', on_update);
});
