package math

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

//todo
// 	makeRotationFromEuler( euler ) {
//
// 		const te = this.elements;
//
// 		const x = euler.x, y = euler.y, z = euler.z;
// 		const a = Math.cos( x ), b = Math.sin( x );
// 		const c = Math.cos( y ), d = Math.sin( y );
// 		const e = Math.cos( z ), f = Math.sin( z );
//
// 		if ( euler.order === 'XYZ' ) {
//
// 			const ae = a * e, af = a * f, be = b * e, bf = b * f;
//
// 			te[ 0 ] = c * e;
// 			te[ 4 ] = - c * f;
// 			te[ 8 ] = d;
//
// 			te[ 1 ] = af + be * d;
// 			te[ 5 ] = ae - bf * d;
// 			te[ 9 ] = - b * c;
//
// 			te[ 2 ] = bf - ae * d;
// 			te[ 6 ] = be + af * d;
// 			te[ 10 ] = a * c;
//
// 		} else if ( euler.order === 'YXZ' ) {
//
// 			const ce = c * e, cf = c * f, de = d * e, df = d * f;
//
// 			te[ 0 ] = ce + df * b;
// 			te[ 4 ] = de * b - cf;
// 			te[ 8 ] = a * d;
//
// 			te[ 1 ] = a * f;
// 			te[ 5 ] = a * e;
// 			te[ 9 ] = - b;
//
// 			te[ 2 ] = cf * b - de;
// 			te[ 6 ] = df + ce * b;
// 			te[ 10 ] = a * c;
//
// 		} else if ( euler.order === 'ZXY' ) {
//
// 			const ce = c * e, cf = c * f, de = d * e, df = d * f;
//
// 			te[ 0 ] = ce - df * b;
// 			te[ 4 ] = - a * f;
// 			te[ 8 ] = de + cf * b;
//
// 			te[ 1 ] = cf + de * b;
// 			te[ 5 ] = a * e;
// 			te[ 9 ] = df - ce * b;
//
// 			te[ 2 ] = - a * d;
// 			te[ 6 ] = b;
// 			te[ 10 ] = a * c;
//
// 		} else if ( euler.order === 'ZYX' ) {
//
// 			const ae = a * e, af = a * f, be = b * e, bf = b * f;
//
// 			te[ 0 ] = c * e;
// 			te[ 4 ] = be * d - af;
// 			te[ 8 ] = ae * d + bf;
//
// 			te[ 1 ] = c * f;
// 			te[ 5 ] = bf * d + ae;
// 			te[ 9 ] = af * d - be;
//
// 			te[ 2 ] = - d;
// 			te[ 6 ] = b * c;
// 			te[ 10 ] = a * c;
//
// 		} else if ( euler.order === 'YZX' ) {
//
// 			const ac = a * c, ad = a * d, bc = b * c, bd = b * d;
//
// 			te[ 0 ] = c * e;
// 			te[ 4 ] = bd - ac * f;
// 			te[ 8 ] = bc * f + ad;
//
// 			te[ 1 ] = f;
// 			te[ 5 ] = a * e;
// 			te[ 9 ] = - b * e;
//
// 			te[ 2 ] = - d * e;
// 			te[ 6 ] = ad * f + bc;
// 			te[ 10 ] = ac - bd * f;
//
// 		} else if ( euler.order === 'XZY' ) {
//
// 			const ac = a * c, ad = a * d, bc = b * c, bd = b * d;
//
// 			te[ 0 ] = c * e;
// 			te[ 4 ] = - f;
// 			te[ 8 ] = d * e;
//
// 			te[ 1 ] = ac * f + bd;
// 			te[ 5 ] = a * e;
// 			te[ 9 ] = ad * f - bc;
//
// 			te[ 2 ] = bc * f - ad;
// 			te[ 6 ] = b * e;
// 			te[ 10 ] = bd * f + ac;
//
// 		}
//
// 		// bottom row
// 		te[ 3 ] = 0;
// 		te[ 7 ] = 0;
// 		te[ 11 ] = 0;
//
// 		// last column
// 		te[ 12 ] = 0;
// 		te[ 13 ] = 0;
// 		te[ 14 ] = 0;
// 		te[ 15 ] = 1;
//
// 		return this;
//
// 	}

