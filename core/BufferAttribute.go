package core

import (
	"math"
	"reflect"
)

type BufferAttribute struct {
	Name       string
	ItemSize   int
	Count      int
	Normalized bool
	Version    int
}

func (a *BufferAttribute) IsBufferAttribute() bool { return a != nil }

type IBufferAttribute interface {
	GetBufferAttribute() *BufferAttribute
}

//class BufferAttribute {
//
//	onUploadCallback() {}
//
//	set needsUpdate( value ) {
//
//		if ( value === true ) this.version ++;
//
//	}
//
//	setUsage( value ) {
//
//		this.usage = value;
//
//		return this;
//
//	}
//
//	addUpdateRange( start, count ) {
//
//		this.updateRanges.push( { start, count } );
//
//	}
//
//	clearUpdateRanges() {
//
//		this.updateRanges.length = 0;
//
//	}
//
//	copy( source ) {
//
//		this.name = source.name;
//		this.array = new source.array.constructor( source.array );
//		this.itemSize = source.itemSize;
//		this.count = source.count;
//		this.normalized = source.normalized;
//
//		this.usage = source.usage;
//		this.gpuType = source.gpuType;
//
//		return this;
//
//	}
//
//	copyAt( index1, attribute, index2 ) {
//
//		index1 *= this.itemSize;
//		index2 *= attribute.itemSize;
//
//		for ( let i = 0, l = this.itemSize; i < l; i ++ ) {
//
//			this.array[ index1 + i ] = attribute.array[ index2 + i ];
//
//		}
//
//		return this;
//
//	}
//
//	copyArray( array ) {
//
//		this.array.set( array );
//
//		return this;
//
//	}
//
//	applyMatrix3( m ) {
//
//		if ( this.itemSize === 2 ) {
//
//			for ( let i = 0, l = this.count; i < l; i ++ ) {
//
//				_vector2.fromBufferAttribute( this, i );
//				_vector2.applyMatrix3( m );
//
//				this.setXY( i, _vector2.x, _vector2.y );
//
//			}
//
//		} else if ( this.itemSize === 3 ) {
//
//			for ( let i = 0, l = this.count; i < l; i ++ ) {
//
//				_vector.fromBufferAttribute( this, i );
//				_vector.applyMatrix3( m );
//
//				this.setXYZ( i, _vector.x, _vector.y, _vector.z );
//
//			}
//
//		}
//
//		return this;
//
//	}
//
//	applyMatrix4( m ) {
//
//		for ( let i = 0, l = this.count; i < l; i ++ ) {
//
//			_vector.fromBufferAttribute( this, i );
//
//			_vector.applyMatrix4( m );
//
//			this.setXYZ( i, _vector.x, _vector.y, _vector.z );
//
//		}
//
//		return this;
//
//	}
//
//	applyNormalMatrix( m ) {
//
//		for ( let i = 0, l = this.count; i < l; i ++ ) {
//
//			_vector.fromBufferAttribute( this, i );
//
//			_vector.applyNormalMatrix( m );
//
//			this.setXYZ( i, _vector.x, _vector.y, _vector.z );
//
//		}
//
//		return this;
//
//	}
//
//	transformDirection( m ) {
//
//		for ( let i = 0, l = this.count; i < l; i ++ ) {
//
//			_vector.fromBufferAttribute( this, i );
//
//			_vector.transformDirection( m );
//
//			this.setXYZ( i, _vector.x, _vector.y, _vector.z );
//
//		}
//
//		return this;
//
//	}
//
//	set( value, offset = 0 ) {
//
//		// Matching BufferAttribute constructor, do not normalize the array.
//		this.array.set( value, offset );
//
//		return this;
//
//	}
//
//	getComponent( index, component ) {
//
//		let value = this.array[ index * this.itemSize + component ];
//
//		if ( this.normalized ) value = denormalize( value, this.array );
//
//		return value;
//
//	}
//
//	setComponent( index, component, value ) {
//
//		if ( this.normalized ) value = normalize( value, this.array );
//
//		this.array[ index * this.itemSize + component ] = value;
//
//		return this;
//
//	}
//
//	getX( index ) {
//
//		let x = this.array[ index * this.itemSize ];
//
//		if ( this.normalized ) x = denormalize( x, this.array );
//
//		return x;
//
//	}
//
//	setX( index, x ) {
//
//		if ( this.normalized ) x = normalize( x, this.array );
//
//		this.array[ index * this.itemSize ] = x;
//
//		return this;
//
//	}
//
//	getY( index ) {
//
//		let y = this.array[ index * this.itemSize + 1 ];
//
//		if ( this.normalized ) y = denormalize( y, this.array );
//
//		return y;
//
//	}
//
//	setY( index, y ) {
//
//		if ( this.normalized ) y = normalize( y, this.array );
//
//		this.array[ index * this.itemSize + 1 ] = y;
//
//		return this;
//
//	}
//
//	getZ( index ) {
//
//		let z = this.array[ index * this.itemSize + 2 ];
//
//		if ( this.normalized ) z = denormalize( z, this.array );
//
//		return z;
//
//	}
//
//	setZ( index, z ) {
//
//		if ( this.normalized ) z = normalize( z, this.array );
//
//		this.array[ index * this.itemSize + 2 ] = z;
//
//		return this;
//
//	}
//
//	getW( index ) {
//
//		let w = this.array[ index * this.itemSize + 3 ];
//
//		if ( this.normalized ) w = denormalize( w, this.array );
//
//		return w;
//
//	}
//
//	setW( index, w ) {
//
//		if ( this.normalized ) w = normalize( w, this.array );
//
//		this.array[ index * this.itemSize + 3 ] = w;
//
//		return this;
//
//	}
//
//	setXY( index, x, y ) {
//
//		index *= this.itemSize;
//
//		if ( this.normalized ) {
//
//			x = normalize( x, this.array );
//			y = normalize( y, this.array );
//
//		}
//
//		this.array[ index + 0 ] = x;
//		this.array[ index + 1 ] = y;
//
//		return this;
//
//	}
//
//	setXYZ( index, x, y, z ) {
//
//		index *= this.itemSize;
//
//		if ( this.normalized ) {
//
//			x = normalize( x, this.array );
//			y = normalize( y, this.array );
//			z = normalize( z, this.array );
//
//		}
//
//		this.array[ index + 0 ] = x;
//		this.array[ index + 1 ] = y;
//		this.array[ index + 2 ] = z;
//
//		return this;
//
//	}
//
//	setXYZW( index, x, y, z, w ) {
//
//		index *= this.itemSize;
//
//		if ( this.normalized ) {
//
//			x = normalize( x, this.array );
//			y = normalize( y, this.array );
//			z = normalize( z, this.array );
//			w = normalize( w, this.array );
//
//		}
//
//		this.array[ index + 0 ] = x;
//		this.array[ index + 1 ] = y;
//		this.array[ index + 2 ] = z;
//		this.array[ index + 3 ] = w;
//
//		return this;
//
//	}
//
//	onUpload( callback ) {
//
//		this.onUploadCallback = callback;
//
//		return this;
//
//	}
//
//	clone() {
//
//		return new this.constructor( this.array, this.itemSize ).copy( this );
//
//	}
//
//	toJSON() {
//
//		const data = {
//			itemSize: this.itemSize,
//			type: this.array.constructor.name,
//			array: Array.from( this.array ),
//			normalized: this.normalized
//		};
//
//		if ( this.name !== '' ) data.name = this.name;
//		if ( this.usage !== StaticDrawUsage ) data.usage = this.usage;
//
//		return data;
//
//	}
//
//}
//
////
//
//class Int8BufferAttribute extends BufferAttribute {
//
//	constructor( array, itemSize, normalized ) {
//
//		super( new Int8Array( array ), itemSize, normalized );
//
//	}
//
//}
//
//class Uint8BufferAttribute extends BufferAttribute {
//
//	constructor( array, itemSize, normalized ) {
//
//		super( new Uint8Array( array ), itemSize, normalized );
//
//	}
//
//}
//
//class Uint8ClampedBufferAttribute extends BufferAttribute {
//
//	constructor( array, itemSize, normalized ) {
//
//		super( new Uint8ClampedArray( array ), itemSize, normalized );
//
//	}
//
//}
//
//class Int16BufferAttribute extends BufferAttribute {
//
//	constructor( array, itemSize, normalized ) {
//
//		super( new Int16Array( array ), itemSize, normalized );
//
//	}
//
//}
//

