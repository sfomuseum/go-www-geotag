L.Control.Layers = L.Control.extend({
    _map: null,
    _layer: null,
    _current: -1,
    options: {
	position: 'bottomright',
	catalog: [],
	on_change: null,
    },
    onAdd: function(map) {

	this._map = map;
	
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
		layer = _this.get_previous_layer();
	    }
	    
	    if (e.keyCode == 39){
		layer = _this.get_next_layer();	    
	    }	

	    _this.on_change(layer);
	});

        L.DomEvent.on(this.select, 'change', this._change, this);
	return this.div;
    },

    'on_change': function(layer){

	if (this._layer != null){
	    this._map.removeLayer(this._layer);
	    this._layer = null;
	}

	if (! layer){
	    return;
	}
	
	var url = layer["url"];
	var args = {};	 // min,max zoom...

	if (layer["min_zoom"]){
	    args["minZoom"] = parseInt(layer["min_zoom"]);
	}

	if (layer["max_zoom"]){
	    args["maxZoom"] = parseInt(layer["max_zoom"]);
	}
	
	var layer = L.tileLayer(url, args);
	layer.addTo(this._map);

	this._layer = layer;

	// user defined stuff
	
	if (this.options.on_change){
	    this.options.on_change(layer);
	}
	
    },
    
    _change: function(e) {

	var idx = this.select.options[this.select.selectedIndex].value;
	idx = parseInt(idx);

	this._current = idx;
	
	var layer = this.get_layer(idx);
	this.on_change(layer);
    },

    'get_layer': function(idx){

	idx = parseInt(idx);	
	
	if (idx < 0){
	    return null;
	}

	if (idx >= this.options.catalog.length){
	    return null;
	}

	return this.options.catalog[idx];
    },
    
    'get_next_layer': function(){

	var current = this._current;
	current = parseInt(current);

	var count_layers = this.options.catalog.length;
	var next = current + 1;

	if (next >= count_layers){
	    next = 0;
	}
	
	this.select.selectedIndex = next + 1;
	this._current = next;
	
	return this.get_layer(next);
    },

    'get_previous_layer': function(){

	var current = this._current;	
	current = parseInt(current);

	var count_layers = this.options.catalog.length;
	var prev = current - 1;

	if (prev < 0){
	    prev = count_layers - 1;
	} 

	this.select.selectedIndex = prev + 1;
	this._current = prev;
	
	return this.get_layer(prev);		
    },
});

// https://leafletjs.com/examples/extending/extending-3-controls.html
// maybe also crib stuff from here: https://github.com/davidchouse/Leaflet.NavBar
