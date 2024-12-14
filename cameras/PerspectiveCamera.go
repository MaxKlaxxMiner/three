package cameras

import (
	math3 "github.com/MaxKlaxxMiner/three/math"
	"math"
)

type PerspectiveCamera struct {
	Camera
	Fov        float64
	Zoom       float64
	Near       float64
	Far        float64
	Focus      float64
	Aspect     float64
	View       *View
	FilmGauge  float64
	FilmOffset float64
}

type View struct {
	Enabled    bool
	FullWidth  float64
	FullHeight float64
	OffsetX    float64
	OffsetY    float64
	Width      float64
	Height     float64
}

func NewPerspectiveCameraDefaults() *PerspectiveCamera {
	return NewPerspectiveCamera(50, 1, 0.1, 2000)
}

func NewPerspectiveCamera(fov, aspect, near, far float64) *PerspectiveCamera {
	this := new(PerspectiveCamera)
	this.Camera = *NewCamera()
	this.This = this
	this.Type = "PerspectiveCamera"
	this.Fov = fov
	this.Zoom = 1
	this.Near = near
	this.Far = far
	this.Focus = 10
	this.Aspect = aspect
	this.View = nil
	this.FilmGauge = 35 // width of the film (default in millimeters)
	this.FilmOffset = 0 // horizontal film offset (same unit as gauge)
	this.UpdateProjectionMatrix()
	return this
}

func (c *PerspectiveCamera) IsPerspectiveCamera() bool { return c != nil }

func (c *PerspectiveCamera) Copy(source *PerspectiveCamera) *PerspectiveCamera {
	return c.CopyRecursive(source, true)
}

func (c *PerspectiveCamera) CopyRecursive(source *PerspectiveCamera, recursive bool) *PerspectiveCamera {
	c.Camera.CopyRecursive(&source.Camera, recursive)

	c.Fov = source.Fov
	c.Zoom = source.Zoom

	c.Near = source.Near
	c.Far = source.Far
	c.Focus = source.Focus

	c.Aspect = source.Aspect
	if source.View != nil {
		c.View = new(View)
		*c.View = *source.View
	} else {
		c.View = nil
	}

	c.FilmGauge = source.FilmGauge
	c.FilmOffset = source.FilmOffset

	return c
}

func (c *PerspectiveCamera) SetFocalLength(focalLength float64) {
	/**
	 * Sets the FOV by focal length in respect to the current .filmGauge.
	 *
	 * The default film gauge is 35, so that the focal length can be specified for
	 * a 35mm (full frame) camera.
	 *
	 * Values for focal length and film gauge must have the same unit.
	 */

	/** see {@link http://www.bobatkins.com/photography/technical/field_of_view.html} */
	vExtentSlope := 0.5 * c.GetFilmHeight() / focalLength

	c.Fov = math3.RAD2DEG * 2 * math.Atan(vExtentSlope)
	c.UpdateProjectionMatrix()
}

func (c *PerspectiveCamera) GetFocalLength() float64 {
	/**
	 * Calculates the focal length from the current .fov and .filmGauge.
	 */

	vExtentSlope := math.Tan(math3.DEG2RAD * 0.5 * c.Fov)

	return 0.5 * c.GetFilmHeight() / vExtentSlope
}

func (c *PerspectiveCamera) GetEffectiveFOV() float64 {
	return math3.RAD2DEG * 2 * math.Atan(math.Tan(math3.DEG2RAD*0.5*c.Fov)/c.Zoom)
}

func (c *PerspectiveCamera) GetFilmWidth() float64 {
	// film not completely covered in portrait format (aspect < 1)
	return c.FilmGauge * math.Min(c.Aspect, 1)
}

func (c *PerspectiveCamera) GetFilmHeight() float64 {
	// film not completely covered in landscape format (aspect > 1)
	return c.FilmGauge / math.Max(c.Aspect, 1)
}

func (c *PerspectiveCamera) GetViewBounds(distance float64, minTarget, maxTarget *math3.Vector2) {
	/**
	 * Computes the 2D bounds of the camera's viewable rectangle at a given distance along the viewing direction.
	 * Sets minTarget and maxTarget to the coordinates of the lower-left and upper-right corners of the view rectangle.
	 */

	_v3PerspectiveCamera.Set(-1, -1, 0.5).ApplyMatrix4(&c.ProjectionMatrixInverse)
	minTarget.Set(_v3PerspectiveCamera.X, _v3PerspectiveCamera.Y).MultiplyScalar(-distance / _v3PerspectiveCamera.Z)

	_v3PerspectiveCamera.Set(1, 1, 0.5).ApplyMatrix4(&c.ProjectionMatrixInverse)
	maxTarget.Set(_v3PerspectiveCamera.X, _v3PerspectiveCamera.Y).MultiplyScalar(-distance / _v3PerspectiveCamera.Z)
}