type Uint16BufferAttribute struct {
	BufferAttribute
	Array []uint16
}

func NewUint16BufferAttribute(array []uint16, itemSize int) *Uint16BufferAttribute {
	return NewUint16BufferAttributeNormalized(array, itemSize, false)
}

func NewUint16BufferAttributeNormalized(array []uint16, itemSize int, normalized bool) *Uint16BufferAttribute {
	this := new(Uint16BufferAttribute)
	this.ItemSize = itemSize
	this.Count = len(array) / itemSize
	this.Normalized = normalized
	this.Array = array

	//todo
	//		this.usage = StaticDrawUsage;
	//		this.updateRanges = [];
	//		this.gpuType = FloatType;

	return this
}

func NewUint16BufferAttributeFromAny(array any, itemSize int) *Uint16BufferAttribute {
	return NewUint16BufferAttributeNormalizedFromAny(array, itemSize, false)
}

func NewUint16BufferAttributeNormalizedFromAny(array any, itemSize int, normalized bool) *Uint16BufferAttribute {
	return NewUint16BufferAttributeNormalized(convertingSliceAuto[uint16](array), itemSize, normalized)
}

func (a *Uint16BufferAttribute) GetBufferAttribute() *BufferAttribute {
	return &a.BufferAttribute
}

//todo
//class Int32BufferAttribute extends BufferAttribute {
//
//	constructor( array, itemSize, normalized ) {
//
//		super( new Int32Array( array ), itemSize, normalized );
//
//	}
//
//}

