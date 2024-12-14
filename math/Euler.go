package math

import (
	"math"
	"strconv"
)

type Euler struct {
	x, y, z           float64
	order             EulerOrderType
	_onChangeCallback func()
}

func NewEulerDefaults() *Euler {
	return &Euler{0, 0, 0, EulerDefaultOrder, func() {}}
}

func NewEuler(x, y, z float64) *Euler {
	return &Euler{x, y, z, EulerDefaultOrder, func() {}}
}

func NewEulerWithOrder(x, y, z float64, order EulerOrderType) *Euler {
	return &Euler{x, y, z, order, func() {}}
}

type EulerOrderType int

var EulerDefaultOrder = EulerOrderXYZ

const (
	EulerOrderXYZ EulerOrderType = iota
	EulerOrderXZY
	EulerOrderYXZ
	EulerOrderYZX
	EulerOrderZXY
	EulerOrderZYX
)

func (e *Euler) IsEuler() bool { return e != nil }

func (e *Euler) GetX() float64                       { return e.x }
func (e *Euler) GetY() float64                       { return e.y }
func (e *Euler) GetZ() float64                       { return e.z }
func (e *Euler) GetOrder() EulerOrderType            { return e.order }
func (e *Euler) GetXYZ() (float64, float64, float64) { return e.x, e.y, e.z }

func (e *Euler) SetX(x float64) {
	e.x = x
	e._onChangeCallback()
}
func (e *Euler) SetY(y float64) {
	e.y = y
	e._onChangeCallback()
}
func (e *Euler) SetZ(z float64) {
	e.z = z
	e._onChangeCallback()
}
func (e *Euler) SetOrder(order EulerOrderType) {
	e.order = order
	e._onChangeCallback()
}
func (e *Euler) SetOrderNoUpdate(order EulerOrderType) {
	e.order = order
}
func (e *Euler) SetXYZ(x, y, z float64) *Euler {
	e.x, e.y, e.z = x, y, z
	e._onChangeCallback()
	return e
}
func (e *Euler) SetXYZOrder(x, y, z float64, order EulerOrderType) *Euler {
	e.x, e.y, e.z = x, y, z
	e.order = order
	e._onChangeCallback()
	return e
}

func (e *Euler) Clone() *Euler {
	return NewEulerWithOrder(e.x, e.y, e.z, e.order)
}

func (e *Euler) Copy(a *Euler) *Euler {
	e.x, e.y, e.z = a.x, a.y, a.z
	e.order = a.order
	e._onChangeCallback()
	return e
}

func (e *Euler) SetFromRotationMatrix(m *Matrix4) *Euler {
	return e.SetFromRotationMatrixOrderUpdate(m, e.order, true)
}

func (e *Euler) SetFromRotationMatrixOrderUpdate(m *Matrix4, order EulerOrderType, update bool) *Euler {
	// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)

	m11, m12, m13 := m.N[0], m.N[4], m.N[8]
	m21, m22, m23 := m.N[1], m.N[5], m.N[9]
	m31, m32, m33 := m.N[2], m.N[6], m.N[10]

	switch order {
	case EulerOrderXYZ:
		e.y = math.Asin(Clamp(m13, -1, 1))
		if math.Abs(m13) < 0.9999999 {
			e.x = math.Atan2(-m23, m33)
			e.z = math.Atan2(-m12, m11)
		} else {
			e.x = math.Atan2(m32, m22)
			e.z = 0
		}
	case EulerOrderYXZ:
		e.x = math.Asin(-Clamp(m23, -1, 1))
		if math.Abs(m23) < 0.9999999 {
			e.y = math.Atan2(m13, m33)
			e.z = math.Atan2(m21, m22)
		} else {
			e.y = math.Atan2(-m31, m11)
			e.z = 0
		}
	case EulerOrderZXY:
		e.x = math.Asin(Clamp(m32, -1, 1))
		if math.Abs(m32) < 0.9999999 {
			e.y = math.Atan2(-m31, m33)
			e.z = math.Atan2(-m12, m22)
		} else {
			e.y = 0
			e.z = math.Atan2(m21, m11)
		}
	case EulerOrderZYX:
		e.y = math.Asin(-Clamp(m31, -1, 1))
		if math.Abs(m31) < 0.9999999 {
			e.x = math.Atan2(m32, m33)
			e.z = math.Atan2(m21, m11)
		} else {
			e.x = 0
			e.z = math.Atan2(-m12, m22)
		}
	case EulerOrderYZX:
		e.z = math.Asin(Clamp(m21, -1, 1))
		if math.Abs(m21) < 0.9999999 {
			e.x = math.Atan2(-m23, m22)
			e.y = math.Atan2(-m31, m11)
		} else {
			e.x = 0
			e.y = math.Atan2(m13, m33)
		}
	case EulerOrderXZY:
		e.z = math.Asin(-Clamp(m12, -1, 1))
		if math.Abs(m12) < 0.9999999 {
			e.x = math.Atan2(m32, m22)
			e.y = math.Atan2(m13, m11)
		} else {
			e.x = math.Atan2(-m23, m33)
			e.y = 0
		}
	default:
		panic("THREE.Euler: .setFromRotationMatrix() encountered an unknown order: " + strconv.Itoa(int(order)))
	}

	e.order = order

	if update {
		e._onChangeCallback()
	}

	return e
}

func (e *Euler) SetFromQuaternion(q *Quaternion, order EulerOrderType, update bool) *Euler {
	_matrixEuler.MakeRotationFromQuaternion(q)
	return e.SetFromRotationMatrixOrderUpdate(_matrixEuler, order, update)
}

func (e *Euler) SetFromVector3(v *Vector3) *Euler {
	return e.SetXYZ(v.X, v.Y, v.Z)
}

func (e *Euler) SetFromVector3Order(v *Vector3, order EulerOrderType) *Euler {
	return e.SetXYZOrder(v.X, v.Y, v.Z, order)
}

func (e *Euler) Reorder(newOrder EulerOrderType) *Euler {
	// WARNING: this discards revolution information -bhouston
	_quaternionEuler.SetFromEuler(e)
	return e.SetFromQuaternion(_quaternionEuler, newOrder, false)
}

func (e *Euler) Equals(euler *Euler) bool {
	return euler.x == e.x && euler.y == e.y && euler.z == e.z && euler.order == e.order
}

func (e *Euler) FromArray(array []float64) *Euler {
	e.x = array[0]
	e.y = array[1]
	e.z = array[2]
	if len(array) > 3 && array[3] > 1000000 {
		e.order = EulerOrderType(int(array[3])/1000000 - 1)
	}
	e._onChangeCallback()
	return e
}

func (e *Euler) ToArray(array []float64) []float64 {
	array[0] = e.x
	array[1] = e.y
	array[2] = e.z
	if len(array) > 3 {
		array[3] = float64((int(e.order) + 1) * 1000001)
	}
	return array
}

func (e *Euler) OverrideOnChange(callback func()) *Euler {
	e._onChangeCallback = callback
	return e
}

func (e *Euler) Append(array []float64) []float64 {
	return append(array, e.x, e.y, e.y, float64((int(e.order)+1)*1000001))
}

var _matrixEuler = NewMatrix4Identity()
var _quaternionEuler = NewQuaternionDefaults()
