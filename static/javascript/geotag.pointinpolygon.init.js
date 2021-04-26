window.addEventListener("load", function load(event){

    var endpoint = document.body.getAttribute("data-point-in-polygon-endpoint");

    if (! endpoint){
	return;
    }
    
    geotag.pointinpolygon.set_endpoint(endpoint);
});
