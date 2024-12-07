package webgl

import (
	"fmt"
	"github.com/MaxKlaxxMiner/three/consts"
	"github.com/MaxKlaxxMiner/three/math"
	"github.com/MaxKlaxxMiner/three/utils"
	"strconv"
	"strings"
)

type State struct {
	gl         Context
	extensions Extensions
	// 	const colorBuffer = new ColorBuffer(); todo
	// 	const depthBuffer = new DepthBuffer(); todo
	// 	const stencilBuffer = new StencilBuffer(); todo
	// 	const uboBindings = new WeakMap(); todo
	// 	const uboProgramMap = new WeakMap(); todo
	enabledCapabilities map[int32]bool
	// 	let currentBoundFramebuffers = {}; todo
	// 	let currentDrawbuffers = new WeakMap(); todo
	// 	let defaultDrawbuffers = []; todo
	// 	let currentProgram = null; todo
	// 	let currentBlendingEnabled = false; todo
	// 	let currentBlending = null; todo
	// 	let currentBlendEquation = null; todo
	// 	let currentBlendSrc = null; todo
	// 	let currentBlendDst = null; todo
	// 	let currentBlendEquationAlpha = null; todo
	// 	let currentBlendSrcAlpha = null; todo
	// 	let currentBlendDstAlpha = null; todo
	// 	let currentBlendColor = new Color( 0, 0, 0 ); todo
	// 	let currentBlendAlpha = 0; todo
	// 	let currentPremultipledAlpha = false; todo
	currentFlipSided bool
	currentCullFace  consts.CullFace
	// 	let currentLineWidth = null; todo
	// 	let currentPolygonOffsetFactor = null; todo
	// 	let currentPolygonOffsetUnits = null; todo
	maxTextures        int
	lineWidthAvailable bool
	// 	let currentTextureSlot = null; todo
	// 	let currentBoundTextures = {}; todo
	// 	const scissorParam = gl.getParameter( gl.SCISSOR_BOX ); todo
	// 	const viewportParam = gl.getParameter( gl.VIEWPORT ); todo
	// 	const currentScissor = new Vector4().fromArray( scissorParam ); todo
	currentViewport math.Vector4
	// 	const emptyTextures = {}; todo
	// 	emptyTextures[ gl.TEXTURE_2D ] = createTexture( gl.TEXTURE_2D, gl.TEXTURE_2D, 1 ); todo
	// 	emptyTextures[ gl.TEXTURE_CUBE_MAP ] = createTexture( gl.TEXTURE_CUBE_MAP, gl.TEXTURE_CUBE_MAP_POSITIVE_X, 6 ); todo
	// 	emptyTextures[ gl.TEXTURE_2D_ARRAY ] = createTexture( gl.TEXTURE_2D_ARRAY, gl.TEXTURE_2D_ARRAY, 1, 1 ); todo
	// 	emptyTextures[ gl.TEXTURE_3D ] = createTexture( gl.TEXTURE_3D, gl.TEXTURE_3D, 1, 1 ); todo

	//todo
	// 		buffers: {
	// 			color: colorBuffer,
	// 			depth: depthBuffer,
	// 			stencil: stencilBuffer
	// 		},
	//
	// 		bindFramebuffer: bindFramebuffer,
	// 		drawBuffers: drawBuffers,
	//
	// 		useProgram: useProgram,
	//
	// 		setBlending: setBlending,
	// 		setMaterial: setMaterial,
	//
	// 		setLineWidth: setLineWidth,
	// 		setPolygonOffset: setPolygonOffset,
	//
	// 		setScissorTest: setScissorTest,
	//
	// 		activeTexture: activeTexture,
	// 		bindTexture: bindTexture,
	// 		unbindTexture: unbindTexture,
	// 		compressedTexImage2D: compressedTexImage2D,
	// 		compressedTexImage3D: compressedTexImage3D,
	// 		texImage2D: texImage2D,
	// 		texImage3D: texImage3D,
	//
	// 		updateUBOMapping: updateUBOMapping,
	// 		uniformBlockBinding: uniformBlockBinding,
	//
	// 		texStorage2D: texStorage2D,
	// 		texStorage3D: texStorage3D,
	// 		texSubImage2D: texSubImage2D,
	// 		texSubImage3D: texSubImage3D,
	// 		compressedTexSubImage2D: compressedTexSubImage2D,
	// 		compressedTexSubImage3D: compressedTexSubImage3D,
	//
	// 		scissor: scissor,
	// 		reset: reset
}

