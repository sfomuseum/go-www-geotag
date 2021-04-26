window.addEventListener("load", function load(event){

    var ed = document.getElementById("editor");

    if (! ed){
	return;
    }
    
    var endpoint = ed.getAttribute("data-point-in-polygon-endpoint");

    if (! endpoint){
	return;
    }

    geotag.pointinpolygon.set_endpoint(endpoint);    
});
