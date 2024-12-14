package math

import (
	"fmt"
	"github.com/MaxKlaxxMiner/three/utils"
	"math"
)

type Quaternion struct {
	x, y, z, w        float64
	_onChangeCallback func()
}

func NewQuaternionDefaults() *Quaternion {
	return &Quaternion{0, 0, 0, 1, func() {}}
}

func NewQuaternion(x, y, z, w float64) *Quaternion {
	return &Quaternion{x, y, z, w, func() {}}
}

func (q *Quaternion) IsQuaternion() bool { return q != nil }

func SlerpFlat(dst, src0, src1 []float64, t float64) {
	// fuzz-free, array-based Quaternion SLERP operation
	_, _, _ = dst[3], src0[3], src1[3]

	x0, y0, z0, w0 := src0[0], src0[1], src0[2], src0[3]
	x1, y1, z1, w1 := src1[0], src1[1], src1[2], src1[3]

	if t == 0 {
		dst[0], dst[1], dst[2], dst[3] = x0, y0, z0, w0
		return
	}

	if t == 1 {
		dst[0], dst[1], dst[2], dst[3] = x1, y1, z1, w1
		return
	}

	if w0 != w1 || x0 != x1 || y0 != y1 || z0 != z1 {
		s := 1 - t
		cos := x0*x1 + y0*y1 + z0*z1 + w0*w1
		dir := utils.If(cos >= 0, 1.0, -1.0)
		sqrSin := 1 - cos*cos

		// Skip the Slerp for tiny steps to avoid numeric problems:
		if sqrSin > EPSILON {
			sin := math.Sqrt(sqrSin)
			length := math.Atan2(sin, cos*dir)
			s = math.Sin(s*length) / sin
			t = math.Sin(t*length) / sin
		}

		tDir := t * dir
		x0 = x0*s + x1*tDir
		y0 = y0*s + y1*tDir
		z0 = z0*s + z1*tDir
		w0 = w0*s + w1*tDir

		// Normalize in case we just did a lerp:
		if s == 1-t {
			f := 1 / math.Sqrt(x0*x0+y0*y0+z0*z0+w0*w0)
			x0 *= f
			y0 *= f
			z0 *= f
			w0 *= f
		}
	}

	dst[0], dst[1], dst[2], dst[3] = x0, y0, z0, w0
}

func MultiplyQuaternionsFlat(dst, src0, src1 []float64) []float64 {
	_, _, _ = dst[3], src0[3], src1[3]

	x0, y0, z0, w0 := src0[0], src0[1], src0[2], src0[3]
	x1, y1, z1, w1 := src1[0], src1[1], src1[2], src1[3]

	dst[0] = x0*w1 + w0*x1 + y0*z1 - z0*y1
	dst[1] = y0*w1 + w0*y1 + z0*x1 - x0*z1
	dst[2] = z0*w1 + w0*z1 + x0*y1 - y0*x1
	dst[3] = w0*w1 - x0*x1 - y0*y1 - z0*z1

	return dst
}

func (q *Quaternion) GetX() float64 { return q.x }
func (q *Quaternion) GetY() float64 { return q.y }
func (q *Quaternion) GetZ() float64 { return q.z }
func (q *Quaternion) GetW() float64 { return q.w }

func (q *Quaternion) SetX(x float64) {
	q.x = x
	q._onChangeCallback()
}
func (q *Quaternion) SetY(y float64) {
	q.y = y
	q._onChangeCallback()
}
func (q *Quaternion) SetZ(z float64) {
	q.z = z
	q._onChangeCallback()
}
func (q *Quaternion) SetW(w float64) {
	q.w = w
	q._onChangeCallback()
}

func (q *Quaternion) Set(x, y, z, w float64) *Quaternion {
	q.x, q.y, q.z, q.w = x, y, z, w
	q._onChangeCallback()
	return q
}

func (q *Quaternion) Clone() *Quaternion {
	return NewQuaternion(q.x, q.y, q.z, q.w)
}

func (q *Quaternion) Copy(quaternion *Quaternion) *Quaternion {
	q.x, q.y, q.z, q.w = quaternion.x, quaternion.y, quaternion.z, quaternion.w
	q._onChangeCallback()
	return q
}