func NewWebGLState(gl Context, extensions Extensions) *State {
	r := &State{gl: gl, extensions: extensions}

	//todo
	// 	const colorBuffer = new ColorBuffer();
	// 	const depthBuffer = new DepthBuffer();
	// 	const stencilBuffer = new StencilBuffer();
	//
	// 	const uboBindings = new WeakMap();
	// 	const uboProgramMap = new WeakMap();

	r.enabledCapabilities = make(map[int32]bool)

	r.maxTextures = gl.GetParameterInt(gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS)

	glVersion := gl.GetParameterStr(gl.VERSION)
	if p := strings.Index(glVersion, "WebGL "); p == 0 {
		glVersion = glVersion[6:]
		version, _ := strconv.Atoi(glVersion[:strings.IndexFunc(glVersion, func(r rune) bool { return r < '0' || r > '9' })])
		r.lineWidthAvailable = version >= 2
	} else if p = strings.Index(glVersion, "OpenGL ES "); p >= 0 {
		glVersion = glVersion[p+10:]
		version, _ := strconv.Atoi(glVersion[:strings.IndexFunc(glVersion, func(r rune) bool { return r < '0' || r > '9' })])
		r.lineWidthAvailable = version >= 3
	}

	//todo
	// 	let currentTextureSlot = null;
	// 	let currentBoundTextures = {};
	//
	//scissorParam := r.gl.Call("getParameter", gl.SCISSOR_BOX)
	viewportParam := r.gl.Call("getParameter", gl.VIEWPORT)

	//
	// 	const currentScissor = new Vector4().fromArray( scissorParam ); todo
	//r.currentViewport.Set(viewportParam.Index(0).Float(), viewportParam.Index(1).Float(), viewportParam.Index(2).Float(), viewportParam.Index(3).Float())
	r.currentViewport.FromArray(utils.ConvertTypedArrayToFloat64Slice(viewportParam))

	// 	const currentViewport = new Vector4().fromArray( viewportParam );
	//
	// 	const emptyTextures = {};
	// 	emptyTextures[ gl.TEXTURE_2D ] = createTexture( gl.TEXTURE_2D, gl.TEXTURE_2D, 1 );
	// 	emptyTextures[ gl.TEXTURE_CUBE_MAP ] = createTexture( gl.TEXTURE_CUBE_MAP, gl.TEXTURE_CUBE_MAP_POSITIVE_X, 6 );
	// 	emptyTextures[ gl.TEXTURE_2D_ARRAY ] = createTexture( gl.TEXTURE_2D_ARRAY, gl.TEXTURE_2D_ARRAY, 1, 1 );
	// 	emptyTextures[ gl.TEXTURE_3D ] = createTexture( gl.TEXTURE_3D, gl.TEXTURE_3D, 1, 1 );
	//
	// 	// init
	//
	// 	colorBuffer.setClear( 0, 0, 0, 1 );
	// 	depthBuffer.setClear( 1 );
	// 	stencilBuffer.setClear( 0 );

	r.Enable(gl.DEPTH_TEST)
	//todo
	// 	depthBuffer.setFunc( LessEqualDepth );

	r.currentFlipSided = true
	r.SetFlipSided(false)
	r.SetCullFace(consts.CullFaceBack)
	r.Enable(gl.CULL_FACE)

	//todo
	// 	setBlending( NoBlending );
	//
	// 	const equationToGL = {
	// 		[ AddEquation ]: gl.FUNC_ADD,
	// 		[ SubtractEquation ]: gl.FUNC_SUBTRACT,
	// 		[ ReverseSubtractEquation ]: gl.FUNC_REVERSE_SUBTRACT
	// 	};
	//
	// 	equationToGL[ MinEquation ] = gl.MIN;
	// 	equationToGL[ MaxEquation ] = gl.MAX;
	//
	// 	const factorToGL = {
	// 		[ ZeroFactor ]: gl.ZERO,
	// 		[ OneFactor ]: gl.ONE,
	// 		[ SrcColorFactor ]: gl.SRC_COLOR,
	// 		[ SrcAlphaFactor ]: gl.SRC_ALPHA,
	// 		[ SrcAlphaSaturateFactor ]: gl.SRC_ALPHA_SATURATE,
	// 		[ DstColorFactor ]: gl.DST_COLOR,
	// 		[ DstAlphaFactor ]: gl.DST_ALPHA,
	// 		[ OneMinusSrcColorFactor ]: gl.ONE_MINUS_SRC_COLOR,
	// 		[ OneMinusSrcAlphaFactor ]: gl.ONE_MINUS_SRC_ALPHA,
	// 		[ OneMinusDstColorFactor ]: gl.ONE_MINUS_DST_COLOR,
	// 		[ OneMinusDstAlphaFactor ]: gl.ONE_MINUS_DST_ALPHA,
	// 		[ ConstantColorFactor ]: gl.CONSTANT_COLOR,
	// 		[ OneMinusConstantColorFactor ]: gl.ONE_MINUS_CONSTANT_COLOR,
	// 		[ ConstantAlphaFactor ]: gl.CONSTANT_ALPHA,
	// 		[ OneMinusConstantAlphaFactor ]: gl.ONE_MINUS_CONSTANT_ALPHA
	// 	};
	//
	// 	return {
	//
	// 		buffers: {
	// 			color: colorBuffer,
	// 			depth: depthBuffer,
	// 			stencil: stencilBuffer
	// 		},
	//
	// 		bindFramebuffer: bindFramebuffer,
	// 		drawBuffers: drawBuffers,
	//
	// 		useProgram: useProgram,
	//
	// 		setBlending: setBlending,
	// 		setMaterial: setMaterial,
	//
	// 		setLineWidth: setLineWidth,
	// 		setPolygonOffset: setPolygonOffset,
	//
	// 		setScissorTest: setScissorTest,
	//
	// 		activeTexture: activeTexture,
	// 		bindTexture: bindTexture,
	// 		unbindTexture: unbindTexture,
	// 		compressedTexImage2D: compressedTexImage2D,
	// 		compressedTexImage3D: compressedTexImage3D,
	// 		texImage2D: texImage2D,
	// 		texImage3D: texImage3D,
	//
	// 		updateUBOMapping: updateUBOMapping,
	// 		uniformBlockBinding: uniformBlockBinding,
	//
	// 		texStorage2D: texStorage2D,
	// 		texStorage3D: texStorage3D,
	// 		texSubImage2D: texSubImage2D,
	// 		texSubImage3D: texSubImage3D,
	// 		compressedTexSubImage2D: compressedTexSubImage2D,
	// 		compressedTexSubImage3D: compressedTexSubImage3D,
	//
	// 		scissor: scissor,
	// 		viewport: viewport,
	//
	// 		reset: reset
	//
	// 	};
	return r
}

