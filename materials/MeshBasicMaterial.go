package materials

import (
	"github.com/MaxKlaxxMiner/three/math"
	"github.com/MaxKlaxxMiner/three/utils"
)

type MeshBasicMaterial struct {
	Material
	Color math.Color
}

type MeshBasicMaterialParameters struct {
	Color *math.Color
}

func NewMeshBasicMaterial(parameters MeshBasicMaterialParameters) *MeshBasicMaterial {
	this := new(MeshBasicMaterial)
	this.Material = *NewMaterial()
	this.Type = "MeshBasicMaterial"

	this.Color = utils.NotNullOrDefault(parameters.Color, *math.NewColorHex(0xffffff)) // emissive

	//todo
	// 		this.map = null;
	//
	// 		this.lightMap = null;
	// 		this.lightMapIntensity = 1.0;
	//
	// 		this.aoMap = null;
	// 		this.aoMapIntensity = 1.0;
	//
	// 		this.specularMap = null;
	//
	// 		this.alphaMap = null;
	//
	// 		this.envMap = null;
	// 		this.envMapRotation = new Euler();
	// 		this.combine = MultiplyOperation;
	// 		this.reflectivity = 1;
	// 		this.refractionRatio = 0.98;
	//
	// 		this.wireframe = false;
	// 		this.wireframeLinewidth = 1;
	// 		this.wireframeLinecap = 'round';
	// 		this.wireframeLinejoin = 'round';
	//
	// 		this.fog = true;

	return this
}

func (m *MeshBasicMaterial) IsMeshBasicMaterial() bool { return m != nil }

// 	copy( source ) {
//
// 		super.copy( source );
//
// 		this.color.copy( source.color );
//
// 		this.map = source.map;
//
// 		this.lightMap = source.lightMap;
// 		this.lightMapIntensity = source.lightMapIntensity;
//
// 		this.aoMap = source.aoMap;
// 		this.aoMapIntensity = source.aoMapIntensity;
//
// 		this.specularMap = source.specularMap;
//
// 		this.alphaMap = source.alphaMap;
//
// 		this.envMap = source.envMap;
// 		this.envMapRotation.copy( source.envMapRotation );
// 		this.combine = source.combine;
// 		this.reflectivity = source.reflectivity;
// 		this.refractionRatio = source.refractionRatio;
//
// 		this.wireframe = source.wireframe;
// 		this.wireframeLinewidth = source.wireframeLinewidth;
// 		this.wireframeLinecap = source.wireframeLinecap;
// 		this.wireframeLinejoin = source.wireframeLinejoin;
//
// 		this.fog = source.fog;
//
// 		return this;
//
// 	}
// import { Euler } from '../math/Euler.js';
