window.addEventListener("load", function load(event){

    var init_lat = document.body.getAttribute("data-initial-latitude");
    var init_lon = document.body.getAttribute("data-initial-longitude");

    console.log("INIT", init_lat, init_lon);
    geotag.camera.init(init_lat, init_lon);
});
