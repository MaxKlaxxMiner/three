package webgl

import . "github.com/MaxKlaxxMiner/three/consts"

type Utils struct {
	gl         Context
	extensions Extensions
}

func NewWebGLUtils(gl Context, extensions Extensions) *Utils {
	return &Utils{gl, extensions}
}

func (u *Utils) Convert(p Format) ConstValue {
	return u.ConvertWithColorspace(p, NoColorSpace)
}

func (u *Utils) ConvertWithColorspace(p Format, colorSpace ColorSpace) ConstValue {
	switch p {
	case UnsignedByteType:
		return u.gl.UNSIGNED_BYTE
	case UnsignedShort4444Type:
		return u.gl.UNSIGNED_SHORT_4_4_4_4
	case UnsignedShort5551Type:
		return u.gl.UNSIGNED_SHORT_5_5_5_1
	case UnsignedInt248Type:
		return u.gl.UNSIGNED_INT_24_8
	case UnsignedInt5999Type:
		return u.gl.UNSIGNED_INT_5_9_9_9_REV

	case ByteType:
		return u.gl.BYTE
	case ShortType:
		return u.gl.SHORT
	case UnsignedShortType:
		return u.gl.UNSIGNED_SHORT
	case IntType:
		return u.gl.INT
	case UnsignedIntType:
		return u.gl.UNSIGNED_INT
	case FloatType:
		return u.gl.FLOAT
	case HalfFloatType:
		return u.gl.HALF_FLOAT

	case AlphaFormat:
		return u.gl.ALPHA
	case RGBFormat:
		return u.gl.RGB
	case RGBAFormat:
		return u.gl.RGBA
	case LuminanceFormat:
		return u.gl.LUMINANCE
	case LuminanceAlphaFormat:
		return u.gl.LUMINANCE_ALPHA
	case DepthFormat:
		return u.gl.DEPTH_COMPONENT
	case DepthStencilFormat:
		return u.gl.DEPTH_STENCIL

	// --- WebGL2 formats ---
	case RedFormat:
		return u.gl.RED
	case RedIntegerFormat:
		return u.gl.RED_INTEGER
	case RGFormat:
		return u.gl.RG
	case RGIntegerFormat:
		return u.gl.RG_INTEGER
	case RGBIntegerFormat:
		return u.gl.RGB_INTEGER
	case RGBAIntegerFormat:
		return u.gl.RGBA_INTEGER

	// --- S3TC ---
	case RGB_S3TC_DXT1_Format, RGBA_S3TC_DXT1_Format, RGBA_S3TC_DXT3_Format, RGBA_S3TC_DXT5_Format:
		panic("todo")
	//todo
	//		const transfer = ColorManagement.getTransfer( colorSpace );
	//			if ( transfer === SRGBTransfer ) {
	//
	//				extension = extensions.get( 'WEBGL_compressed_texture_s3tc_srgb' );
	//
	//				if ( extension !== null ) {
	//
	//					if ( p === RGB_S3TC_DXT1_Format ) return extension.COMPRESSED_SRGB_S3TC_DXT1_EXT;
	//					if ( p === RGBA_S3TC_DXT1_Format ) return extension.COMPRESSED_SRGB_ALPHA_S3TC_DXT1_EXT;
	//					if ( p === RGBA_S3TC_DXT3_Format ) return extension.COMPRESSED_SRGB_ALPHA_S3TC_DXT3_EXT;
	//					if ( p === RGBA_S3TC_DXT5_Format ) return extension.COMPRESSED_SRGB_ALPHA_S3TC_DXT5_EXT;
	//
	//				} else {
	//
	//					return null;
	//
	//				}
	//
	//			} else {
	//
	//				extension = extensions.get( 'WEBGL_compressed_texture_s3tc' );
	//
	//				if ( extension !== null ) {
	//
	//					if ( p === RGB_S3TC_DXT1_Format ) return extension.COMPRESSED_RGB_S3TC_DXT1_EXT;
	//					if ( p === RGBA_S3TC_DXT1_Format ) return extension.COMPRESSED_RGBA_S3TC_DXT1_EXT;
	//					if ( p === RGBA_S3TC_DXT3_Format ) return extension.COMPRESSED_RGBA_S3TC_DXT3_EXT;
	//					if ( p === RGBA_S3TC_DXT5_Format ) return extension.COMPRESSED_RGBA_S3TC_DXT5_EXT;
	//
	//				} else {
	//
	//					return null;
	//
	//				}
	//
	//			}
	//

	// --- PVRTC ---
	case RGB_PVRTC_4BPPV1_Format, RGB_PVRTC_2BPPV1_Format, RGBA_PVRTC_4BPPV1_Format, RGBA_PVRTC_2BPPV1_Format:
		extension := u.extensions.Get("WEBGL_compressed_texture_pvrtc")
		if extension.IsNull() {
			return 0
		}
		switch p {
		case RGB_PVRTC_4BPPV1_Format:
			return extension.Const("COMPRESSED_RGB_PVRTC_4BPPV1_IMG")
		case RGB_PVRTC_2BPPV1_Format:
			return extension.Const("COMPRESSED_RGB_PVRTC_2BPPV1_IMG")
		case RGBA_PVRTC_4BPPV1_Format:
			return extension.Const("COMPRESSED_RGBA_PVRTC_4BPPV1_IMG")
		default: // case RGBA_PVRTC_2BPPV1_Format:
			return extension.Const("COMPRESSED_RGBA_PVRTC_2BPPV1_IMG")
		}

	// --- ETC ---
	case RGB_ETC1_Format, RGB_ETC2_Format, RGBA_ETC2_EAC_Format:
		extension := u.extensions.Get("WEBGL_compressed_texture_etc")
		if extension.IsNull() {
			return 0
		}
		panic("todo")
	//todo
	//		const transfer = ColorManagement.getTransfer( colorSpace );
	//	if p == RGB_ETC1_Format || p == RGB_ETC2_Format {
	//		//		todo return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ETC2 : extension.COMPRESSED_RGB8_ETC2;
	//	}
	//	if p == RGBA_ETC2_EAC_Format
	//	todo			return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ETC2_EAC : extension.COMPRESSED_RGBA8_ETC2_EAC;

	//		// ASTC
	//todo
	//		if ( p === RGBA_ASTC_4x4_Format || p === RGBA_ASTC_5x4_Format || p === RGBA_ASTC_5x5_Format ||
	//			p === RGBA_ASTC_6x5_Format || p === RGBA_ASTC_6x6_Format || p === RGBA_ASTC_8x5_Format ||
	//			p === RGBA_ASTC_8x6_Format || p === RGBA_ASTC_8x8_Format || p === RGBA_ASTC_10x5_Format ||
	//			p === RGBA_ASTC_10x6_Format || p === RGBA_ASTC_10x8_Format || p === RGBA_ASTC_10x10_Format ||
	//			p === RGBA_ASTC_12x10_Format || p === RGBA_ASTC_12x12_Format ) {
	//
	//			extension = extensions.get( 'WEBGL_compressed_texture_astc' );
	//
	//			if ( extension !== null ) {
	//
	//				if ( p === RGBA_ASTC_4x4_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR : extension.COMPRESSED_RGBA_ASTC_4x4_KHR;
	//				if ( p === RGBA_ASTC_5x4_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR : extension.COMPRESSED_RGBA_ASTC_5x4_KHR;
	//				if ( p === RGBA_ASTC_5x5_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR : extension.COMPRESSED_RGBA_ASTC_5x5_KHR;
	//				if ( p === RGBA_ASTC_6x5_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR : extension.COMPRESSED_RGBA_ASTC_6x5_KHR;
	//				if ( p === RGBA_ASTC_6x6_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR : extension.COMPRESSED_RGBA_ASTC_6x6_KHR;
	//				if ( p === RGBA_ASTC_8x5_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR : extension.COMPRESSED_RGBA_ASTC_8x5_KHR;
	//				if ( p === RGBA_ASTC_8x6_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR : extension.COMPRESSED_RGBA_ASTC_8x6_KHR;
	//				if ( p === RGBA_ASTC_8x8_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR : extension.COMPRESSED_RGBA_ASTC_8x8_KHR;
	//				if ( p === RGBA_ASTC_10x5_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR : extension.COMPRESSED_RGBA_ASTC_10x5_KHR;
	//				if ( p === RGBA_ASTC_10x6_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR : extension.COMPRESSED_RGBA_ASTC_10x6_KHR;
	//				if ( p === RGBA_ASTC_10x8_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR : extension.COMPRESSED_RGBA_ASTC_10x8_KHR;
	//				if ( p === RGBA_ASTC_10x10_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR : extension.COMPRESSED_RGBA_ASTC_10x10_KHR;
	//				if ( p === RGBA_ASTC_12x10_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR : extension.COMPRESSED_RGBA_ASTC_12x10_KHR;
	//				if ( p === RGBA_ASTC_12x12_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR : extension.COMPRESSED_RGBA_ASTC_12x12_KHR;
	//
	//			} else {
	//
	//				return null;
	//
	//			}
	//
	//		}
	//
	//		// BPTC
	//
	//		if ( p === RGBA_BPTC_Format || p === RGB_BPTC_SIGNED_Format || p === RGB_BPTC_UNSIGNED_Format ) {
	//
	//			extension = extensions.get( 'EXT_texture_compression_bptc' );
	//
	//			if ( extension !== null ) {
	//
	//				if ( p === RGBA_BPTC_Format ) return ( transfer === SRGBTransfer ) ? extension.COMPRESSED_SRGB_ALPHA_BPTC_UNORM_EXT : extension.COMPRESSED_RGBA_BPTC_UNORM_EXT;
	//				if ( p === RGB_BPTC_SIGNED_Format ) return extension.COMPRESSED_RGB_BPTC_SIGNED_FLOAT_EXT;
	//				if ( p === RGB_BPTC_UNSIGNED_Format ) return extension.COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_EXT;
	//
	//			} else {
	//
	//				return null;
	//
	//			}
	//
	//		}
	//
	//		// RGTC
	//
	//		if ( p === RED_RGTC1_Format || p === SIGNED_RED_RGTC1_Format || p === RED_GREEN_RGTC2_Format || p === SIGNED_RED_GREEN_RGTC2_Format ) {
	//
	//			extension = extensions.get( 'EXT_texture_compression_rgtc' );
	//
	//			if ( extension !== null ) {
	//
	//				if ( p === RGBA_BPTC_Format ) return extension.COMPRESSED_RED_RGTC1_EXT;
	//				if ( p === SIGNED_RED_RGTC1_Format ) return extension.COMPRESSED_SIGNED_RED_RGTC1_EXT;
	//				if ( p === RED_GREEN_RGTC2_Format ) return extension.COMPRESSED_RED_GREEN_RGTC2_EXT;
	//				if ( p === SIGNED_RED_GREEN_RGTC2_Format ) return extension.COMPRESSED_SIGNED_RED_GREEN_RGTC2_EXT;
	//
	//			} else {
	//
	//				return null;
	//
	//			}
	//
	//		}
	//
	//		//
	//
	//		if ( p === UnsignedInt248Type ) return gl.UNSIGNED_INT_24_8;
	//
	//
	//
	default:
		panic("todo")
		// if "p" can't be resolved, assume the user defines a WebGL constant as a string (fallback/workaround for packed RGB formats)
		//		return ( gl[ p ] !== undefined ) ? gl[ p ] : null; todo
		return 0
	}
}
