window.addEventListener("load", function load(event){

    var init_lat = document.body.getAttribute("data-initial-latitude");
    var init_lon = document.body.getAttribute("data-initial-longitude");

    geotag.camera.init(init_lat, init_lon);
});
