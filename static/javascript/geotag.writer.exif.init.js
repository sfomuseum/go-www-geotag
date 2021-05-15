window.addEventListener("load", function load(event){

    if (! WebAssembly.instantiateStreaming){
	 
	WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer();
            return await WebAssembly.instantiate(source, importObject);
	};
    }

    const update_go = new Go();
    let update_mod, update_inst;

    WebAssembly.instantiateStreaming(fetch("/wasm/update_exif.wasm"), update_go.importObject).then(
	
	async result => {
	    
	    enable();
	    
            update_mod = result.module;
            update_inst = result.instance;
	    await update_go.run(update_inst);
	}
    );
    
    async function enable() {

	var save = document.getElementById("writer-save");

	if (! save){
	    console.log("Unable to load save button");
	    return;
	}
	
	save.onclick = function(){

	    var wrapper_el = document.getElementById("feature");	
	    var body_el = document.getElementById("feature-body");
	
	    var str_f = body_el.innerText;
	    var f = JSON.parse(str_f);
	    
	    var uri = document.body.getAttribute("data-geotag-uri", uri);
	    
	    if (! uri){
		console.log("Missing data-geotag-uri attribute");
		// return false;
	    }

	    var base = uri.split('/').reverse()[0];
	    
	    geotag.writer.exif.write_geotag(base, f);
	    return false;
	};

	save.removeAttribute("disabled");	
    }


});
