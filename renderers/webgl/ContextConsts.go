package webgl

//goland:noinspection GoSnakeCaseUsage
type ContextConsts struct {
	ALPHA                            int32
	BACK                             int32
	BYTE                             int32
	CCW                              int32
	CULL_FACE                        int32
	CW                               int32
	DEPTH_COMPONENT                  int32
	DEPTH_STENCIL                    int32
	DEPTH_TEST                       int32
	FLOAT                            int32
	FRONT                            int32
	FRONT_AND_BACK                   int32
	HALF_FLOAT                       int32
	INT                              int32
	LINE_LOOP                        int32
	LINE_STRIP                       int32
	LINES                            int32
	LUMINANCE                        int32
	LUMINANCE_ALPHA                  int32
	MAX_COMBINED_TEXTURE_IMAGE_UNITS int32
	POINTS                           int32
	RED                              int32
	RED_INTEGER                      int32
	RG                               int32
	RG_INTEGER                       int32
	RGB                              int32
	RGB_INTEGER                      int32
	RGBA                             int32
	RGBA_INTEGER                     int32
	SHORT                            int32
	TRIANGLES                        int32
	UNSIGNED_BYTE                    int32
	UNSIGNED_SHORT                   int32
	UNSIGNED_SHORT_4_4_4_4           int32
	UNSIGNED_SHORT_5_5_5_1           int32
	UNSIGNED_INT                     int32
	UNSIGNED_INT_24_8                int32
	UNSIGNED_INT_5_9_9_9_REV         int32
	VERSION                          int32
}

func (gl *Context) InitContextConsts() {
	gl.ALPHA = gl.Const("ALPHA")
	gl.BACK = gl.Const("BACK")
	gl.BYTE = gl.Const("BYTE")
	gl.CCW = gl.Const("CCW")
	gl.CULL_FACE = gl.Const("CULL_FACE")
	gl.CW = gl.Const("CW")
	gl.DEPTH_COMPONENT = gl.Const("DEPTH_COMPONENT")
	gl.DEPTH_STENCIL = gl.Const("DEPTH_STENCIL")
	gl.DEPTH_TEST = gl.Const("DEPTH_TEST")
	gl.FLOAT = gl.Const("FLOAT")
	gl.FRONT = gl.Const("FRONT")
	gl.FRONT_AND_BACK = gl.Const("FRONT_AND_BACK")
	gl.HALF_FLOAT = gl.Const("HALF_FLOAT")
	gl.INT = gl.Const("INT")
	gl.LINE_LOOP = gl.Const("LINE_LOOP")
	gl.LINE_STRIP = gl.Const("LINE_STRIP")
	gl.LINES = gl.Const("LINES")
	gl.LUMINANCE = gl.Const("LUMINANCE")
	gl.LUMINANCE_ALPHA = gl.Const("LUMINANCE_ALPHA")
	gl.MAX_COMBINED_TEXTURE_IMAGE_UNITS = gl.Const("MAX_COMBINED_TEXTURE_IMAGE_UNITS")
	gl.POINTS = gl.Const("POINTS")
	gl.RED = gl.Const("RED")
	gl.RED_INTEGER = gl.Const("RED_INTEGER")
	gl.RG = gl.Const("RG")
	gl.RG_INTEGER = gl.Const("RG_INTEGER")
	gl.RGB = gl.Const("RGB")
	gl.RGB_INTEGER = gl.Const("RGB_INTEGER")
	gl.RGBA = gl.Const("RGBA")
	gl.RGBA_INTEGER = gl.Const("RGBA_INTEGER")
	gl.SHORT = gl.Const("SHORT")
	gl.TRIANGLES = gl.Const("TRIANGLES")
	gl.UNSIGNED_BYTE = gl.Const("UNSIGNED_BYTE")
	gl.UNSIGNED_SHORT = gl.Const("UNSIGNED_SHORT")
	gl.UNSIGNED_SHORT_4_4_4_4 = gl.Const("UNSIGNED_SHORT_4_4_4_4")
	gl.UNSIGNED_SHORT_5_5_5_1 = gl.Const("UNSIGNED_SHORT_5_5_5_1")
	gl.UNSIGNED_INT = gl.Const("UNSIGNED_INT")
	gl.UNSIGNED_INT_24_8 = gl.Const("UNSIGNED_INT_24_8")
	gl.UNSIGNED_INT_5_9_9_9_REV = gl.Const("UNSIGNED_INT_5_9_9_9_REV")
	gl.VERSION = gl.Const("VERSION")
}
