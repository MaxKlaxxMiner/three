package math

import "math"

/**
 * Ref: https://en.wikipedia.org/wiki/Spherical_coordinate_system
 *
 * phi (the polar angle) is measured from the positive y-axis. The positive y-axis is up.
 * theta (the azimuthal angle) is measured from the positive z-axis.
 */

type Spherical struct {
	Radius, Phi, Theta float64
}

func NewSphericalDefaults() *Spherical {
	return &Spherical{1, 0, 0}
}

func NewSpherical(radius, phi, theta float64) *Spherical {
	return &Spherical{radius, phi, theta}
}

func (s *Spherical) Set(radius, phi, theta float64) *Spherical {
	s.Radius, s.Phi, s.Theta = radius, phi, theta
	return s
}

func (s *Spherical) Copy(other *Spherical) *Spherical {
	s.Radius, s.Phi, s.Theta = other.Radius, other.Phi, other.Theta
	return s
}

// MakeSafe restrict phi to be between EPS and PI-EPS
func (s *Spherical) MakeSafe() *Spherical {
	const EPS = 0.000001
	s.Phi = Clamp(s.Phi, EPS, math.Pi-EPS)
	return s
}

func (s *Spherical) SetFromVector3(v *Vector3) *Spherical {
	return s.SetFromCartesianCoords(v.X, v.Y, v.Z)
}

func (s *Spherical) SetFromCartesianCoords(x, y, z float64) *Spherical {
	s.Radius = math.Sqrt(x*x + y*y + z*z)

	if s.Radius == 0 {
		s.Theta = 0
		s.Phi = 0
	} else {
		s.Theta = math.Atan2(x, z)
		s.Phi = math.Acos(Clamp(y/s.Radius, -1, 1))
	}

	return s
}

func (s *Spherical) Clone() *Spherical {
	return NewSphericalDefaults().Copy(s)
}
