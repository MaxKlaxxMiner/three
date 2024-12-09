package math

import "math"

type Matrix3 struct {
	N [3 * 3]float64
}

func (m *Matrix3) IsMatrix3() bool { return m != nil }

func NewMatrix3Identity() *Matrix3 {
	return new(Matrix3).Identity()
}

func NewMatrix3(n11, n12, n13, n21, n22, n23, n31, n32, n33 float64) *Matrix3 {
	return &Matrix3{
		N: [3 * 3]float64{
			n11, n21, n31,
			n12, n22, n32,
			n13, n23, n33,
		},
	}
}

func (m *Matrix3) Set(n11, n12, n13, n21, n22, n23, n31, n32, n33 float64) *Matrix3 {
	m.N = [3 * 3]float64{
		n11, n21, n31,
		n12, n22, n32,
		n13, n23, n33,
	}
	return m
}

var zeroMatrix3 Matrix3

var identityMatrix3 = Matrix3{
	N: [3 * 3]float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	},
}

func (m *Matrix3) Identity() *Matrix3 {
	m.N = identityMatrix3.N
	return m
}

func (m *Matrix3) Copy(v *Matrix3) *Matrix3 {
	m.N = v.N
	return m
}

func (m *Matrix3) ExtractBasis(xAxis, yAxis, zAxis *Vector3) *Matrix3 {
	xAxis.SetFromMatrix3Column(m, 0)
	yAxis.SetFromMatrix3Column(m, 1)
	zAxis.SetFromMatrix3Column(m, 2)
	return m
}

func (m *Matrix3) SetFromMatrix4(v *Matrix4) *Matrix3 {
	return m.Set(
		v.N[0], v.N[4], v.N[8],
		v.N[1], v.N[5], v.N[9],
		v.N[2], v.N[6], v.N[10],
	)
}

func (m *Matrix3) Multiply(v *Matrix3) *Matrix3 {
	return m.MultiplyMatrices(m, v)
}

func (m *Matrix3) Premultiply(v *Matrix3) *Matrix3 {
	return m.MultiplyMatrices(v, m)
}

func (m *Matrix3) MultiplyMatrices(a, b *Matrix3) *Matrix3 {
	a11, a12, a13 := a.N[0], a.N[3], a.N[6]
	a21, a22, a23 := a.N[1], a.N[4], a.N[7]
	a31, a32, a33 := a.N[2], a.N[5], a.N[8]

	b11, b12, b13 := b.N[0], b.N[3], b.N[6]
	b21, b22, b23 := b.N[1], b.N[4], b.N[7]
	b31, b32, b33 := b.N[2], b.N[5], b.N[8]

	m.N[0] = a11*b11 + a12*b21 + a13*b31
	m.N[3] = a11*b12 + a12*b22 + a13*b32
	m.N[6] = a11*b13 + a12*b23 + a13*b33

	m.N[1] = a21*b11 + a22*b21 + a23*b31
	m.N[4] = a21*b12 + a22*b22 + a23*b32
	m.N[7] = a21*b13 + a22*b23 + a23*b33

	m.N[2] = a31*b11 + a32*b21 + a33*b31
	m.N[5] = a31*b12 + a32*b22 + a33*b32
	m.N[8] = a31*b13 + a32*b23 + a33*b33

	return m
}

func (m *Matrix3) MultiplyScalar(s float64) *Matrix3 {
	m.N[0] *= s
	m.N[1] *= s
	m.N[2] *= s
	m.N[3] *= s
	m.N[4] *= s
	m.N[5] *= s
	m.N[6] *= s
	m.N[7] *= s
	m.N[8] *= s
	return m
}

func (m *Matrix3) Determinant() float64 {
	a, b, c := m.N[0], m.N[1], m.N[2]
	d, e, f := m.N[3], m.N[4], m.N[5]
	g, h, i := m.N[6], m.N[7], m.N[8]
	return a*e*i - a*f*h - b*d*i + b*f*g + c*d*h - c*e*g
}

