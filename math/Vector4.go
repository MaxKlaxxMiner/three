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

//todo
// 	applyMatrix4( m ) {
//
// 		const x = this.x, y = this.y, z = this.z, w = this.w;
// 		const e = m.elements;
//
// 		this.x = e[ 0 ] * x + e[ 4 ] * y + e[ 8 ] * z + e[ 12 ] * w;
// 		this.y = e[ 1 ] * x + e[ 5 ] * y + e[ 9 ] * z + e[ 13 ] * w;
// 		this.z = e[ 2 ] * x + e[ 6 ] * y + e[ 10 ] * z + e[ 14 ] * w;
// 		this.w = e[ 3 ] * x + e[ 7 ] * y + e[ 11 ] * z + e[ 15 ] * w;
//
// 		return this;
//
// 	}

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

//todo
// 	setAxisAngleFromQuaternion( q ) {
//
// 		// http://www.euclideanspace.com/maths/geometry/rotations/conversions/quaternionToAngle/index.htm
//
// 		// q is assumed to be normalized
//
// 		this.w = 2 * Math.acos( q.w );
//
// 		const s = Math.sqrt( 1 - q.w * q.w );
//
// 		if ( s < 0.0001 ) {
//
// 			this.x = 1;
// 			this.y = 0;
// 			this.z = 0;
//
// 		} else {
//
// 			this.x = q.x / s;
// 			this.y = q.y / s;
// 			this.z = q.z / s;
//
// 		}
//
// 		return this;
//
// 	}
//
// 	setAxisAngleFromRotationMatrix( m ) {
//
// 		// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToAngle/index.htm
//
// 		// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)
//
// 		let angle, x, y, z; // variables for result
// 		const epsilon = 0.01,		// margin to allow for rounding errors
// 			epsilon2 = 0.1,		// margin to distinguish between 0 and 180 degrees
//
// 			te = m.elements,
//
// 			m11 = te[ 0 ], m12 = te[ 4 ], m13 = te[ 8 ],
// 			m21 = te[ 1 ], m22 = te[ 5 ], m23 = te[ 9 ],
// 			m31 = te[ 2 ], m32 = te[ 6 ], m33 = te[ 10 ];
//
// 		if ( ( Math.abs( m12 - m21 ) < epsilon ) &&
// 		     ( Math.abs( m13 - m31 ) < epsilon ) &&
// 		     ( Math.abs( m23 - m32 ) < epsilon ) ) {
//
// 			// singularity found
// 			// first check for identity matrix which must have +1 for all terms
// 			// in leading diagonal and zero in other terms
//
// 			if ( ( Math.abs( m12 + m21 ) < epsilon2 ) &&
// 			     ( Math.abs( m13 + m31 ) < epsilon2 ) &&
// 			     ( Math.abs( m23 + m32 ) < epsilon2 ) &&
// 			     ( Math.abs( m11 + m22 + m33 - 3 ) < epsilon2 ) ) {
//
// 				// this singularity is identity matrix so angle = 0
//
// 				this.set( 1, 0, 0, 0 );
//
// 				return this; // zero angle, arbitrary axis
//
// 			}
//
// 			// otherwise this singularity is angle = 180
//
// 			angle = Math.PI;
//
// 			const xx = ( m11 + 1 ) / 2;
// 			const yy = ( m22 + 1 ) / 2;
// 			const zz = ( m33 + 1 ) / 2;
// 			const xy = ( m12 + m21 ) / 4;
// 			const xz = ( m13 + m31 ) / 4;
// 			const yz = ( m23 + m32 ) / 4;
//
// 			if ( ( xx > yy ) && ( xx > zz ) ) {
//
// 				// m11 is the largest diagonal term
//
// 				if ( xx < epsilon ) {
//
// 					x = 0;
// 					y = 0.707106781;
// 					z = 0.707106781;
//
// 				} else {
//
// 					x = Math.sqrt( xx );
// 					y = xy / x;
// 					z = xz / x;
//
// 				}
//
// 			} else if ( yy > zz ) {
//
// 				// m22 is the largest diagonal term
//
// 				if ( yy < epsilon ) {
//
// 					x = 0.707106781;
// 					y = 0;
// 					z = 0.707106781;
//
// 				} else {
//
// 					y = Math.sqrt( yy );
// 					x = xy / y;
// 					z = yz / y;
//
// 				}
//
// 			} else {
//
// 				// m33 is the largest diagonal term so base result on this
//
// 				if ( zz < epsilon ) {
//
// 					x = 0.707106781;
// 					y = 0.707106781;
// 					z = 0;
//
// 				} else {
//
// 					z = Math.sqrt( zz );
// 					x = xz / z;
// 					y = yz / z;
//
// 				}
//
// 			}
//
// 			this.set( x, y, z, angle );
//
// 			return this; // return 180 deg rotation
//
// 		}
//
// 		// as we have reached here there are no singularities so we can handle normally
//
// 		let s = Math.sqrt( ( m32 - m23 ) * ( m32 - m23 ) +
// 			( m13 - m31 ) * ( m13 - m31 ) +
// 			( m21 - m12 ) * ( m21 - m12 ) ); // used to normalize
//
// 		if ( Math.abs( s ) < 0.001 ) s = 1;
//
// 		// prevent divide by zero, should not happen if matrix is orthogonal and should be
// 		// caught by singularity test above, but I've left it in just in case
//
// 		this.x = ( m32 - m23 ) / s;
// 		this.y = ( m13 - m31 ) / s;
// 		this.z = ( m21 - m12 ) / s;
// 		this.w = Math.acos( ( m11 + m22 + m33 - 1 ) / 2 );
//
// 		return this;
//
// 	}
//
// 	setFromMatrixPosition( m ) {
//
// 		const e = m.elements;
//
// 		this.x = e[ 12 ];
// 		this.y = e[ 13 ];
// 		this.z = e[ 14 ];
// 		this.w = e[ 15 ];
//
// 		return this;
//
// 	}

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
//
// 		this.x = attribute.getX( index );
// 		this.y = attribute.getY( index );
// 		this.z = attribute.getZ( index );
// 		this.w = attribute.getW( index );
//
// 		return this;
//
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
