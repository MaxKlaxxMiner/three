package mathutils

import (
	"math"
	"time"
)

var lut = "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f404142434445464748494a4b4c4d4e4f505152535455565758595a5b5c5d5e5f606162636465666768696a6b6c6d6e6f707172737475767778797a7b7c7d7e7f808182838485868788898a8b8c8d8e8f909192939495969798999a9b9c9d9e9fa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebfc0c1c2c3c4c5c6c7c8c9cacbcccdcecfd0d1d2d3d4d5d6d7d8d9dadbdcdddedfe0e1e2e3e4e5e6e7e8e9eaebecedeeeff0f1f2f3f4f5f6f7f8f9fafbfcfdfeff"
var rnd = QuickRandom(time.Now().UnixNano())

const DEG2RAD = math.Pi / 180
const RAD2DEG = 180 / math.Pi

type UUID struct {
	d0, d1, d2, d3 uint32
}

func GenerateUUID() (r UUID) {
	rnd = rnd.Next()
	r.d0 = uint32(rnd)
	rnd = rnd.Next()
	r.d1 = uint32(rnd)
	rnd = rnd.Next()
	r.d2 = uint32(rnd)
	rnd = rnd.Next()
	r.d3 = uint32(rnd)
	return
}

func (u UUID) String() string {
	return lut[u.d0&0xff:][:2] + lut[(u.d0>>8)&0xff:][:2] + lut[(u.d0>>16)&0xff:][:2] + lut[(u.d0>>24)&0xff:][:2] + "-" +
		lut[u.d1&0xff:][:2] + lut[(u.d1>>8)&0xff:][:2] + "-" + lut[(u.d1>>16)&0x0f|0x40:][:2] + lut[(u.d1>>24)&0xff:][:2] + "-" +
		lut[(u.d2&0x3f)|0x80:][:2] + lut[(u.d2>>8)&0xff:][:2] + "-" + lut[(u.d2>>16)&0xff:][:2] + lut[(u.d2>>24)&0xff:][:2] +
		lut[u.d3&0xff:][:2] + lut[(u.d3>>8)&0xff:][:2] + lut[(u.d3>>16)&0xff:][:2] + lut[(u.d3>>24)&0xff:][:2]
}

// --- clamp( value, min, max ) ---

//func Clamp(value, min, max float64) float64 {
//	return math.Max(min, math.Min(max, value))
//}

