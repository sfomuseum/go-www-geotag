window.addEventListener("load", function load(event){

    var api_key = document.body.getAttribute("data-nextzen-api-key");
    var style_url = document.body.getAttribute("data-nextzen-style-url");
    var tile_url = document.body.getAttribute("data-nextzen-tile-url");    

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
});
