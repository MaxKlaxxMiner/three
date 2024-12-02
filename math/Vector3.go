package math

import (
	"github.com/MaxKlaxxMiner/three/utils"
	"math"
	"strconv"
)

type Vector3 struct {
	X, Y, Z float64
}

func NewVector3Defaults() *Vector3 {
	return &Vector3{}
}

func NewVector3(x, y, z float64) *Vector3 {
	return &Vector3{x, y, z}
}

func (v *Vector3) IsVector3() bool { return v != nil }

func (v *Vector3) Set(x, y, z float64) *Vector3 {
	v.X, v.Y, v.Z = x, y, z
	return v
}

func (v *Vector3) SetScalar(scalar float64) *Vector3 {
	v.X, v.Y, v.Z = scalar, scalar, scalar
	return v
}

func (v *Vector3) SetX(x float64) *Vector3 {
	v.X = x
	return v
}

func (v *Vector3) SetY(y float64) *Vector3 {
	v.Y = y
	return v
}

func (v *Vector3) SetZ(z float64) *Vector3 {
	v.Z = z
	return v
}

func (v *Vector3) SetComponent(index int, value float64) *Vector3 {
	switch index {
	case 0, 'x', 'X':
		v.X = value
	case 1, 'y', 'Y':
		v.Y = value
	case 2, 'z', 'Z':
		v.Z = value
	default:
		panic("index is out of range: " + strconv.Itoa(index))
	}
	return v
}

func (v *Vector3) GetComponent(index int) float64 {
	switch index {
	case 0, 'x', 'X':
		return v.X
	case 1, 'y', 'Y':
		return v.Y
	case 2, 'z', 'Z':
		return v.Z
	default:
		panic("index is out of range: " + strconv.Itoa(index))
	}
}

func (v *Vector3) Clone() *Vector3 {
	return NewVector3(v.X, v.Y, v.Z)
}

func (v *Vector3) Copy(src *Vector3) *Vector3 {
	v.X, v.Y, v.Z = src.X, src.Y, src.Z
	return v
}

func (v *Vector3) Add(a *Vector3) *Vector3 {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
	return v
}

func (v *Vector3) AddScalar(s float64) *Vector3 {
	v.X += s
	v.Y += s
	v.Z += s
	return v
}

func (v *Vector3) AddVectors(a, b *Vector3) *Vector3 {
	v.X, v.Y, v.Z = a.X+b.X, a.Y+b.Y, a.Z+b.Z
	return v
}

func (v *Vector3) AddScaledVector(a *Vector3, s float64) *Vector3 {
	v.X += a.X * s
	v.Y += a.Y * s
	v.Z += a.Z * s
	return v
}

func (v *Vector3) Sub(a *Vector3) *Vector3 {
	v.X -= a.X
	v.Y -= a.Y
	v.Z -= a.Z
	return v
}

func (v *Vector3) SubScalar(s float64) *Vector3 {
	v.X -= s
	v.Y -= s
	v.Z -= s
	return v
}

func (v *Vector3) SubVectors(a, b *Vector3) *Vector3 {
	v.X -= a.X - b.X
	v.Y -= a.Y - b.Y
	v.Z -= a.Z - b.Z
	return v
}

func (v *Vector3) Multiply(a *Vector3) *Vector3 {
	v.X *= a.X
	v.Y *= a.Y
	v.Z *= a.Z
	return v
}

func (v *Vector3) MultiplyScalar(s float64) *Vector3 {
	v.X *= s
	v.Y *= s
	v.Z *= s
	return v
}

func (v *Vector3) MultiplyVectors(a, b *Vector3) *Vector3 {
	v.X, v.Y, v.Z = a.X*b.X, a.Y*b.Y, a.Z*b.Z
	return v
}

