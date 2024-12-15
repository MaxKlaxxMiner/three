package webgl

type RenderStates struct {
	extensions *Extensions
}

func NewWebGLRenderStates(extensions *Extensions) *RenderStates {
	this := &RenderStates{extensions: extensions}
	//todo
	// 	const lights = new WebGLLights( extensions );
	//
	// 	const lightsArray = [];
	// 	const shadowsArray = [];
	//
	// 		get: get,
	// 		dispose: dispose
	return this
}

//todo
// 	function init( camera ) {
//
// 		state.camera = camera;
//
// 		lightsArray.length = 0;
// 		shadowsArray.length = 0;
//
// 	}
//
// 	function pushLight( light ) {
//
// 		lightsArray.push( light );
//
// 	}
//
// 	function pushShadow( shadowLight ) {
//
// 		shadowsArray.push( shadowLight );
//
// 	}
//
// 	function setupLights() {
//
// 		lights.setup( lightsArray );
//
// 	}
//
// 	function setupLightsView( camera ) {
//
// 		lights.setupView( lightsArray, camera );
//
// 	}
//
// 	const state = {
// 		lightsArray: lightsArray,
// 		shadowsArray: shadowsArray,
//
// 		camera: null,
//
// 		lights: lights,
//
// 		transmissionRenderTarget: {}
// 	};
//
// 	return {
// 		init: init,
// 		state: state,
// 		setupLights: setupLights,
// 		setupLightsView: setupLightsView,
//
// 		pushLight: pushLight,
// 		pushShadow: pushShadow
// 	};
//
// }
//
// function WebGLRenderStates( extensions ) {
//
// 	let renderStates = new WeakMap();
//
// 	function get( scene, renderCallDepth = 0 ) {
//
// 		const renderStateArray = renderStates.get( scene );
// 		let renderState;
//
// 		if ( renderStateArray === undefined ) {
//
// 			renderState = new WebGLRenderState( extensions );
// 			renderStates.set( scene, [ renderState ] );
//
// 		} else {
//
// 			if ( renderCallDepth >= renderStateArray.length ) {
//
// 				renderState = new WebGLRenderState( extensions );
// 				renderStateArray.push( renderState );
//
// 			} else {
//
// 				renderState = renderStateArray[ renderCallDepth ];
//
// 			}
//
// 		}
//
// 		return renderState;
//
// 	}
//
// 	function dispose() {
//
// 		renderStates = new WeakMap();
//
// 	}
// import { WebGLLights } from './WebGLLights.js';
