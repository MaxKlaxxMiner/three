package math

import (
	"github.com/MaxKlaxxMiner/three/utils"
	"math"
	"strconv"
)

type Vector2 struct {
	X, Y float64
}

func NewVector2Defaults() *Vector2 {
	return new(Vector2)
}

func NewVector2(x, y float64) *Vector2 {
	return &Vector2{x, y}
}

func (v *Vector2) IsVector2() bool { return v != nil }

func (v *Vector2) GetWidth() float64       { return v.X }
func (v *Vector2) SetWidth(value float64)  { v.X = value }
func (v *Vector2) GetHeight() float64      { return v.Y }
func (v *Vector2) SetHeight(value float64) { v.Y = value }

func (v *Vector2) Set(x, y float64) *Vector2 {
	v.X, v.Y = x, y
	return v
}

func (v *Vector2) SetScalar(scalar float64) *Vector2 {
	v.X, v.Y = scalar, scalar
	return v
}

func (v *Vector2) SetX(x float64) *Vector2 {
	v.X = x
	return v
}

func (v *Vector2) SetY(y float64) *Vector2 {
	v.Y = y
	return v
}

func (v *Vector2) SetComponent(index int, value float64) *Vector2 {
	switch index {
	case 0, 'x', 'X':
		v.X = value
	case 1, 'y', 'Y':
		v.Y = value
	default:
		panic("index is out of range: " + strconv.Itoa(index))
	}
	return v
}

func (v *Vector2) GetComponent(index int) float64 {
	switch index {
	case 0, 'x', 'X':
		return v.X
	case 1, 'y', 'Y':
		return v.Y
	default:
		panic("index is out of range: " + strconv.Itoa(index))
	}
}

func (v *Vector2) Clone() *Vector2 {
	return NewVector2(v.X, v.Y)
}

func (v *Vector2) Copy(a *Vector2) *Vector2 {
	v.X, v.Y = a.X, a.Y
	return v
}

func (v *Vector2) Add(a *Vector2) *Vector2 {
	v.X += a.X
	v.Y += a.Y
	return v
}

func (v *Vector2) AddScalar(s float64) *Vector2 {
	v.X += s
	v.Y += s
	return v
}

func (v *Vector2) AddVectors(a, b *Vector2) *Vector2 {
	return v.Set(a.X+b.X, a.Y+b.Y)
}

func (v *Vector2) AddScaledVector(a *Vector2, s float64) *Vector2 {
	v.X += a.X * s
	v.Y += a.Y * s
	return v
}

func (v *Vector2) Sub(a *Vector2) *Vector2 {
	v.X -= a.X
	v.Y -= a.Y
	return v
}

func (v *Vector2) SubScalar(s float64) *Vector2 {
	v.X -= s
	v.Y -= s
	return v
}

func (v *Vector2) SubVectors(a, b *Vector2) *Vector2 {
	return v.Set(a.X-b.X, a.Y-b.Y)
}

func (v *Vector2) Multiply(a *Vector2) *Vector2 {
	v.X *= a.X
	v.Y *= a.Y
	return v
}

func (v *Vector2) MultiplyScalar(s float64) *Vector2 {
	v.X *= s
	v.Y *= s
	return v
}

func (v *Vector2) Divide(a *Vector2) *Vector2 {
	v.X /= a.X
	v.Y /= a.Y
	return v
}

func (v *Vector2) DivideScalar(s float64) *Vector2 {
	return v.MultiplyScalar(1 / s)
}

func (v *Vector2) ApplyMatrix3(m *Matrix3) *Vector2 {
	x, y := v.X, v.Y
	v.X = m.N[0]*x + m.N[3]*y + m.N[6]
	v.Y = m.N[1]*x + m.N[4]*y + m.N[7]
	return v
}

func (v *Vector2) Min(a *Vector2) *Vector2 {
	v.X = math.Min(v.X, a.X)
	v.Y = math.Min(v.Y, a.Y)
	return v
}

func (v *Vector2) Max(a *Vector2) *Vector2 {
	v.X = math.Max(v.X, a.X)
	v.Y = math.Max(v.Y, a.Y)
	return v
}

func (v *Vector2) Clamp(min, max *Vector2) *Vector2 {
	// assumes min < max, componentwise
	v.X = Clamp(v.X, min.X, max.X)
	v.Y = Clamp(v.Y, min.Y, max.Y)
	return v
}

func (v *Vector2) ClampScalar(minVal, maxVal float64) *Vector2 {
	v.X = Clamp(v.X, minVal, maxVal)
	v.Y = Clamp(v.Y, minVal, maxVal)
	return v
}

