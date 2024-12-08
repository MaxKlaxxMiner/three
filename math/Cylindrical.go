package math

import "math"

/**
 * Ref: https://en.wikipedia.org/wiki/Cylindrical_coordinate_system
 */

type Cylindrical struct {
	Radius, Theta, Y float64
}

func NewCylindricalDefaults() *Cylindrical {
	return &Cylindrical{1, 0, 0}
}

func NewCylindrical(radius, theta, y float64) *Cylindrical {
	return &Cylindrical{
		radius, // distance from the origin to a point in the x-z plane
		theta,  // counterclockwise angle in the x-z plane measured in radians from the positive z-axis
		y,      // height above the x-z plane
	}
}

func (c *Cylindrical) Set(radius, theta, y float64) *Cylindrical {
	c.Radius, c.Theta, c.Y = radius, theta, y
	return c
}

func (c *Cylindrical) Copy(other *Cylindrical) *Cylindrical {
	c.Radius, c.Theta, c.Y = other.Radius, other.Theta, other.Y
	return c
}

func (c *Cylindrical) SetFromVector3(v *Vector3) *Cylindrical {
	return c.SetFromCartesianCoords(v.X, v.Y, v.Z)
}

func (c *Cylindrical) SetFromCartesianCoords(x, y, z float64) *Cylindrical {
	c.Radius = math.Sqrt(x*x + z*z)
	c.Theta = math.Atan2(x, z)
	c.Y = y
	return c
}

func (c *Cylindrical) Clone() *Cylindrical {
	return NewCylindricalDefaults().Copy(c)
}
