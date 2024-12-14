package scenes

import (
	"github.com/MaxKlaxxMiner/three/core"
	"github.com/MaxKlaxxMiner/three/math"
)

type Scene struct {
	core.Object3D
	BackgroundBlurriness float64
	BackgroundIntensity  float64
	BackgroundRotation   math.Euler
	EnvironmentIntensity float64
	EnvironmentRotation  math.Euler
}

func NewScene() *Scene {
	this := new(Scene)
	this.Object3D = *core.NewObject3D()
	this.This = this
	this.Type = "Scene"
	// 		this.background = null; todo
	// 		this.environment = null; todo
	// 		this.fog = null; todo
	this.BackgroundBlurriness = 0
	this.BackgroundIntensity = 1
	this.BackgroundRotation = *math.NewEulerDefaults()
	this.EnvironmentIntensity = 1
	this.EnvironmentRotation = *math.NewEulerDefaults()
	// 		this.overrideMaterial = null; todo
	return this
}

func (s *Scene) IsScene() bool { return s != nil }

func (s *Scene) Copy(source *Scene) *Scene {
	return s.CopyRecursive(source, true)
}

func (s *Scene) CopyRecursive(source *Scene, recursive bool) *Scene {
	s.Object3D.CopyRecursive(&source.Object3D, recursive)

	// 		if ( source.background !== null ) this.background = source.background.clone(); todo
	// 		if ( source.environment !== null ) this.environment = source.environment.clone(); todo
	// 		if ( source.fog !== null ) this.fog = source.fog.clone(); todo

	s.BackgroundBlurriness = source.BackgroundBlurriness
	s.BackgroundIntensity = source.BackgroundIntensity
	s.BackgroundRotation.Copy(&source.BackgroundRotation)

	s.EnvironmentIntensity = source.EnvironmentIntensity
	s.EnvironmentRotation.Copy(&source.EnvironmentRotation)

	//		if ( source.overrideMaterial !== null ) this.overrideMaterial = source.overrideMaterial.clone(); todo

	s.MatrixAutoUpdate = source.MatrixAutoUpdate

	return s
}

// todo
// 	toJSON( meta ) {
// 		const data = super.toJSON( meta );
//
// 		if ( this.fog !== null ) data.object.fog = this.fog.toJSON();
//
// 		if ( this.backgroundBlurriness > 0 ) data.object.backgroundBlurriness = this.backgroundBlurriness;
// 		if ( this.backgroundIntensity !== 1 ) data.object.backgroundIntensity = this.backgroundIntensity;
// 		data.object.backgroundRotation = this.backgroundRotation.toArray();
//
// 		if ( this.environmentIntensity !== 1 ) data.object.environmentIntensity = this.environmentIntensity;
// 		data.object.environmentRotation = this.environmentRotation.toArray();
//
// 		return data;
// 	}
