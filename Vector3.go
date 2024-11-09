package three

import "strconv"

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// NewVector3Zero Creates a new Vector3 with 0 for X, Y and Z.
func NewVector3Zero() *Vector3 {
	return &Vector3{}
}

// NewVector3 Creates a new Vector3.
func NewVector3(x, y, z float64) *Vector3 {
	return &Vector3{x, y, z}
}

// Set Sets the x, y and z components of this vector.
func (v *Vector3) Set(x, y, z float64) *Vector3 {
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

// SetScalar Set the x, y and z values of this vector both equal to scalar.
func (v *Vector3) SetScalar(scalar float64) *Vector3 {
	v.X = scalar
	v.Y = scalar
	v.Z = scalar
	return v
}

// SetX Replace this vector's x value with x.
func (v *Vector3) SetX(x float64) *Vector3 {
	v.X = x
	return v
}

// SetY Replace this vector's y value with y.
func (v *Vector3) SetY(y float64) *Vector3 {
	v.Y = y
	return v
}

// SetZ Replace this vector's z value with z.
func (v *Vector3) SetZ(z float64) *Vector3 {
	v.Z = z
	return v
}

// SetComponent Sets the indexed components of this vector. index: 0 = x, 1 = y or 2 = z
func (v *Vector3) SetComponent(index int, value float64) *Vector3 {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	default:
		panic("index is out of range: " + strconv.Itoa(index))
	}
	return v
}

// GetComponent Sets the indexed components of this vector. index: 0 = x, 1 = y or 2 = z
func (v *Vector3) GetComponent(index int) float64 {
	switch index {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	default:
		panic("index is out of range: " + strconv.Itoa(index))
	}
}

// Clone Returns a new vector3 with the same x, y and z values as this one.
func (v *Vector3) Clone() *Vector3 {
	return NewVector3(v.X, v.Y, v.Z)
}

// Copy Copies the values of the passed vector3's x, y and z properties to this vector3.
func (v *Vector3) Copy(a *Vector3) *Vector3 {
	v.X = a.X
	v.Y = a.Y
	v.Z = a.Z
	return v
}

// Add Adds vector a to this vector.
func (v *Vector3) Add(a *Vector3) *Vector3 {
	v.X += a.X
	v.Y += a.Y
	v.Z += a.Z
	return v
}

// AddScalar Adds the scalar value s to this vector's x, y and z values.
func (v *Vector3) AddScalar(s float64) *Vector3 {
	v.X += s
	v.Y += s
	v.Z += s
	return v
}

// AddVectors Sets this vector to a + b.
func (v *Vector3) AddVectors(a, b *Vector3) *Vector3 {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
	return v
}

// AddScaledVector Adds the multiple of a and s to this vector.
func (v *Vector3) AddScaledVector(a *Vector3, s float64) *Vector3 {
	v.X += a.X * s
	v.Y += a.Y * s
	v.Z += a.Z * s
	return v
}

// Sub Subtracts a from this vector.
func (v *Vector3) Sub(a *Vector3) *Vector3 {
	v.X -= a.X
	v.Y -= a.Y
	v.Z -= a.Z
	return v
}

// SubScalar Subtracts s from this vector's x, y and z components.
func (v *Vector3) SubScalar(s float64) *Vector3 {
	v.X -= s
	v.Y -= s
	v.Z -= s
	return v
}

// SubVectors Sets this vector to a - b.
func (v *Vector3) SubVectors(a, b *Vector3) *Vector3 {
	v.X -= a.X - b.X
	v.Y -= a.Y - b.Y
	v.Z -= a.Z - b.Z
	return v
}

// Multiply Multiplies this vector by a.
func (v *Vector3) Multiply(a *Vector3) *Vector3 {
	v.X *= a.X
	v.Y *= a.Y
	v.Z *= a.Z
	return v
}

