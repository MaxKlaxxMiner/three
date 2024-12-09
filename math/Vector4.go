package math

import (
	"github.com/MaxKlaxxMiner/three/utils"
	"math"
	"strconv"
)

type Vector4 struct {
	X, Y, Z, W float64
}

func NewVector4Defaults() *Vector4 {
	return &Vector4{W: 1}
}

func NewVector4(x, y, z, w float64) *Vector4 {
	return &Vector4{x, y, z, w}
}

func (v *Vector4) IsVector4() bool { return v != nil }

func (v *Vector4) GetWidth() float64       { return v.Z }
func (v *Vector4) SetWidth(value float64)  { v.Z = value }
func (v *Vector4) GetHeight() float64      { return v.W }
func (v *Vector4) SetHeight(value float64) { v.W = value }

func (v *Vector4) Set(x, y, z, w float64) *Vector4 {
	v.X, v.Y, v.Z, v.W = x, y, z, w
	return v
}

func (v *Vector4) SetScalar(scalar float64) *Vector4 {
	v.X, v.Y, v.Z, v.W = scalar, scalar, scalar, scalar
	return v
}

func (v *Vector4) SetX(x float64) *Vector4 {
	v.X = x
	return v
}

func (v *Vector4) SetY(y float64) *Vector4 {
	v.Y = y
	return v
}

func (v *Vector4) SetZ(z float64) *Vector4 {
	v.Z = z
	return v
}

func (v *Vector4) SetW(w float64) *Vector4 {
	v.W = w
	return v
}

func (v *Vector4) SetComponent(index int, value float64) *Vector4 {
	switch index {
	case 0, 'x', 'X':
		v.X = value
	case 1, 'y', 'Y':
		v.Y = value
	case 2, 'z', 'Z':
		v.Z = value
	case 3, 'w', 'W':
		v.W = value
	default:
		panic("index is out of range: " + strconv.Itoa(index))
	}
	return v
}

func (v *Vector4) GetComponent(index int) float64 {
	switch index {
	case 0, 'x', 'X':
		return v.X
	case 1, 'y', 'Y':
		return v.Y
	case 2, 'z', 'Z':
		return v.Z
	case 3, 'w', 'W':
		return v.Z
	default:
		panic("index is out of range: " + strconv.Itoa(index))
	}
}

func (v *Vector4) Clone() *Vector4 {
	return NewVector4(v.X, v.Y, v.Z, v.W)
}

func (v *Vector4) Copy(src *Vector4) *Vector4 {
	v.X, v.Y, v.Z, v.W = src.X, src.Y, src.Z, src.W
	return v
}

func (v *Vector4) CopyVector3(src *Vector3) *Vector4 {
	v.X, v.Y, v.Z, v.W = src.X, src.Y, src.Z, 1
	return v
}

func (v *Vector4) Add(a *Vector4) *Vector4 {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
	v.W += a.W
	return v
}

func (v *Vector4) AddScalar(s float64) *Vector4 {
	v.X += s
	v.Y += s
	v.Z += s
	v.W += s
	return v
}

func (v *Vector4) AddVectors(a, b *Vector4) *Vector4 {
	v.X, v.Y, v.Z, v.W = a.X+b.X, a.Y+b.Y, a.Z+b.Z, a.W+b.W
	return v
}

func (v *Vector4) AddScaledVector(a *Vector4, s float64) *Vector4 {
	v.X += a.X * s
	v.Y += a.Y * s
	v.Z += a.Z * s
	v.W += a.W * s
	return v
}

func (v *Vector4) Sub(a *Vector4) *Vector4 {
	v.X -= a.X
	v.Y -= a.Y
	v.Z -= a.Z
	v.W -= a.W
	return v
}

func (v *Vector4) SubScalar(s float64) *Vector4 {
	v.X -= s
	v.Y -= s
	v.Z -= s
	v.W -= s
	return v
}

func (v *Vector4) SubVectors(a, b *Vector4) *Vector4 {
	v.X -= a.X - b.X
	v.Y -= a.Y - b.Y
	v.Z -= a.Z - b.Z
	v.W -= a.W - b.W
	return v
}

func (v *Vector4) Multiply(a *Vector4) *Vector4 {
	v.X *= a.X
	v.Y *= a.Y
	v.Z *= a.Z
	v.W *= a.W
	return v
}

func (v *Vector4) MultiplyScalar(s float64) *Vector4 {
	v.X *= s
	v.Y *= s
	v.Z *= s
	v.W *= s
	return v
}

func (v *Vector4) MultiplyVectors(a, b *Vector4) *Vector4 {
	v.X, v.Y, v.Z, v.W = a.X*b.X, a.Y*b.Y, a.Z*b.Z, a.W*b.W
	return v
}

