package math

import "math"

type Matrix4 struct {
	N [4 * 4]float64
}

func (m *Matrix4) IsMatrix4() bool { return m != nil }

func NewMatrix4Identity() *Matrix4 {
	return new(Matrix4).Identity()
}

func NewMatrix4(n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44 float64) *Matrix4 {
	return &Matrix4{
		N: [4 * 4]float64{
			n11, n21, n31, n41,
			n12, n22, n32, n42,
			n13, n23, n33, n43,
			n14, n24, n34, n44,
		},
	}
}

func (m *Matrix4) Set(n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44 float64) *Matrix4 {
	m.N = [4 * 4]float64{
		n11, n21, n31, n41,
		n12, n22, n32, n42,
		n13, n23, n33, n43,
		n14, n24, n34, n44,
	}
	return m
}

var zeroMatrix4 Matrix4

var identityMatrix4 = Matrix4{
	N: [4 * 4]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	},
}

func (m *Matrix4) Identity() *Matrix4 {
	m.N = identityMatrix4.N
	return m
}

func (m *Matrix4) Clone() *Matrix4 {
	return new(Matrix4).Copy(m)
}

func (m *Matrix4) Copy(v *Matrix4) *Matrix4 {
	m.N = v.N
	return m
}

func (m *Matrix4) CopyPosition(v *Matrix4) *Matrix4 {
	m.N[12] = v.N[12]
	m.N[13] = v.N[13]
	m.N[14] = v.N[14]
	return m
}

