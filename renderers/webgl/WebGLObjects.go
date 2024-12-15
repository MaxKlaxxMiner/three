package webgl

type Objects struct {
	gl         *Context
	geometries *Geometries
	attributes *Attributes
	info       *Info
}

func NewWebGLObjects(gl *Context, geometries *Geometries, attributes *Attributes, info *Info) *Objects {
	this := &Objects{gl: gl, geometries: geometries, attributes: attributes, info: info}
	//todo
	// 	let updateMap = new WeakMap();
	//
	// 		update: update,
	// 		dispose: dispose
	return this
}

//todo
// 	function update( object ) {
//
// 		const frame = info.render.frame;
//
// 		const geometry = object.geometry;
// 		const buffergeometry = geometries.get( object, geometry );
//
// 		// Update once per frame
//
// 		if ( updateMap.get( buffergeometry ) !== frame ) {
//
// 			geometries.update( buffergeometry );
//
// 			updateMap.set( buffergeometry, frame );
//
// 		}
//
// 		if ( object.isInstancedMesh ) {
//
// 			if ( object.hasEventListener( 'dispose', onInstancedMeshDispose ) === false ) {
//
// 				object.addEventListener( 'dispose', onInstancedMeshDispose );
//
// 			}
//
// 			if ( updateMap.get( object ) !== frame ) {
//
// 				attributes.update( object.instanceMatrix, gl.ARRAY_BUFFER );
//
// 				if ( object.instanceColor !== null ) {
//
// 					attributes.update( object.instanceColor, gl.ARRAY_BUFFER );
//
// 				}
//
// 				updateMap.set( object, frame );
//
// 			}
//
// 		}
//
// 		if ( object.isSkinnedMesh ) {
//
// 			const skeleton = object.skeleton;
//
// 			if ( updateMap.get( skeleton ) !== frame ) {
//
// 				skeleton.update();
//
// 				updateMap.set( skeleton, frame );
//
// 			}
//
// 		}
//
// 		return buffergeometry;
//
// 	}
//
// 	function dispose() {
//
// 		updateMap = new WeakMap();
//
// 	}
//
// 	function onInstancedMeshDispose( event ) {
//
// 		const instancedMesh = event.target;
//
// 		instancedMesh.removeEventListener( 'dispose', onInstancedMeshDispose );
//
// 		attributes.remove( instancedMesh.instanceMatrix );
//
// 		if ( instancedMesh.instanceColor !== null ) attributes.remove( instancedMesh.instanceColor );
//
// 	}