func (v *Vector4) ApplyMatrix4(m *Matrix4) *Vector4 {
	x, y, z, w := v.X, v.Y, v.Z, v.W
	v.X = m.N[0]*x + m.N[4]*y + m.N[8]*z + m.N[12]*w
	v.Y = m.N[1]*x + m.N[5]*y + m.N[9]*z + m.N[13]*w
	v.Z = m.N[2]*x + m.N[6]*y + m.N[10]*z + m.N[14]*w
	v.W = m.N[3]*x + m.N[7]*y + m.N[11]*z + m.N[15]*w
	return v
}

func (v *Vector4) Divide(a *Vector4) *Vector4 {
	v.X /= a.X
	v.Y /= a.Y
	v.Z /= a.Z
	v.W /= a.W
	return v
}

func (v *Vector4) DivideScalar(s float64) *Vector4 {
	return v.MultiplyScalar(1 / s)
}

func (v *Vector4) SetAxisAngleFromQuaternion(q *Quaternion) *Vector4 {
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/quaternionToAngle/index.htm

	// q is assumed to be normalized
	v.W = 2 * math.Acos(q.w)

	s := math.Sqrt(1 - q.w*q.w)
	if s < 0.0001 {
		v.X, v.Y, v.Z = 1, 0, 0
	} else {
		v.X, v.Y, v.Z = q.x/s, q.y/s, q.z/s
	}
	return v
}

func (v *Vector4) SetAxisAngleFromRotationMatrix(m *Matrix4) *Vector4 {
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToAngle/index.htm

	// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)

	const epsilon = 0.01 // margin to allow for rounding errors
	const epsilon2 = 0.1 // margin to distinguish between 0 and 180 degrees

	m11, m12, m13 := m.N[0], m.N[4], m.N[8]
	m21, m22, m23 := m.N[1], m.N[5], m.N[9]
	m31, m32, m33 := m.N[2], m.N[6], m.N[10]

	if math.Abs(m12-m21) < epsilon && math.Abs(m13-m31) < epsilon && math.Abs(m23-m32) < epsilon {
		var angle, x, y, z float64 // variables for result

		// singularity found
		// first check for identity matrix which must have +1 for all terms
		// in leading diagonal and zero in other terms

		if math.Abs(m12+m21) < epsilon2 && math.Abs(m13+m31) < epsilon2 && math.Abs(m23+m32) < epsilon2 && math.Abs(m11+m22+m33-3) < epsilon2 {
			// this singularity is identity matrix so angle = 0
			return v.Set(1, 0, 0, 0) // zero angle, arbitrary axis
		}

		// otherwise this singularity is angle = 180
		angle = math.Pi

		xx := (m11 + 1) / 2
		yy := (m22 + 1) / 2
		zz := (m33 + 1) / 2
		xy := (m12 + m21) / 4
		xz := (m13 + m31) / 4
		yz := (m23 + m32) / 4

		if xx > yy && xx > zz {
			// m11 is the largest diagonal term
			if xx < epsilon {
				x = 0
				y = 0.707106781
				z = 0.707106781
			} else {
				x = math.Sqrt(xx)
				y = xy / x
				z = xz / x
			}
		} else if yy > zz {
			// m22 is the largest diagonal term
			if yy < epsilon {
				x = 0.707106781
				y = 0
				z = 0.707106781
			} else {
				y = math.Sqrt(yy)
				x = xy / y
				z = yz / y
			}
		} else {
			// m33 is the largest diagonal term so base result on this
			if zz < epsilon {
				x = 0.707106781
				y = 0.707106781
				z = 0
			} else {
				z = math.Sqrt(zz)
				x = xz / z
				y = yz / z
			}
		}
		return v.Set(x, y, z, angle) // return 180 deg rotation
	}

	// as we have reached here there are no singularities so we can handle normally

	s := math.Sqrt((m32-m23)*(m32-m23) + (m13-m31)*(m13-m31) + (m21-m12)*(m21-m12)) // used to normalize

	if math.Abs(s) < 0.001 {
		s = 1
	}

	// prevent divide by zero, should not happen if matrix is orthogonal and should be
	// caught by singularity test above, but I've left it in just in case

	v.X = (m32 - m23) / s
	v.Y = (m13 - m31) / s
	v.Z = (m21 - m12) / s
	v.W = math.Acos((m11 + m22 + m33 - 1) / 2)

	return v
}

func (v *Vector4) SetFromMatrixPosition(m *Matrix4) *Vector4 {
	v.X, v.Y, v.Z, v.W = m.N[12], m.N[13], m.N[14], m.N[15]
	return v
}

func (v *Vector4) Min(a *Vector4) *Vector4 {
	v.X = math.Min(v.X, a.X)
	v.Y = math.Min(v.Y, a.Y)
	v.Z = math.Min(v.Z, a.Z)
	v.W = math.Min(v.W, a.W)
	return v
}

