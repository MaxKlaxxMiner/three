package webgl

type CubeUVMaps struct {
	renderer *Renderer
}

func NewWebGLCubeUVMaps(renderer *Renderer) *CubeUVMaps {
	this := &CubeUVMaps{renderer: renderer}
	// 	let cubeUVmaps = new WeakMap(); todo
	// 	let pmremGenerator = null;
	// 		get: get,
	// 		dispose: dispose
	return this
}

//todo
// 	function get( texture ) {
//
// 		if ( texture && texture.isTexture ) {
//
// 			const mapping = texture.mapping;
//
// 			const isEquirectMap = ( mapping === EquirectangularReflectionMapping || mapping === EquirectangularRefractionMapping );
// 			const isCubeMap = ( mapping === CubeReflectionMapping || mapping === CubeRefractionMapping );
//
// 			// equirect/cube map to cubeUV conversion
//
// 			if ( isEquirectMap || isCubeMap ) {
//
// 				let renderTarget = cubeUVmaps.get( texture );
//
// 				const currentPMREMVersion = renderTarget !== undefined ? renderTarget.texture.pmremVersion : 0;
//
// 				if ( texture.isRenderTargetTexture && texture.pmremVersion !== currentPMREMVersion ) {
//
// 					if ( pmremGenerator === null ) pmremGenerator = new PMREMGenerator( renderer );
//
// 					renderTarget = isEquirectMap ? pmremGenerator.fromEquirectangular( texture, renderTarget ) : pmremGenerator.fromCubemap( texture, renderTarget );
// 					renderTarget.texture.pmremVersion = texture.pmremVersion;
//
// 					cubeUVmaps.set( texture, renderTarget );
//
// 					return renderTarget.texture;
//
// 				} else {
//
// 					if ( renderTarget !== undefined ) {
//
// 						return renderTarget.texture;
//
// 					} else {
//
// 						const image = texture.image;
//
// 						if ( ( isEquirectMap && image && image.height > 0 ) || ( isCubeMap && image && isCubeTextureComplete( image ) ) ) {
//
// 							if ( pmremGenerator === null ) pmremGenerator = new PMREMGenerator( renderer );
//
// 							renderTarget = isEquirectMap ? pmremGenerator.fromEquirectangular( texture ) : pmremGenerator.fromCubemap( texture );
// 							renderTarget.texture.pmremVersion = texture.pmremVersion;
//
// 							cubeUVmaps.set( texture, renderTarget );
//
// 							texture.addEventListener( 'dispose', onTextureDispose );
//
// 							return renderTarget.texture;
//
// 						} else {
//
// 							// image not yet ready. try the conversion next frame
//
// 							return null;
//
// 						}
//
// 					}
//
// 				}
//
// 			}
//
// 		}
//
// 		return texture;
//
// 	}
//
// 	function isCubeTextureComplete( image ) {
//
// 		let count = 0;
// 		const length = 6;
//
// 		for ( let i = 0; i < length; i ++ ) {
//
// 			if ( image[ i ] !== undefined ) count ++;
//
// 		}
//
// 		return count === length;
//
//
// 	}
//
// 	function onTextureDispose( event ) {
//
// 		const texture = event.target;
//
// 		texture.removeEventListener( 'dispose', onTextureDispose );
//
// 		const cubemapUV = cubeUVmaps.get( texture );
//
// 		if ( cubemapUV !== undefined ) {
//
// 			cubeUVmaps.delete( texture );
// 			cubemapUV.dispose();
//
// 		}
//
// 	}
//
// 	function dispose() {
//
// 		cubeUVmaps = new WeakMap();
//
// 		if ( pmremGenerator !== null ) {
//
// 			pmremGenerator.dispose();
// 			pmremGenerator = null;
//
// 		}
//
// 	}
// import { CubeReflectionMapping, CubeRefractionMapping, EquirectangularReflectionMapping, EquirectangularRefractionMapping } from '../../constants.js';
// import { PMREMGenerator } from '../../extras/PMREMGenerator.js';
