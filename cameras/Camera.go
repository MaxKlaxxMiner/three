package cameras

import (
	"github.com/MaxKlaxxMiner/three/core"
	"github.com/MaxKlaxxMiner/three/math"
)

type Camera struct {
	core.Object3D
	MatrixWorldInverse      math.Matrix4
	ProjectionMatrix        math.Matrix4
	ProjectionMatrixInverse math.Matrix4
	WebGLCoordinateSystem   bool
}

func NewCamera() *Camera {
	this := new(Camera)
	this.Object3D = *core.NewObject3D()
	this.This = this
	this.Type = "Camera"
	this.MatrixWorldInverse = *math.NewMatrix4Identity()

	this.ProjectionMatrix = *math.NewMatrix4Identity()
	this.ProjectionMatrixInverse = *math.NewMatrix4Identity()

	this.WebGLCoordinateSystem = true
	return this
}

func (c *Camera) IsCamera() bool { return c != nil }

func (c *Camera) Copy(source *Camera) *Camera {
	return c.CopyRecursive(source, true)
}

func (c *Camera) CopyRecursive(source *Camera, recursive bool) *Camera {
	c.Object3D.CopyRecursive(&source.Object3D, recursive)

	c.MatrixWorldInverse.Copy(&source.MatrixWorldInverse)

	c.ProjectionMatrix.Copy(&source.ProjectionMatrix)
	c.ProjectionMatrixInverse.Copy(&source.ProjectionMatrixInverse)

	c.WebGLCoordinateSystem = source.WebGLCoordinateSystem
	return c
}

func (c *Camera) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return c.Object3D.GetWorldDirection(target).Negate()
}

func (c *Camera) UpdateMatrixWorld() {
	c.UpdateMatrixWorldForce(false)
}

func (c *Camera) UpdateMatrixWorldForce(force bool) {
	c.Object3D.UpdateMatrixWorldForce(force)

	c.MatrixWorldInverse.Copy(&c.MatrixWorld).Invert()
}

func (c *Camera) UpdateWorldMatrix(updateParents, updateChildren bool) {
	c.Object3D.UpdateWorldMatrix(updateParents, updateChildren)

	c.MatrixWorldInverse.Copy(&c.MatrixWorld).Invert()
}

func (c *Camera) Clone() *Camera {
	return NewCamera().Copy(c)
}
