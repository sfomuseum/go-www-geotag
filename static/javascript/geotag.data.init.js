window.addEventListener("load", function load(event){

    var ed = document.getElementById("editor");

    if (! ed){
	return;
    }
    
    var endpoint = ed.getAttribute("data-point-in-polygon-data-endpoint");

    if (! endpoint){
	return;
    }

    geotag.data.set_endpoint(endpoint);    
});