// MultiplyScalar Multiplies this vector by scalar s.
func (v *Vector3) MultiplyScalar(s float64) *Vector3 {
	v.X *= s
	v.Y *= s
	v.Z *= s
	return v
}

// MultiplyVectors Sets this vector equal to a * b, component-wise.
func (v *Vector3) MultiplyVectors(a, b *Vector3) *Vector3 {
	v.X = a.X * b.X
	v.Y = a.Y * b.Y
	v.Z = a.Z * b.Z
	return v
}

//// todo
////	applyEuler( euler ) {
////
////		return this.applyQuaternion( _quaternion.setFromEuler( euler ) );
////
////	}
////
////	applyAxisAngle( axis, angle ) {
////
////		return this.applyQuaternion( _quaternion.setFromAxisAngle( axis, angle ) );
////
////	}
////
////	applyMatrix3( m ) {
////
////		const x = this.x, y = this.y, z = this.z;
////		const e = m.elements;
////
////		this.x = e[ 0 ] * x + e[ 3 ] * y + e[ 6 ] * z;
////		this.y = e[ 1 ] * x + e[ 4 ] * y + e[ 7 ] * z;
////		this.z = e[ 2 ] * x + e[ 5 ] * y + e[ 8 ] * z;
////
////		return this;
////
////	}
////
////	applyNormalMatrix( m ) {
////
////		return this.applyMatrix3( m ).normalize();
////
////	}
////
////	applyMatrix4( m ) {
////
////		const x = this.x, y = this.y, z = this.z;
////		const e = m.elements;
////
////		const w = 1 / ( e[ 3 ] * x + e[ 7 ] * y + e[ 11 ] * z + e[ 15 ] );
////
////		this.x = ( e[ 0 ] * x + e[ 4 ] * y + e[ 8 ] * z + e[ 12 ] ) * w;
////		this.y = ( e[ 1 ] * x + e[ 5 ] * y + e[ 9 ] * z + e[ 13 ] ) * w;
////		this.z = ( e[ 2 ] * x + e[ 6 ] * y + e[ 10 ] * z + e[ 14 ] ) * w;
////
////		return this;
////
////	}
////
////	applyQuaternion( q ) {
////
////		// quaternion q is assumed to have unit length
////
////		const vx = this.x, vy = this.y, vz = this.z;
////		const qx = q.x, qy = q.y, qz = q.z, qw = q.w;
////
////		// t = 2 * cross( q.xyz, v );
////		const tx = 2 * ( qy * vz - qz * vy );
////		const ty = 2 * ( qz * vx - qx * vz );
////		const tz = 2 * ( qx * vy - qy * vx );
////
////		// v + q.w * t + cross( q.xyz, t );
////		this.x = vx + qw * tx + qy * tz - qz * ty;
////		this.y = vy + qw * ty + qz * tx - qx * tz;
////		this.z = vz + qw * tz + qx * ty - qy * tx;
////
////		return this;
////
////	}
////
////	project( camera ) {
////
////		return this.applyMatrix4( camera.matrixWorldInverse ).applyMatrix4( camera.projectionMatrix );
////
////	}
////
////	unproject( camera ) {
////
////		return this.applyMatrix4( camera.projectionMatrixInverse ).applyMatrix4( camera.matrixWorld );
////
////	}
////
////	transformDirection( m ) {
////
////		// input: THREE.Matrix4 affine matrix
////		// vector interpreted as a direction
////
////		const x = this.x, y = this.y, z = this.z;
////		const e = m.elements;
////
////		this.x = e[ 0 ] * x + e[ 4 ] * y + e[ 8 ] * z;
////		this.y = e[ 1 ] * x + e[ 5 ] * y + e[ 9 ] * z;
////		this.z = e[ 2 ] * x + e[ 6 ] * y + e[ 10 ] * z;
////
////		return this.normalize();
////
////	}
////
//
//// --- divide( v ) ---
//
//func (v *Vector3) Divide(a *Vector3) *Vector3 {
//	v.X /= a.X
//	v.Y /= a.Y
//	v.Z /= a.Z
//	return v
//}
//
//// --- divideScalar( scalar ) ---
//
//func (v *Vector3) DivideScalar(s float64) *Vector3 {
//	return v.MultiplyScalar(1 / s)
//}
//
//// --- min( v ) ---
//
//func (v *Vector3) Min(a *Vector3) *Vector3 {
//	v.X = math.Min(v.X, a.X)
//	v.Y = math.Min(v.Y, a.Y)
//	v.Z = math.Min(v.Z, a.Z)
//	return v
//}
//
//// --- max( v ) ---
//
//func (v *Vector3) Max(a *Vector3) *Vector3 {
//	v.X = math.Max(v.X, a.X)
//	v.Y = math.Max(v.Y, a.Y)
//	v.Z = math.Max(v.Z, a.Z)
//	return v
//}
//
//// --- clamp( min, max ) ---
//
//func (v *Vector3) Clamp(min, max *Vector3) *Vector3 {
//	v.X = mathutils.Clamp(v.X, min.X, max.X)
//	v.Y = mathutils.Clamp(v.Y, min.Y, max.Y)
//	v.Z = mathutils.Clamp(v.Z, min.Z, max.Z)
//	return v
//}
//
//// --- clampScalar( minVal, maxVal ) ---
//
//func (v *Vector3) ClampScalar(minVal, maxVal float64) *Vector3 {
//	v.X = mathutils.Clamp(v.X, minVal, maxVal)
//	v.Y = mathutils.Clamp(v.Y, minVal, maxVal)
//	v.Z = mathutils.Clamp(v.Z, minVal, maxVal)
//	return v
//}
//
//// --- clampLength( min, max ) ---
//
//func (v *Vector3) ClampLength(minVal, maxVal float64) *Vector3 {
//	length := v.Length()
//	return v.DivideScalar(tool.If(length > 0, length, 1)).MultiplyScalar(mathutils.Clamp(length, minVal, maxVal))
//}
//
//// --- floor() ---
//
//func (v *Vector3) Floor() *Vector3 {
//	v.X = math.Floor(v.X)
//	v.Y = math.Floor(v.Y)
//	v.Z = math.Floor(v.Z)
//	return v
//}
//
//// --- ceil() ---
//
//func (v *Vector3) Ceil() *Vector3 {
//	v.X = math.Ceil(v.X)
//	v.Y = math.Ceil(v.Y)
//	v.Z = math.Ceil(v.Z)
//	return v
//}
//
//// --- round() ---
//
//func (v *Vector3) Round() *Vector3 {
//	v.X = math.Round(v.X)
//	v.Y = math.Round(v.Y)
//	v.Z = math.Round(v.Z)
//	return v
//}
//
//// --- roundToZero() ---
//
//func (v *Vector3) RoundToZero() *Vector3 {
//	v.X = math.Trunc(v.X)
//	v.Y = math.Trunc(v.Y)
//	v.Z = math.Trunc(v.Z)
//	return v
//}
//
//// --- negate() ---
//
//func (v *Vector3) Negate() *Vector3 {
//	v.X = -v.X
//	v.Y = -v.Y
//	v.Z = -v.Z
//	return v
//}
//
//// --- dot( v ) ---
//
//func (v *Vector3) Dot(a *Vector3) float64 {
//	return v.X*a.X + v.Y*a.Y + v.Z*a.Z
//}
//
//// --- lengthSq() ---
//
//func (v *Vector3) LengthSq() float64 {
//	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
//}
//
//// --- length() ---
//
//func (v *Vector3) Length() float64 {
//	return math.Sqrt(v.LengthSq())
//}
//
//func (v *Vector3) LengthOrOne() float64 {
//	if l := v.Length(); l > 0 {
//		return l
//	}
//	return 1
//}
//
//// --- manhattanLength() ---
//
//func (v *Vector3) ManhattanLength() float64 {
//	return math.Abs(v.X) + math.Abs(v.Y) + math.Abs(v.Z)
//}
//
//// --- normalize() ---
//
//func (v *Vector3) Normalize() *Vector3 {
//	return v.DivideScalar(v.LengthOrOne())
//}
//
//// --- setLength( length ) ---
//
//func (v *Vector3) SetLength(length float64) *Vector3 {
//	return v.Normalize().MultiplyScalar(length)
//}
//
//// --- lerp( v, alpha ) ---
//
//func (v *Vector3) Lerp(a *Vector3, alpha float64) *Vector3 {
//	v.X += (a.X - v.X) * alpha
//	v.Y += (a.Y - v.Y) * alpha
//	v.Z += (a.Z - v.Z) * alpha
//	return v
//}
//
//// --- lerpVectors( v1, v2, alpha ) ---
//
//func (v *Vector3) LerpVectors(a, b *Vector3, alpha float64) *Vector3 {
//	v.X = a.X + (b.X-a.X)*alpha
//	v.Y = a.Y + (b.Y-a.Y)*alpha
//	v.Z = a.Z + (b.Z-a.Z)*alpha
//	return v
//}
//
//// --- cross( v ) ---
//
//func (v *Vector3) Cross(a *Vector3) *Vector3 {
//	return v.CrossVectors(v, a)
//}
//
//// --- crossVectors( a, b ) ---
//
//func (v *Vector3) CrossVectors(a, b *Vector3) *Vector3 {
//	ax, ay, az := a.X, a.Y, a.Z
//	bx, by, bz := b.X, b.Y, b.Z
//	v.X = ay*bz - az*by
//	v.Y = az*bx - ax*bz
//	v.Z = ax*by - ay*bx
//	return v
//}
//
//// --- projectOnVector( v ) ---
//
//func (v *Vector3) ProjectOnVector(a *Vector3) *Vector3 {
//	if denominator := a.LengthSq(); denominator == 0 {
//		return v.SetScalar(0)
//	} else {
//		scalar := a.Dot(v) / denominator
//		return v.Copy(a).MultiplyScalar(scalar)
//	}
//}
//
//// --- projectOnPlane( planeNormal ) ---
//
//func (v *Vector3) ProjectOnPlane(planeNormal *Vector3) *Vector3 {
//	return v.Sub(_vector.Copy(v).ProjectOnVector(planeNormal))
//}
//
//// --- reflect( normal ) ---
//
//func (v *Vector3) Reflect(normal *Vector3) *Vector3 {
//	return v.Sub(_vector.Copy(normal).MultiplyScalar(2 * v.Dot(normal)))
//}
//
//// --- angleTo( v ) ---
//
//func (v *Vector3) AngleTo(a *Vector3) float64 {
//	if denominator := math.Sqrt(v.LengthSq() * a.LengthSq()); denominator == 0 {
//		return math.Pi / 2
//	} else {
//		return math.Acos(mathutils.Clamp(v.Dot(a)/denominator, -1, 1))
//	}
//}
//
//// --- distanceTo( v ) ---
//
//func (v *Vector3) DistanceTo(a *Vector3) float64 {
//	return math.Sqrt(v.DistanceToSquared(a))
//}
//
//// --- distanceToSquared( v ) ---
//
//func (v *Vector3) DistanceToSquared(a *Vector3) float64 {
//	dx, dy, dz := v.X-a.X, v.Y-a.Y, v.Z-a.Z
//	return dx*dx + dy*dy + dz*dz
//}
//
//// --- manhattanDistanceTo( v ) ---
//
//func (v *Vector3) ManhattanDistanceTo(a *Vector3) float64 {
//	return math.Abs(v.X-a.X) + math.Abs(v.Y-a.Y) + math.Abs(v.Z-a.Z)
//}
//
//// todo
////	setFromSpherical( s ) {
////
////		return this.setFromSphericalCoords( s.radius, s.phi, s.theta );
////
////	}
//
//// --- setFromSphericalCoords( radius, phi, theta ) ---
//
//func (v *Vector3) SetFromSphericalCoords(radius, phi, theta float64) *Vector3 {
//	sinPhiRadius := math.Sin(phi) * radius
//	v.X = sinPhiRadius * math.Sin(theta)
//	v.Y = math.Cos(phi) * radius
//	v.Z = sinPhiRadius * math.Cos(theta)
//	return v
//}
//
//// todo
////	setFromCylindrical( c ) {
////
////		return this.setFromCylindricalCoords( c.radius, c.theta, c.y );
////
////	}
////
//
//// --- setFromCylindricalCoords( radius, theta, y ) ---
//
//func (v *Vector3) SetFromCylindricalCoords(radius, theta, y float64) *Vector3 {
//	v.X = radius * math.Sin(theta)
//	v.Y = y
//	v.Z = radius * math.Cos(theta)
//	return v
//}
//
//// todo
////	setFromMatrixPosition( m ) {
////
////		const e = m.elements;
////
////		this.x = e[ 12 ];
////		this.y = e[ 13 ];
////		this.z = e[ 14 ];
////
////		return this;
////
////	}
////
////	setFromMatrixScale( m ) {
////
////		const sx = this.setFromMatrixColumn( m, 0 ).length();
////		const sy = this.setFromMatrixColumn( m, 1 ).length();
////		const sz = this.setFromMatrixColumn( m, 2 ).length();
////
////		this.x = sx;
////		this.y = sy;
////		this.z = sz;
////
////		return this;
////
////	}
////
////	setFromMatrixColumn( m, index ) {
////
////		return this.fromArray( m.elements, index * 4 );
////
////	}
////
////	setFromMatrix3Column( m, index ) {
////
////		return this.fromArray( m.elements, index * 3 );
////
////	}
////
////	setFromEuler( e ) {
////
////		this.x = e._x;
////		this.y = e._y;
////		this.z = e._z;
////
////		return this;
////
////	}
////
////	setFromColor( c ) {
////
////		this.x = c.r;
////		this.y = c.g;
////		this.z = c.b;
////
////		return this;
////
////	}
//
//// --- equals( v ) ---
//
//func (v *Vector3) Equals(a *Vector3) bool {
//	return v.X == a.X && v.Y == a.Y && v.Z == a.Z
//}
//
//// --- fromArray( array, offset = 0 ) ---
//
//func (v *Vector3) FromArray(array []float64) *Vector3 {
//	_ = array[2]
//	v.X = array[0]
//	v.Y = array[1]
//	v.Z = array[2]
//	return v
//}
//
//// --- toArray( array = [], offset = 0 ) ---
//
//func (v *Vector3) ToArray(array []float64) []float64 {
//	_ = array[2]
//	array[0] = v.X
//	array[1] = v.Y
//	array[2] = v.Z
//	return array
//}
//
//// todo
////	fromBufferAttribute( attribute, index ) {
////
////		this.x = attribute.getX( index );
////		this.y = attribute.getY( index );
////		this.z = attribute.getZ( index );
////
////		return this;
////
////	}
////
//
//// --- random() ---
//
//func (v *Vector3) Random() *Vector3 {
//	v.X = rand2.Float64()
//	v.Y = rand2.Float64()
//	v.Z = rand2.Float64()
//	return v
//}
//
//// --- randomDirection() ---
//
//func (v *Vector3) RandomDirection() *Vector3 {
//	// https://mathworld.wolfram.com/SpherePointPicking.html
//
//	theta := rand2.Float64() * math.Pi * 2
//	u := rand2.Float64()*2 - 1
//	c := math.Sqrt(1 - u*u)
//
//	v.X = c * math.Cos(theta)
//	v.Y = u
//	v.Z = c * math.Sin(theta)
//	return v
//}
//
//// --- temp variables ---
//
//var _vector = &Vector3{}