//todo
// 	applyEuler( euler ) {
//
// 		return this.applyQuaternion( _quaternion.setFromEuler( euler ) );
//
// 	}
//
// 	applyAxisAngle( axis, angle ) {
//
// 		return this.applyQuaternion( _quaternion.setFromAxisAngle( axis, angle ) );
//
// 	}
//
// 	applyMatrix3( m ) {
//
// 		const x = this.x, y = this.y, z = this.z;
// 		const e = m.elements;
//
// 		this.x = e[ 0 ] * x + e[ 3 ] * y + e[ 6 ] * z;
// 		this.y = e[ 1 ] * x + e[ 4 ] * y + e[ 7 ] * z;
// 		this.z = e[ 2 ] * x + e[ 5 ] * y + e[ 8 ] * z;
//
// 		return this;
//
// 	}
//
// 	applyNormalMatrix( m ) {
//
// 		return this.applyMatrix3( m ).normalize();
//
// 	}
//
// 	applyMatrix4( m ) {
//
// 		const x = this.x, y = this.y, z = this.z;
// 		const e = m.elements;
//
// 		const w = 1 / ( e[ 3 ] * x + e[ 7 ] * y + e[ 11 ] * z + e[ 15 ] );
//
// 		this.x = ( e[ 0 ] * x + e[ 4 ] * y + e[ 8 ] * z + e[ 12 ] ) * w;
// 		this.y = ( e[ 1 ] * x + e[ 5 ] * y + e[ 9 ] * z + e[ 13 ] ) * w;
// 		this.z = ( e[ 2 ] * x + e[ 6 ] * y + e[ 10 ] * z + e[ 14 ] ) * w;
//
// 		return this;
//
// 	}
//
// 	applyQuaternion( q ) {
//
// 		// quaternion q is assumed to have unit length
//
// 		const vx = this.x, vy = this.y, vz = this.z;
// 		const qx = q.x, qy = q.y, qz = q.z, qw = q.w;
//
// 		// t = 2 * cross( q.xyz, v );
// 		const tx = 2 * ( qy * vz - qz * vy );
// 		const ty = 2 * ( qz * vx - qx * vz );
// 		const tz = 2 * ( qx * vy - qy * vx );
//
// 		// v + q.w * t + cross( q.xyz, t );
// 		this.x = vx + qw * tx + qy * tz - qz * ty;
// 		this.y = vy + qw * ty + qz * tx - qx * tz;
// 		this.z = vz + qw * tz + qx * ty - qy * tx;
//
// 		return this;
//
// 	}
//
// 	project( camera ) {
//
// 		return this.applyMatrix4( camera.matrixWorldInverse ).applyMatrix4( camera.projectionMatrix );
//
// 	}
//
// 	unproject( camera ) {
//
// 		return this.applyMatrix4( camera.projectionMatrixInverse ).applyMatrix4( camera.matrixWorld );
//
// 	}
//
// 	transformDirection( m ) {
//
// 		// input: THREE.Matrix4 affine matrix
// 		// vector interpreted as a direction
//
// 		const x = this.x, y = this.y, z = this.z;
// 		const e = m.elements;
//
// 		this.x = e[ 0 ] * x + e[ 4 ] * y + e[ 8 ] * z;
// 		this.y = e[ 1 ] * x + e[ 5 ] * y + e[ 9 ] * z;
// 		this.z = e[ 2 ] * x + e[ 6 ] * y + e[ 10 ] * z;
//
// 		return this.normalize();
//
// 	}
//

func (v *Vector3) Divide(a *Vector3) *Vector3 {
	v.X /= a.X
	v.Y /= a.Y
	v.Z /= a.Z
	return v
}

func (v *Vector3) DivideScalar(s float64) *Vector3 {
	return v.MultiplyScalar(1 / s)
}

func (v *Vector3) Min(a *Vector3) *Vector3 {
	v.X = math.Min(v.X, a.X)
	v.Y = math.Min(v.Y, a.Y)
	v.Z = math.Min(v.Z, a.Z)
	return v
}

func (v *Vector3) Max(a *Vector3) *Vector3 {
	v.X = math.Max(v.X, a.X)
	v.Y = math.Max(v.Y, a.Y)
	v.Z = math.Max(v.Z, a.Z)
	return v
}

func (v *Vector3) Clamp(min, max *Vector3) *Vector3 {
	v.X = Clamp(v.X, min.X, max.X)
	v.Y = Clamp(v.Y, min.Y, max.Y)
	v.Z = Clamp(v.Z, min.Z, max.Z)
	return v
}

