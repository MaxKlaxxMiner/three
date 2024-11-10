package three

import (
	"github.com/MaxKlaxxMiner/three/mathutils"
	"math"
)

type PerspectiveCamera struct {
	Camera
	Fov        float64 // Camera frustum vertical field of view, from bottom to top of view, in degrees. Default is 50.
	Zoom       float64 // Gets or sets the zoom factor of the camera. Default is 1.
	Near       float64 // Camera frustum near plane. Default is 0.1. The valid range is greater than 0 and less than the current value of the far plane. Note that, unlike for the OrthographicCamera, 0 is not a valid value for a PerspectiveCamera's near plane.
	Far        float64 // Camera frustum far plane. Default is 2000.
	Focus      float64 // Object distance used for stereoscopy and depth-of-field effects. This parameter does not influence the projection matrix unless a StereoCamera is being used. Default is 10.
	Aspect     float64 // Camera frustum aspect ratio, usually the canvas width / canvas height. Default is 1 (square canvas).
	FilmGauge  float64 // Film size used for the larger axis. Default is 35 (millimeters). This parameter does not influence the projection matrix unless FilmOffset is set to a nonzero value.
	FilmOffset float64 // Horizontal off-center offset in the same unit as FilmGauge. Default is 0.
}

// NewPerspectiveCamera Camera that uses perspective projection.
//
//	fov: Camera frustum vertical field of view.
//	aspect: Camera frustum aspect ratio.
//	near: Camera frustum near plane.
//	far: Camera frustum far plane.
func NewPerspectiveCamera(fov, aspect, near, far float64) *PerspectiveCamera {
	r := &PerspectiveCamera{
		Fov:       fov,
		Zoom:      1,
		Near:      near,
		Far:       far,
		Focus:     10,
		Aspect:    aspect,
		FilmGauge: 35,
	}
	r.UpdateProjectionMatrix()
	return r
}

