package math

type Euler struct {
	x, y, z           float64
	order             EulerOrderType
	_onChangeCallback func()
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

//todo
// 	setFromRotationMatrix( m, order = this._order, update = true ) {
//
// 		// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)
//
// 		const te = m.elements;
// 		const m11 = te[ 0 ], m12 = te[ 4 ], m13 = te[ 8 ];
// 		const m21 = te[ 1 ], m22 = te[ 5 ], m23 = te[ 9 ];
// 		const m31 = te[ 2 ], m32 = te[ 6 ], m33 = te[ 10 ];
//
// 		switch ( order ) {
//
// 			case 'XYZ':
//
// 				this._y = Math.asin( clamp( m13, - 1, 1 ) );
//
// 				if ( Math.abs( m13 ) < 0.9999999 ) {
//
// 					this._x = Math.atan2( - m23, m33 );
// 					this._z = Math.atan2( - m12, m11 );
//
// 				} else {
//
// 					this._x = Math.atan2( m32, m22 );
// 					this._z = 0;
//
// 				}
//
// 				break;
//
// 			case 'YXZ':
//
// 				this._x = Math.asin( - clamp( m23, - 1, 1 ) );
//
// 				if ( Math.abs( m23 ) < 0.9999999 ) {
//
// 					this._y = Math.atan2( m13, m33 );
// 					this._z = Math.atan2( m21, m22 );
//
// 				} else {
//
// 					this._y = Math.atan2( - m31, m11 );
// 					this._z = 0;
//
// 				}
//
// 				break;
//
// 			case 'ZXY':
//
// 				this._x = Math.asin( clamp( m32, - 1, 1 ) );
//
// 				if ( Math.abs( m32 ) < 0.9999999 ) {
//
// 					this._y = Math.atan2( - m31, m33 );
// 					this._z = Math.atan2( - m12, m22 );
//
// 				} else {
//
// 					this._y = 0;
// 					this._z = Math.atan2( m21, m11 );
//
// 				}
//
// 				break;
//
// 			case 'ZYX':
//
// 				this._y = Math.asin( - clamp( m31, - 1, 1 ) );
//
// 				if ( Math.abs( m31 ) < 0.9999999 ) {
//
// 					this._x = Math.atan2( m32, m33 );
// 					this._z = Math.atan2( m21, m11 );
//
// 				} else {
//
// 					this._x = 0;
// 					this._z = Math.atan2( - m12, m22 );
//
// 				}
//
// 				break;
//
// 			case 'YZX':
//
// 				this._z = Math.asin( clamp( m21, - 1, 1 ) );
//
// 				if ( Math.abs( m21 ) < 0.9999999 ) {
//
// 					this._x = Math.atan2( - m23, m22 );
// 					this._y = Math.atan2( - m31, m11 );
//
// 				} else {
//
// 					this._x = 0;
// 					this._y = Math.atan2( m13, m33 );
//
// 				}
//
// 				break;
//
// 			case 'XZY':
//
// 				this._z = Math.asin( - clamp( m12, - 1, 1 ) );
//
// 				if ( Math.abs( m12 ) < 0.9999999 ) {
//
// 					this._x = Math.atan2( m32, m22 );
// 					this._y = Math.atan2( m13, m11 );
//
// 				} else {
//
// 					this._x = Math.atan2( - m23, m33 );
// 					this._y = 0;
//
// 				}
//
// 				break;
//
// 			default:
//
// 				console.warn( 'THREE.Euler: .setFromRotationMatrix() encountered an unknown order: ' + order );
//
// 		}
//
// 		this._order = order;
//
// 		if ( update === true ) this._onChangeCallback();
//
// 		return this;
//
// 	}
//
// 	setFromQuaternion( q, order, update ) {
//
// 		_matrix.makeRotationFromQuaternion( q );
//
// 		return this.setFromRotationMatrix( _matrix, order, update );
//
// 	}
//
// 	setFromVector3( v, order = this._order ) {
//
// 		return this.set( v.x, v.y, v.z, order );
//
// 	}
//
// 	reorder( newOrder ) {
//
// 		// WARNING: this discards revolution information -bhouston
//
// 		_quaternion.setFromEuler( this );
//
// 		return this.setFromQuaternion( _quaternion, newOrder );
//
// 	}
//
// 	equals( euler ) {
//
// 		return ( euler._x === this._x ) && ( euler._y === this._y ) && ( euler._z === this._z ) && ( euler._order === this._order );
//
// 	}
//
// 	fromArray( array ) {
//
// 		this._x = array[ 0 ];
// 		this._y = array[ 1 ];
// 		this._z = array[ 2 ];
// 		if ( array[ 3 ] !== undefined ) this._order = array[ 3 ];
//
// 		this._onChangeCallback();
//
// 		return this;
//
// 	}
//
// 	toArray( array = [], offset = 0 ) {
//
// 		array[ offset ] = this._x;
// 		array[ offset + 1 ] = this._y;
// 		array[ offset + 2 ] = this._z;
// 		array[ offset + 3 ] = this._order;
//
// 		return array;
//
// 	}
//
// 	_onChange( callback ) {
//
// 		this._onChangeCallback = callback;
//
// 		return this;
//
// 	}
//
// 	*[ Symbol.iterator ]() {
//
// 		yield this._x;
// 		yield this._y;
// 		yield this._z;
// 		yield this._order;
//
// 	}
//
// const _matrix = /*@__PURE__*/ new Matrix4();
// const _quaternion = /*@__PURE__*/ new Quaternion();
