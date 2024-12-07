package webgl

//goland:noinspection GoSnakeCaseUsage
type ContextConsts struct {
	ALPHA                            ConstValue
	BYTE                             ConstValue
	DEPTH_COMPONENT                  ConstValue
	DEPTH_STENCIL                    ConstValue
	FLOAT                            ConstValue
	HALF_FLOAT                       ConstValue
	INT                              ConstValue
	LUMINANCE                        ConstValue
	LUMINANCE_ALPHA                  ConstValue
	MAX_COMBINED_TEXTURE_IMAGE_UNITS ConstValue
	RED                              ConstValue
	RED_INTEGER                      ConstValue
	RG                               ConstValue
	RG_INTEGER                       ConstValue
	RGB                              ConstValue
	RGB_INTEGER                      ConstValue
	RGBA                             ConstValue
	RGBA_INTEGER                     ConstValue
	SHORT                            ConstValue
	UNSIGNED_BYTE                    ConstValue
	UNSIGNED_SHORT                   ConstValue
	UNSIGNED_SHORT_4_4_4_4           ConstValue
	UNSIGNED_SHORT_5_5_5_1           ConstValue
	UNSIGNED_INT                     ConstValue
	UNSIGNED_INT_24_8                ConstValue
	UNSIGNED_INT_5_9_9_9_REV         ConstValue
	VERSION                          ConstValue
}

func (gl *Context) InitContextConsts() {
	gl.ALPHA = gl.Const("ALPHA")
	gl.BYTE = gl.Const("BYTE")
	gl.DEPTH_COMPONENT = gl.Const("DEPTH_COMPONENT")
	gl.DEPTH_STENCIL = gl.Const("DEPTH_STENCIL")
	gl.FLOAT = gl.Const("FLOAT")
	gl.HALF_FLOAT = gl.Const("HALF_FLOAT")
	gl.INT = gl.Const("INT")
	gl.LUMINANCE = gl.Const("LUMINANCE")
	gl.LUMINANCE_ALPHA = gl.Const("LUMINANCE_ALPHA")
	gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS = gl.Const("MAX_COMBINED_TEXTURE_IMAGE_UNITS")
	gl.RED = gl.Const("RED")
	gl.RED_INTEGER = gl.Const("RED_INTEGER")
	gl.RG = gl.Const("RG")
	gl.RG_INTEGER = gl.Const("RG_INTEGER")
	gl.RGB = gl.Const("RGB")
	gl.RGB_INTEGER = gl.Const("RGB_INTEGER")
	gl.RGBA = gl.Const("RGBA")
	gl.RGBA_INTEGER = gl.Const("RGBA_INTEGER")
	gl.SHORT = gl.Const("SHORT")
	gl.UNSIGNED_BYTE = gl.Const("UNSIGNED_BYTE")
	gl.UNSIGNED_SHORT = gl.Const("UNSIGNED_SHORT")
	gl.UNSIGNED_SHORT_4_4_4_4 = gl.Const("UNSIGNED_SHORT_4_4_4_4")
	gl.UNSIGNED_SHORT_5_5_5_1 = gl.Const("UNSIGNED_SHORT_5_5_5_1")
	gl.UNSIGNED_INT = gl.Const("UNSIGNED_INT")
	gl.UNSIGNED_INT_24_8 = gl.Const("UNSIGNED_INT_24_8")
	gl.UNSIGNED_INT_5_9_9_9_REV = gl.Const("UNSIGNED_INT_5_9_9_9_REV")
	gl.VERSION = gl.Const("VERSION")
}
