var $ = function(id) {
  return document.getElementById(id);
};
var num;
var max;
var sp;
var frc;
var fade = -0.02;
var fin  = 0.06;  
var fout = 0.03; 
set(); 

function set() {
	var c = $("canv");
	var $$ = c.getContext("2d");
  c.width = window.innerWidth;
  c.height = window.innerHeight;
	data();
	var nodes = [];
	var edges = [];

	function draw() {
    window.requestAnimationFrame(draw);
		nodes = n_upd(c.width, c.height, nodes);
		edges = e_upd(nodes, edges);
		_draw(c, $$, nodes, edges);
	}
	draw();
	for (var i = 0; i < 400; i++) 
    create(nodes);
    edges = [];
    draw();  
    nodes.concat(edges).forEach(function(item) {  
    item.opac = 1;
	});
	_draw(c, $$, nodes, edges);
}

function n_upd(nw, nh, nodes) {
	var _w  = nw  / Math.max(nw, nh);
	var _h = nh / Math.max(nw, nh);
	var narr = [];
	nodes.forEach(function(node, idx) {
		node.px += node.vx * sp;
		node.py += node.vy * sp;
		node.vx = node.vx * 0.99 + (Math.random() - 0.5) * 0.3;
		node.vy = node.vy * 0.99 + (Math.random() - 0.5) * 0.3;
		if (idx >= num || node.px < fade || _w - node.px < fade
				|| node.py < fade || _h - node.py < fade)
			node.opac = Math.max(node.opac - fout, 0);
		else 
			node.opac = Math.min(node.opac + fin, 1);
		if (node.opac > 0)
			narr.push(node);
	});

	for (var i = narr.length; i < num; i++) {
		narr.push({ 
			px: Math.random() * _w,
			py: Math.random() * _h,
			rad: (Math.pow(Math.random(), 2) + 0.35) * 0.015,  
			vx: 0.0,
			vy: 0.0,
			opac: 0.0,
		});
	}
	create(narr);
	return narr;
}

function create(nodes) {
	var del = [];
	for (var i = 0; i < nodes.length * 2; i++)
		del.push(0.0);

	for (var i = 0; i < nodes.length; i++) {
		var a = nodes[i];
		for (var j = 0; j < i; j++) {
			var b = nodes[j];
			var dx = a.px - b.px;
			var dy = a.py - b.py;
			var dist = dx * dx + dy * dy;
			var calc = frc / (Math.sqrt(dist) * (dist + 0.00001));
			dx *= calc;
			dy *= calc;
			del[i * 2 + 0] += dx;
			del[i * 2 + 1] += dy;
			del[j * 2 + 0] -= dx;
			del[j * 2 + 1] -= dy;
		}
	}
	for (var i = 0; i < nodes.length; i++) {
		nodes[i].px += del[i * 2 + 0];
		nodes[i].py += del[i * 2 + 1];
	}
}