func (q *Quaternion) SetFromEuler(euler *Euler) *Quaternion {
	return q.SetFromEulerUpdate(euler, true)
}

func (q *Quaternion) SetFromEulerUpdate(euler *Euler, update bool) *Quaternion {
	x, y, z, order := euler.x, euler.y, euler.z, euler.order

	// http://www.mathworks.com/matlabcentral/fileexchange/
	// 	20696-function-to-convert-between-dcm-euler-angles-quaternions-and-euler-vectors/
	//	content/SpinCalc.m

	c1 := math.Cos(x / 2)
	c2 := math.Cos(y / 2)
	c3 := math.Cos(z / 2)

	s1 := math.Sin(x / 2)
	s2 := math.Sin(y / 2)
	s3 := math.Sin(z / 2)

	switch order {
	case EulerOrderXYZ:
		q.x = s1*c2*c3 + c1*s2*s3
		q.y = c1*s2*c3 - s1*c2*s3
		q.z = c1*c2*s3 + s1*s2*c3
		q.w = c1*c2*c3 - s1*s2*s3
	case EulerOrderYXZ:
		q.x = s1*c2*c3 + c1*s2*s3
		q.y = c1*s2*c3 - s1*c2*s3
		q.z = c1*c2*s3 - s1*s2*c3
		q.w = c1*c2*c3 + s1*s2*s3
	case EulerOrderZXY:
		q.x = s1*c2*c3 - c1*s2*s3
		q.y = c1*s2*c3 + s1*c2*s3
		q.z = c1*c2*s3 + s1*s2*c3
		q.w = c1*c2*c3 - s1*s2*s3
	case EulerOrderZYX:
		q.x = s1*c2*c3 - c1*s2*s3
		q.y = c1*s2*c3 + s1*c2*s3
		q.z = c1*c2*s3 - s1*s2*c3
		q.w = c1*c2*c3 + s1*s2*s3
	case EulerOrderYZX:
		q.x = s1*c2*c3 + c1*s2*s3
		q.y = c1*s2*c3 + s1*c2*s3
		q.z = c1*c2*s3 - s1*s2*c3
		q.w = c1*c2*c3 - s1*s2*s3
	case EulerOrderXZY:
		q.x = s1*c2*c3 - c1*s2*s3
		q.y = c1*s2*c3 - s1*c2*s3
		q.z = c1*c2*s3 + s1*s2*c3
		q.w = c1*c2*c3 + s1*s2*s3
	default:
		fmt.Println("THREE.Quaternion: .setFromEuler() encountered an unknown order: ", order)
	}

	if update {
		q._onChangeCallback()
	}
	return q
}

func (q *Quaternion) SetFromAxisAngle(axis *Vector3, angle float64) *Quaternion {
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/angleToQuaternion/index.htm
	// assumes axis is normalized

	halfAngle := angle / 2
	s := math.Sin(halfAngle)

	q.x = axis.X * s
	q.y = axis.Y * s
	q.z = axis.Z * s
	q.w = math.Cos(halfAngle)

	q._onChangeCallback()
	return q
}

func (q *Quaternion) SetFromRotationMatrix(m *Matrix4) *Quaternion {
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/index.htm
	// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)

	m11, m12, m13 := m.N[0], m.N[4], m.N[8]
	m21, m22, m23 := m.N[1], m.N[5], m.N[9]
	m31, m32, m33 := m.N[2], m.N[6], m.N[10]

	trace := m11 + m22 + m33

	if trace > 0 {
		s := 0.5 / math.Sqrt(trace+1.0)
		q.w = 0.25 / s
		q.x = (m32 - m23) * s
		q.y = (m13 - m31) * s
		q.z = (m21 - m12) * s
	} else if m11 > m22 && m11 > m33 {
		s := 2.0 * math.Sqrt(1.0+m11-m22-m33)
		q.w = (m32 - m23) / s
		q.x = 0.25 * s
		q.y = (m12 + m21) / s
		q.z = (m13 + m31) / s
	} else if m22 > m33 {
		s := 2.0 * math.Sqrt(1.0+m22-m11-m33)
		q.w = (m13 - m31) / s
		q.x = (m12 + m21) / s
		q.y = 0.25 * s
		q.z = (m23 + m32) / s
	} else {
		s := 2.0 * math.Sqrt(1.0+m33-m11-m22)
		q.w = (m21 - m12) / s
		q.x = (m13 + m31) / s
		q.y = (m23 + m32) / s
		q.z = 0.25 * s
	}

	q._onChangeCallback()
	return q
}

