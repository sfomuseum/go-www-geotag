var geotag = geotag || {};
geotag.writer = geotag.writer || {};

geotag.writer.exif = (function(){

    var self = {
	
	'write_geotag': function(f){

	    var geoms = f.geometry.geometries;
	    var pt = geoms[0];
	    
	    var update = {
		"X-Latitude": pt.coordinates[0],
		"X-Longitude": pt.coordinates[1],
	    };
	    
	    var enc_update;
	    
	    try {
		enc_update = JSON.stringify(update);
	    } catch(e){
		console.log("Failed to marhsal update", update, e);
		return false;
	    }
    
	    var img = document.getElementById("image");
	    
	    var canvas = document.createElement("canvas");
	    canvas.width = img.width;
	    canvas.height = img.height;
	    var ctx = canvas.getContext("2d");
	    ctx.drawImage(img, 0, 0);
	    var b64_img = canvas.toDataURL("image/jpeg", 1.0);
	    
	    update_exif(b64_img, enc_update).then(data => {
		
		var blob = self.dataURLToBlob(data);
		
		if (! blob){
		    return false;
		}

		var fname = "example.jpg";	// FIX ME
		saveAs(blob, fname);
	
	    }).catch(err => {
		console.log("Failed to update EXIF data", err);
	    });
	},

	dataURLToBlob: function(dataURL){

	    var BASE64_MARKER = ";base64,";
	    
	    if (dataURL.indexOf(BASE64_MARKER) == -1){
		
		var parts = dataURL.split(",");
		var contentType = parts[0].split(":")[1];
		var raw = decodeURIComponent(parts[1]);
		
		return new Blob([raw], {type: contentType});
	    }
	    
	    var parts = dataURL.split(BASE64_MARKER);
	    var contentType = parts[0].split(":")[1];
	    var raw = window.atob(parts[1]);
	    var rawLength = raw.length;
	    
	    var uInt8Array = new Uint8Array(rawLength);
	    
	    for (var i = 0; i < rawLength; ++i) {
		uInt8Array[i] = raw.charCodeAt(i);
	    }
	    
	    return new Blob([uInt8Array], {type: contentType});
	}

    };

    return self;

})();
