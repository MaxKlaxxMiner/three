package cameras

import "github.com/MaxKlaxxMiner/three/core"

type Camera struct {
	core.Object3D
}

func NewCamera() *Camera {
	this := new(Camera)
	this.Object3D = *core.NewObject3D()
	this.This = this
	this.Type = "Camera"
	// 		this.matrixWorldInverse = new Matrix4(); todo
	//
	// 		this.projectionMatrix = new Matrix4(); todo
	// 		this.projectionMatrixInverse = new Matrix4(); todo
	//
	// 		this.coordinateSystem = WebGLCoordinateSystem; todo
	return this
}

func (c *Camera) IsCamera() bool { return c != nil }

// 	copy( source, recursive ) {
//
// 		super.copy( source, recursive );
//
// 		this.matrixWorldInverse.copy( source.matrixWorldInverse );
//
// 		this.projectionMatrix.copy( source.projectionMatrix );
// 		this.projectionMatrixInverse.copy( source.projectionMatrixInverse );
//
// 		this.coordinateSystem = source.coordinateSystem;
//
// 		return this;
//
// 	}
//
// 	getWorldDirection( target ) {
//
// 		return super.getWorldDirection( target ).negate();
//
// 	}
//
// 	updateMatrixWorld( force ) {
//
// 		super.updateMatrixWorld( force );
//
// 		this.matrixWorldInverse.copy( this.matrixWorld ).invert();
//
// 	}
//
// 	updateWorldMatrix( updateParents, updateChildren ) {
//
// 		super.updateWorldMatrix( updateParents, updateChildren );
//
// 		this.matrixWorldInverse.copy( this.matrixWorld ).invert();
//
// 	}
//
// 	clone() {
//
// 		return new this.constructor().copy( this );
//
// 	}