//todo
// const reversedFuncs = {
// 	[ NeverDepth ]: AlwaysDepth,
// 	[ LessDepth ]: GreaterDepth,
// 	[ EqualDepth ]: NotEqualDepth,
// 	[ LessEqualDepth ]: GreaterEqualDepth,
//
// 	[ AlwaysDepth ]: NeverDepth,
// 	[ GreaterDepth ]: LessDepth,
// 	[ NotEqualDepth ]: EqualDepth,
// 	[ GreaterEqualDepth ]: LessEqualDepth,
// };
// 	function ColorBuffer() {
//
// 		let locked = false;
//
// 		const color = new Vector4();
// 		let currentColorMask = null;
// 		const currentColorClear = new Vector4( 0, 0, 0, 0 );
//
// 		return {
//
// 			setMask: function ( colorMask ) {
//
// 				if ( currentColorMask !== colorMask && ! locked ) {
//
// 					gl.colorMask( colorMask, colorMask, colorMask, colorMask );
// 					currentColorMask = colorMask;
//
// 				}
//
// 			},
//
// 			setLocked: function ( lock ) {
//
// 				locked = lock;
//
// 			},
//
// 			setClear: function ( r, g, b, a, premultipliedAlpha ) {
//
// 				if ( premultipliedAlpha === true ) {
//
// 					r *= a; g *= a; b *= a;
//
// 				}
//
// 				color.set( r, g, b, a );
//
// 				if ( currentColorClear.equals( color ) === false ) {
//
// 					gl.clearColor( r, g, b, a );
// 					currentColorClear.copy( color );
//
// 				}
//
// 			},
//
// 			reset: function () {
//
// 				locked = false;
//
// 				currentColorMask = null;
// 				currentColorClear.set( - 1, 0, 0, 0 ); // set to invalid state
//
// 			}
//
// 		};
//
// 	}
//
// 	function DepthBuffer() {
//
// 		let locked = false;
// 		let reversed = false;
//
// 		let currentDepthMask = null;
// 		let currentDepthFunc = null;
// 		let currentDepthClear = null;
//
// 		return {
//
// 			setReversed: function ( value ) {
//
// 				if ( reversed !== value ) {
//
// 					const ext = extensions.get( 'EXT_clip_control' );
//
// 					if ( reversed ) {
//
// 						ext.clipControlEXT( ext.LOWER_LEFT_EXT, ext.ZERO_TO_ONE_EXT );
//
// 					} else {
//
// 						ext.clipControlEXT( ext.LOWER_LEFT_EXT, ext.NEGATIVE_ONE_TO_ONE_EXT );
//
// 					}
//
// 					const oldDepth = currentDepthClear;
// 					currentDepthClear = null;
// 					this.setClear( oldDepth );
//
// 				}
//
// 				reversed = value;
//
// 			},
//
// 			getReversed: function () {
//
// 				return reversed;
//
// 			},
//
// 			setTest: function ( depthTest ) {
//
// 				if ( depthTest ) {
//
// 					enable( gl.DEPTH_TEST );
//
// 				} else {
//
// 					disable( gl.DEPTH_TEST );
//
// 				}
//
// 			},
//
// 			setMask: function ( depthMask ) {
//
// 				if ( currentDepthMask !== depthMask && ! locked ) {
//
// 					gl.depthMask( depthMask );
// 					currentDepthMask = depthMask;
//
// 				}
//
// 			},
//
// 			setFunc: function ( depthFunc ) {
//
// 				if ( reversed ) depthFunc = reversedFuncs[ depthFunc ];
//
// 				if ( currentDepthFunc !== depthFunc ) {
//
// 					switch ( depthFunc ) {
//
// 						case NeverDepth:
//
// 							gl.depthFunc( gl.NEVER );
// 							break;
//
// 						case AlwaysDepth:
//
// 							gl.depthFunc( gl.ALWAYS );
// 							break;
//
// 						case LessDepth:
//
// 							gl.depthFunc( gl.LESS );
// 							break;
//
// 						case LessEqualDepth:
//
// 							gl.depthFunc( gl.LEQUAL );
// 							break;
//
// 						case EqualDepth:
//
// 							gl.depthFunc( gl.EQUAL );
// 							break;
//
// 						case GreaterEqualDepth:
//
// 							gl.depthFunc( gl.GEQUAL );
// 							break;
//
// 						case GreaterDepth:
//
// 							gl.depthFunc( gl.GREATER );
// 							break;
//
// 						case NotEqualDepth:
//
// 							gl.depthFunc( gl.NOTEQUAL );
// 							break;
//
// 						default:
//
// 							gl.depthFunc( gl.LEQUAL );
//
// 					}
//
// 					currentDepthFunc = depthFunc;
//
// 				}
//
// 			},
//
// 			setLocked: function ( lock ) {
//
// 				locked = lock;
//
// 			},
//
// 			setClear: function ( depth ) {
//
// 				if ( currentDepthClear !== depth ) {
//
// 					if ( reversed ) {
//
// 						depth = 1 - depth;
//
// 					}
//
// 					gl.clearDepth( depth );
// 					currentDepthClear = depth;
//
// 				}
//
// 			},
//
// 			reset: function () {
//
// 				locked = false;
//
// 				currentDepthMask = null;
// 				currentDepthFunc = null;
// 				currentDepthClear = null;
// 				reversed = false;
//
// 			}
//
// 		};
//
// 	}
//
// 	function StencilBuffer() {
//
// 		let locked = false;
//
// 		let currentStencilMask = null;
// 		let currentStencilFunc = null;
// 		let currentStencilRef = null;
// 		let currentStencilFuncMask = null;
// 		let currentStencilFail = null;
// 		let currentStencilZFail = null;
// 		let currentStencilZPass = null;
// 		let currentStencilClear = null;
//
// 		return {
//
// 			setTest: function ( stencilTest ) {
//
// 				if ( ! locked ) {
//
// 					if ( stencilTest ) {
//
// 						enable( gl.STENCIL_TEST );
//
// 					} else {
//
// 						disable( gl.STENCIL_TEST );
//
// 					}
//
// 				}
//
// 			},
//
// 			setMask: function ( stencilMask ) {
//
// 				if ( currentStencilMask !== stencilMask && ! locked ) {
//
// 					gl.stencilMask( stencilMask );
// 					currentStencilMask = stencilMask;
//
// 				}
//
// 			},
//
// 			setFunc: function ( stencilFunc, stencilRef, stencilMask ) {
//
// 				if ( currentStencilFunc !== stencilFunc ||
// 				     currentStencilRef !== stencilRef ||
// 				     currentStencilFuncMask !== stencilMask ) {
//
// 					gl.stencilFunc( stencilFunc, stencilRef, stencilMask );
//
// 					currentStencilFunc = stencilFunc;
// 					currentStencilRef = stencilRef;
// 					currentStencilFuncMask = stencilMask;
//
// 				}
//
// 			},
//
// 			setOp: function ( stencilFail, stencilZFail, stencilZPass ) {
//
// 				if ( currentStencilFail !== stencilFail ||
// 				     currentStencilZFail !== stencilZFail ||
// 				     currentStencilZPass !== stencilZPass ) {
//
// 					gl.stencilOp( stencilFail, stencilZFail, stencilZPass );
//
// 					currentStencilFail = stencilFail;
// 					currentStencilZFail = stencilZFail;
// 					currentStencilZPass = stencilZPass;
//
// 				}
//
// 			},
//
// 			setLocked: function ( lock ) {
//
// 				locked = lock;
//
// 			},
//
// 			setClear: function ( stencil ) {
//
// 				if ( currentStencilClear !== stencil ) {
//
// 					gl.clearStencil( stencil );
// 					currentStencilClear = stencil;
//
// 				}
//
// 			},
//
// 			reset: function () {
//
// 				locked = false;
//
// 				currentStencilMask = null;
// 				currentStencilFunc = null;
// 				currentStencilRef = null;
// 				currentStencilFuncMask = null;
// 				currentStencilFail = null;
// 				currentStencilZFail = null;
// 				currentStencilZPass = null;
// 				currentStencilClear = null;
//
// 			}
//
// 		};
//
// 	}
//
// 	function createTexture( type, target, count, dimensions ) {
//
// 		const data = new Uint8Array( 4 ); // 4 is required to match default unpack alignment of 4.
// 		const texture = gl.createTexture();
//
// 		gl.bindTexture( type, texture );
// 		gl.texParameteri( type, gl.TEXTURE_MIN_FILTER, gl.NEAREST );
// 		gl.texParameteri( type, gl.TEXTURE_MAG_FILTER, gl.NEAREST );
//
// 		for ( let i = 0; i < count; i ++ ) {
//
// 			if ( type === gl.TEXTURE_3D || type === gl.TEXTURE_2D_ARRAY ) {
//
// 				gl.texImage3D( target, 0, gl.RGBA, 1, 1, dimensions, 0, gl.RGBA, gl.UNSIGNED_BYTE, data );
//
// 			} else {
//
// 				gl.texImage2D( target + i, 0, gl.RGBA, 1, 1, 0, gl.RGBA, gl.UNSIGNED_BYTE, data );
//
// 			}
//
// 		}
//
// 		return texture;
//
// 	}
//

