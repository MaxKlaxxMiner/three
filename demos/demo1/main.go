package demo1

import (
	"github.com/MaxKlaxxMiner/three"
	"syscall/js"
)

// --- Creating a scene ---

func Main() {
	// const scene = new THREE.Scene();
	scene := three.NewScene()

	// const camera = new THREE.PerspectiveCamera( 75, window.innerWidth / window.innerHeight, 0.1, 1000 );
	windowWidth := js.Global().Get("innerWidth").Int()
	windowHeight := js.Global().Get("innerHeight").Int()
	camera := three.NewPerspectiveCamera(75, float64(windowWidth)/float64(windowHeight), 0.1, 1000)

	// const renderer = new THREE.WebGLRenderer();
	renderer := three.NewWebGLRendererDefaults()
	// renderer.setSize( window.innerWidth, window.innerHeight );
	renderer.SetSize(windowWidth, windowHeight)
	// document.body.appendChild( renderer.domElement );
	js.Global().Get("document").Get("body").Call("appendChild", renderer.DomElement)

	// const geometry = new THREE.BoxGeometry( 1, 1, 1 );
	geometry := three.NewBoxGeometry(1, 1, 1)
	//const material = new THREE.MeshBasicMaterial( { color: 0x00ff00 } );
	//const cube = new THREE.Mesh( geometry, material );
	//scene.add( cube );
	//
	//camera.position.z = 5;
	//
	//function animate() {
	//
	//	cube.rotation.x += 0.01;
	//	cube.rotation.y += 0.01;
	//
	//	renderer.render( scene, camera );
	//
	//}

	_, _, _, _ = scene, camera, renderer, geometry
}