func (v *Vector4) Max(a *Vector4) *Vector4 {
	v.X = math.Max(v.X, a.X)
	v.Y = math.Max(v.Y, a.Y)
	v.Z = math.Max(v.Z, a.Z)
	v.W = math.Max(v.W, a.W)
	return v
}

func (v *Vector4) Clamp(min, max *Vector4) *Vector4 {
	v.X = Clamp(v.X, min.X, max.X)
	v.Y = Clamp(v.Y, min.Y, max.Y)
	v.Z = Clamp(v.Z, min.Z, max.Z)
	v.W = Clamp(v.W, min.W, max.W)
	return v
}

func (v *Vector4) ClampScalar(minVal, maxVal float64) *Vector4 {
	v.X = Clamp(v.X, minVal, maxVal)
	v.Y = Clamp(v.Y, minVal, maxVal)
	v.Z = Clamp(v.Z, minVal, maxVal)
	v.W = Clamp(v.W, minVal, maxVal)
	return v
}

func (v *Vector4) ClampLength(minVal, maxVal float64) *Vector4 {
	length := v.Length()
	return v.DivideScalar(utils.If(length > 0, length, 1)).MultiplyScalar(Clamp(length, minVal, maxVal))
}

func (v *Vector4) Floor() *Vector4 {
	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)
	v.Z = math.Floor(v.Z)
	v.W = math.Floor(v.W)
	return v
}

func (v *Vector4) Ceil() *Vector4 {
	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)
	v.Z = math.Ceil(v.Z)
	v.W = math.Ceil(v.W)
	return v
}

func (v *Vector4) Round() *Vector4 {
	v.X = math.Round(v.X)
	v.Y = math.Round(v.Y)
	v.Z = math.Round(v.Z)
	v.W = math.Round(v.W)
	return v
}

func (v *Vector4) RoundToZero() *Vector4 {
	v.X = math.Trunc(v.X)
	v.Y = math.Trunc(v.Y)
	v.Z = math.Trunc(v.Z)
	v.W = math.Trunc(v.W)
	return v
}

func (v *Vector4) Negate() *Vector4 {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	v.W = -v.W
	return v
}

func (v *Vector4) Dot(a *Vector4) float64 {
	return v.X*a.X + v.Y*a.Y + v.Z*a.Z + v.W*a.W
}

func (v *Vector4) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W
}

func (v *Vector4) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

func (v *Vector4) LengthOrOne() float64 {
	if l := v.Length(); l > 0 {
		return l
	}
	return 1
}

func (v *Vector4) ManhattanLength() float64 {
	return math.Abs(v.X) + math.Abs(v.Y) + math.Abs(v.Z) + math.Abs(v.W)
}

func (v *Vector4) Normalize() *Vector4 {
	return v.DivideScalar(v.LengthOrOne())
}

func (v *Vector4) SetLength(length float64) *Vector4 {
	return v.Normalize().MultiplyScalar(length)
}

func (v *Vector4) Lerp(a *Vector4, alpha float64) *Vector4 {
	v.X += (a.X - v.X) * alpha
	v.Y += (a.Y - v.Y) * alpha
	v.Z += (a.Z - v.Z) * alpha
	v.W += (a.W - v.W) * alpha
	return v
}

func (v *Vector4) LerpVectors(a, b *Vector4, alpha float64) *Vector4 {
	v.X = a.X + (b.X-a.X)*alpha
	v.Y = a.Y + (b.Y-a.Y)*alpha
	v.Z = a.Z + (b.Z-a.Z)*alpha
	v.W = a.W + (b.W-a.W)*alpha
	return v
}

func (v *Vector4) Equals(a *Vector4) bool {
	return *v == *a
}

func (v *Vector4) FromArray(array []float64) *Vector4 {
	_ = array[3]
	v.X = array[0]
	v.Y = array[1]
	v.Z = array[2]
	v.W = array[3]
	return v
}

func (v *Vector4) ToArray(array []float64) []float64 {
	_ = array[3]
	array[0] = v.X
	array[1] = v.Y
	array[2] = v.Z
	array[3] = v.W
	return array
}

//todo
// 	fromBufferAttribute( attribute, index ) {
// 		this.x = attribute.getX( index );
// 		this.y = attribute.getY( index );
// 		this.z = attribute.getZ( index );
// 		this.w = attribute.getW( index );
// 		return this;
// 	}

func (v *Vector4) Random() *Vector4 {
	v.X = RandomFloat()
	v.Y = RandomFloat()
	v.Z = RandomFloat()
	v.W = RandomFloat()
	return v
}

func (v *Vector4) Append(buf []float64) []float64 {
	return append(buf, v.X, v.Y, v.Z, v.W)
}
