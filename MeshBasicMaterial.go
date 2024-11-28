package three

import "github.com/MaxKlaxxMiner/three/util"

type MeshBasicMaterial struct {
	Material
	Color Color
}

type MeshBasicMaterialParameters struct {
	Color *Color
}

func NewMeshBasicMaterial(parameters *MeshBasicMaterialParameters) *MeshBasicMaterial {
	r := &MeshBasicMaterial{}
	r.Material = *NewMaterial()

	r.Color = util.NotNullOrDefault(parameters.Color, *NewColorHex(0xffffff))

	//		this.map = null;
	//
	//		this.lightMap = null;
	//		this.lightMapIntensity = 1.0;
	//
	//		this.aoMap = null;
	//		this.aoMapIntensity = 1.0;
	//
	//		this.specularMap = null;
	//
	//		this.alphaMap = null;
	//
	//		this.envMap = null;
	//		this.envMapRotation = new Euler();
	//		this.combine = MultiplyOperation;
	//		this.reflectivity = 1;
	//		this.refractionRatio = 0.98;
	//
	//		this.wireframe = false;
	//		this.wireframeLinewidth = 1;
	//		this.wireframeLinecap = 'round';
	//		this.wireframeLinejoin = 'round';
	//
	//		this.fog = true;
	//
	//		this.setValues( parameters );
	//	}
	//
	return r
}

//class MeshBasicMaterial extends Material {
//
//	copy( source ) {
//
//		super.copy( source );
//
//		this.color.copy( source.color );
//
//		this.map = source.map;
//
//		this.lightMap = source.lightMap;
//		this.lightMapIntensity = source.lightMapIntensity;
//
//		this.aoMap = source.aoMap;
//		this.aoMapIntensity = source.aoMapIntensity;
//
//		this.specularMap = source.specularMap;
//
//		this.alphaMap = source.alphaMap;
//
//		this.envMap = source.envMap;
//		this.envMapRotation.copy( source.envMapRotation );
//		this.combine = source.combine;
//		this.reflectivity = source.reflectivity;
//		this.refractionRatio = source.refractionRatio;
//
//		this.wireframe = source.wireframe;
//		this.wireframeLinewidth = source.wireframeLinewidth;
//		this.wireframeLinecap = source.wireframeLinecap;
//		this.wireframeLinejoin = source.wireframeLinejoin;
//
//		this.fog = source.fog;
//
//		return this;
//
//	}
//
//}
