package consts

type Format int

// --- Formats WebGL 1 ---
const (
	UnsignedByteType      Format = 1009
	ByteType              Format = 1010
	ShortType             Format = 1011
	UnsignedShortType     Format = 1012
	IntType               Format = 1013
	UnsignedIntType       Format = 1014
	FloatType             Format = 1015
	HalfFloatType         Format = 1016
	UnsignedShort4444Type Format = 1017
	UnsignedShort5551Type Format = 1018
	UnsignedInt248Type    Format = 1020
	UnsignedInt5999Type   Format = 35902
	AlphaFormat           Format = 1021
	RGBFormat             Format = 1022
	RGBAFormat            Format = 1023
	LuminanceFormat       Format = 1024
	LuminanceAlphaFormat  Format = 1025
	DepthFormat           Format = 1026
	DepthStencilFormat    Format = 1027
)

// --- Formats WebGL 2 ---
const (
	RedFormat         Format = 1028
	RedIntegerFormat  Format = 1029
	RGFormat          Format = 1030
	RGIntegerFormat   Format = 1031
	RGBIntegerFormat  Format = 1032
	RGBAIntegerFormat Format = 1033
)

// --- Formats PVRTC ---
//
//goland:noinspection GoSnakeCaseUsage
const (
	RGB_PVRTC_4BPPV1_Format  Format = 35840
	RGB_PVRTC_2BPPV1_Format  Format = 35841
	RGBA_PVRTC_4BPPV1_Format Format = 35842
	RGBA_PVRTC_2BPPV1_Format Format = 35843
)

// --- Formats S3TC ---
//
//goland:noinspection GoSnakeCaseUsage
const (
	RGB_S3TC_DXT1_Format  Format = 33776
	RGBA_S3TC_DXT1_Format Format = 33777
	RGBA_S3TC_DXT3_Format Format = 33778
	RGBA_S3TC_DXT5_Format Format = 33779
)

// --- Formats ETC ---
//
//goland:noinspection GoSnakeCaseUsage
const (
	RGB_ETC1_Format      Format = 36196
	RGB_ETC2_Format      Format = 37492
	RGBA_ETC2_EAC_Format Format = 37496
)

// --- Formats ASTC ---
//
//goland:noinspection GoSnakeCaseUsage
const (
	RGBA_ASTC_4x4_Format   Format = 37808
	RGBA_ASTC_5x4_Format   Format = 37809
	RGBA_ASTC_5x5_Format   Format = 37810
	RGBA_ASTC_6x5_Format   Format = 37811
	RGBA_ASTC_6x6_Format   Format = 37812
	RGBA_ASTC_8x5_Format   Format = 37813
	RGBA_ASTC_8x6_Format   Format = 37814
	RGBA_ASTC_8x8_Format   Format = 37815
	RGBA_ASTC_10x5_Format  Format = 37816
	RGBA_ASTC_10x6_Format  Format = 37817
	RGBA_ASTC_10x8_Format  Format = 37818
	RGBA_ASTC_10x10_Format Format = 37819
	RGBA_ASTC_12x10_Format Format = 37820
	RGBA_ASTC_12x12_Format Format = 37821
)

// --- Formats BPTC ---
//
//goland:noinspection GoSnakeCaseUsage
const (
	RGBA_BPTC_Format         Format = 36492
	RGB_BPTC_SIGNED_Format   Format = 36494
	RGB_BPTC_UNSIGNED_Format Format = 36495
)

// --- Formats RGTC ---
//
//goland:noinspection GoSnakeCaseUsage
const (
	RED_RGTC1_Format              Format = 36283
	SIGNED_RED_RGTC1_Format       Format = 36284
	RED_GREEN_RGTC2_Format        Format = 36285
	SIGNED_RED_GREEN_RGTC2_Format Format = 36286
)