func (v *Vector3) ClampScalar(minVal, maxVal float64) *Vector3 {
	v.X = Clamp(v.X, minVal, maxVal)
	v.Y = Clamp(v.Y, minVal, maxVal)
	v.Z = Clamp(v.Z, minVal, maxVal)
	return v
}

func (v *Vector3) ClampLength(minVal, maxVal float64) *Vector3 {
	length := v.Length()
	return v.DivideScalar(utils.If(length > 0, length, 1)).MultiplyScalar(Clamp(length, minVal, maxVal))
}

func (v *Vector3) Floor() *Vector3 {
	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)
	v.Z = math.Floor(v.Z)
	return v
}

func (v *Vector3) Ceil() *Vector3 {
	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)
	v.Z = math.Ceil(v.Z)
	return v
}

func (v *Vector3) Round() *Vector3 {
	v.X = math.Round(v.X)
	v.Y = math.Round(v.Y)
	v.Z = math.Round(v.Z)
	return v
}

func (v *Vector3) RoundToZero() *Vector3 {
	v.X = math.Trunc(v.X)
	v.Y = math.Trunc(v.Y)
	v.Z = math.Trunc(v.Z)
	return v
}

func (v *Vector3) Negate() *Vector3 {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
	return v
}

func (v *Vector3) Dot(a *Vector3) float64 {
	return v.X*a.X + v.Y*a.Y + v.Z*a.Z
}

func (v *Vector3) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vector3) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

func (v *Vector3) LengthOrOne() float64 {
	if l := v.Length(); l > 0 {
		return l
	}
	return 1
}

func (v *Vector3) ManhattanLength() float64 {
	return math.Abs(v.X) + math.Abs(v.Y) + math.Abs(v.Z)
}

func (v *Vector3) Normalize() *Vector3 {
	return v.DivideScalar(v.LengthOrOne())
}

func (v *Vector3) SetLength(length float64) *Vector3 {
	return v.Normalize().MultiplyScalar(length)
}

func (v *Vector3) Lerp(a *Vector3, alpha float64) *Vector3 {
	v.X += (a.X - v.X) * alpha
	v.Y += (a.Y - v.Y) * alpha
	v.Z += (a.Z - v.Z) * alpha
	return v
}

func (v *Vector3) LerpVectors(a, b *Vector3, alpha float64) *Vector3 {
	v.X = a.X + (b.X-a.X)*alpha
	v.Y = a.Y + (b.Y-a.Y)*alpha
	v.Z = a.Z + (b.Z-a.Z)*alpha
	return v
}

func (v *Vector3) Cross(a *Vector3) *Vector3 {
	return v.CrossVectors(v, a)
}

func (v *Vector3) CrossVectors(a, b *Vector3) *Vector3 {
	ax, ay, az := a.X, a.Y, a.Z
	bx, by, bz := b.X, b.Y, b.Z
	v.X = ay*bz - az*by
	v.Y = az*bx - ax*bz
	v.Z = ax*by - ay*bx
	return v
}

func (v *Vector3) ProjectOnVector(a *Vector3) *Vector3 {
	if denominator := a.LengthSq(); denominator == 0 {
		return v.SetScalar(0)
	} else {
		scalar := a.Dot(v) / denominator
		return v.Copy(a).MultiplyScalar(scalar)
	}
}

func (v *Vector3) ProjectOnPlane(planeNormal *Vector3) *Vector3 {
	return v.Sub(_vector.Copy(v).ProjectOnVector(planeNormal))
}

func (v *Vector3) Reflect(normal *Vector3) *Vector3 {
	// reflect incident vector off plane orthogonal to normal
	// normal is assumed to have unit length
	return v.Sub(_vector.Copy(normal).MultiplyScalar(2 * v.Dot(normal)))
}

func (v *Vector3) AngleTo(a *Vector3) float64 {
	if denominator := math.Sqrt(v.LengthSq() * a.LengthSq()); denominator == 0 {
		return math.Pi / 2
	} else {
		theta := v.Dot(a) / denominator
		// clamp, to handle numerical problems
		return math.Acos(Clamp(theta, -1, 1))
	}
}

func (v *Vector3) DistanceTo(a *Vector3) float64 {
	return math.Sqrt(v.DistanceToSquared(a))
}