func (m *Matrix4) SetFromMatrix3(v *Matrix3) *Matrix4 {
	return m.Set(
		v.N[0], v.N[3], v.N[6], 0,
		v.N[1], v.N[4], v.N[7], 0,
		v.N[2], v.N[5], v.N[8], 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) ExtractBasis(xAxis, yAxis, zAxis *Vector3) *Matrix4 {
	xAxis.X, yAxis.X, zAxis.X = m.N[0], m.N[1], m.N[2]
	xAxis.Y, yAxis.Y, zAxis.Y = m.N[4], m.N[5], m.N[6]
	xAxis.Z, yAxis.Z, zAxis.Z = m.N[8], m.N[9], m.N[10]
	return m
}

func (m *Matrix4) MakeBasis(xAxis, yAxis, zAxis *Vector3) *Matrix4 {
	return m.Set(
		xAxis.X, yAxis.X, zAxis.X, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, yAxis.Z, zAxis.Z, 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) ExtractRotation(v *Matrix4) *Matrix4 {
	// this method does not support reflection matrices

	scaleX := 1 / _v1Matrix4.SetFromMatrixColumn(m, 0).Length()
	scaleY := 1 / _v1Matrix4.SetFromMatrixColumn(m, 1).Length()
	scaleZ := 1 / _v1Matrix4.SetFromMatrixColumn(m, 2).Length()

	m.N[0] = v.N[0] * scaleX
	m.N[1] = v.N[1] * scaleX
	m.N[2] = v.N[2] * scaleX
	m.N[3] = 0

	m.N[4] = v.N[4] * scaleY
	m.N[5] = v.N[5] * scaleY
	m.N[6] = v.N[6] * scaleY
	m.N[7] = 0

	m.N[8] = v.N[8] * scaleZ
	m.N[9] = v.N[9] * scaleZ
	m.N[10] = v.N[10] * scaleZ
	m.N[11] = 0

	m.N[12] = 0
	m.N[13] = 0
	m.N[14] = 0
	m.N[15] = 1

	return m
}

func (m *Matrix4) MakeRotationFromEuler(euler *Euler) *Matrix4 {
	x, y, z := euler.x, euler.y, euler.z
	a, b := math.Cos(x), math.Sin(x)
	c, d := math.Cos(y), math.Sin(y)
	e, f := math.Cos(z), math.Sin(z)

	switch euler.order {

	case EulerOrderXYZ:
		ae, af, be, bf := a*e, a*f, b*e, b*f
		m.N[0] = c * e
		m.N[4] = -c * f
		m.N[8] = d
		m.N[1] = af + be*d
		m.N[5] = ae - bf*d
		m.N[9] = -b * c
		m.N[2] = bf - ae*d
		m.N[6] = be + af*d
		m.N[10] = a * c
	case EulerOrderYXZ:
		ce, cf, de, df := c*e, c*f, d*e, d*f
		m.N[0] = ce + df*b
		m.N[4] = de*b - cf
		m.N[8] = a * d
		m.N[1] = a * f
		m.N[5] = a * e
		m.N[9] = -b
		m.N[2] = cf*b - de
		m.N[6] = df + ce*b
		m.N[10] = a * c
	case EulerOrderZXY:
		ce, cf, de, df := c*e, c*f, d*e, d*f
		m.N[0] = ce - df*b
		m.N[4] = -a * f
		m.N[8] = de + cf*b
		m.N[1] = cf + de*b
		m.N[5] = a * e
		m.N[9] = df - ce*b
		m.N[2] = -a * d
		m.N[6] = b
		m.N[10] = a * c
	case EulerOrderZYX:
		ae, af, be, bf := a*e, a*f, b*e, b*f
		m.N[0] = c * e
		m.N[4] = be*d - af
		m.N[8] = ae*d + bf
		m.N[1] = c * f
		m.N[5] = bf*d + ae
		m.N[9] = af*d - be
		m.N[2] = -d
		m.N[6] = b * c
		m.N[10] = a * c
	case EulerOrderYZX:
		ac, ad, bc, bd := a*c, a*d, b*c, b*d
		m.N[0] = c * e
		m.N[4] = bd - ac*f
		m.N[8] = bc*f + ad
		m.N[1] = f
		m.N[5] = a * e
		m.N[9] = -b * e
		m.N[2] = -d * e
		m.N[6] = ad*f + bc
		m.N[10] = ac - bd*f
	case EulerOrderXZY:
		ac, ad, bc, bd := a*c, a*d, b*c, b*d
		m.N[0] = c * e
		m.N[4] = -f
		m.N[8] = d * e
		m.N[1] = ac*f + bd
		m.N[5] = a * e
		m.N[9] = ad*f - bc
		m.N[2] = bc*f - ad
		m.N[6] = b * e
		m.N[10] = bd*f + ac
	}

	// bottom row
	m.N[3] = 0
	m.N[7] = 0
	m.N[11] = 0

	// last column
	m.N[12] = 0
	m.N[13] = 0
	m.N[14] = 0
	m.N[15] = 1

	return m
}

func (m *Matrix4) MakeRotationFromQuaternion(q *Quaternion) *Matrix4 {
	return m.Compose(_zeroMatrix4, q, _oneMatrix4)
}

func (m *Matrix4) LookAt(eye, target, up *Vector3) *Matrix4 {
	_zMatrix4.SubVectors(eye, target)
	if _zMatrix4.LengthSq() == 0 {
		// eye and target are in the same position
		_zMatrix4.Z = 1
	}
	_zMatrix4.Normalize()
	_xMatrix4.CrossVectors(up, _zMatrix4)

	if _xMatrix4.LengthSq() == 0 {
		// up and z are parallel
		if math.Abs(up.Z) == 1 {
			_zMatrix4.X += 0.0001
		} else {
			_zMatrix4.Z += 0.0001
		}
		_zMatrix4.Normalize()
		_xMatrix4.CrossVectors(up, _zMatrix4)
	}

	_xMatrix4.Normalize()
	_yMatrix4.CrossVectors(_zMatrix4, _xMatrix4)

	m.N[0], m.N[4], m.N[8] = _xMatrix4.X, _yMatrix4.X, _zMatrix4.X
	m.N[1], m.N[5], m.N[9] = _xMatrix4.Y, _yMatrix4.Y, _zMatrix4.Y
	m.N[2], m.N[6], m.N[10] = _xMatrix4.Z, _yMatrix4.Z, _zMatrix4.Z

	return m
}

func (m *Matrix4) Multiply(v *Matrix4) *Matrix4 {
	return m.MultiplyMatrices(m, v)
}

func (m *Matrix4) Premultiply(v *Matrix4) *Matrix4 {
	return m.MultiplyMatrices(v, m)
}

func (m *Matrix4) MultiplyMatrices(a, b *Matrix4) *Matrix4 {
	a11, a12, a13, a14 := a.N[0], a.N[4], a.N[8], a.N[12]
	a21, a22, a23, a24 := a.N[1], a.N[5], a.N[9], a.N[13]
	a31, a32, a33, a34 := a.N[2], a.N[6], a.N[10], a.N[14]
	a41, a42, a43, a44 := a.N[3], a.N[7], a.N[11], a.N[15]

	b11, b12, b13, b14 := b.N[0], b.N[4], b.N[8], b.N[12]
	b21, b22, b23, b24 := b.N[1], b.N[5], b.N[9], b.N[13]
	b31, b32, b33, b34 := b.N[2], b.N[6], b.N[10], b.N[14]
	b41, b42, b43, b44 := b.N[3], b.N[7], b.N[11], b.N[15]

	m.N[0] = a11*b11 + a12*b21 + a13*b31 + a14*b41
	m.N[4] = a11*b12 + a12*b22 + a13*b32 + a14*b42
	m.N[8] = a11*b13 + a12*b23 + a13*b33 + a14*b43
	m.N[12] = a11*b14 + a12*b24 + a13*b34 + a14*b44

	m.N[1] = a21*b11 + a22*b21 + a23*b31 + a24*b41
	m.N[5] = a21*b12 + a22*b22 + a23*b32 + a24*b42
	m.N[9] = a21*b13 + a22*b23 + a23*b33 + a24*b43
	m.N[13] = a21*b14 + a22*b24 + a23*b34 + a24*b44

	m.N[2] = a31*b11 + a32*b21 + a33*b31 + a34*b41
	m.N[6] = a31*b12 + a32*b22 + a33*b32 + a34*b42
	m.N[10] = a31*b13 + a32*b23 + a33*b33 + a34*b43
	m.N[14] = a31*b14 + a32*b24 + a33*b34 + a34*b44

	m.N[3] = a41*b11 + a42*b21 + a43*b31 + a44*b41
	m.N[7] = a41*b12 + a42*b22 + a43*b32 + a44*b42
	m.N[11] = a41*b13 + a42*b23 + a43*b33 + a44*b43
	m.N[15] = a41*b14 + a42*b24 + a43*b34 + a44*b44

	return m
}

func (m *Matrix4) MultiplyScalar(s float64) *Matrix4 {
	m.N[0] *= s
	m.N[1] *= s
	m.N[2] *= s
	m.N[3] *= s
	m.N[4] *= s
	m.N[5] *= s
	m.N[6] *= s
	m.N[7] *= s
	m.N[8] *= s
	m.N[9] *= s
	m.N[10] *= s
	m.N[11] *= s
	m.N[12] *= s
	m.N[13] *= s
	m.N[14] *= s
	m.N[15] *= s
	return m
}

func (m *Matrix4) Determinant() float64 {
	n11, n12, n13, n14 := m.N[0], m.N[4], m.N[8], m.N[12]
	n21, n22, n23, n24 := m.N[1], m.N[5], m.N[9], m.N[13]
	n31, n32, n33, n34 := m.N[2], m.N[6], m.N[10], m.N[14]
	n41, n42, n43, n44 := m.N[3], m.N[7], m.N[11], m.N[15]

	//TOO: make this more efficient
	//( based on http://www.euclideanspace.com/maths/algebra/matrix/functions/inverse/fourD/index.htm )

	return n41*(+n14*n23*n32-n13*n24*n32-n14*n22*n33+n12*n24*n33+n13*n22*n34-n12*n23*n34) +
		n42*(+n11*n23*n34-n11*n24*n33+n14*n21*n33-n13*n21*n34+n13*n24*n31-n14*n23*n31) +
		n43*(+n11*n24*n32-n11*n22*n34-n14*n21*n32+n12*n21*n34+n14*n22*n31-n12*n24*n31) +
		n44*(-n13*n22*n31-n11*n23*n32+n11*n22*n33+n13*n21*n32-n12*n21*n33+n12*n23*n31)
}

func (m *Matrix4) Transpose() *Matrix4 {
	m.N[1], m.N[4] = m.N[4], m.N[1]
	m.N[2], m.N[8] = m.N[8], m.N[2]
	m.N[6], m.N[9] = m.N[9], m.N[6]

	m.N[3], m.N[12] = m.N[12], m.N[3]
	m.N[7], m.N[13] = m.N[13], m.N[7]
	m.N[11], m.N[14] = m.N[14], m.N[11]

	return m
}

func (m *Matrix4) SetPosition(x, y, z float64) *Matrix4 {
	m.N[12], m.N[13], m.N[14] = x, y, z
	return m
}

func (m *Matrix4) SetPositionVector3(v *Vector3) *Matrix4 {
	m.N[12], m.N[13], m.N[14] = v.X, v.Y, v.Z
	return m

}

func (m *Matrix4) Invert() *Matrix4 {
	// based on http://www.euclideanspace.com/maths/algebra/matrix/functions/inverse/fourD/index.htm
	n11, n21, n31, n41 := m.N[0], m.N[1], m.N[2], m.N[3]
	n12, n22, n32, n42 := m.N[4], m.N[5], m.N[6], m.N[7]
	n13, n23, n33, n43 := m.N[8], m.N[9], m.N[10], m.N[11]
	n14, n24, n34, n44 := m.N[12], m.N[13], m.N[14], m.N[15]

	t11 := n23*n34*n42 - n24*n33*n42 + n24*n32*n43 - n22*n34*n43 - n23*n32*n44 + n22*n33*n44
	t12 := n14*n33*n42 - n13*n34*n42 - n14*n32*n43 + n12*n34*n43 + n13*n32*n44 - n12*n33*n44
	t13 := n13*n24*n42 - n14*n23*n42 + n14*n22*n43 - n12*n24*n43 - n13*n22*n44 + n12*n23*n44
	t14 := n14*n23*n32 - n13*n24*n32 - n14*n22*n33 + n12*n24*n33 + n13*n22*n34 - n12*n23*n34

	det := n11*t11 + n21*t12 + n31*t13 + n41*t14

	if det == 0 {
		return m.Copy(&zeroMatrix4)
	}

	detInv := 1 / det

	m.N[0] = t11 * detInv
	m.N[1] = (n24*n33*n41 - n23*n34*n41 - n24*n31*n43 + n21*n34*n43 + n23*n31*n44 - n21*n33*n44) * detInv
	m.N[2] = (n22*n34*n41 - n24*n32*n41 + n24*n31*n42 - n21*n34*n42 - n22*n31*n44 + n21*n32*n44) * detInv
	m.N[3] = (n23*n32*n41 - n22*n33*n41 - n23*n31*n42 + n21*n33*n42 + n22*n31*n43 - n21*n32*n43) * detInv

	m.N[4] = t12 * detInv
	m.N[5] = (n13*n34*n41 - n14*n33*n41 + n14*n31*n43 - n11*n34*n43 - n13*n31*n44 + n11*n33*n44) * detInv
	m.N[6] = (n14*n32*n41 - n12*n34*n41 - n14*n31*n42 + n11*n34*n42 + n12*n31*n44 - n11*n32*n44) * detInv
	m.N[7] = (n12*n33*n41 - n13*n32*n41 + n13*n31*n42 - n11*n33*n42 - n12*n31*n43 + n11*n32*n43) * detInv

	m.N[8] = t13 * detInv
	m.N[9] = (n14*n23*n41 - n13*n24*n41 - n14*n21*n43 + n11*n24*n43 + n13*n21*n44 - n11*n23*n44) * detInv
	m.N[10] = (n12*n24*n41 - n14*n22*n41 + n14*n21*n42 - n11*n24*n42 - n12*n21*n44 + n11*n22*n44) * detInv
	m.N[11] = (n13*n22*n41 - n12*n23*n41 - n13*n21*n42 + n11*n23*n42 + n12*n21*n43 - n11*n22*n43) * detInv

	m.N[12] = t14 * detInv
	m.N[13] = (n13*n24*n31 - n14*n23*n31 + n14*n21*n33 - n11*n24*n33 - n13*n21*n34 + n11*n23*n34) * detInv
	m.N[14] = (n14*n22*n31 - n12*n24*n31 - n14*n21*n32 + n11*n24*n32 + n12*n21*n34 - n11*n22*n34) * detInv
	m.N[15] = (n12*n23*n31 - n13*n22*n31 + n13*n21*n32 - n11*n23*n32 - n12*n21*n33 + n11*n22*n33) * detInv

	return m
}

func (m *Matrix4) Scale(v *Vector3) *Matrix4 {
	m.N[0] *= v.X
	m.N[1] *= v.X
	m.N[2] *= v.X
	m.N[3] *= v.X
	m.N[4] *= v.Y
	m.N[5] *= v.Y
	m.N[6] *= v.Y
	m.N[7] *= v.Y
	m.N[8] *= v.Z
	m.N[9] *= v.Z
	m.N[10] *= v.Z
	m.N[11] *= v.Z
	return m
}

func (m *Matrix4) GetMaxScaleOnAxis() float64 {
	scaleXSq := m.N[0]*m.N[0] + m.N[1]*m.N[1] + m.N[2]*m.N[2]
	scaleYSq := m.N[4]*m.N[4] + m.N[5]*m.N[5] + m.N[6]*m.N[6]
	scaleZSq := m.N[8]*m.N[8] + m.N[9]*m.N[9] + m.N[10]*m.N[10]
	return math.Sqrt(math.Max(math.Max(scaleXSq, scaleYSq), scaleZSq))
}

func (m *Matrix4) MakeTranslation(x, y, z float64) *Matrix4 {
	return m.Set(
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) MakeTranslationVector3(v *Vector3) *Matrix4 {
	return m.Set(
		1, 0, 0, v.X,
		0, 1, 0, v.Y,
		0, 0, 1, v.Z,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) MakeRotationX(theta float64) *Matrix4 {
	c, s := math.Cos(theta), math.Sin(theta)

	return m.Set(
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) MakeRotationY(theta float64) *Matrix4 {
	c, s := math.Cos(theta), math.Sin(theta)

	return m.Set(
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) MakeRotationZ(theta float64) *Matrix4 {
	c, s := math.Cos(theta), math.Sin(theta)

	return m.Set(
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) MakeRotationAxis(axis *Vector3, angle float64) *Matrix4 {
	// Based on http://www.gamedev.net/reference/articles/article1199.asp

	c, s := math.Cos(angle), math.Sin(angle)
	t := 1 - c
	x, y, z := axis.X, axis.Y, axis.Z
	tx, ty := t*x, t*y

	return m.Set(
		tx*x+c, tx*y-s*z, tx*z+s*y, 0,
		tx*y+s*z, ty*y+c, ty*z-s*x, 0,
		tx*z-s*y, ty*z+s*x, t*z*z+c, 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) MakeScale(x, y, z float64) *Matrix4 {
	return m.Set(
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) MakeShear(xy, xz, yx, yz, zx, zy float64) *Matrix4 {
	return m.Set(
		1, yx, zx, 0,
		xy, 1, zy, 0,
		xz, yz, 1, 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix4) Compose(position *Vector3, quaternion *Quaternion, scale *Vector3) *Matrix4 {
	x, y, z, w := quaternion.x, quaternion.y, quaternion.z, quaternion.w
	x2, y2, z2 := x+x, y+y, z+z
	xx, xy, xz := x*x2, x*y2, x*z2
	yy, yz, zz := y*y2, y*z2, z*z2
	wx, wy, wz := w*x2, w*y2, w*z2

	sx, sy, sz := scale.X, scale.Y, scale.Z

	m.N[0] = (1 - (yy + zz)) * sx
	m.N[1] = (xy + wz) * sx
	m.N[2] = (xz - wy) * sx
	m.N[3] = 0

	m.N[4] = (xy - wz) * sy
	m.N[5] = (1 - (xx + zz)) * sy
	m.N[6] = (yz + wx) * sy
	m.N[7] = 0

	m.N[8] = (xz + wy) * sz
	m.N[9] = (yz - wx) * sz
	m.N[10] = (1 - (xx + yy)) * sz
	m.N[11] = 0

	m.N[12] = position.X
	m.N[13] = position.Y
	m.N[14] = position.Z
	m.N[15] = 1

	return m
}

func (m *Matrix4) Decompose(position *Vector3, quaternion *Quaternion, scale *Vector3) *Matrix4 {
	sx := _v1Matrix4.Set(m.N[0], m.N[1], m.N[2]).Length()
	sy := _v1Matrix4.Set(m.N[4], m.N[5], m.N[6]).Length()
	sz := _v1Matrix4.Set(m.N[8], m.N[9], m.N[10]).Length()

	// if determine is negative, we need to invert one scale
	det := m.Determinant()
	if det < 0 {
		sx = -sx
	}

	position.X, position.Y, position.Z = m.N[12], m.N[13], m.N[14]

	// scale the rotation part
	_m1Matrix4.Copy(m)

	invSX, invSY, invSZ := 1/sx, 1/sy, 1/sz

	_m1Matrix4.N[0] *= invSX
	_m1Matrix4.N[1] *= invSX
	_m1Matrix4.N[2] *= invSX

	_m1Matrix4.N[4] *= invSY
	_m1Matrix4.N[5] *= invSY
	_m1Matrix4.N[6] *= invSY

	_m1Matrix4.N[8] *= invSZ
	_m1Matrix4.N[9] *= invSZ
	_m1Matrix4.N[10] *= invSZ

	quaternion.SetFromRotationMatrix(_m1Matrix4)

	scale.X, scale.Y, scale.Z = sx, sy, sz

	return m
}

func (m *Matrix4) MakePerspective(left, right, top, bottom, near, far float64, webGLCoordinateSystem bool) *Matrix4 {
	x := 2 * near / (right - left)
	y := 2 * near / (top - bottom)

	a := (right + left) / (right - left)
	b := (top + bottom) / (top - bottom)

	var c, d float64
	if webGLCoordinateSystem {
		c = -(far + near) / (far - near)
		d = (-2 * far * near) / (far - near)
	} else {
		c = -far / (far - near)
		d = (-far * near) / (far - near)
	}

	m.N[0], m.N[4], m.N[8], m.N[12] = x, 0, a, 0
	m.N[1], m.N[5], m.N[9], m.N[13] = 0, y, b, 0
	m.N[2], m.N[6], m.N[10], m.N[14] = 0, 0, c, d
	m.N[3], m.N[7], m.N[11], m.N[15] = 0, 0, -1, 0
	return m
}

func (m *Matrix4) MakeOrthographic(left, right, top, bottom, near, far float64, webGLCoordinateSystem bool) *Matrix4 {
	w := 1.0 / (right - left)
	h := 1.0 / (top - bottom)
	p := 1.0 / (far - near)

	x := (right + left) * w
	y := (top + bottom) * h

	var z, zInv float64
	if webGLCoordinateSystem {
		z = (far + near) * p
		zInv = -2 * p
	} else {
		z = near * p
		zInv = -1 * p
	}

	m.N[0], m.N[4], m.N[8], m.N[12] = 2*w, 0, 0, -x
	m.N[1], m.N[5], m.N[9], m.N[13] = 0, 2*h, 0, -y
	m.N[2], m.N[6], m.N[10], m.N[14] = 0, 0, zInv, -z
	m.N[3], m.N[7], m.N[11], m.N[15] = 0, 0, 0, 1
	return m
}

func (m *Matrix4) Equals(matrix *Matrix4) bool {
	return matrix.N == m.N
}

func (m *Matrix4) FromArray(array []float64) *Matrix4 {
	copy(m.N[:], array[:16])
	return m
}

func (m *Matrix4) ToArray(array []float64) *Matrix4 {
	copy(array[:16], m.N[:])
	return m
}

var _v1Matrix4 = NewVector3Defaults()
var _m1Matrix4 = NewMatrix4Identity()
var _zeroMatrix4 = NewVector3Defaults()
var _oneMatrix4 = NewVector3(1, 1, 1)
var _xMatrix4 = NewVector3Defaults()
var _yMatrix4 = NewVector3Defaults()
var _zMatrix4 = NewVector3Defaults()