type Uint32BufferAttribute struct {
	BufferAttribute
	Array []uint32
}

func NewUint32BufferAttribute(array []uint32, itemSize int) *Uint32BufferAttribute {
	return NewUint32BufferAttributeNormalized(array, itemSize, false)
}

func NewUint32BufferAttributeNormalized(array []uint32, itemSize int, normalized bool) *Uint32BufferAttribute {
	this := new(Uint32BufferAttribute)
	this.ItemSize = itemSize
	this.Count = len(array) / itemSize
	this.Normalized = normalized
	this.Array = array

	//todo
	//		this.usage = StaticDrawUsage;
	//		this.updateRanges = [];
	//		this.gpuType = FloatType;

	return this
}

func NewUint32BufferAttributeFromAny(array any, itemSize int) *Uint32BufferAttribute {
	return NewUint32BufferAttributeNormalizedFromAny(array, itemSize, false)
}

func NewUint32BufferAttributeNormalizedFromAny(array any, itemSize int, normalized bool) *Uint32BufferAttribute {
	return NewUint32BufferAttributeNormalized(convertingSliceAuto[uint32](array), itemSize, normalized)
}

func (a *Uint32BufferAttribute) AutoDownToUint16() IBufferAttribute {
	if len(a.Array) > math.MaxUint16 { // quick skip
		return a
	}
	tmp := make([]uint16, len(a.Array))
	for i, v := range a.Array {
		if v > math.MaxUint16 {
			return a // skip
		}
		tmp[i] = uint16(v)
	}
	return NewUint16BufferAttributeNormalized(tmp, a.ItemSize, a.Normalized)
}

func (a *Uint32BufferAttribute) GetBufferAttribute() *BufferAttribute {
	return &a.BufferAttribute
}

//todo
//class Float16BufferAttribute extends BufferAttribute {
//
//	constructor( array, itemSize, normalized ) {
//
//		super( new Uint16Array( array ), itemSize, normalized );
//
//		this.isFloat16BufferAttribute = true;
//
//	}
//
//	getX( index ) {
//
//		let x = fromHalfFloat( this.array[ index * this.itemSize ] );
//
//		if ( this.normalized ) x = denormalize( x, this.array );
//
//		return x;
//
//	}
//
//	setX( index, x ) {
//
//		if ( this.normalized ) x = normalize( x, this.array );
//
//		this.array[ index * this.itemSize ] = toHalfFloat( x );
//
//		return this;
//
//	}
//
//	getY( index ) {
//
//		let y = fromHalfFloat( this.array[ index * this.itemSize + 1 ] );
//
//		if ( this.normalized ) y = denormalize( y, this.array );
//
//		return y;
//
//	}
//
//	setY( index, y ) {
//
//		if ( this.normalized ) y = normalize( y, this.array );
//
//		this.array[ index * this.itemSize + 1 ] = toHalfFloat( y );
//
//		return this;
//
//	}
//
//	getZ( index ) {
//
//		let z = fromHalfFloat( this.array[ index * this.itemSize + 2 ] );
//
//		if ( this.normalized ) z = denormalize( z, this.array );
//
//		return z;
//
//	}
//
//	setZ( index, z ) {
//
//		if ( this.normalized ) z = normalize( z, this.array );
//
//		this.array[ index * this.itemSize + 2 ] = toHalfFloat( z );
//
//		return this;
//
//	}
//
//	getW( index ) {
//
//		let w = fromHalfFloat( this.array[ index * this.itemSize + 3 ] );
//
//		if ( this.normalized ) w = denormalize( w, this.array );
//
//		return w;
//
//	}
//
//	setW( index, w ) {
//
//		if ( this.normalized ) w = normalize( w, this.array );
//
//		this.array[ index * this.itemSize + 3 ] = toHalfFloat( w );
//
//		return this;
//
//	}
//
//	setXY( index, x, y ) {
//
//		index *= this.itemSize;
//
//		if ( this.normalized ) {
//
//			x = normalize( x, this.array );
//			y = normalize( y, this.array );
//
//		}
//
//		this.array[ index + 0 ] = toHalfFloat( x );
//		this.array[ index + 1 ] = toHalfFloat( y );
//
//		return this;
//
//	}
//
//	setXYZ( index, x, y, z ) {
//
//		index *= this.itemSize;
//
//		if ( this.normalized ) {
//
//			x = normalize( x, this.array );
//			y = normalize( y, this.array );
//			z = normalize( z, this.array );
//
//		}
//
//		this.array[ index + 0 ] = toHalfFloat( x );
//		this.array[ index + 1 ] = toHalfFloat( y );
//		this.array[ index + 2 ] = toHalfFloat( z );
//
//		return this;
//
//	}
//
//	setXYZW( index, x, y, z, w ) {
//
//		index *= this.itemSize;
//
//		if ( this.normalized ) {
//
//			x = normalize( x, this.array );
//			y = normalize( y, this.array );
//			z = normalize( z, this.array );
//			w = normalize( w, this.array );
//
//		}
//
//		this.array[ index + 0 ] = toHalfFloat( x );
//		this.array[ index + 1 ] = toHalfFloat( y );
//		this.array[ index + 2 ] = toHalfFloat( z );
//		this.array[ index + 3 ] = toHalfFloat( w );
//
//		return this;
//
//	}
//
//}