func (v *Vector3) DistanceToSquared(a *Vector3) float64 {
	dx, dy, dz := v.X-a.X, v.Y-a.Y, v.Z-a.Z
	return dx*dx + dy*dy + dz*dz
}

func (v *Vector3) ManhattanDistanceTo(a *Vector3) float64 {
	return math.Abs(v.X-a.X) + math.Abs(v.Y-a.Y) + math.Abs(v.Z-a.Z)
}

//todo
// setFromSpherical( s ) {
//
//	return this.setFromSphericalCoords( s.radius, s.phi, s.theta );
//
// }
//

func (v *Vector3) SetFromSphericalCoords(radius, phi, theta float64) *Vector3 {
	sinPhiRadius := math.Sin(phi) * radius
	v.X = sinPhiRadius * math.Sin(theta)
	v.Y = math.Cos(phi) * radius
	v.Z = sinPhiRadius * math.Cos(theta)
	return v
}

//todo
// setFromCylindrical( c ) {
//
//	return this.setFromCylindricalCoords( c.radius, c.theta, c.y );
//
// }
//

func (v *Vector3) SetFromCylindricalCoords(radius, theta, y float64) *Vector3 {
	v.X = radius * math.Sin(theta)
	v.Y = y
	v.Z = radius * math.Cos(theta)
	return v
}

//todo
// setFromCylindricalCoords( radius, theta, y ) {
//
//	this.x = radius * Math.sin( theta );
//	this.y = y;
//	this.z = radius * Math.cos( theta );
//
//	return this;
//
// }
//
// setFromMatrixPosition( m ) {
//
//	const e = m.elements;
//
//	this.x = e[ 12 ];
//	this.y = e[ 13 ];
//	this.z = e[ 14 ];
//
//	return this;
//
// }
//
// setFromMatrixScale( m ) {
//
//	const sx = this.setFromMatrixColumn( m, 0 ).length();
//	const sy = this.setFromMatrixColumn( m, 1 ).length();
//	const sz = this.setFromMatrixColumn( m, 2 ).length();
//
//	this.x = sx;
//	this.y = sy;
//	this.z = sz;
//
//	return this;
//
// }
//
// setFromMatrixColumn( m, index ) {
//
//	return this.fromArray( m.elements, index * 4 );
//
// }
//
// setFromMatrix3Column( m, index ) {
//
//	return this.fromArray( m.elements, index * 3 );
//
// }
//
// setFromEuler( e ) {
//
//	this.x = e._x;
//	this.y = e._y;
//	this.z = e._z;
//
//	return this;
//
// }
//

func (v *Vector3) SetFromColor(color *Color) *Vector3 {
	v.X, v.Y, v.Z = color.R, color.G, color.B
	return v
}

func (v *Vector3) Equals(a *Vector3) bool {
	return *v == *a
}

func (v *Vector3) FromArray(array []float64) *Vector3 {
	_ = array[2]
	v.X = array[0]
	v.Y = array[1]
	v.Z = array[2]
	return v
}

func (v *Vector3) ToArray(array []float64) []float64 {
	_ = array[2]
	array[0] = v.X
	array[1] = v.Y
	array[2] = v.Z
	return array
}

//todo
// fromBufferAttribute( attribute, index ) {
//	this.x = attribute.getX( index );
//	this.y = attribute.getY( index );
//	this.z = attribute.getZ( index );
//	return this;
// }

func (v *Vector3) Random() *Vector3 {
	v.X = RandomFloat()
	v.Y = RandomFloat()
	v.Z = RandomFloat()
	return v
}

func (v *Vector3) RandomDirection() *Vector3 {
	// https://mathworld.wolfram.com/SpherePointPicking.html
	theta := RandomFloat() * math.Pi * 2
	u := RandomFloat()*2 - 1
	c := math.Sqrt(1 - u*u)

	v.X = c * math.Cos(theta)
	v.Y = u
	v.Z = c * math.Sin(theta)
	return v
}

func (v *Vector3) Append(buf []float64) []float64 {
	return append(buf, v.X, v.Y, v.Z)
}

var _vector = NewVector3Defaults()