func (s *State) Enable(id int32) {
	if s.enabledCapabilities[id] {
		return
	}
	if id == 0 {
		fmt.Println("Enable: Capabilities-ID = 0")
		panic("Enable: Capabilities-ID is 0")
	}
	s.gl.Call("enable", id)
	s.enabledCapabilities[id] = true
}

func (s *State) Disable(id int32) {
	if !s.enabledCapabilities[id] {
		return
	}
	if id == 0 {
		fmt.Println("Disable: Capabilities-ID = 0")
		panic("Disable: Capabilities-ID is 0")
	}
	s.gl.Call("disable", id)
	s.enabledCapabilities[id] = false
}

func (s *State) IsEnabled(id int32) bool {
	if b, ok := s.enabledCapabilities[id]; ok {
		return b
	} else {
		b = s.IsEnabledDirect(id)
		s.enabledCapabilities[id] = b
		return b
	}
}

func (s *State) IsEnabledDirect(id int32) bool {
	return s.gl.Call("isEnabled", id).Bool()
}

//todo
// 	function bindFramebuffer( target, framebuffer ) {
//
// 		if ( currentBoundFramebuffers[ target ] !== framebuffer ) {
//
// 			gl.bindFramebuffer( target, framebuffer );
//
// 			currentBoundFramebuffers[ target ] = framebuffer;
//
// 			// gl.DRAW_FRAMEBUFFER is equivalent to gl.FRAMEBUFFER
//
// 			if ( target === gl.DRAW_FRAMEBUFFER ) {
//
// 				currentBoundFramebuffers[ gl.FRAMEBUFFER ] = framebuffer;
//
// 			}
//
// 			if ( target === gl.FRAMEBUFFER ) {
//
// 				currentBoundFramebuffers[ gl.DRAW_FRAMEBUFFER ] = framebuffer;
//
// 			}
//
// 			return true;
//
// 		}
//
// 		return false;
//
// 	}
//
// 	function drawBuffers( renderTarget, framebuffer ) {
//
// 		let drawBuffers = defaultDrawbuffers;
//
// 		let needsUpdate = false;
//
// 		if ( renderTarget ) {
//
// 			drawBuffers = currentDrawbuffers.get( framebuffer );
//
// 			if ( drawBuffers === undefined ) {
//
// 				drawBuffers = [];
// 				currentDrawbuffers.set( framebuffer, drawBuffers );
//
// 			}
//
// 			const textures = renderTarget.textures;
//
// 			if ( drawBuffers.length !== textures.length || drawBuffers[ 0 ] !== gl.COLOR_ATTACHMENT0 ) {
//
// 				for ( let i = 0, il = textures.length; i < il; i ++ ) {
//
// 					drawBuffers[ i ] = gl.COLOR_ATTACHMENT0 + i;
//
// 				}
//
// 				drawBuffers.length = textures.length;
//
// 				needsUpdate = true;
//
// 			}
//
// 		} else {
//
// 			if ( drawBuffers[ 0 ] !== gl.BACK ) {
//
// 				drawBuffers[ 0 ] = gl.BACK;
//
// 				needsUpdate = true;
//
// 			}
//
// 		}
//
// 		if ( needsUpdate ) {
//
// 			gl.drawBuffers( drawBuffers );
//
// 		}
//
// 	}
//
// 	function useProgram( program ) {
//
// 		if ( currentProgram !== program ) {
//
// 			gl.useProgram( program );
//
// 			currentProgram = program;
//
// 			return true;
//
// 		}
//
// 		return false;
//
// 	}
//
// 	function setBlending( blending, blendEquation, blendSrc, blendDst, blendEquationAlpha, blendSrcAlpha, blendDstAlpha, blendColor, blendAlpha, premultipliedAlpha ) {
//
// 		if ( blending === NoBlending ) {
//
// 			if ( currentBlendingEnabled === true ) {
//
// 				disable( gl.BLEND );
// 				currentBlendingEnabled = false;
//
// 			}
//
// 			return;
//
// 		}
//
// 		if ( currentBlendingEnabled === false ) {
//
// 			enable( gl.BLEND );
// 			currentBlendingEnabled = true;
//
// 		}
//
// 		if ( blending !== CustomBlending ) {
//
// 			if ( blending !== currentBlending || premultipliedAlpha !== currentPremultipledAlpha ) {
//
// 				if ( currentBlendEquation !== AddEquation || currentBlendEquationAlpha !== AddEquation ) {
//
// 					gl.blendEquation( gl.FUNC_ADD );
//
// 					currentBlendEquation = AddEquation;
// 					currentBlendEquationAlpha = AddEquation;
//
// 				}
//
// 				if ( premultipliedAlpha ) {
//
// 					switch ( blending ) {
//
// 						case NormalBlending:
// 							gl.blendFuncSeparate( gl.ONE, gl.ONE_MINUS_SRC_ALPHA, gl.ONE, gl.ONE_MINUS_SRC_ALPHA );
// 							break;
//
// 						case AdditiveBlending:
// 							gl.blendFunc( gl.ONE, gl.ONE );
// 							break;
//
// 						case SubtractiveBlending:
// 							gl.blendFuncSeparate( gl.ZERO, gl.ONE_MINUS_SRC_COLOR, gl.ZERO, gl.ONE );
// 							break;
//
// 						case MultiplyBlending:
// 							gl.blendFuncSeparate( gl.ZERO, gl.SRC_COLOR, gl.ZERO, gl.SRC_ALPHA );
// 							break;
//
// 						default:
// 							console.error( 'THREE.WebGLState: Invalid blending: ', blending );
// 							break;
//
// 					}
//
// 				} else {
//
// 					switch ( blending ) {
//
// 						case NormalBlending:
// 							gl.blendFuncSeparate( gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA, gl.ONE, gl.ONE_MINUS_SRC_ALPHA );
// 							break;
//
// 						case AdditiveBlending:
// 							gl.blendFunc( gl.SRC_ALPHA, gl.ONE );
// 							break;
//
// 						case SubtractiveBlending:
// 							gl.blendFuncSeparate( gl.ZERO, gl.ONE_MINUS_SRC_COLOR, gl.ZERO, gl.ONE );
// 							break;
//
// 						case MultiplyBlending:
// 							gl.blendFunc( gl.ZERO, gl.SRC_COLOR );
// 							break;
//
// 						default:
// 							console.error( 'THREE.WebGLState: Invalid blending: ', blending );
// 							break;
//
// 					}
//
// 				}
//
// 				currentBlendSrc = null;
// 				currentBlendDst = null;
// 				currentBlendSrcAlpha = null;
// 				currentBlendDstAlpha = null;
// 				currentBlendColor.set( 0, 0, 0 );
// 				currentBlendAlpha = 0;
//
// 				currentBlending = blending;
// 				currentPremultipledAlpha = premultipliedAlpha;
//
// 			}
//
// 			return;
//
// 		}
//
// 		// custom blending
//
// 		blendEquationAlpha = blendEquationAlpha || blendEquation;
// 		blendSrcAlpha = blendSrcAlpha || blendSrc;
// 		blendDstAlpha = blendDstAlpha || blendDst;
//
// 		if ( blendEquation !== currentBlendEquation || blendEquationAlpha !== currentBlendEquationAlpha ) {
//
// 			gl.blendEquationSeparate( equationToGL[ blendEquation ], equationToGL[ blendEquationAlpha ] );
//
// 			currentBlendEquation = blendEquation;
// 			currentBlendEquationAlpha = blendEquationAlpha;
//
// 		}
//
// 		if ( blendSrc !== currentBlendSrc || blendDst !== currentBlendDst || blendSrcAlpha !== currentBlendSrcAlpha || blendDstAlpha !== currentBlendDstAlpha ) {
//
// 			gl.blendFuncSeparate( factorToGL[ blendSrc ], factorToGL[ blendDst ], factorToGL[ blendSrcAlpha ], factorToGL[ blendDstAlpha ] );
//
// 			currentBlendSrc = blendSrc;
// 			currentBlendDst = blendDst;
// 			currentBlendSrcAlpha = blendSrcAlpha;
// 			currentBlendDstAlpha = blendDstAlpha;
//
// 		}
//
// 		if ( blendColor.equals( currentBlendColor ) === false || blendAlpha !== currentBlendAlpha ) {
//
// 			gl.blendColor( blendColor.r, blendColor.g, blendColor.b, blendAlpha );
//
// 			currentBlendColor.copy( blendColor );
// 			currentBlendAlpha = blendAlpha;
//
// 		}
//
// 		currentBlending = blending;
// 		currentPremultipledAlpha = false;
//
// 	}
//
// 	function setMaterial( material, frontFaceCW ) {
//
// 		material.side === DoubleSide
// 			? disable( gl.CULL_FACE )
// 			: enable( gl.CULL_FACE );
//
// 		let flipSided = ( material.side === BackSide );
// 		if ( frontFaceCW ) flipSided = ! flipSided;
//
// 		setFlipSided( flipSided );
//
// 		( material.blending === NormalBlending && material.transparent === false )
// 			? setBlending( NoBlending )
// 			: setBlending( material.blending, material.blendEquation, material.blendSrc, material.blendDst, material.blendEquationAlpha, material.blendSrcAlpha, material.blendDstAlpha, material.blendColor, material.blendAlpha, material.premultipliedAlpha );
//
// 		depthBuffer.setFunc( material.depthFunc );
// 		depthBuffer.setTest( material.depthTest );
// 		depthBuffer.setMask( material.depthWrite );
// 		colorBuffer.setMask( material.colorWrite );
//
// 		const stencilWrite = material.stencilWrite;
// 		stencilBuffer.setTest( stencilWrite );
// 		if ( stencilWrite ) {
//
// 			stencilBuffer.setMask( material.stencilWriteMask );
// 			stencilBuffer.setFunc( material.stencilFunc, material.stencilRef, material.stencilFuncMask );
// 			stencilBuffer.setOp( material.stencilFail, material.stencilZFail, material.stencilZPass );
//
// 		}
//
// 		setPolygonOffset( material.polygonOffset, material.polygonOffsetFactor, material.polygonOffsetUnits );
//
// 		material.alphaToCoverage === true
// 			? enable( gl.SAMPLE_ALPHA_TO_COVERAGE )
// 			: disable( gl.SAMPLE_ALPHA_TO_COVERAGE );
//
// 	}
//

