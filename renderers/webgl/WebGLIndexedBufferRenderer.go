package webgl

type IndexedBufferRenderer struct {
	gl         *Context
	extensions *Extensions
	info       *Info
}

func NewWebGLIndexedBufferRenderer(gl *Context, extensions *Extensions, info *Info) *IndexedBufferRenderer {
	this := &IndexedBufferRenderer{gl: gl, extensions: extensions, info: info}
	//todo
	// 	let mode;
	//
	// 	let type, bytesPerElement;
	//
	// 	this.setMode = setMode;
	// 	this.setIndex = setIndex;
	// 	this.render = render;
	// 	this.renderInstances = renderInstances;
	// 	this.renderMultiDraw = renderMultiDraw;
	// 	this.renderMultiDrawInstances = renderMultiDrawInstances;
	return this
}

//todo
// 	function setMode( value ) {
//
// 		mode = value;
//
// 	}
//
// 	function setIndex( value ) {
//
// 		type = value.type;
// 		bytesPerElement = value.bytesPerElement;
//
// 	}
//
// 	function render( start, count ) {
//
// 		gl.drawElements( mode, count, type, start * bytesPerElement );
//
// 		info.update( count, mode, 1 );
//
// 	}
//
// 	function renderInstances( start, count, primcount ) {
//
// 		if ( primcount === 0 ) return;
//
// 		gl.drawElementsInstanced( mode, count, type, start * bytesPerElement, primcount );
//
// 		info.update( count, mode, primcount );
//
// 	}
//
// 	function renderMultiDraw( starts, counts, drawCount ) {
//
// 		if ( drawCount === 0 ) return;
//
// 		const extension = extensions.get( 'WEBGL_multi_draw' );
// 		extension.multiDrawElementsWEBGL( mode, counts, 0, type, starts, 0, drawCount );
//
// 		let elementCount = 0;
// 		for ( let i = 0; i < drawCount; i ++ ) {
//
// 			elementCount += counts[ i ];
//
// 		}
//
// 		info.update( elementCount, mode, 1 );
//
//
// 	}
//
// 	function renderMultiDrawInstances( starts, counts, drawCount, primcount ) {
//
// 		if ( drawCount === 0 ) return;
//
// 		const extension = extensions.get( 'WEBGL_multi_draw' );
//
// 		if ( extension === null ) {
//
// 			for ( let i = 0; i < starts.length; i ++ ) {
//
// 				renderInstances( starts[ i ] / bytesPerElement, counts[ i ], primcount[ i ] );
//
// 			}
//
// 		} else {
//
// 			extension.multiDrawElementsInstancedWEBGL( mode, counts, 0, type, starts, 0, primcount, 0, drawCount );
//
// 			let elementCount = 0;
// 			for ( let i = 0; i < drawCount; i ++ ) {
//
// 				elementCount += counts[ i ] * primcount[ i ];
//
// 			}
//
// 			info.update( elementCount, mode, 1 );
//
// 		}
//
// 	}