// todo
//// compute euclidean modulo of m % n
//// https://en.wikipedia.org/wiki/Modulo_operation
//function euclideanModulo( n, m ) {
//
//	return ( ( n % m ) + m ) % m;
//
//}
//
//// Linear mapping from range <a1, a2> to range <b1, b2>
//function mapLinear( x, a1, a2, b1, b2 ) {
//
//	return b1 + ( x - a1 ) * ( b2 - b1 ) / ( a2 - a1 );
//
//}
//
//// https://www.gamedev.net/tutorials/programming/general-and-gameplay-programming/inverse-lerp-a-super-useful-yet-often-overlooked-function-r5230/
//function inverseLerp( x, y, value ) {
//
//	if ( x !== y ) {
//
//		return ( value - x ) / ( y - x );
//
//	} else {
//
//		return 0;
//
//	}
//
//}
//
//// https://en.wikipedia.org/wiki/Linear_interpolation
//function lerp( x, y, t ) {
//
//	return ( 1 - t ) * x + t * y;
//
//}
//
//// http://www.rorydriscoll.com/2016/03/07/frame-rate-independent-damping-using-lerp/
//function damp( x, y, lambda, dt ) {
//
//	return lerp( x, y, 1 - Math.exp( - lambda * dt ) );
//
//}
//
//// https://www.desmos.com/calculator/vcsjnyz7x4
//function pingpong( x, length = 1 ) {
//
//	return length - Math.abs( euclideanModulo( x, length * 2 ) - length );
//
//}
//
//// http://en.wikipedia.org/wiki/Smoothstep
//function smoothstep( x, min, max ) {
//
//	if ( x <= min ) return 0;
//	if ( x >= max ) return 1;
//
//	x = ( x - min ) / ( max - min );
//
//	return x * x * ( 3 - 2 * x );
//
//}
//
//function smootherstep( x, min, max ) {
//
//	if ( x <= min ) return 0;
//	if ( x >= max ) return 1;
//
//	x = ( x - min ) / ( max - min );
//
//	return x * x * x * ( x * ( x * 6 - 15 ) + 10 );
//
//}
//
//// Random integer from <low, high> interval
//function randInt( low, high ) {
//
//	return low + Math.floor( Math.random() * ( high - low + 1 ) );
//
//}
//
//// Random float from <low, high> interval
//function randFloat( low, high ) {
//
//	return low + Math.random() * ( high - low );
//
//}
//
//// Random float from <-range/2, range/2> interval
//function randFloatSpread( range ) {
//
//	return range * ( 0.5 - Math.random() );
//
//}
//
//// Deterministic pseudo-random float in the interval [ 0, 1 ]
//function seededRandom( s ) {
//
//	if ( s !== undefined ) _seed = s;
//
//	// Mulberry32 generator
//
//	let t = _seed += 0x6D2B79F5;
//
//	t = Math.imul( t ^ t >>> 15, t | 1 );
//
//	t ^= t + Math.imul( t ^ t >>> 7, t | 61 );
//
//	return ( ( t ^ t >>> 14 ) >>> 0 ) / 4294967296;
//
//}
//
//function degToRad( degrees ) {
//
//	return degrees * DEG2RAD;
//
//}
//
//function radToDeg( radians ) {
//
//	return radians * RAD2DEG;
//
//}
//
//function isPowerOfTwo( value ) {
//
//	return ( value & ( value - 1 ) ) === 0 && value !== 0;
//
//}
//
//function ceilPowerOfTwo( value ) {
//
//	return Math.pow( 2, Math.ceil( Math.log( value ) / Math.LN2 ) );
//
//}
//
//function floorPowerOfTwo( value ) {
//
//	return Math.pow( 2, Math.floor( Math.log( value ) / Math.LN2 ) );
//
//}
//
//function setQuaternionFromProperEuler( q, a, b, c, order ) {
//
//	// Intrinsic Proper Euler Angles - see https://en.wikipedia.org/wiki/Euler_angles
//
//	// rotations are applied to the axes in the order specified by 'order'
//	// rotation by angle 'a' is applied first, then by angle 'b', then by angle 'c'
//	// angles are in radians
//
//	const cos = Math.cos;
//	const sin = Math.sin;
//
//	const c2 = cos( b / 2 );
//	const s2 = sin( b / 2 );
//
//	const c13 = cos( ( a + c ) / 2 );
//	const s13 = sin( ( a + c ) / 2 );
//
//	const c1_3 = cos( ( a - c ) / 2 );
//	const s1_3 = sin( ( a - c ) / 2 );
//
//	const c3_1 = cos( ( c - a ) / 2 );
//	const s3_1 = sin( ( c - a ) / 2 );
//
//	switch ( order ) {
//
//		case 'XYX':
//			q.set( c2 * s13, s2 * c1_3, s2 * s1_3, c2 * c13 );
//			break;
//
//		case 'YZY':
//			q.set( s2 * s1_3, c2 * s13, s2 * c1_3, c2 * c13 );
//			break;
//
//		case 'ZXZ':
//			q.set( s2 * c1_3, s2 * s1_3, c2 * s13, c2 * c13 );
//			break;
//
//		case 'XZX':
//			q.set( c2 * s13, s2 * s3_1, s2 * c3_1, c2 * c13 );
//			break;
//
//		case 'YXY':
//			q.set( s2 * c3_1, c2 * s13, s2 * s3_1, c2 * c13 );
//			break;
//
//		case 'ZYZ':
//			q.set( s2 * s3_1, s2 * c3_1, c2 * s13, c2 * c13 );
//			break;
//
//		default:
//			console.warn( 'THREE.MathUtils: .setQuaternionFromProperEuler() encountered an unknown order: ' + order );
//
//	}
//
//}
//
//function denormalize( value, array ) {
//
//	switch ( array.constructor ) {
//
//		case Float32Array:
//
//			return value;
//
//		case Uint32Array:
//
//			return value / 4294967295.0;
//
//		case Uint16Array:
//
//			return value / 65535.0;
//
//		case Uint8Array:
//
//			return value / 255.0;
//
//		case Int32Array:
//
//			return Math.max( value / 2147483647.0, - 1.0 );
//
//		case Int16Array:
//
//			return Math.max( value / 32767.0, - 1.0 );
//
//		case Int8Array:
//
//			return Math.max( value / 127.0, - 1.0 );
//
//		default:
//
//			throw new Error( 'Invalid component type.' );
//
//	}
//
//}
//
//function normalize( value, array ) {
//
//	switch ( array.constructor ) {
//
//		case Float32Array:
//
//			return value;
//
//		case Uint32Array:
//
//			return Math.round( value * 4294967295.0 );
//
//		case Uint16Array:
//
//			return Math.round( value * 65535.0 );
//
//		case Uint8Array:
//
//			return Math.round( value * 255.0 );
//
//		case Int32Array:
//
//			return Math.round( value * 2147483647.0 );
//
//		case Int16Array:
//
//			return Math.round( value * 32767.0 );
//
//		case Int8Array:
//
//			return Math.round( value * 127.0 );
//
//		default:
//
//			throw new Error( 'Invalid component type.' );
//
//	}
//
//}