func (v *Vector2) ClampLength(minVal, maxVal float64) *Vector2 {
	length := v.Length()
	return v.DivideScalar(utils.If(length > 0, length, 1)).MultiplyScalar(Clamp(length, minVal, maxVal))
}

func (v *Vector2) Floor() *Vector2 {
	v.X = math.Floor(v.X)
	v.Y = math.Floor(v.Y)
	return v
}

func (v *Vector2) Ceil() *Vector2 {
	v.X = math.Ceil(v.X)
	v.Y = math.Ceil(v.Y)
	return v
}

func (v *Vector2) Round() *Vector2 {
	v.X = math.Round(v.X)
	v.Y = math.Round(v.Y)
	return v
}

func (v *Vector2) RoundToZero() *Vector2 {
	v.X = math.Trunc(v.X)
	v.Y = math.Trunc(v.Y)
	return v
}

func (v *Vector2) Negate() *Vector2 {
	v.X = -v.X
	v.Y = -v.Y
	return v
}

func (v *Vector2) Dot(a *Vector2) float64 {
	return v.X*a.X + v.Y*a.Y
}

func (v *Vector2) Cross(a *Vector2) float64 {
	return v.X*a.Y - v.Y*a.X
}

func (v *Vector2) LengthSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vector2) Length() float64 {
	return math.Sqrt(v.LengthSq())
}

func (v *Vector2) LengthOrOne() float64 {
	if l := v.Length(); l > 0 {
		return l
	}
	return 1
}

func (v *Vector2) ManhattanLength() float64 {
	return math.Abs(v.X) + math.Abs(v.Y)
}

func (v *Vector2) Normalize() *Vector2 {
	return v.DivideScalar(v.LengthOrOne())
}

func (v *Vector2) Angle() float64 {
	// computes the angle in radians with respect to the positive x-axis
	return math.Atan2(-v.Y, -v.X) + math.Pi
}

func (v *Vector2) AngleTo(a *Vector2) float64 {
	denominator := math.Sqrt(v.LengthSq() * a.LengthSq())
	if denominator == 0 {
		return math.Pi / 2
	}
	theta := v.Dot(v) / denominator

	// clamp, to handle numerical problems
	return math.Acos(Clamp(theta, -1, 1))
}

func (v *Vector2) DistanceTo(a *Vector2) float64 {
	return math.Sqrt(v.DistanceToSquared(a))
}

func (v *Vector2) DistanceToSquared(a *Vector2) float64 {
	dx, dy := v.X-a.X, v.Y-a.Y
	return dx*dx + dy*dy
}

func (v *Vector2) ManhattanDistanceTo(a *Vector2) float64 {
	return math.Abs(v.X-a.X) + math.Abs(v.Y-a.Y)
}

func (v *Vector2) SetLength(length float64) *Vector2 {
	return v.Normalize().MultiplyScalar(length)
}

func (v *Vector2) Lerp(a *Vector2, alpha float64) *Vector2 {
	v.X += (a.X - v.X) * alpha
	v.Y += (a.Y - v.Y) * alpha
	return v
}

func (v *Vector2) LerpVectors(a, b *Vector2, alpha float64) *Vector2 {
	v.X = a.X + (b.X-a.X)*alpha
	v.Y = a.Y + (b.Y-a.Y)*alpha
	return v
}

func (v *Vector2) Equals(a *Vector2) bool {
	return *v == *a
}

func (v *Vector2) FromArray(array []float64) *Vector2 {
	_ = array[1]
	v.X, v.Y = array[0], array[1]
	return v
}

func (v *Vector2) ToArray(array []float64) []float64 {
	_ = array[1]
	array[0], array[1] = v.X, v.Y
	return array
}

//todo
// 	fromBufferAttribute( attribute, index ) {
// 		this.x = attribute.getX( index );
// 		this.y = attribute.getY( index );
// 		return this;
// 	}

func (v *Vector2) RotateAround(center *Vector2, angle float64) *Vector2 {
	c, s := math.Cos(angle), math.Sin(angle)
	x, y := v.X-center.X, v.Y-center.Y
	v.X = x*c - y*s + center.X
	v.Y = x*s + y*c + center.Y
	return v
}

func (v *Vector2) Random() *Vector2 {
	v.X = RandomFloat()
	v.Y = RandomFloat()
	return v
}

func (v *Vector2) Append(buf []float64) []float64 {
	return append(buf, v.X, v.Y)
}
