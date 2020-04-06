window.addEventListener("load", function load(event){

    var endpoint = document.body.getAttribute("data-search-endpoint");

    if (! endpoint){
	return;
    }
    
    var q = document.getElementById("placeholder-query");

    if (! q){
	console.log("Missing placeholder-query element.");
	return;
    }

    var r = document.getElementById("placeholder-results");

    if (! r){
	console.log("Missing placeholder-results element.");
	return;
    }

    geotag.placeholder.set_endpoint(endpoint);
    
    q.onkeyup = function(e){

	var el = e.target;
	var text = el.value;
	
	if (text.length < 3){
	    return;
	}

	var on_success = function(rsp){

	    console.log("SUCCESS", text, rsp);	    
	    var el = geotag.placeholder.render_query_results(rsp);
	    
	    r.appendChild(el);
	    r.style.display = "block";	    
	};

	var on_error = function(err){
	    console.log("SAD", err);
	};

	geotag.placeholder.query(text, on_success, on_error);
	r.style.display = "none";
	r.innerHTML = "";
    };
});