// todo
//
//	copy( source, recursive ) {
//
//		super.copy( source, recursive );
//
//		this.fov = source.fov;
//		this.zoom = source.zoom;
//
//		this.near = source.near;
//		this.far = source.far;
//		this.focus = source.focus;
//
//		this.aspect = source.aspect;
//		this.view = source.view === null ? null : Object.assign( {}, source.view );
//
//		this.filmGauge = source.filmGauge;
//		this.filmOffset = source.filmOffset;
//
//		return this;
//
//	}
//
//	**
//	 * Sets the FOV by focal length in respect to the current .filmGauge.
//	 *
//	 * The default film gauge is 35, so that the focal length can be specified for
//	 * a 35mm (full frame) camera.
//	 *
//	 * Values for focal length and film gauge must have the same unit.
//	 *
//	setFocalLength( focalLength ) {
//
//		** see {@link http://www.bobatkins.com/photography/technical/field_of_view.html} *
//		const vExtentSlope = 0.5 * this.getFilmHeight() / focalLength;
//
//		this.fov = MathUtils.RAD2DEG * 2 * Math.atan( vExtentSlope );
//		this.updateProjectionMatrix();
//
//	}
//
//	**
//	 * Calculates the focal length from the current .fov and .filmGauge.
//	 *
//	getFocalLength() {
//
//		const vExtentSlope = Math.tan( MathUtils.DEG2RAD * 0.5 * this.fov );
//
//		return 0.5 * this.getFilmHeight() / vExtentSlope;
//
//	}
//
//	getEffectiveFOV() {
//
//		return MathUtils.RAD2DEG * 2 * Math.atan(
//			Math.tan( MathUtils.DEG2RAD * 0.5 * this.fov ) / this.zoom );
//
//	}
//
//	getFilmWidth() {
//
//		// film not completely covered in portrait format (aspect < 1)
//		return this.filmGauge * Math.min( this.aspect, 1 );
//
//	}
//
//	getFilmHeight() {
//
//		// film not completely covered in landscape format (aspect > 1)
//		return this.filmGauge / Math.max( this.aspect, 1 );
//
//	}
//
//	**
//	 * Computes the 2D bounds of the camera's viewable rectangle at a given distance along the viewing direction.
//	 * Sets minTarget and maxTarget to the coordinates of the lower-left and upper-right corners of the view rectangle.
//	 *
//	getViewBounds( distance, minTarget, maxTarget ) {
//
//		_v3.set( - 1, - 1, 0.5 ).applyMatrix4( this.projectionMatrixInverse );
//
//		minTarget.set( _v3.x, _v3.y ).multiplyScalar( - distance / _v3.z );
//
//		_v3.set( 1, 1, 0.5 ).applyMatrix4( this.projectionMatrixInverse );
//
//		maxTarget.set( _v3.x, _v3.y ).multiplyScalar( - distance / _v3.z );
//
//	}
//
//	**
//	 * Computes the width and height of the camera's viewable rectangle at a given distance along the viewing direction.
//	 * Copies the result into the target Vector2, where x is width and y is height.
//	 *
//	getViewSize( distance, target ) {
//
//		this.getViewBounds( distance, _minTarget, _maxTarget );
//
//		return target.subVectors( _maxTarget, _minTarget );
//
//	}
//
//	**
//	 * Sets an offset in a larger frustum. This is useful for multi-window or
//	 * multi-monitor/multi-machine setups.
//	 *
//	 * For example, if you have 3x2 monitors and each monitor is 1920x1080 and
//	 * the monitors are in grid like this
//	 *
//	 *   +---+---+---+
//	 *   | A | B | C |
//	 *   +---+---+---+
//	 *   | D | E | F |
//	 *   +---+---+---+
//	 *
//	 * then for each monitor you would call it like this
//	 *
//	 *   const w = 1920;
//	 *   const h = 1080;
//	 *   const fullWidth = w * 3;
//	 *   const fullHeight = h * 2;
//	 *
//	 *   --A--
//	 *   camera.setViewOffset( fullWidth, fullHeight, w * 0, h * 0, w, h );
//	 *   --B--
//	 *   camera.setViewOffset( fullWidth, fullHeight, w * 1, h * 0, w, h );
//	 *   --C--
//	 *   camera.setViewOffset( fullWidth, fullHeight, w * 2, h * 0, w, h );
//	 *   --D--
//	 *   camera.setViewOffset( fullWidth, fullHeight, w * 0, h * 1, w, h );
//	 *   --E--
//	 *   camera.setViewOffset( fullWidth, fullHeight, w * 1, h * 1, w, h );
//	 *   --F--
//	 *   camera.setViewOffset( fullWidth, fullHeight, w * 2, h * 1, w, h );
//	 *
//	 *   Note there is no reason monitors have to be the same size or in a grid.
//	 *
//	setViewOffset( fullWidth, fullHeight, x, y, width, height ) {
//
//		this.aspect = fullWidth / fullHeight;
//
//		if ( this.view === null ) {
//
//			this.view = {
//				enabled: true,
//				fullWidth: 1,
//				fullHeight: 1,
//				offsetX: 0,
//				offsetY: 0,
//				width: 1,
//				height: 1
//			};
//
//		}
//
//		this.view.enabled = true;
//		this.view.fullWidth = fullWidth;
//		this.view.fullHeight = fullHeight;
//		this.view.offsetX = x;
//		this.view.offsetY = y;
//		this.view.width = width;
//		this.view.height = height;
//
//		this.updateProjectionMatrix();
//
//	}
//
//	clearViewOffset() {
//
//		if ( this.view !== null ) {
//
//			this.view.enabled = false;
//
//		}
//
//		this.updateProjectionMatrix();
//
//	}
//

func (c *PerspectiveCamera) UpdateProjectionMatrix() {
	near := c.Near
	top := near * math.Tan(mathutils.DEG2RAD*0.5*c.Fov) / c.Zoom
	height := 2 * top
	width := c.Aspect * height
	left := -0.5 * width

	//todo
	//		const view = this.view;
	//		if ( this.view !== null && this.view.enabled ) {
	//			const fullWidth = view.fullWidth,
	//				fullHeight = view.fullHeight;
	//			left += view.offsetX * width / fullWidth;
	//			top -= view.offsetY * height / fullHeight;
	//			width *= view.width / fullWidth;
	//			height *= view.height / fullHeight;
	//		}
	//		const skew = this.filmOffset;
	//		if ( skew !== 0 ) left += near * skew / this.getFilmWidth();
	//

	//		this.projectionMatrix.makePerspective( left, left + width, top, top - height, near, this.far, this.coordinateSystem );
	//
	//		this.projectionMatrixInverse.copy( this.projectionMatrix ).invert();
	_ = left
}

//todo

//	toJSON( meta ) {
//
//		const data = super.toJSON( meta );
//
//		data.object.fov = this.fov;
//		data.object.zoom = this.zoom;
//
//		data.object.near = this.near;
//		data.object.far = this.far;
//		data.object.focus = this.focus;
//
//		data.object.aspect = this.aspect;
//
//		if ( this.view !== null ) data.object.view = Object.assign( {}, this.view );
//
//		data.object.filmGauge = this.filmGauge;
//		data.object.filmOffset = this.filmOffset;
//
//		return data;
//
//	}
//
//}