func (m *Matrix4) MakeRotationFromQuaternion(q *Quaternion) *Matrix4 {
	return m.Compose(_zeroMatrix4, q, _oneMatrix4)
}

// 	lookAt( eye, target, up ) {
//
// 		const te = this.elements;
//
// 		_z.subVectors( eye, target );
//
// 		if ( _z.lengthSq() === 0 ) {
//
// 			// eye and target are in the same position
//
// 			_z.z = 1;
//
// 		}
//
// 		_z.normalize();
// 		_x.crossVectors( up, _z );
//
// 		if ( _x.lengthSq() === 0 ) {
//
// 			// up and z are parallel
//
// 			if ( Math.abs( up.z ) === 1 ) {
//
// 				_z.x += 0.0001;
//
// 			} else {
//
// 				_z.z += 0.0001;
//
// 			}
//
// 			_z.normalize();
// 			_x.crossVectors( up, _z );
//
// 		}
//
// 		_x.normalize();
// 		_y.crossVectors( _z, _x );
//
// 		te[ 0 ] = _x.x; te[ 4 ] = _y.x; te[ 8 ] = _z.x;
// 		te[ 1 ] = _x.y; te[ 5 ] = _y.y; te[ 9 ] = _z.y;
// 		te[ 2 ] = _x.z; te[ 6 ] = _y.z; te[ 10 ] = _z.z;
//
// 		return this;
//
// 	}
//
// 	multiply( m ) {
//
// 		return this.multiplyMatrices( this, m );
//
// 	}
//
// 	premultiply( m ) {
//
// 		return this.multiplyMatrices( m, this );
//
// 	}

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

//todo
// 	multiplyScalar( s ) {
//
// 		const te = this.elements;
//
// 		te[ 0 ] *= s; te[ 4 ] *= s; te[ 8 ] *= s; te[ 12 ] *= s;
// 		te[ 1 ] *= s; te[ 5 ] *= s; te[ 9 ] *= s; te[ 13 ] *= s;
// 		te[ 2 ] *= s; te[ 6 ] *= s; te[ 10 ] *= s; te[ 14 ] *= s;
// 		te[ 3 ] *= s; te[ 7 ] *= s; te[ 11 ] *= s; te[ 15 ] *= s;
//
// 		return this;
//
// 	}
//
// 	determinant() {
//
// 		const te = this.elements;
//
// 		const n11 = te[ 0 ], n12 = te[ 4 ], n13 = te[ 8 ], n14 = te[ 12 ];
// 		const n21 = te[ 1 ], n22 = te[ 5 ], n23 = te[ 9 ], n24 = te[ 13 ];
// 		const n31 = te[ 2 ], n32 = te[ 6 ], n33 = te[ 10 ], n34 = te[ 14 ];
// 		const n41 = te[ 3 ], n42 = te[ 7 ], n43 = te[ 11 ], n44 = te[ 15 ];
//
// 		//TOO: make this more efficient
// 		//( based on http://www.euclideanspace.com/maths/algebra/matrix/functions/inverse/fourD/index.htm )
//
// 		return (
// 			n41 * (
// 				+ n14 * n23 * n32
// 				 - n13 * n24 * n32
// 				 - n14 * n22 * n33
// 				 + n12 * n24 * n33
// 				 + n13 * n22 * n34
// 				 - n12 * n23 * n34
// 			) +
// 			n42 * (
// 				+ n11 * n23 * n34
// 				 - n11 * n24 * n33
// 				 + n14 * n21 * n33
// 				 - n13 * n21 * n34
// 				 + n13 * n24 * n31
// 				 - n14 * n23 * n31
// 			) +
// 			n43 * (
// 				+ n11 * n24 * n32
// 				 - n11 * n22 * n34
// 				 - n14 * n21 * n32
// 				 + n12 * n21 * n34
// 				 + n14 * n22 * n31
// 				 - n12 * n24 * n31
// 			) +
// 			n44 * (
// 				- n13 * n22 * n31
// 				 - n11 * n23 * n32
// 				 + n11 * n22 * n33
// 				 + n13 * n21 * n32
// 				 - n12 * n21 * n33
// 				 + n12 * n23 * n31
// 			)
//
// 		);
//
// 	}
//
// 	transpose() {
//
// 		const te = this.elements;
// 		let tmp;
//
// 		tmp = te[ 1 ]; te[ 1 ] = te[ 4 ]; te[ 4 ] = tmp;
// 		tmp = te[ 2 ]; te[ 2 ] = te[ 8 ]; te[ 8 ] = tmp;
// 		tmp = te[ 6 ]; te[ 6 ] = te[ 9 ]; te[ 9 ] = tmp;
//
// 		tmp = te[ 3 ]; te[ 3 ] = te[ 12 ]; te[ 12 ] = tmp;
// 		tmp = te[ 7 ]; te[ 7 ] = te[ 13 ]; te[ 13 ] = tmp;
// 		tmp = te[ 11 ]; te[ 11 ] = te[ 14 ]; te[ 14 ] = tmp;
//
// 		return this;
//
// 	}
//
// 	setPosition( x, y, z ) {
//
// 		const te = this.elements;
//
// 		if ( x.isVector3 ) {
//
// 			te[ 12 ] = x.x;
// 			te[ 13 ] = x.y;
// 			te[ 14 ] = x.z;
//
// 		} else {
//
// 			te[ 12 ] = x;
// 			te[ 13 ] = y;
// 			te[ 14 ] = z;
//
// 		}
//
// 		return this;
//
// 	}

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

