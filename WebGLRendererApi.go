package three

//		// API
//
//
// todo
//		this.getContextAttributes = function () {
//
//			return _gl.getContextAttributes();
//
//		};
//
//		this.forceContextLoss = function () {
//
//			const extension = extensions.get( 'WEBGL_lose_context' );
//			if ( extension ) extension.loseContext();
//
//		};
//
//		this.forceContextRestore = function () {
//
//			const extension = extensions.get( 'WEBGL_lose_context' );
//			if ( extension ) extension.restoreContext();
//
//		};
//
//		this.getPixelRatio = function () {
//
//			return _pixelRatio;
//
//		};
//
//		this.setPixelRatio = function ( value ) {
//
//			if ( value === undefined ) return;
//
//			_pixelRatio = value;
//
//			this.setSize( _width, _height, false );
//
//		};
//
//		this.getSize = function ( target ) {
//
//			return target.set( _width, _height );
//
//		};
//
//		this.setSize = function ( width, height, updateStyle = true ) {
//
//			if ( xr.isPresenting ) {
//
//				console.warn( 'THREE.WebGLRenderer: Can\'t change size while VR device is presenting.' );
//				return;
//
//			}
//
//			_width = width;
//			_height = height;
//
//			canvas.width = Math.floor( width * _pixelRatio );
//			canvas.height = Math.floor( height * _pixelRatio );
//
//			if ( updateStyle === true ) {
//
//				canvas.style.width = width + 'px';
//				canvas.style.height = height + 'px';
//
//			}
//
//			this.setViewport( 0, 0, width, height );
//
//		};
//
//		this.getDrawingBufferSize = function ( target ) {
//
//			return target.set( _width * _pixelRatio, _height * _pixelRatio ).floor();
//
//		};
//
//		this.setDrawingBufferSize = function ( width, height, pixelRatio ) {
//
//			_width = width;
//			_height = height;
//
//			_pixelRatio = pixelRatio;
//
//			canvas.width = Math.floor( width * pixelRatio );
//			canvas.height = Math.floor( height * pixelRatio );
//
//			this.setViewport( 0, 0, width, height );
//
//		};
//
//		this.getCurrentViewport = function ( target ) {
//
//			return target.copy( _currentViewport );
//
//		};
//
//		this.getViewport = function ( target ) {
//
//			return target.copy( _viewport );
//
//		};
//
//		this.setViewport = function ( x, y, width, height ) {
//
//			if ( x.isVector4 ) {
//
//				_viewport.set( x.x, x.y, x.z, x.w );
//
//			} else {
//
//				_viewport.set( x, y, width, height );
//
//			}
//
//			state.viewport( _currentViewport.copy( _viewport ).multiplyScalar( _pixelRatio ).round() );
//
//		};
//
//		this.getScissor = function ( target ) {
//
//			return target.copy( _scissor );
//
//		};
//
//		this.setScissor = function ( x, y, width, height ) {
//
//			if ( x.isVector4 ) {
//
//				_scissor.set( x.x, x.y, x.z, x.w );
//
//			} else {
//
//				_scissor.set( x, y, width, height );
//
//			}
//
//			state.scissor( _currentScissor.copy( _scissor ).multiplyScalar( _pixelRatio ).round() );
//
//		};
//
//		this.getScissorTest = function () {
//
//			return _scissorTest;
//
//		};
//
//		this.setScissorTest = function ( boolean ) {
//
//			state.setScissorTest( _scissorTest = boolean );
//
//		};
//
//		this.setOpaqueSort = function ( method ) {
//
//			_opaqueSort = method;
//
//		};
//
//		this.setTransparentSort = function ( method ) {
//
//			_transparentSort = method;
//
//		};
//
//		// Clearing
//
//		this.getClearColor = function ( target ) {
//
//			return target.copy( background.getClearColor() );
//
//		};
//
//		this.setClearColor = function () {
//
//			background.setClearColor.apply( background, arguments );
//
//		};
//
//		this.getClearAlpha = function () {
//
//			return background.getClearAlpha();
//
//		};
//
//		this.setClearAlpha = function () {
//
//			background.setClearAlpha.apply( background, arguments );
//
//		};
//
//		this.clear = function ( color = true, depth = true, stencil = true ) {
//
//			let bits = 0;
//
//			if ( color ) {
//
//				// check if we're trying to clear an integer target
//				let isIntegerFormat = false;
//				if ( _currentRenderTarget !== null ) {
//
//					const targetFormat = _currentRenderTarget.texture.format;
//					isIntegerFormat = targetFormat === RGBAIntegerFormat ||
//						targetFormat === RGIntegerFormat ||
//						targetFormat === RedIntegerFormat;
//
//				}
//
//				// use the appropriate clear functions to clear the target if it's a signed
//				// or unsigned integer target
//				if ( isIntegerFormat ) {
//
//					const targetType = _currentRenderTarget.texture.type;
//					const isUnsignedType = targetType === UnsignedByteType ||
//						targetType === UnsignedIntType ||
//						targetType === UnsignedShortType ||
//						targetType === UnsignedInt248Type ||
//						targetType === UnsignedShort4444Type ||
//						targetType === UnsignedShort5551Type;
//
//					const clearColor = background.getClearColor();
//					const a = background.getClearAlpha();
//					const r = clearColor.r;
//					const g = clearColor.g;
//					const b = clearColor.b;
//
//					if ( isUnsignedType ) {
//
//						uintClearColor[ 0 ] = r;
//						uintClearColor[ 1 ] = g;
//						uintClearColor[ 2 ] = b;
//						uintClearColor[ 3 ] = a;
//						_gl.clearBufferuiv( _gl.COLOR, 0, uintClearColor );
//
//					} else {
//
//						intClearColor[ 0 ] = r;
//						intClearColor[ 1 ] = g;
//						intClearColor[ 2 ] = b;
//						intClearColor[ 3 ] = a;
//						_gl.clearBufferiv( _gl.COLOR, 0, intClearColor );
//
//					}
//
//				} else {
//
//					bits |= _gl.COLOR_BUFFER_BIT;
//
//				}
//
//			}
//
//			if ( depth ) {
//
//				bits |= _gl.DEPTH_BUFFER_BIT;
//
//			}
//
//			if ( stencil ) {
//
//				bits |= _gl.STENCIL_BUFFER_BIT;
//				this.state.buffers.stencil.setMask( 0xffffffff );
//
//			}
//
//			_gl.clear( bits );
//
//		};
//
//		this.clearColor = function () {
//
//			this.clear( true, false, false );
//
//		};
//
//		this.clearDepth = function () {
//
//			this.clear( false, true, false );
//
//		};
//
//		this.clearStencil = function () {
//
//			this.clear( false, false, true );
//
//		};
//
//		//
//
//		this.dispose = function () {
//
//			canvas.removeEventListener( 'webglcontextlost', onContextLost, false );
//			canvas.removeEventListener( 'webglcontextrestored', onContextRestore, false );
//			canvas.removeEventListener( 'webglcontextcreationerror', onContextCreationError, false );
//
//			renderLists.dispose();
//			renderStates.dispose();
//			properties.dispose();
//			cubemaps.dispose();
//			cubeuvmaps.dispose();
//			objects.dispose();
//			bindingStates.dispose();
//			uniformsGroups.dispose();
//			programCache.dispose();
//
//			xr.dispose();
//
//			xr.removeEventListener( 'sessionstart', onXRSessionStart );
//			xr.removeEventListener( 'sessionend', onXRSessionEnd );
//
//			animation.stop();
//
//		};
