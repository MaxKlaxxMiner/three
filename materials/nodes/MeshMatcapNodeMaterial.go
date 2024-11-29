package nodes

// const _defaultValues = /*@__PURE__*/ new MeshMatcapMaterial();
//
// class MeshMatcapNodeMaterial extends NodeMaterial {
//
// 	static get type() {
//
// 		return 'MeshMatcapNodeMaterial';
//
// 	}
//
// 	constructor( parameters ) {
//
// 		super();
//
// 		this.lights = false;
//
// 		this.isMeshMatcapNodeMaterial = true;
//
// 		this.setDefaultValues( _defaultValues );
//
// 		this.setValues( parameters );
//
// 	}
//
// 	setupVariants( builder ) {
//
// 		const uv = matcapUV;
//
// 		let matcapColor;
//
// 		if ( builder.material.matcap ) {
//
// 			matcapColor = materialReference( 'matcap', 'texture' ).context( { getUV: () => uv } );
//
// 		} else {
//
// 			matcapColor = vec3( mix( 0.2, 0.8, uv.y ) ); // default if matcap is missing
//
// 		}
//
// 		diffuseColor.rgb.mulAssign( matcapColor.rgb );
//
// 	}
//
// }