func (m *Matrix3) Invert() *Matrix3 {
	n11, n21, n31 := m.N[0], m.N[1], m.N[2]
	n12, n22, n32 := m.N[3], m.N[4], m.N[5]
	n13, n23, n33 := m.N[6], m.N[7], m.N[8]

	t11 := n33*n22 - n32*n23
	t12 := n32*n13 - n33*n12
	t13 := n23*n12 - n22*n13

	det := n11*t11 + n21*t12 + n31*t13

	if det == 0 {
		return m.Set(0, 0, 0, 0, 0, 0, 0, 0, 0)
	}

	detInv := 1 / det

	m.N[0] = t11 * detInv
	m.N[1] = (n31*n23 - n33*n21) * detInv
	m.N[2] = (n32*n21 - n31*n22) * detInv

	m.N[3] = t12 * detInv
	m.N[4] = (n33*n11 - n31*n13) * detInv
	m.N[5] = (n31*n12 - n32*n11) * detInv

	m.N[6] = t13 * detInv
	m.N[7] = (n21*n13 - n23*n11) * detInv
	m.N[8] = (n22*n11 - n21*n12) * detInv

	return m
}

func (m *Matrix3) Transpose() *Matrix3 {
	m.N[1], m.N[3] = m.N[3], m.N[1]
	m.N[2], m.N[6] = m.N[6], m.N[2]
	m.N[5], m.N[7] = m.N[7], m.N[5]
	return m
}

func (m *Matrix3) GetNormalMatrix(matrix4 *Matrix4) *Matrix3 {
	return m.SetFromMatrix4(matrix4).Invert().Transpose()
}

func (m *Matrix3) TransposeIntoArray(r []float64) *Matrix3 {
	_ = r[8]
	r[0] = m.N[0]
	r[1] = m.N[3]
	r[2] = m.N[6]
	r[3] = m.N[1]
	r[4] = m.N[4]
	r[5] = m.N[7]
	r[6] = m.N[2]
	r[7] = m.N[5]
	r[8] = m.N[8]
	return m
}

func (m *Matrix3) SetUvTransform(tx, ty, sx, sy, rotation, cx, cy float64) *Matrix3 {
	c := math.Cos(rotation)
	s := math.Sin(rotation)

	return m.Set(
		sx*c, sx*s, -sx*(c*cx+s*cy)+cx+tx,
		-sy*s, sy*c, -sy*(-s*cx+c*cy)+cy+ty,
		0, 0, 1,
	)
}

func (m *Matrix3) Scale(sx, sy float64) *Matrix3 {
	return m.Premultiply(_m3Matrix3.MakeScale(sx, sy))
}

func (m *Matrix3) Rotate(theta float64) *Matrix3 {
	return m.Premultiply(_m3Matrix3.MakeRotation(-theta))
}

func (m *Matrix3) Translate(tx, ty float64) *Matrix3 {
	return m.Premultiply(_m3Matrix3.MakeTranslation(tx, ty))
}

// --- for 2D Transforms ---

func (m *Matrix3) MakeTranslationVector2(v *Vector2) *Matrix3 {
	return m.MakeTranslation(v.X, v.Y)
}

func (m *Matrix3) MakeTranslation(x, y float64) *Matrix3 {
	return m.Set(
		1, 0, x,
		0, 1, y,
		0, 0, 1,
	)
}

func (m *Matrix3) MakeRotation(theta float64) *Matrix3 {
	// counterclockwise
	c := math.Cos(theta)
	s := math.Sin(theta)

	return m.Set(
		c, -s, 0,
		s, c, 0,
		0, 0, 1,
	)
}

func (m *Matrix3) MakeScale(x, y float64) *Matrix3 {
	return m.Set(
		x, 0, 0,
		0, y, 0,
		0, 0, 1,
	)
}

func (m *Matrix3) Equals(matrix *Matrix3) bool {
	return matrix.N == m.N
}

func (m *Matrix3) FromArray(array []float64) *Matrix3 {
	copy(m.N[:], array[:9])
	return m
}

func (m *Matrix3) ToArray(array []float64) *Matrix3 {
	copy(array[:9], m.N[:])
	return m
}

func (m *Matrix3) Clone() *Matrix3 {
	return new(Matrix3).Copy(m)
}

var _m3Matrix3 = NewMatrix3Identity()
