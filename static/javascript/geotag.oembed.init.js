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
    
    f.onclick = function(e){

	var url = q.value;
	
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
