package webgl

import "fmt"

type Extensions struct {
	gl         Context
	extensions map[string]GLExtension
}

func NewWebGLExtensions(gl Context) *Extensions {
	return &Extensions{gl, make(map[string]GLExtension)}
}

func (e *Extensions) getExtension(name string) (ext GLExtension) {
	if ext, ok := e.extensions[name]; ok {
		return ext
	}

	ext = e.gl.getExtension(name)
	if !ext.Truthy() {
		switch name {
		case "WEBGL_depth_texture":
			ext = e.gl.getExtension("MOZ_WEBGL_depth_texture")
			if !ext.Truthy() {
				ext = e.gl.getExtension("WEBKIT_WEBGL_depth_texture")
			}
		case "EXT_texture_filter_anisotropic":
			ext = e.gl.getExtension("MOZ_EXT_texture_filter_anisotropic")
			if !ext.Truthy() {
				ext = e.gl.getExtension("WEBKIT_EXT_texture_filter_anisotropic")
			}
		case "WEBGL_compressed_texture_s3tc":
			ext = e.gl.getExtension("MOZ_WEBGL_compressed_texture_s3tc")
			if !ext.Truthy() {
				ext = e.gl.getExtension("WEBKIT_WEBGL_compressed_texture_s3tc")
			}
		case "WEBGL_compressed_texture_pvrtc":
			ext = e.gl.getExtension("WEBKIT_WEBGL_compressed_texture_pvrtc")
		}
	}

	e.extensions[name] = ext
	return
}

func (e *Extensions) Has(name string) bool {
	return !e.getExtension(name).IsNull()
}

func (e *Extensions) Init() {
	if e.extensions == nil {
		panic("use NewWebGLExtensions() for initalization")
	}

	e.getExtension("EXT_color_buffer_float")
	e.getExtension("WEBGL_clip_cull_distance")
	e.getExtension("OES_texture_float_linear")
	e.getExtension("EXT_color_buffer_half_float")
	e.getExtension("WEBGL_multisampled_render_to_texture")
	e.getExtension("WEBGL_render_shared_exponent")
}

func (e *Extensions) Get(name string) (ext GLExtension) {
	ext = e.getExtension(name)

	if ext.IsNull() {
		fmt.Println("THREE.WebGLRenderer: " + name + " extension not supported.")
	}

	return
}