func (q *Quaternion) setFromUnitVectors(vFrom, vTo *Vector3) *Quaternion {
	// assumes direction vectors vFrom and vTo are normalized

	r := vFrom.Dot(vTo) + 1

	if r < EPSILON {
		// vFrom and vTo point in opposite directions
		r = 0

		if math.Abs(vFrom.X) > math.Abs(vFrom.Z) {
			q.x = -vFrom.Y
			q.y = vFrom.X
			q.z = 0
			q.w = r
		} else {
			q.x = 0
			q.y = -vFrom.Z
			q.z = vFrom.Y
			q.w = r
		}
	} else {
		// crossVectors( vFrom, vTo ); // inlined to avoid cyclic dependency on Vector3

		q.x = vFrom.Y*vTo.Z - vFrom.Z*vTo.Y
		q.y = vFrom.Z*vTo.X - vFrom.X*vTo.Z
		q.z = vFrom.X*vTo.Y - vFrom.Y*vTo.X
		q.w = r
	}

	return q.Normalize()
}

func (q *Quaternion) AngleTo(v *Quaternion) float64 {
	return 2 * math.Acos(math.Abs(Clamp(q.Dot(v), -1, 1)))
}

func (q *Quaternion) RotateTowards(v *Quaternion, step float64) *Quaternion {
	angle := q.AngleTo(v)
	if angle == 0 {
		return q
	}

	t := math.Min(1, step/angle)
	q.Slerp(v, t)

	return q
}

func (q *Quaternion) Identity() *Quaternion {
	return q.Set(0, 0, 0, 1)
}

func (q *Quaternion) Invert() *Quaternion {
	// quaternion is assumed to have unit length
	return q.Conjugate()
}

func (q *Quaternion) Conjugate() *Quaternion {
	q.x *= -1
	q.y *= -1
	q.z *= -1
	q._onChangeCallback()
	return q
}

func (q *Quaternion) Dot(v *Quaternion) float64 {
	return q.x*v.x + q.y*v.y + q.z*v.z + q.w*v.w
}

func (q *Quaternion) LengthSq() float64 {
	return q.x*q.x + q.y*q.y + q.z*q.z + q.w*q.w
}

func (q *Quaternion) Length() float64 {
	return math.Sqrt(q.LengthSq())
}

func (q *Quaternion) Normalize() *Quaternion {
	l := q.Length()

	if l == 0 {
		q.x = 0
		q.y = 0
		q.z = 0
		q.w = 1
	} else {
		l = 1 / l
		q.x *= l
		q.y *= l
		q.z *= l
		q.w *= l

	}

	q._onChangeCallback()
	return q
}

func (q *Quaternion) Multiply(v *Quaternion) *Quaternion {
	return q.MultiplyQuaternions(q, v)
}

func (q *Quaternion) Premultiply(v *Quaternion) *Quaternion {
	return q.MultiplyQuaternions(v, q)
}

func (q *Quaternion) MultiplyQuaternions(a, b *Quaternion) *Quaternion {
	// from http://www.euclideanspace.com/maths/algebra/realNormedAlgebra/quaternions/code/index.htm

	qax, qay, qaz, qaw := a.x, a.y, a.z, a.w
	qbx, qby, qbz, qbw := b.x, b.y, b.z, b.w

	q.x = qax*qbw + qaw*qbx + qay*qbz - qaz*qby
	q.y = qay*qbw + qaw*qby + qaz*qbx - qax*qbz
	q.z = qaz*qbw + qaw*qbz + qax*qby - qay*qbx
	q.w = qaw*qbw - qax*qbx - qay*qby - qaz*qbz

	q._onChangeCallback()
	return q
}

