package scenes

type Scene struct {
	// Object3D todo
}

func NewScene() *Scene {
	this := new(Scene)
	// 		super(); // Object3D todo
	// 		this.type = 'Scene'; todo
	// 		this.background = null; todo
	// 		this.environment = null; todo
	// 		this.fog = null; todo
	// 		this.backgroundBlurriness = 0; todo
	// 		this.backgroundIntensity = 1; todo
	// 		this.backgroundRotation = new Euler(); todo
	// 		this.environmentIntensity = 1; todo
	// 		this.environmentRotation = new Euler(); todo
	// 		this.overrideMaterial = null; todo
	return this
}

func (s *Scene) IsScene() bool { return s != nil }

// todo
// 	copy( source, recursive ) {
//
// 		super.copy( source, recursive );
//
// 		if ( source.background !== null ) this.background = source.background.clone();
// 		if ( source.environment !== null ) this.environment = source.environment.clone();
// 		if ( source.fog !== null ) this.fog = source.fog.clone();
//
// 		this.backgroundBlurriness = source.backgroundBlurriness;
// 		this.backgroundIntensity = source.backgroundIntensity;
// 		this.backgroundRotation.copy( source.backgroundRotation );
//
// 		this.environmentIntensity = source.environmentIntensity;
// 		this.environmentRotation.copy( source.environmentRotation );
//
// 		if ( source.overrideMaterial !== null ) this.overrideMaterial = source.overrideMaterial.clone();
//
// 		this.matrixAutoUpdate = source.matrixAutoUpdate;
//
// 		return this;
//
// 	}
//
// 	toJSON( meta ) {
//
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
//
// 	}
//
// import { Object3D } from '../core/Object3D.js';
// import { Euler } from '../math/Euler.js';
