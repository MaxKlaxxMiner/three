package math

type Matrix2 struct {
	N [2 * 2]float64
}

func (m *Matrix2) IsMatrix2() bool { return m != nil }

func NewMatrix2Identity() *Matrix2 {
	return new(Matrix2).Identity()
}

func NewMatrix2(n11, n12, n21, n22 float64) *Matrix2 {
	return &Matrix2{
		N: [2 * 2]float64{
			n11, n21,
			n12, n22,
		},
	}
}

func (m *Matrix2) Set(n11, n12, n21, n22 float64) *Matrix2 {
	m.N = [2 * 2]float64{
		n11, n21,
		n12, n22,
	}
	return m
}

var identityMatrix2 = Matrix2{
	N: [2 * 2]float64{
		1, 0,
		0, 1,
	},
}

func (m *Matrix2) Identity() *Matrix2 {
	m.N = identityMatrix2.N
	return m
}

func (m *Matrix2) Equals(matrix *Matrix2) bool {
	return matrix.N == m.N
}

func (m *Matrix2) FromArray(array []float64) *Matrix2 {
	copy(m.N[:], array[:4])
	return m
}

func (m *Matrix2) ToArray(array []float64) *Matrix2 {
	copy(array[:4], m.N[:])
	return m
}

func (m *Matrix2) Copy(v *Matrix2) *Matrix2 {
	m.N = v.N
	return m
}

func (m *Matrix2) Clone() *Matrix2 {
	return new(Matrix2).Copy(m)
}