func (s *State) SetFlipSided(flipSided bool) {
	if s.currentFlipSided == flipSided {
		return
	}
	if flipSided {
		s.gl.Call("frontFace", s.gl.CW)
	} else {
		s.gl.Call("frontFace", s.gl.CCW)
	}
	s.currentFlipSided = flipSided
}

func (s *State) SetCullFace(cullFace consts.CullFace) {
	if s.currentCullFace == cullFace {
		return
	}

	if cullFace != consts.CullFaceNone {
		s.Enable(s.gl.CULL_FACE)

		if cullFace == consts.CullFaceBack {
			s.gl.Call("cullFace", s.gl.BACK)
		} else if cullFace == consts.CullFaceFront {
			s.gl.Call("cullFace", s.gl.FRONT)
		} else {
			s.gl.Call("cullFace", s.gl.FRONT_AND_BACK)
		}
	} else {
		s.Disable(s.gl.CULL_FACE)
	}

	s.currentCullFace = cullFace
}

//todo
// 	function setLineWidth( width ) {
//
// 		if ( width !== currentLineWidth ) {
//
// 			if ( lineWidthAvailable ) gl.lineWidth( width );
//
// 			currentLineWidth = width;
//
// 		}
//
// 	}
//
// 	function setPolygonOffset( polygonOffset, factor, units ) {
//
// 		if ( polygonOffset ) {
//
// 			enable( gl.POLYGON_OFFSET_FILL );
//
// 			if ( currentPolygonOffsetFactor !== factor || currentPolygonOffsetUnits !== units ) {
//
// 				gl.polygonOffset( factor, units );
//
// 				currentPolygonOffsetFactor = factor;
// 				currentPolygonOffsetUnits = units;
//
// 			}
//
// 		} else {
//
// 			disable( gl.POLYGON_OFFSET_FILL );
//
// 		}
//
// 	}
//
// 	function setScissorTest( scissorTest ) {
//
// 		if ( scissorTest ) {
//
// 			enable( gl.SCISSOR_TEST );
//
// 		} else {
//
// 			disable( gl.SCISSOR_TEST );
//
// 		}
//
// 	}
//
// 	// texture
//
// 	function activeTexture( webglSlot ) {
//
// 		if ( webglSlot === undefined ) webglSlot = gl.TEXTURE0 + maxTextures - 1;
//
// 		if ( currentTextureSlot !== webglSlot ) {
//
// 			gl.activeTexture( webglSlot );
// 			currentTextureSlot = webglSlot;
//
// 		}
//
// 	}
//
// 	function bindTexture( webglType, webglTexture, webglSlot ) {
//
// 		if ( webglSlot === undefined ) {
//
// 			if ( currentTextureSlot === null ) {
//
// 				webglSlot = gl.TEXTURE0 + maxTextures - 1;
//
// 			} else {
//
// 				webglSlot = currentTextureSlot;
//
// 			}
//
// 		}
//
// 		let boundTexture = currentBoundTextures[ webglSlot ];
//
// 		if ( boundTexture === undefined ) {
//
// 			boundTexture = { type: undefined, texture: undefined };
// 			currentBoundTextures[ webglSlot ] = boundTexture;
//
// 		}
//
// 		if ( boundTexture.type !== webglType || boundTexture.texture !== webglTexture ) {
//
// 			if ( currentTextureSlot !== webglSlot ) {
//
// 				gl.activeTexture( webglSlot );
// 				currentTextureSlot = webglSlot;
//
// 			}
//
// 			gl.bindTexture( webglType, webglTexture || emptyTextures[ webglType ] );
//
// 			boundTexture.type = webglType;
// 			boundTexture.texture = webglTexture;
//
// 		}
//
// 	}
//
// 	function unbindTexture() {
//
// 		const boundTexture = currentBoundTextures[ currentTextureSlot ];
//
// 		if ( boundTexture !== undefined && boundTexture.type !== undefined ) {
//
// 			gl.bindTexture( boundTexture.type, null );
//
// 			boundTexture.type = undefined;
// 			boundTexture.texture = undefined;
//
// 		}
//
// 	}
//
// 	function compressedTexImage2D() {
//
// 		try {
//
// 			gl.compressedTexImage2D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	function compressedTexImage3D() {
//
// 		try {
//
// 			gl.compressedTexImage3D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	function texSubImage2D() {
//
// 		try {
//
// 			gl.texSubImage2D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	function texSubImage3D() {
//
// 		try {
//
// 			gl.texSubImage3D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	function compressedTexSubImage2D() {
//
// 		try {
//
// 			gl.compressedTexSubImage2D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	function compressedTexSubImage3D() {
//
// 		try {
//
// 			gl.compressedTexSubImage3D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	function texStorage2D() {
//
// 		try {
//
// 			gl.texStorage2D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	function texStorage3D() {
//
// 		try {
//
// 			gl.texStorage3D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	function texImage2D() {
//
// 		try {
//
// 			gl.texImage2D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	function texImage3D() {
//
// 		try {
//
// 			gl.texImage3D.apply( gl, arguments );
//
// 		} catch ( error ) {
//
// 			console.error( 'THREE.WebGLState:', error );
//
// 		}
//
// 	}
//
// 	//
//
// 	function scissor( scissor ) {
//
// 		if ( currentScissor.equals( scissor ) === false ) {
//
// 			gl.scissor( scissor.x, scissor.y, scissor.z, scissor.w );
// 			currentScissor.copy( scissor );
//
// 		}
//
// 	}
//