type Float32BufferAttribute struct {
	BufferAttribute
	Array []float32
}

func NewFloat32BufferAttribute(array []float32, itemSize int) *Float32BufferAttribute {
	return NewFloat32BufferAttributeNormalized(array, itemSize, false)
}

func NewFloat32BufferAttributeNormalized(array []float32, itemSize int, normalized bool) *Float32BufferAttribute {
	this := new(Float32BufferAttribute)
	this.ItemSize = itemSize
	this.Count = len(array) / itemSize
	this.Normalized = normalized
	this.Array = array

	//todo
	//		this.usage = StaticDrawUsage;
	//		this.updateRanges = [];
	//		this.gpuType = FloatType;

	return this
}

func NewFloat32BufferAttributeFromAny(array any, itemSize int) *Float32BufferAttribute {
	return NewFloat32BufferAttributeNormalizedFromAny(array, itemSize, false)
}

func NewFloat32BufferAttributeNormalizedFromAny(array any, itemSize int, normalized bool) *Float32BufferAttribute {
	return NewFloat32BufferAttributeNormalized(convertingSliceAuto[float32](array), itemSize, normalized)
}

func (a *Float32BufferAttribute) GetBufferAttribute() *BufferAttribute {
	return &a.BufferAttribute
}


type numberTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func convertingSliceAuto[TDst numberTypes](array any) []TDst {
	if array == nil {
		return nil
	}
	switch arr := array.(type) {
	case []int:
		return convertingSlice[TDst, int](arr)
	case []int8:
		return convertingSlice[TDst, int8](arr)
	case []int16:
		return convertingSlice[TDst, int16](arr)
	case []int32:
		return convertingSlice[TDst, int32](arr)
	case []int64:
		return convertingSlice[TDst, int64](arr)
	case []uint:
		return convertingSlice[TDst, uint](arr)
	case []uint8:
		return convertingSlice[TDst, uint8](arr)
	case []uint16:
		return convertingSlice[TDst, uint16](arr)
	case []uint32:
		return convertingSlice[TDst, uint32](arr)
	case []uint64:
		return convertingSlice[TDst, uint64](arr)
	case []float32:
		return convertingSlice[TDst, float32](arr)
	case []float64:
		return convertingSlice[TDst, float64](arr)
	default:
		panic("invalid array type:" + reflect.TypeOf(array).String())
	}
}

func convertingSlice[TDst, TSrc numberTypes](src []TSrc) (dst []TDst) {
	dst = make([]TDst, len(src))
	for i := range src {
		dst[i] = TDst(src[i])
	}
	return
}

//import { denormalize, normalize } from '../math/MathUtils.js';
//import { StaticDrawUsage, FloatType } from '../constants.js';
//import { fromHalfFloat, toHalfFloat } from '../extras/DataUtils.js';
//
//const _vector = /*@__PURE__*/ new Vector3();
//const _vector2 = /*@__PURE__*/ new Vector2();