func (q *Quaternion) Slerp(qb *Quaternion, t float64) *Quaternion {
	if t == 0 {
		return q
	}
	if t == 1 {
		return q.Copy(qb)
	}

	x, y, z, w := q.x, q.y, q.z, q.w

	// http://www.euclideanspace.com/maths/algebra/realNormedAlgebra/quaternions/slerp/

	cosHalfTheta := w*qb.w + x*qb.x + y*qb.y + z*qb.z

	if cosHalfTheta < 0 {
		q.x, q.y, q.z, q.w = -qb.x, -qb.y, -qb.z, -qb.w
		cosHalfTheta = -cosHalfTheta
	} else {
		q.x, q.y, q.z, q.w = qb.x, qb.y, qb.z, qb.w // todo issue fix: q.Copy(qb) without trigger q._onChangeCallback()
	}

	if cosHalfTheta >= 1.0 {
		q.w, q.x, q.y, q.z = w, x, y, z
		return q //todo issue: missing q._onChangeCallback()?
	}

	sqrSinHalfTheta := 1.0 - cosHalfTheta*cosHalfTheta

	if sqrSinHalfTheta <= EPSILON {
		s := 1 - t
		q.w, q.x, q.y, q.z = s*w+t*q.w, s*x+t*q.x, s*y+t*q.y, s*z+t*q.z
		return q.Normalize() // normalize calls _onChangeCallback()
	}

	sinHalfTheta := math.Sqrt(sqrSinHalfTheta)
	halfTheta := math.Atan2(sinHalfTheta, cosHalfTheta)
	ratioA := math.Sin((1-t)*halfTheta) / sinHalfTheta
	ratioB := math.Sin(t*halfTheta) / sinHalfTheta
	q.x, q.y, q.z, q.w = x*ratioA+q.x*ratioB, y*ratioA+q.y*ratioB, z*ratioA+q.z*ratioB, w*ratioA+q.w*ratioB

	q._onChangeCallback()
	return q
}

func (q *Quaternion) SlerpQuaternions(qa, qb *Quaternion, t float64) *Quaternion {
	return q.Copy(qa).Slerp(qb, t)
}

func (q *Quaternion) Random() *Quaternion {
	// sets this quaternion to a uniform random unit quaternnion

	// Ken Shoemake
	// Uniform random rotations
	// D. Kirk, editor, Graphics Gems III, pages 124-132. Academic Press, New York, 1992.

	theta1 := 2 * math.Pi * RandomFloat()
	theta2 := 2 * math.Pi * RandomFloat()

	x0 := RandomFloat()
	r1 := math.Sqrt(1 - x0)
	r2 := math.Sqrt(x0)

	return q.Set(
		r1*math.Sin(theta1),
		r1*math.Cos(theta1),
		r2*math.Sin(theta2),
		r2*math.Cos(theta2))
}

func (q *Quaternion) Equals(quaternion *Quaternion) bool {
	return quaternion.x == q.x && quaternion.y == q.y && quaternion.z == q.z && quaternion.w == q.w
}

func (q *Quaternion) FromArray(array []float64) *Quaternion {
	_ = array[3]
	q.x, q.y, q.z, q.w = array[0], array[1], array[2], array[3]
	q._onChangeCallback()
	return q
}

func (q *Quaternion) ToArray(array []float64) []float64 {
	_ = array[3]
	array[0], array[1], array[2], array[3] = q.x, q.y, q.z, q.w
	return array
}

//todo
// 	fromBufferAttribute( attribute, index ) {
// 		this._x = attribute.getX( index );
// 		this._y = attribute.getY( index );
// 		this._z = attribute.getZ( index );
// 		this._w = attribute.getW( index );
// 		this._onChangeCallback();
// 		return this;
// 	}

func (q *Quaternion) OverrideOnChange(callback func()) *Quaternion {
	q._onChangeCallback = callback
	return q
}

func (q *Quaternion) Append(buf []float64) []float64 {
	return append(buf, q.x, q.y, q.z, q.w)
}