// todo
// scale( v ) {
//
//	const te = this.elements;
//	const x = v.x, y = v.y, z = v.z;
//
//	te[ 0 ] *= x; te[ 4 ] *= y; te[ 8 ] *= z;
//	te[ 1 ] *= x; te[ 5 ] *= y; te[ 9 ] *= z;
//	te[ 2 ] *= x; te[ 6 ] *= y; te[ 10 ] *= z;
//	te[ 3 ] *= x; te[ 7 ] *= y; te[ 11 ] *= z;
//
//	return this;
//
// }
//
// getMaxScaleOnAxis() {
//
//	const te = this.elements;
//
//	const scaleXSq = te[ 0 ] * te[ 0 ] + te[ 1 ] * te[ 1 ] + te[ 2 ] * te[ 2 ];
//	const scaleYSq = te[ 4 ] * te[ 4 ] + te[ 5 ] * te[ 5 ] + te[ 6 ] * te[ 6 ];
//	const scaleZSq = te[ 8 ] * te[ 8 ] + te[ 9 ] * te[ 9 ] + te[ 10 ] * te[ 10 ];
//
//	return Math.sqrt( Math.max( scaleXSq, scaleYSq, scaleZSq ) );
//
// }
//
// makeTranslation( x, y, z ) {
//
//	if ( x.isVector3 ) {
//
//		this.set(
//
//			1, 0, 0, x.x,
//			0, 1, 0, x.y,
//			0, 0, 1, x.z,
//			0, 0, 0, 1
//
//		);
//
//	} else {
//
//		this.set(
//
//			1, 0, 0, x,
//			0, 1, 0, y,
//			0, 0, 1, z,
//			0, 0, 0, 1
//
//		);
//
//	}
//
//	return this;
//
// }
//
// makeRotationX( theta ) {
//
//	const c = Math.cos( theta ), s = Math.sin( theta );
//
//	this.set(
//
//		1, 0, 0, 0,
//		0, c, - s, 0,
//		0, s, c, 0,
//		0, 0, 0, 1
//
//	);
//
//	return this;
//
// }
//
// makeRotationY( theta ) {
//
//	const c = Math.cos( theta ), s = Math.sin( theta );
//
//	this.set(
//
//		 c, 0, s, 0,
//		 0, 1, 0, 0,
//		- s, 0, c, 0,
//		 0, 0, 0, 1
//
//	);
//
//	return this;
//
// }
//
// makeRotationZ( theta ) {
//
//	const c = Math.cos( theta ), s = Math.sin( theta );
//
//	this.set(
//
//		c, - s, 0, 0,
//		s, c, 0, 0,
//		0, 0, 1, 0,
//		0, 0, 0, 1
//
//	);
//
//	return this;
//
// }
//
// makeRotationAxis( axis, angle ) {
//
//	// Based on http://www.gamedev.net/reference/articles/article1199.asp
//
//	const c = Math.cos( angle );
//	const s = Math.sin( angle );
//	const t = 1 - c;
//	const x = axis.x, y = axis.y, z = axis.z;
//	const tx = t * x, ty = t * y;
//
//	this.set(
//
//		tx * x + c, tx * y - s * z, tx * z + s * y, 0,
//		tx * y + s * z, ty * y + c, ty * z - s * x, 0,
//		tx * z - s * y, ty * z + s * x, t * z * z + c, 0,
//		0, 0, 0, 1
//
//	);
//
//	return this;
//
// }
//
// makeScale( x, y, z ) {
//
//	this.set(
//
//		x, 0, 0, 0,
//		0, y, 0, 0,
//		0, 0, z, 0,
//		0, 0, 0, 1
//
//	);
//
//	return this;
//
// }
//
// makeShear( xy, xz, yx, yz, zx, zy ) {
//
//	this.set(
//
//		1, yx, zx, 0,
//		xy, 1, zy, 0,
//		xz, yz, 1, 0,
//		0, 0, 0, 1
//
//	);
//
//	return this;
//
// }

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