function e_upd(nodes, edges) {
	var sum = calcsum(nodes);
	var est = calcest(sum, nodes);
	for (var i = 0; i < sum.length && est.length < nodes.length - 1 + max; i++) {
		var edge = {a:nodes[sum[i][1]], b:nodes[sum[i][2]]};  
		if (!isEdge(est, edge))
			est.push(edge);
	}
	sum = null; 
	var tree = [];
	edges.forEach(function(edge) {
		if (isEdge(est, edge))
			edge.opac = Math.min(edge.opac + fin, 1);
		else
			edge.opac = Math.max(edge.opac - fout, 0);
		if (edge.opac > 0 && edge.a.opac > 0 && edge.b.opac > 0)
			tree.push(edge);
	});
	
	for (var i = 0; i < est.length && tree.length < nodes.length - 1 + max; i++) {
		var edge = est[i];
		if (!isEdge(tree, edge)) {
			edge.opac = 0.0;  
			tree.push(edge);
		}
	}
	return tree;
}
function _draw(c, $$, nodes, edges) {
	var w  = c.width;
	var h = c.height;
	var sz = Math.max(w, h);
	$$.fillStyle = '#C5CFC6';
	$$.fillRect(0, 0, w, h);
	nodes.forEach(function(node) {
   var g1 = $$.createRadialGradient(node.px * sz, node.py * sz,  0, node.px * sz, node.py * sz,node.rad * sz);
    g1.addColorStop(0.0, '#C5CFC6');
    g1.addColorStop(0.3, '#C5CFC6');
    g1.addColorStop(0.4, '#F8EDD1');
    g1.addColorStop(0.6, '#F8EDD1');
    g1.addColorStop(1.0, '#9D9D93');
    $$.fillStyle = g1;
		$$.beginPath();
		$$.arc(node.px * sz, node.py * sz, node.rad * sz, 0, Math.PI * 2);
		$$.fill();
    
	});

	$$.lineWidth = sz / 800;
	edges.forEach(function(edge) {
		var a = edge.a;
		var b = edge.b;
		var dx = a.px - b.px;
		var dy = a.py - b.py;
		var fx = Math.hypot(dx, dy);
		if (fx > a.rad + b.rad) {  
			dx /= fx;  
			dy /= fx;
   var g2 = $$.createLinearGradient((a.px - dx * a.rad) * sz, (a.py - dy * a.rad) * sz,  (b.px + dx * b.rad) * sz, (b.py + dy * b.rad) * sz);
      g2.addColorStop(0.0, '#F8EDD1');
      g2.addColorStop(0.5, '#C5CFC6');
      g2.addColorStop(0.8, '#F8EDD1');
      g2.addColorStop(1.0, '#9D9D93');
      $$.fillStyle = g2;
			$$.strokeStyle = g2;
			$$.beginPath();
			$$.moveTo((a.px - dx * a.rad) * sz, (a.py - dy * a.rad) * sz);
			$$.lineTo((b.px + dx * b.rad) * sz, (b.py + dy * b.rad) * sz);
			$$.stroke();
		}
	});
}
function calcsum(nodes) {
	var end = [];
	for (var i = 0; i < nodes.length; i++) {  
		var a = nodes[i];
		for (var j = 0; j < i; j++) {
			var b = nodes[j];
			var nwght = Math.hypot(a.px - b.px, a.py - b.py);  
			nwght /= Math.pow(a.rad * b.rad, wght); 
			end.push([nwght, i, j]);
		}
	}
	end.sort(function(a, b) {
		var x = a[0], y = b[0];
		return x < y ? -1 : (x > y ? 1 : 0);
	});
	return end;
}

function calcest(sum, nodes) {
	var end = [];
	var ds = new disconn(nodes.length);
	for (var i = 0; i < sum.length && end.length < nodes.length - 1; i++) {
		var edge = sum[i];
		var j = edge[1];
		var k = edge[2];
		if (ds.conn(j, k))
			end.push({a:nodes[j], b:nodes[k]});
	}
	return end;
}

function isEdge(arr, edge) {
	for (var i = 0; i < arr.length; i++) {
		var elem = arr[i];
		if (elem.a == edge.a && elem.b == edge.b ||
		    elem.a == edge.b && elem.b == edge.a)
			return true;
	}
	return false;
}


function disconn(sz) {
	var parr = [];
	var carr = [];
	for (var i = 0; i < sz; i++) {
		parr.push(i);
		carr.push(0);
	}
	
	function rep(i) {
		if (parr[i] != i)
			parr[i] = rep(parr[i]);
		return parr[i];
	}
	
	this.conn = function(i, j) {
		var A = rep(i);
		var B = rep(j);
		if (A == B)
			return false;
		var cidx = carr[A] - carr[B];
		if (cidx >= 0) {
			if (cidx == 0)
				carr[A]++;
			parr[B] = A;
		} else
			parr[A] = B;
		return true;
	};
	
}
function data() {
	var add = $("addedge");
	add.oninput = function() {
		max = Math.round(parseFloat(this.value) / 100 * num);
	};
	add.oninput();
	
	var numNodes = $("numnodes");
	numNodes.oninput = function() {
		num = parseInt(this.value, 10);
		max = Math.round(parseFloat(add.value) / 100 * num);
	};
	numNodes.oninput();
	
	var model= $("type");
	model.onchange = function() {
		wght = parseFloat(this.value);
	};
	model.onchange();
	
	var speed = $("speed");
	speed.oninput = function() {
		var temp = parseFloat(this.value);
		if (!isNaN(temp))
			sp = temp * 0.0001;
	};
	speed.oninput();
	
	var force = $("force");
	force.oninput = function() {
		var temp = parseFloat(this.value);
		if (!isNaN(temp))
			frc = temp * 0.000001;
	};
	force.oninput();
}

window.addEventListener('resize',function(){
  c.width = window.innerWidth;
  c.height = window.innerHeight;
}, false);