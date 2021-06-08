// version v0.0.1

L.Control.Layers = L.Control.extend({
    _map: null,
    _hash: null,
    _tile_layer: null,	// a leaflet.js tilelayer
    _layer: null,	// layer details (a row in options.catalog)
    _current: -1,
    _max_zoom: 18,
    _min_zoom: 1,
    options: {
	position: 'bottomright',
	catalog: [],
	on_change: null,
    },
    onAdd: function(map) {

	this._map = map;

	this._map.min_zoom = map.getMinZoom();
	this._map.max_zoom = map.getMaxZoom();

	this.div = L.DomUtil.create('div','leaflet-layers-container');

	this.select = L.DomUtil.create('select','leaflet-layers-select',this.div);

	var opt = L.DomUtil.create('option', '', this.select);
	opt.setAttribute("value", -1);
	opt.innerText = "";
	
	for (var idx in this.options.catalog){

	    var layer = this.options.catalog[idx];
	    var label = layer["label"];
	    
	    var opt = L.DomUtil.create('option', '', this.select);
	    opt.setAttribute("value", idx);
	    opt.innerText = label;
	}

	var _this = this;

	document.addEventListener('keydown', function(e){
		
	    var layer;
	    var offset;

	    if (e.keyCode == 37){
		layer = _this._getPreviousLayer();
	    }
	    
	    if (e.keyCode == 39){
		layer = _this._getNextLayer();	    
	    }	

	    if (layer){
		_this.setLayer(layer);
	    }
	});

        L.DomEvent.on(this.select, 'change', this._change, this);
	return this.div;
    },

    'setLayer': function(layer){

	if (this._tile_layer != null){
	    this._map.removeLayer(this._tile_layer);
	    this._tile_layer = null;

	    this._map.setMaxZoom(this._max_zoom);
	}

	var url = layer["url"];
	var args = {};	 // min,max zoom...

	if (layer["min_zoom"]){
	    args["minZoom"] = parseInt(layer["min_zoom"]);
	}

	if (layer["max_zoom"]){
	    args["maxZoom"] = parseInt(layer["max_zoom"]);
	}
	
	var tile_layer = L.tileLayer(url, args);
	tile_layer.addTo(this._map);

	var idx = this._getIndexWithLayer(layer);

	this._current = idx;
	this._layer = layer;
	this._tile_layer = tile_layer;

	if (this._map.getZoom() > args["maxZoom"]){
	    this._map.setZoom(args["maxZoom"]);
	}

	if (args["maxZoom"] <= this._max_zoom){

	    this._map.setMaxZoom(args["maxZoom"]);
	}

	this.select.selectedIndex = idx + 1;
	this._updatePermalink();

	// user defined stuff
	
	if (this.options.on_change){
	    this.options.on_change(layer);
	}
	
	return idx;
    },

    'addHash': function(on_parse, on_format){
	
	// see also:
	// https://github.com/mlevans/leaflet-hash/pull/10
	
	if (! L.hash){
	    return;
	}

	if (this._hash){
	    return this._hash;
	}

	var map = this._map;

	var hash = L.hash(map);
	var _this = this;

	hash.formatHash = function(map) {
	    
	    var center = map.getCenter();
	    var zoom = map.getZoom();
	    var precision = Math.max(0, Math.ceil(Math.log(zoom) / Math.LN2));
	    
	    var parts = [
		zoom,
		center.lat.toFixed(precision),
		center.lng.toFixed(precision),
	    ];

	    var current = _this._getCurrentLayer();

	    if (current){
		parts.unshift(current['label']);
	    }

	    if (on_format){
		
		var extras = on_format();

		if ((extras) && (extras.length > 0)){
		    
		    extras = extras.reverse();
		    
		    var count = extras.length;
		    
		    for (var i=0; i < count; i++){
			parts.unshift(extras[i]);
		    }
		}
	    }
	    
	    return "#" + parts.join("/");
	};
	    
	hash.parseHash = function(hash_str) {

	    if (hash_str.indexOf('#') === 0) {
		    hash_str = hash_str.substr(1);
	    }

	    var h = _this.parseHashString(hash_str);
	    
	    if (! h){
		return false;
	    }

	    var lat = h['latitude'];
	    var lon = h['longitude'];
	    var zoom = h['zoom'];
	    var label = h['label'];

	    if (label) {

		var layer = _this._getLayerWithLabel(label);

		if (layer){
		    _this.setLayer(layer);
		}
	    }
	    
	    if (zoom){
		_this._map.setZoom(zoom);
	    }

	    if ((! lat) || (! lon)){
		return;
	    }

	    return {
		center: new L.LatLng(lat, lon),
		zoom: zoom
	    };
	    
	};

	this._hash = hash;
	return hash;
    },

    'parseHashString': function(hash_str){
	
	if (hash_str.indexOf('#') === 0) {
	    hash_str = hash_str.substr(1);
	}

	var map = this._map;

	var center = map.getCenter();
	var zoom = map.getZoom();

	var lat = center[1];
	var lon = center[0];
	var label;

	var update_position = false;

	var args = hash_str.split("/");
		
	switch (args.length){
	case 4:
	    label = args[0];
	    zoom = args[1];
	    lat = args[2];
	    lon = args[3];			
	    update_position = true;
	    break;
	case 3:
	    zoom = args[0];
	    lat = args[1];
	    lon = args[2];
	    update_position = true;			
	    break;
	case 2:
	    label = args[0];
	    zoom = args[1];
	    update_position = true;
	    break;
	case 1:
	    label = args[0];
	    break;
	default:
	    // console.log("Unrecognized hash string", hash_str);
	    return null;
	}

	// console.log(label, zoom, lat, lon);

	if (zoom){

	    zoom = parseInt(zoom, 10);

	    if (isNaN(zoom)){
		console.log("Invalid zoom");
		return null;
	    }
	}
	
	if (lat) {

	    lat = parseFloat(lat);

	    if (isNaN(lat)){
		console.log("Invalid latitude");
		return false;
	    }
	}

	if (lon){

	    lon = parseFloat(lon);		

	    if (isNaN(lat)){
		console.log("Invalid longitude");
		return false;
	    }
	}
	
	var h = {
	    'latitude': lat, 
	    'longitude': lon,
	    'zoom': zoom,
	    'label': label,
	    'update_position': update_position,
	};
	
	return h;
    },
    
    _change: function(e) {

	var idx = this.select.options[this.select.selectedIndex].value;
	idx = parseInt(idx);

	this._current = idx;
	
	var layer = this._getLayerWithIndex(idx);
	this.setLayer(layer);
    },

    '_getCurrentLayer': function(){
	var current = this._current;
	return this._getLayerWithIndex(current);
    },

    '_getIndexWithLayer': function(layer){

	var count = this.options.catalog.length;
	var idx = -1;

	for (var i=0; i < count; i++){

	    var l = this.options.catalog[i];

	    if (l['label'] == layer['label']){
		idx = i;
		break;
	    }
	}

	return idx;
    },

    '_getLayerWithIndex': function(idx){

	idx = parseInt(idx);	
	
	if (idx < 0){
	    return null;
	}

	if (idx >= this.options.catalog.length){
	    return null;
	}

	return this.options.catalog[idx];
    },
    
    '_getLayerWithLabel': function(label){

	var count = this.options.catalog.length;
	var layer = null;

	for (var i=0; i < count; i++){

	    var l = this.options.catalog[i];

	    if (l['label'] == label){
		layer = l;
		break;
	    }
	}

	return layer;
    },

    '_getNextLayer': function(){

	var current = this._current;
	current = parseInt(current);

	var count_layers = this.options.catalog.length;
	var next = current + 1;

	if (next >= count_layers){
	    next = 0;
	}
	
	this.select.selectedIndex = next + 1;
	this._current = next;
	
	return this._getLayerWithIndex(next);
    },

    '_getPreviousLayer': function(){

	var current = this._current;	
	current = parseInt(current);

	var count_layers = this.options.catalog.length;
	var prev = current - 1;

	if (prev < 0){
	    prev = count_layers - 1;
	} 

	this.select.selectedIndex = prev + 1;
	this._current = prev;
	
	return this._getLayerWithIndex(prev);		
    },

    '_updatePermalink': function(){

	if (! this._hash){
	    return;
	}

	this._hash.onMapMove();
    },

});

// https://leafletjs.com/examples/extending/extending-3-controls.html
// maybe also crib stuff from here: https://github.com/davidchouse/Leaflet.NavBar
