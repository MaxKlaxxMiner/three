package three

type Scene struct {
	Object3D
}

func NewScene() *Scene {
	s := new(Scene)
	//todo s.Object3D.Init()
	//		this.background = null;
	//		this.environment = null;
	//		this.fog = null;
	//		this.backgroundBlurriness = 0;
	//		this.backgroundIntensity = 1;
	//		this.backgroundRotation = new Euler();
	//		this.environmentIntensity = 1;
	//		this.environmentRotation = new Euler();
	//		this.overrideMaterial = null;
	return s
}

//class Scene extends Object3D {
//
//	copy( source, recursive ) {
//
//		super.copy( source, recursive );
//
//		if ( source.background !== null ) this.background = source.background.clone();
//		if ( source.environment !== null ) this.environment = source.environment.clone();
//		if ( source.fog !== null ) this.fog = source.fog.clone();
//
//		this.backgroundBlurriness = source.backgroundBlurriness;
//		this.backgroundIntensity = source.backgroundIntensity;
//		this.backgroundRotation.copy( source.backgroundRotation );
//
//		this.environmentIntensity = source.environmentIntensity;
//		this.environmentRotation.copy( source.environmentRotation );
//
//		if ( source.overrideMaterial !== null ) this.overrideMaterial = source.overrideMaterial.clone();
//
//		this.matrixAutoUpdate = source.matrixAutoUpdate;
//
//		return this;
//
//	}
//
//	toJSON( meta ) {
//
//		const data = super.toJSON( meta );
//
//		if ( this.fog !== null ) data.object.fog = this.fog.toJSON();
//
//		if ( this.backgroundBlurriness > 0 ) data.object.backgroundBlurriness = this.backgroundBlurriness;
//		if ( this.backgroundIntensity !== 1 ) data.object.backgroundIntensity = this.backgroundIntensity;
//		data.object.backgroundRotation = this.backgroundRotation.toArray();
//
//		if ( this.environmentIntensity !== 1 ) data.object.environmentIntensity = this.environmentIntensity;
//		data.object.environmentRotation = this.environmentRotation.toArray();
//
//		return data;
//
//	}
//
//}