func (s *State) Viewport(viewport *math.Vector4) {
	if !s.currentViewport.Equals(viewport) {
		s.gl.Call("viewport", viewport.X, viewport.Y, viewport.Z, viewport.W)
		s.currentViewport.Copy(viewport)
	}
}

//todo
// 	function updateUBOMapping( uniformsGroup, program ) {
//
// 		let mapping = uboProgramMap.get( program );
//
// 		if ( mapping === undefined ) {
//
// 			mapping = new WeakMap();
//
// 			uboProgramMap.set( program, mapping );
//
// 		}
//
// 		let blockIndex = mapping.get( uniformsGroup );
//
// 		if ( blockIndex === undefined ) {
//
// 			blockIndex = gl.getUniformBlockIndex( program, uniformsGroup.name );
//
// 			mapping.set( uniformsGroup, blockIndex );
//
// 		}
//
// 	}
//
// 	function uniformBlockBinding( uniformsGroup, program ) {
//
// 		const mapping = uboProgramMap.get( program );
// 		const blockIndex = mapping.get( uniformsGroup );
//
// 		if ( uboBindings.get( program ) !== blockIndex ) {
//
// 			// bind shader specific block index to global block point
// 			gl.uniformBlockBinding( program, blockIndex, uniformsGroup.__bindingPointIndex );
//
// 			uboBindings.set( program, blockIndex );
//
// 		}
//
// 	}
//
// 	//
//
// 	function reset() {
//
// 		// reset state
//
// 		gl.disable( gl.BLEND );
// 		gl.disable( gl.CULL_FACE );
// 		gl.disable( gl.DEPTH_TEST );
// 		gl.disable( gl.POLYGON_OFFSET_FILL );
// 		gl.disable( gl.SCISSOR_TEST );
// 		gl.disable( gl.STENCIL_TEST );
// 		gl.disable( gl.SAMPLE_ALPHA_TO_COVERAGE );
//
// 		gl.blendEquation( gl.FUNC_ADD );
// 		gl.blendFunc( gl.ONE, gl.ZERO );
// 		gl.blendFuncSeparate( gl.ONE, gl.ZERO, gl.ONE, gl.ZERO );
// 		gl.blendColor( 0, 0, 0, 0 );
//
// 		gl.colorMask( true, true, true, true );
// 		gl.clearColor( 0, 0, 0, 0 );
//
// 		gl.depthMask( true );
// 		gl.depthFunc( gl.LESS );
//
// 		depthBuffer.setReversed( false );
//
// 		gl.clearDepth( 1 );
//
// 		gl.stencilMask( 0xffffffff );
// 		gl.stencilFunc( gl.ALWAYS, 0, 0xffffffff );
// 		gl.stencilOp( gl.KEEP, gl.KEEP, gl.KEEP );
// 		gl.clearStencil( 0 );
//
// 		gl.cullFace( gl.BACK );
// 		gl.frontFace( gl.CCW );
//
// 		gl.polygonOffset( 0, 0 );
//
// 		gl.activeTexture( gl.TEXTURE0 );
//
// 		gl.bindFramebuffer( gl.FRAMEBUFFER, null );
// 		gl.bindFramebuffer( gl.DRAW_FRAMEBUFFER, null );
// 		gl.bindFramebuffer( gl.READ_FRAMEBUFFER, null );
//
// 		gl.useProgram( null );
//
// 		gl.lineWidth( 1 );
//
// 		gl.scissor( 0, 0, gl.canvas.width, gl.canvas.height );
// 		gl.viewport( 0, 0, gl.canvas.width, gl.canvas.height );
//
// 		// reset internals
//
// 		enabledCapabilities = {};
//
// 		currentTextureSlot = null;
// 		currentBoundTextures = {};
//
// 		currentBoundFramebuffers = {};
// 		currentDrawbuffers = new WeakMap();
// 		defaultDrawbuffers = [];
//
// 		currentProgram = null;
//
// 		currentBlendingEnabled = false;
// 		currentBlending = null;
// 		currentBlendEquation = null;
// 		currentBlendSrc = null;
// 		currentBlendDst = null;
// 		currentBlendEquationAlpha = null;
// 		currentBlendSrcAlpha = null;
// 		currentBlendDstAlpha = null;
// 		currentBlendColor = new Color( 0, 0, 0 );
// 		currentBlendAlpha = 0;
// 		currentPremultipledAlpha = false;
//
// 		currentFlipSided = null;
// 		currentCullFace = null;
//
// 		currentLineWidth = null;
//
// 		currentPolygonOffsetFactor = null;
// 		currentPolygonOffsetUnits = null;
//
// 		currentScissor.set( 0, 0, gl.canvas.width, gl.canvas.height );
// 		currentViewport.set( 0, 0, gl.canvas.width, gl.canvas.height );
//
// 		colorBuffer.reset();
// 		depthBuffer.reset();
// 		stencilBuffer.reset();
//
// 	}
//