package three

type Camera struct {
	Object3D
	MatrixWorldInverse      Matrix4
	ProjectionMatrix        Matrix4
	ProjectionMatrixInverse Matrix4
	// coordinateSystem = WebGLCoordinateSystem hardcoded
}

func NewCamera() *Camera {
	c := &Camera{}
	c.Init()
	return c
}

func (c *Camera) Init() {
	c.MatrixWorldInverse.Identity()
	c.ProjectionMatrix.Identity()
	c.ProjectionMatrixInverse.Identity()
}

func (c *Camera) Copy(source *Camera, recursive bool) *Camera {
	//todo: c.Object3D.copy

	c.MatrixWorldInverse.Copy(&source.MatrixWorldInverse)

	c.ProjectionMatrix.Copy(&source.ProjectionMatrix)
	c.ProjectionMatrixInverse.Copy(&source.ProjectionMatrixInverse)

	return c
}

// todo
//	getWorldDirection( target ) {
//
//		return super.getWorldDirection( target ).negate();
//
//	}
//
//	updateMatrixWorld( force ) {
//
//		super.updateMatrixWorld( force );
//
//		this.matrixWorldInverse.copy( this.matrixWorld ).invert();
//
//	}
//
//	updateWorldMatrix( updateParents, updateChildren ) {
//
//		super.updateWorldMatrix( updateParents, updateChildren );
//
//		this.matrixWorldInverse.copy( this.matrixWorld ).invert();
//
//	}
//
//	clone() {
//
//		return new this.constructor().copy( this );
//
//	}
//
//}
