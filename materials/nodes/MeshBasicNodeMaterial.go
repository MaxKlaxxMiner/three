package nodes

// const _defaultValues = /*@__PURE__*/ new MeshBasicMaterial();
//
// class MeshBasicNodeMaterial extends NodeMaterial {
//
// 	static get type() {
//
// 		return 'MeshBasicNodeMaterial';
//
// 	}
//
// 	constructor( parameters ) {
//
// 		super();
//
// 		this.isMeshBasicNodeMaterial = true;
//
// 		this.lights = true;
//
// 		this.setDefaultValues( _defaultValues );
//
// 		this.setValues( parameters );
//
// 	}
//
// 	setupNormal() {
//
// 		return normalView; // see #28839
//
// 	}
//
// 	setupEnvironment( builder ) {
//
// 		const envNode = super.setupEnvironment( builder );
//
// 		return envNode ? new BasicEnvironmentNode( envNode ) : null;
//
// 	}
//
// 	setupLightMap( builder ) {
//
// 		let node = null;
//
// 		if ( builder.material.lightMap ) {
//
// 			node = new BasicLightMapNode( materialLightMap );
//
// 		}
//
// 		return node;
//
// 	}
//
// 	setupOutgoingLight() {
//
// 		return diffuseColor.rgb;
//
// 	}
//
// 	setupLightingModel() {
//
// 		return new BasicLightingModel();
//
// 	}
//
// }
