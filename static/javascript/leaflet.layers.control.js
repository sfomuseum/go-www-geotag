L.Control.Layers = L.Control.extend({
    options: {
	position: 'bottomright',
	catalog: [],
	current_layer: null,
	on_change: null,
    },
    onAdd: function(map) {

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

	    if (idx == this.options.current_layer){
		opt.setAttribute("selected", "selected");
	    }
	}

	var _this = this;

	document.addEventListener('keydown', function(e){
		
	    var layer;
	    var offset;

	    if (e.keyCode == 37){
		layer = self.get_previous_layer();
	    }
	    
	    if (e.keyCode == 39){
		layer = self.map.get_next_layer();	    
	    }	

	    if (! layer){
		return;
	    }
	    
	    for (var idx in _this.options.layers){

		idx = parseInt(idx);

		if (layer == _this.options.layers[idx]){

		    // console.log("FUU", idx, _this.options.layers[idx], _this.select.options[idx+1].value);
		    _this.select.selectedIndex = idx + 1;
		}
	    }

	    _this.options.on_change(layer);
	});

        L.DomEvent.on(this.select, 'change', this._change, this);
	return this.div;
    },

    'on_change': function(layer){

	console.log("CHANGE", layer);
	
	if (this.options.on_change){
	    this.options.on_change(layer);
	}
	
    },
    
    _change: function(e) {

	var idx = this.select.options[this.select.selectedIndex].value;	
	var layer = this.get_layer(idx);
	this.on_change(layer);
    },

    'get_layer': function(idx){

	idx = parseInt(idx);	
	console.log("GET LAYER", idx);
	
	if (idx < 0){
	    return null;
	}

	if (idx >= this.options.catalog.length){
	    return null;
	}

	return this.options.catalog[idx];
    },
    
    'get_next_layer': function(){

    },

    'get_previous_layer': function(){

    },
});

// https://leafletjs.com/examples/extending/extending-3-controls.html
// maybe also crib stuff from here: https://github.com/davidchouse/Leaflet.NavBar