// todo
// decompose( position, quaternion, scale ) {
//
//	const te = this.elements;
//
//	let sx = _v1.set( te[ 0 ], te[ 1 ], te[ 2 ] ).length();
//	const sy = _v1.set( te[ 4 ], te[ 5 ], te[ 6 ] ).length();
//	const sz = _v1.set( te[ 8 ], te[ 9 ], te[ 10 ] ).length();
//
//	// if determine is negative, we need to invert one scale
//	const det = this.determinant();
//	if ( det < 0 ) sx = - sx;
//
//	position.x = te[ 12 ];
//	position.y = te[ 13 ];
//	position.z = te[ 14 ];
//
//	// scale the rotation part
//	_m1.copy( this );
//
//	const invSX = 1 / sx;
//	const invSY = 1 / sy;
//	const invSZ = 1 / sz;
//
//	_m1.elements[ 0 ] *= invSX;
//	_m1.elements[ 1 ] *= invSX;
//	_m1.elements[ 2 ] *= invSX;
//
//	_m1.elements[ 4 ] *= invSY;
//	_m1.elements[ 5 ] *= invSY;
//	_m1.elements[ 6 ] *= invSY;
//
//	_m1.elements[ 8 ] *= invSZ;
//	_m1.elements[ 9 ] *= invSZ;
//	_m1.elements[ 10 ] *= invSZ;
//
//	quaternion.setFromRotationMatrix( _m1 );
//
//	scale.x = sx;
//	scale.y = sy;
//	scale.z = sz;
//
//	return this;
//
// }
//
// makePerspective( left, right, top, bottom, near, far, coordinateSystem = WebGLCoordinateSystem ) {
//
//	const te = this.elements;
//	const x = 2 * near / ( right - left );
//	const y = 2 * near / ( top - bottom );
//
//	const a = ( right + left ) / ( right - left );
//	const b = ( top + bottom ) / ( top - bottom );
//
//	let c, d;
//
//	if ( coordinateSystem === WebGLCoordinateSystem ) {
//
//		c = - ( far + near ) / ( far - near );
//		d = ( - 2 * far * near ) / ( far - near );
//
//	} else if ( coordinateSystem === WebGPUCoordinateSystem ) {
//
//		c = - far / ( far - near );
//		d = ( - far * near ) / ( far - near );
//
//	} else {
//
//		throw new Error( 'THREE.Matrix4.makePerspective(): Invalid coordinate system: ' + coordinateSystem );
//
//	}
//
//	te[ 0 ] = x;	te[ 4 ] = 0;	te[ 8 ] = a; 	te[ 12 ] = 0;
//	te[ 1 ] = 0;	te[ 5 ] = y;	te[ 9 ] = b; 	te[ 13 ] = 0;
//	te[ 2 ] = 0;	te[ 6 ] = 0;	te[ 10 ] = c; 	te[ 14 ] = d;
//	te[ 3 ] = 0;	te[ 7 ] = 0;	te[ 11 ] = - 1;	te[ 15 ] = 0;
//
//	return this;
//
// }
//
// makeOrthographic( left, right, top, bottom, near, far, coordinateSystem = WebGLCoordinateSystem ) {
//
//	const te = this.elements;
//	const w = 1.0 / ( right - left );
//	const h = 1.0 / ( top - bottom );
//	const p = 1.0 / ( far - near );
//
//	const x = ( right + left ) * w;
//	const y = ( top + bottom ) * h;
//
//	let z, zInv;
//
//	if ( coordinateSystem === WebGLCoordinateSystem ) {
//
//		z = ( far + near ) * p;
//		zInv = - 2 * p;
//
//	} else if ( coordinateSystem === WebGPUCoordinateSystem ) {
//
//		z = near * p;
//		zInv = - 1 * p;
//
//	} else {
//
//		throw new Error( 'THREE.Matrix4.makeOrthographic(): Invalid coordinate system: ' + coordinateSystem );
//
//	}
//
//	te[ 0 ] = 2 * w;	te[ 4 ] = 0;		te[ 8 ] = 0; 		te[ 12 ] = - x;
//	te[ 1 ] = 0; 		te[ 5 ] = 2 * h;	te[ 9 ] = 0; 		te[ 13 ] = - y;
//	te[ 2 ] = 0; 		te[ 6 ] = 0;		te[ 10 ] = zInv;	te[ 14 ] = - z;
//	te[ 3 ] = 0; 		te[ 7 ] = 0;		te[ 11 ] = 0;		te[ 15 ] = 1;
//
//	return this;
//
// }
//
// equals( matrix ) {
//
//	const te = this.elements;
//	const me = matrix.elements;
//
//	for ( let i = 0; i < 16; i ++ ) {
//
//		if ( te[ i ] !== me[ i ] ) return false;
//
//	}
//
//	return true;
//
// }
//
// fromArray( array, offset = 0 ) {
//
//	for ( let i = 0; i < 16; i ++ ) {
//
//		this.elements[ i ] = array[ i + offset ];
//
//	}
//
//	return this;
//
// }
//
// toArray( array = [], offset = 0 ) {
//
//	const te = this.elements;
//
//	array[ offset ] = te[ 0 ];
//	array[ offset + 1 ] = te[ 1 ];
//	array[ offset + 2 ] = te[ 2 ];
//	array[ offset + 3 ] = te[ 3 ];
//
//	array[ offset + 4 ] = te[ 4 ];
//	array[ offset + 5 ] = te[ 5 ];
//	array[ offset + 6 ] = te[ 6 ];
//	array[ offset + 7 ] = te[ 7 ];
//
//	array[ offset + 8 ] = te[ 8 ];
//	array[ offset + 9 ] = te[ 9 ];
//	array[ offset + 10 ] = te[ 10 ];
//	array[ offset + 11 ] = te[ 11 ];
//
//	array[ offset + 12 ] = te[ 12 ];
//	array[ offset + 13 ] = te[ 13 ];
//	array[ offset + 14 ] = te[ 14 ];
//	array[ offset + 15 ] = te[ 15 ];
//
//	return array;
//
// }

var _v1Matrix4 = NewVector3Defaults()
var _m1Matrix4 = NewMatrix4Identity()
var _zeroMatrix4 = NewVector3Defaults()
var _oneMatrix4 = NewVector3(1, 1, 1)
var _xMatrix4 = NewVector3Defaults()
var _yMatrix4 = NewVector3Defaults()
var _zMatrix4 = NewVector3Defaults()