func (c *PerspectiveCamera) GetViewSize(distance float64, target *math3.Vector2) *math3.Vector2 {
	/**
	 * Computes the width and height of the camera's viewable rectangle at a given distance along the viewing direction.
	 * Copies the result into the target Vector2, where x is width and y is height.
	 */

	c.GetViewBounds(distance, _minTargetPerspectiveCamera, _maxTargetPerspectiveCamera)

	return target.SubVectors(_maxTargetPerspectiveCamera, _minTargetPerspectiveCamera)
}

func (c *PerspectiveCamera) SetViewOffset(fullWidth, fullHeight, x, y, width, height float64) {
	/**
	 * Sets an offset in a larger frustum. This is useful for multi-window or
	 * multi-monitor/multi-machine setups.
	 *
	 * For example, if you have 3x2 monitors and each monitor is 1920x1080 and
	 * the monitors are in grid like this
	 *
	 *   +---+---+---+
	 *   | A | B | C |
	 *   +---+---+---+
	 *   | D | E | F |
	 *   +---+---+---+
	 *
	 * then for each monitor you would call it like this
	 *
	 *   const w = 1920;
	 *   const h = 1080;
	 *   const fullWidth = w * 3;
	 *   const fullHeight = h * 2;
	 *
	 *   --A--
	 *   camera.setViewOffset( fullWidth, fullHeight, w * 0, h * 0, w, h );
	 *   --B--
	 *   camera.setViewOffset( fullWidth, fullHeight, w * 1, h * 0, w, h );
	 *   --C--
	 *   camera.setViewOffset( fullWidth, fullHeight, w * 2, h * 0, w, h );
	 *   --D--
	 *   camera.setViewOffset( fullWidth, fullHeight, w * 0, h * 1, w, h );
	 *   --E--
	 *   camera.setViewOffset( fullWidth, fullHeight, w * 1, h * 1, w, h );
	 *   --F--
	 *   camera.setViewOffset( fullWidth, fullHeight, w * 2, h * 1, w, h );
	 *
	 *   Note there is no reason monitors have to be the same size or in a grid.
	 */

	c.Aspect = fullWidth / fullHeight

	if c.View == nil {
		c.View = &View{
			Enabled:    true,
			FullWidth:  1,
			FullHeight: 1,
			OffsetX:    0,
			OffsetY:    0,
			Width:      1,
			Height:     1,
		}
	}

	c.View.Enabled = true
	c.View.FullWidth = fullWidth
	c.View.FullHeight = fullHeight
	c.View.OffsetX = x
	c.View.OffsetY = y
	c.View.Width = width
	c.View.Height = height

	c.UpdateProjectionMatrix()
}

func (c *PerspectiveCamera) ClearViewOffset() {
	if c.View != nil {
		c.View.Enabled = false
	}

	c.UpdateProjectionMatrix()
}

func (c *PerspectiveCamera) UpdateProjectionMatrix() {
	near := c.Near
	top := near * math.Tan(math3.DEG2RAD*0.5*c.Fov) / c.Zoom
	height := 2 * top
	width := c.Aspect * height
	left := -0.5 * width

	if c.View != nil && c.View.Enabled {
		left += c.View.OffsetX * width / c.View.FullWidth
		top -= c.View.OffsetY * height / c.View.FullHeight
		width *= c.View.Width / c.View.FullWidth
		height *= c.View.Height / c.View.FullHeight
	}

	skew := c.FilmOffset
	if skew != 0 {
		left += near * skew / c.GetFilmWidth()
	}

	c.ProjectionMatrix.MakePerspective(left, left+width, top, top-height, near, c.Far, c.WebGLCoordinateSystem)

	c.ProjectionMatrixInverse.Copy(&c.ProjectionMatrix).Invert()
}

//todo
// 	toJSON( meta ) {
// 		const data = super.toJSON( meta );
//
// 		data.object.fov = this.fov;
// 		data.object.zoom = this.zoom;
//
// 		data.object.near = this.near;
// 		data.object.far = this.far;
// 		data.object.focus = this.focus;
//
// 		data.object.aspect = this.aspect;
//
// 		if ( this.view !== null ) data.object.view = Object.assign( {}, this.view );
//
// 		data.object.filmGauge = this.filmGauge;
// 		data.object.filmOffset = this.filmOffset;
//
// 		return data;
// 	}

var _v3PerspectiveCamera = math3.NewVector3Defaults()
var _minTargetPerspectiveCamera = math3.NewVector2Defaults()
var _maxTargetPerspectiveCamera = math3.NewVector2Defaults()
