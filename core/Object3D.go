package core

import "github.com/MaxKlaxxMiner/three/math"

type Object3D struct {
	EventDispatcher
	Id                     int
	Uuid                   math.UUID
	Name                   string
	Type                   string
	Parent                 *Object3D
	Children               []*Object3D
	Up                     math.Vector3
	Position               math.Vector3
	Rotation               math.Euler
	Quaternion             math.Quaternion
	ModelViewMatrix        math.Matrix4
	NormalMatrix           math.Matrix3
	Scale                  math.Vector3
	Matrix                 math.Matrix4
	MatrixWorld            math.Matrix4
	MatrixAutoUpdate       bool
	MatrixWorldAutoUpdate  bool
	MatrixWorldNeedsUpdate bool
	Layers                 Layers
	Visible                bool
	CastShadow             bool
	ReceiveShadow          bool
	FrustumCulled          bool
	RenderOrder            int

	OnBeforeShadow func() // renderer, object, camera, shadowCamera, geometry, depthMaterial, group
	OnAfterShadow  func() // renderer, object, camera, shadowCamera, geometry, depthMaterial, group
	OnBeforeRender func() // renderer, scene, camera, geometry, material, group
	OnAfterRender  func() // renderer, scene, camera, geometry, material, group
}

var _object3DId = 0

func NewObject3D() *Object3D {
	this := new(Object3D)
	this.This = this
	this.Id = _object3DId
	_object3DId++
	this.Uuid = math.GenerateUUID()
	this.Name = ""
	this.Type = "Object3D"

	this.Parent = nil
	this.Children = nil

	this.Up.Copy(Object3dDefaultUp)
	this.Position = math.Vector3{}
	this.Rotation = *math.NewEuler(0, 0, 0)
	this.Quaternion = *math.NewQuaternionDefaults()
	this.Scale.SetScalar(1)

	this.Rotation.OverrideOnChange(func() { this.Quaternion.SetFromEulerUpdate(&this.Rotation, false) })
	this.Quaternion.OverrideOnChange(func() { this.Rotation.SetFromQuaternion(&this.Quaternion, this.Rotation.GetOrder(), false) })

	this.ModelViewMatrix = *math.NewMatrix4Identity()
	this.NormalMatrix = *math.NewMatrix3Identity()

	this.Matrix = *math.NewMatrix4Identity()
	this.MatrixWorld = *math.NewMatrix4Identity()

	this.MatrixAutoUpdate = Object3dDefaultMatrixAutoUpdate

	this.MatrixWorldAutoUpdate = Object3dDefaultMatrixWorldAutoUpdate // checked by the renderer
	this.MatrixWorldNeedsUpdate = false

	this.Layers = *NewLayers()
	this.Visible = true

	this.CastShadow = false
	this.ReceiveShadow = false

	this.FrustumCulled = true
	this.RenderOrder = 0

	// 		this.animations = []; todo
	// 		this.userData = {}; todo

	this.OnBeforeShadow = func() {}
	this.OnAfterShadow = func() {}
	this.OnBeforeRender = func() {}
	this.OnAfterRender = func() {}

	return this
}

func (o *Object3D) IsObject3D() bool { return o != nil }

//todo
// 	applyMatrix4( matrix ) {
//
// 		if ( this.matrixAutoUpdate ) this.updateMatrix();
//
// 		this.matrix.premultiply( matrix );
//
// 		this.matrix.decompose( this.position, this.quaternion, this.scale );
//
// 	}
//
// 	applyQuaternion( q ) {
//
// 		this.quaternion.premultiply( q );
//
// 		return this;
//
// 	}
//
// 	setRotationFromAxisAngle( axis, angle ) {
//
// 		// assumes axis is normalized
//
// 		this.quaternion.setFromAxisAngle( axis, angle );
//
// 	}
//
// 	setRotationFromEuler( euler ) {
//
// 		this.quaternion.setFromEuler( euler, true );
//
// 	}
//
// 	setRotationFromMatrix( m ) {
//
// 		// assumes the upper 3x3 of m is a pure rotation matrix (i.e, unscaled)
//
// 		this.quaternion.setFromRotationMatrix( m );
//
// 	}
//
// 	setRotationFromQuaternion( q ) {
//
// 		// assumes q is normalized
//
// 		this.quaternion.copy( q );
//
// 	}
//
// 	rotateOnAxis( axis, angle ) {
//
// 		// rotate object on axis in object space
// 		// axis is assumed to be normalized
//
// 		_q1.setFromAxisAngle( axis, angle );
//
// 		this.quaternion.multiply( _q1 );
//
// 		return this;
//
// 	}
//
// 	rotateOnWorldAxis( axis, angle ) {
//
// 		// rotate object on axis in world space
// 		// axis is assumed to be normalized
// 		// method assumes no rotated parent
//
// 		_q1.setFromAxisAngle( axis, angle );
//
// 		this.quaternion.premultiply( _q1 );
//
// 		return this;
//
// 	}
//
// 	rotateX( angle ) {
//
// 		return this.rotateOnAxis( _xAxis, angle );
//
// 	}
//
// 	rotateY( angle ) {
//
// 		return this.rotateOnAxis( _yAxis, angle );
//
// 	}
//
// 	rotateZ( angle ) {
//
// 		return this.rotateOnAxis( _zAxis, angle );
//
// 	}
//
// 	translateOnAxis( axis, distance ) {
//
// 		// translate object by distance along axis in object space
// 		// axis is assumed to be normalized
//
// 		_v1.copy( axis ).applyQuaternion( this.quaternion );
//
// 		this.position.add( _v1.multiplyScalar( distance ) );
//
// 		return this;
//
// 	}
//
// 	translateX( distance ) {
//
// 		return this.translateOnAxis( _xAxis, distance );
//
// 	}
//
// 	translateY( distance ) {
//
// 		return this.translateOnAxis( _yAxis, distance );
//
// 	}
//
// 	translateZ( distance ) {
//
// 		return this.translateOnAxis( _zAxis, distance );
//
// 	}
//
// 	localToWorld( vector ) {
//
// 		this.updateWorldMatrix( true, false );
//
// 		return vector.applyMatrix4( this.matrixWorld );
//
// 	}
//
// 	worldToLocal( vector ) {
//
// 		this.updateWorldMatrix( true, false );
//
// 		return vector.applyMatrix4( _m1.copy( this.matrixWorld ).invert() );
//
// 	}
//
// 	lookAt( x, y, z ) {
//
// 		// This method does not support objects having non-uniformly-scaled parent(s)
//
// 		if ( x.isVector3 ) {
//
// 			_target.copy( x );
//
// 		} else {
//
// 			_target.set( x, y, z );
//
// 		}
//
// 		const parent = this.parent;
//
// 		this.updateWorldMatrix( true, false );
//
// 		_position.setFromMatrixPosition( this.matrixWorld );
//
// 		if ( this.isCamera || this.isLight ) {
//
// 			_m1.lookAt( _position, _target, this.up );
//
// 		} else {
//
// 			_m1.lookAt( _target, _position, this.up );
//
// 		}
//
// 		this.quaternion.setFromRotationMatrix( _m1 );
//
// 		if ( parent ) {
//
// 			_m1.extractRotation( parent.matrixWorld );
// 			_q1.setFromRotationMatrix( _m1 );
// 			this.quaternion.premultiply( _q1.invert() );
//
// 		}
//
// 	}

func (o *Object3D) Add(objects ...*Object3D) *Object3D {
	for _, obj := range objects {
		if obj == o {
			panic("THREE.Object3D.add: object can't be added as a child of itself.")
		}
		if !obj.IsObject3D() {
			panic("THREE.Object3D.add: object not an instance of THREE.Object3D.")
		}
		// object.removeFromParent(); todo
		obj.Parent = o
		o.Children = append(o.Children, obj)
		// object.dispatchEvent( _addedEvent ); todo
		// _childaddedEvent.child = object; todo
		// this.dispatchEvent( _childaddedEvent ); todo
		// _childaddedEvent.child = null; todo
	}
	return o
}

//todo
// 	remove( object ) {
//
// 		if ( arguments.length > 1 ) {
//
// 			for ( let i = 0; i < arguments.length; i ++ ) {
//
// 				this.remove( arguments[ i ] );
//
// 			}
//
// 			return this;
//
// 		}
//
// 		const index = this.children.indexOf( object );
//
// 		if ( index !== - 1 ) {
//
// 			object.parent = null;
// 			this.children.splice( index, 1 );
//
// 			object.dispatchEvent( _removedEvent );
//
// 			_childremovedEvent.child = object;
// 			this.dispatchEvent( _childremovedEvent );
// 			_childremovedEvent.child = null;
//
// 		}
//
// 		return this;
//
// 	}
//
// 	removeFromParent() {
//
// 		const parent = this.parent;
//
// 		if ( parent !== null ) {
//
// 			parent.remove( this );
//
// 		}
//
// 		return this;
//
// 	}
//
// 	clear() {
//
// 		return this.remove( ... this.children );
//
// 	}
//
// 	attach( object ) {
//
// 		// adds object as a child of this, while maintaining the object's world transform
//
// 		// Note: This method does not support scene graphs having non-uniformly-scaled nodes(s)
//
// 		this.updateWorldMatrix( true, false );
//
// 		_m1.copy( this.matrixWorld ).invert();
//
// 		if ( object.parent !== null ) {
//
// 			object.parent.updateWorldMatrix( true, false );
//
// 			_m1.multiply( object.parent.matrixWorld );
//
// 		}
//
// 		object.applyMatrix4( _m1 );
//
// 		object.removeFromParent();
// 		object.parent = this;
// 		this.children.push( object );
//
// 		object.updateWorldMatrix( false, true );
//
// 		object.dispatchEvent( _addedEvent );
//
// 		_childaddedEvent.child = object;
// 		this.dispatchEvent( _childaddedEvent );
// 		_childaddedEvent.child = null;
//
// 		return this;
//
// 	}
//
// 	getObjectById( id ) {
//
// 		return this.getObjectByProperty( 'id', id );
//
// 	}
//
// 	getObjectByName( name ) {
//
// 		return this.getObjectByProperty( 'name', name );
//
// 	}
//
// 	getObjectByProperty( name, value ) {
//
// 		if ( this[ name ] === value ) return this;
//
// 		for ( let i = 0, l = this.children.length; i < l; i ++ ) {
//
// 			const child = this.children[ i ];
// 			const object = child.getObjectByProperty( name, value );
//
// 			if ( object !== undefined ) {
//
// 				return object;
//
// 			}
//
// 		}
//
// 		return undefined;
//
// 	}
//
// 	getObjectsByProperty( name, value, result = [] ) {
//
// 		if ( this[ name ] === value ) result.push( this );
//
// 		const children = this.children;
//
// 		for ( let i = 0, l = children.length; i < l; i ++ ) {
//
// 			children[ i ].getObjectsByProperty( name, value, result );
//
// 		}
//
// 		return result;
//
// 	}
//
// 	getWorldPosition( target ) {
//
// 		this.updateWorldMatrix( true, false );
//
// 		return target.setFromMatrixPosition( this.matrixWorld );
//
// 	}
//
// 	getWorldQuaternion( target ) {
//
// 		this.updateWorldMatrix( true, false );
//
// 		this.matrixWorld.decompose( _position, target, _scale );
//
// 		return target;
//
// 	}
//
// 	getWorldScale( target ) {
//
// 		this.updateWorldMatrix( true, false );
//
// 		this.matrixWorld.decompose( _position, _quaternion, target );
//
// 		return target;
//
// 	}
//
// 	getWorldDirection( target ) {
//
// 		this.updateWorldMatrix( true, false );
//
// 		const e = this.matrixWorld.elements;
//
// 		return target.set( e[ 8 ], e[ 9 ], e[ 10 ] ).normalize();
//
// 	}
//
// 	raycast( /* raycaster, intersects */ ) {}
//
// 	traverse( callback ) {
//
// 		callback( this );
//
// 		const children = this.children;
//
// 		for ( let i = 0, l = children.length; i < l; i ++ ) {
//
// 			children[ i ].traverse( callback );
//
// 		}
//
// 	}
//
// 	traverseVisible( callback ) {
//
// 		if ( this.visible === false ) return;
//
// 		callback( this );
//
// 		const children = this.children;
//
// 		for ( let i = 0, l = children.length; i < l; i ++ ) {
//
// 			children[ i ].traverseVisible( callback );
//
// 		}
//
// 	}
//
// 	traverseAncestors( callback ) {
//
// 		const parent = this.parent;
//
// 		if ( parent !== null ) {
//
// 			callback( parent );
//
// 			parent.traverseAncestors( callback );
//
// 		}
//
// 	}

func (o *Object3D) UpdateMatrix() {
	o.Matrix.Compose(&o.Position, &o.Quaternion, &o.Scale)
	o.MatrixWorldNeedsUpdate = true
}

func (o *Object3D) UpdateMatrixWorld() {
	o.UpdateMatrixWorldForce(false)
}

func (o *Object3D) UpdateMatrixWorldForce(force bool) {
	if o.MatrixAutoUpdate {
		o.UpdateMatrix()
	}

	if o.MatrixWorldNeedsUpdate || force {
		if o.MatrixWorldAutoUpdate {
			if o.Parent == nil {
				o.MatrixWorld.Copy(&o.Matrix)
			} else {
				o.MatrixWorld.MultiplyMatrices(&o.Parent.MatrixWorld, &o.Matrix)
			}
		}
		o.MatrixWorldNeedsUpdate = false
		force = true
	}

	// --- make sure descendants are updated if required ---

	for _, child := range o.Children {
		child.UpdateMatrixWorldForce(force)
	}
}

//todo
// 	updateWorldMatrix( updateParents, updateChildren ) {
//
// 		const parent = this.parent;
//
// 		if ( updateParents === true && parent !== null ) {
//
// 			parent.updateWorldMatrix( true, false );
//
// 		}
//
// 		if ( this.matrixAutoUpdate ) this.updateMatrix();
//
// 		if ( this.matrixWorldAutoUpdate === true ) {
//
// 			if ( this.parent === null ) {
//
// 				this.matrixWorld.copy( this.matrix );
//
// 			} else {
//
// 				this.matrixWorld.multiplyMatrices( this.parent.matrixWorld, this.matrix );
//
// 			}
//
// 		}
//
// 		// make sure descendants are updated
//
// 		if ( updateChildren === true ) {
//
// 			const children = this.children;
//
// 			for ( let i = 0, l = children.length; i < l; i ++ ) {
//
// 				const child = children[ i ];
//
// 				child.updateWorldMatrix( false, true );
//
// 			}
//
// 		}
//
// 	}
//
// 	toJSON( meta ) {
//
// 		// meta is a string when called from JSON.stringify
// 		const isRootObject = ( meta === undefined || typeof meta === 'string' );
//
// 		const output = {};
//
// 		// meta is a hash used to collect geometries, materials.
// 		// not providing it implies that this is the root object
// 		// being serialized.
// 		if ( isRootObject ) {
//
// 			// initialize meta obj
// 			meta = {
// 				geometries: {},
// 				materials: {},
// 				textures: {},
// 				images: {},
// 				shapes: {},
// 				skeletons: {},
// 				animations: {},
// 				nodes: {}
// 			};
//
// 			output.metadata = {
// 				version: 4.6,
// 				type: 'Object',
// 				generator: 'Object3D.toJSON'
// 			};
//
// 		}
//
// 		// standard Object3D serialization
//
// 		const object = {};
//
// 		object.uuid = this.uuid;
// 		object.type = this.type;
//
// 		if ( this.name !== '' ) object.name = this.name;
// 		if ( this.castShadow === true ) object.castShadow = true;
// 		if ( this.receiveShadow === true ) object.receiveShadow = true;
// 		if ( this.visible === false ) object.visible = false;
// 		if ( this.frustumCulled === false ) object.frustumCulled = false;
// 		if ( this.renderOrder !== 0 ) object.renderOrder = this.renderOrder;
// 		if ( Object.keys( this.userData ).length > 0 ) object.userData = this.userData;
//
// 		object.layers = this.layers.mask;
// 		object.matrix = this.matrix.toArray();
// 		object.up = this.up.toArray();
//
// 		if ( this.matrixAutoUpdate === false ) object.matrixAutoUpdate = false;
//
// 		// object specific properties
//
// 		if ( this.isInstancedMesh ) {
//
// 			object.type = 'InstancedMesh';
// 			object.count = this.count;
// 			object.instanceMatrix = this.instanceMatrix.toJSON();
// 			if ( this.instanceColor !== null ) object.instanceColor = this.instanceColor.toJSON();
//
// 		}
//
// 		if ( this.isBatchedMesh ) {
//
// 			object.type = 'BatchedMesh';
// 			object.perObjectFrustumCulled = this.perObjectFrustumCulled;
// 			object.sortObjects = this.sortObjects;
//
// 			object.drawRanges = this._drawRanges;
// 			object.reservedRanges = this._reservedRanges;
//
// 			object.visibility = this._visibility;
// 			object.active = this._active;
// 			object.bounds = this._bounds.map( bound => ( {
// 				boxInitialized: bound.boxInitialized,
// 				boxMin: bound.box.min.toArray(),
// 				boxMax: bound.box.max.toArray(),
//
// 				sphereInitialized: bound.sphereInitialized,
// 				sphereRadius: bound.sphere.radius,
// 				sphereCenter: bound.sphere.center.toArray()
// 			} ) );
//
// 			object.maxInstanceCount = this._maxInstanceCount;
// 			object.maxVertexCount = this._maxVertexCount;
// 			object.maxIndexCount = this._maxIndexCount;
//
// 			object.geometryInitialized = this._geometryInitialized;
// 			object.geometryCount = this._geometryCount;
//
// 			object.matricesTexture = this._matricesTexture.toJSON( meta );
//
// 			if ( this._colorsTexture !== null ) object.colorsTexture = this._colorsTexture.toJSON( meta );
//
// 			if ( this.boundingSphere !== null ) {
//
// 				object.boundingSphere = {
// 					center: object.boundingSphere.center.toArray(),
// 					radius: object.boundingSphere.radius
// 				};
//
// 			}
//
// 			if ( this.boundingBox !== null ) {
//
// 				object.boundingBox = {
// 					min: object.boundingBox.min.toArray(),
// 					max: object.boundingBox.max.toArray()
// 				};
//
// 			}
//
// 		}
//
// 		//
//
// 		function serialize( library, element ) {
//
// 			if ( library[ element.uuid ] === undefined ) {
//
// 				library[ element.uuid ] = element.toJSON( meta );
//
// 			}
//
// 			return element.uuid;
//
// 		}
//
// 		if ( this.isScene ) {
//
// 			if ( this.background ) {
//
// 				if ( this.background.isColor ) {
//
// 					object.background = this.background.toJSON();
//
// 				} else if ( this.background.isTexture ) {
//
// 					object.background = this.background.toJSON( meta ).uuid;
//
// 				}
//
// 			}
//
// 			if ( this.environment && this.environment.isTexture && this.environment.isRenderTargetTexture !== true ) {
//
// 				object.environment = this.environment.toJSON( meta ).uuid;
//
// 			}
//
// 		} else if ( this.isMesh || this.isLine || this.isPoints ) {
//
// 			object.geometry = serialize( meta.geometries, this.geometry );
//
// 			const parameters = this.geometry.parameters;
//
// 			if ( parameters !== undefined && parameters.shapes !== undefined ) {
//
// 				const shapes = parameters.shapes;
//
// 				if ( Array.isArray( shapes ) ) {
//
// 					for ( let i = 0, l = shapes.length; i < l; i ++ ) {
//
// 						const shape = shapes[ i ];
//
// 						serialize( meta.shapes, shape );
//
// 					}
//
// 				} else {
//
// 					serialize( meta.shapes, shapes );
//
// 				}
//
// 			}
//
// 		}
//
// 		if ( this.isSkinnedMesh ) {
//
// 			object.bindMode = this.bindMode;
// 			object.bindMatrix = this.bindMatrix.toArray();
//
// 			if ( this.skeleton !== undefined ) {
//
// 				serialize( meta.skeletons, this.skeleton );
//
// 				object.skeleton = this.skeleton.uuid;
//
// 			}
//
// 		}
//
// 		if ( this.material !== undefined ) {
//
// 			if ( Array.isArray( this.material ) ) {
//
// 				const uuids = [];
//
// 				for ( let i = 0, l = this.material.length; i < l; i ++ ) {
//
// 					uuids.push( serialize( meta.materials, this.material[ i ] ) );
//
// 				}
//
// 				object.material = uuids;
//
// 			} else {
//
// 				object.material = serialize( meta.materials, this.material );
//
// 			}
//
// 		}
//
// 		//
//
// 		if ( this.children.length > 0 ) {
//
// 			object.children = [];
//
// 			for ( let i = 0; i < this.children.length; i ++ ) {
//
// 				object.children.push( this.children[ i ].toJSON( meta ).object );
//
// 			}
//
// 		}
//
// 		//
//
// 		if ( this.animations.length > 0 ) {
//
// 			object.animations = [];
//
// 			for ( let i = 0; i < this.animations.length; i ++ ) {
//
// 				const animation = this.animations[ i ];
//
// 				object.animations.push( serialize( meta.animations, animation ) );
//
// 			}
//
// 		}
//
// 		if ( isRootObject ) {
//
// 			const geometries = extractFromCache( meta.geometries );
// 			const materials = extractFromCache( meta.materials );
// 			const textures = extractFromCache( meta.textures );
// 			const images = extractFromCache( meta.images );
// 			const shapes = extractFromCache( meta.shapes );
// 			const skeletons = extractFromCache( meta.skeletons );
// 			const animations = extractFromCache( meta.animations );
// 			const nodes = extractFromCache( meta.nodes );
//
// 			if ( geometries.length > 0 ) output.geometries = geometries;
// 			if ( materials.length > 0 ) output.materials = materials;
// 			if ( textures.length > 0 ) output.textures = textures;
// 			if ( images.length > 0 ) output.images = images;
// 			if ( shapes.length > 0 ) output.shapes = shapes;
// 			if ( skeletons.length > 0 ) output.skeletons = skeletons;
// 			if ( animations.length > 0 ) output.animations = animations;
// 			if ( nodes.length > 0 ) output.nodes = nodes;
//
// 		}
//
// 		output.object = object;
//
// 		return output;
//
// 		// extract data from the cache hash
// 		// remove metadata on each item
// 		// and return as array
// 		function extractFromCache( cache ) {
//
// 			const values = [];
// 			for ( const key in cache ) {
//
// 				const data = cache[ key ];
// 				delete data.metadata;
// 				values.push( data );
//
// 			}
//
// 			return values;
//
// 		}
//
// 	}

func (o *Object3D) Clone() *Object3D {
	return NewObject3D().Copy(o)
}

func (o *Object3D) CloneRecursive(recursive bool) *Object3D {
	return NewObject3D().CopyRecursive(o, recursive)
}

func (o *Object3D) Copy(source *Object3D) *Object3D {
	return o.CopyRecursive(source, true)
}

func (o *Object3D) CopyRecursive(source *Object3D, recursive bool) *Object3D {
	o.Name = source.Name

	o.Up.Copy(&source.Up)

	o.Position.Copy(&source.Position)
	o.Rotation.SetOrderNoUpdate(source.Rotation.GetOrder())
	o.Quaternion.Copy(&source.Quaternion)
	o.Scale.Copy(&source.Scale)

	o.Matrix.Copy(&source.Matrix)
	o.MatrixWorld.Copy(&source.MatrixWorld)

	o.MatrixAutoUpdate = source.MatrixAutoUpdate

	o.MatrixWorldAutoUpdate = source.MatrixWorldAutoUpdate
	o.MatrixWorldNeedsUpdate = source.MatrixWorldNeedsUpdate

	o.Layers.Mask = source.Layers.Mask
	o.Visible = source.Visible

	o.CastShadow = source.CastShadow
	o.ReceiveShadow = source.ReceiveShadow

	o.FrustumCulled = source.FrustumCulled
	o.RenderOrder = source.RenderOrder

	// 		this.animations = source.animations.slice(); todo

	// 		this.userData = JSON.parse( JSON.stringify( source.userData ) ); todo

	if recursive {
		for i := range source.Children {
			o.Add(source.Children[i].Clone())
		}
	}

	return o
}

var Object3dDefaultUp = math.NewVector3(0, 1, 0)
var Object3dDefaultMatrixAutoUpdate = true
var Object3dDefaultMatrixWorldAutoUpdate = true

//todo
// const _v1 = /*@__PURE__*/ new Vector3();
// const _q1 = /*@__PURE__*/ new Quaternion();
// const _m1 = /*@__PURE__*/ new Matrix4();
// const _target = /*@__PURE__*/ new Vector3();
//
// const _position = /*@__PURE__*/ new Vector3();
// const _scale = /*@__PURE__*/ new Vector3();
// const _quaternion = /*@__PURE__*/ new Quaternion();
//
// const _xAxis = /*@__PURE__*/ new Vector3( 1, 0, 0 );
// const _yAxis = /*@__PURE__*/ new Vector3( 0, 1, 0 );
// const _zAxis = /*@__PURE__*/ new Vector3( 0, 0, 1 );
//
// const _addedEvent = { type: 'added' };
// const _removedEvent = { type: 'removed' };
//
// const _childaddedEvent = { type: 'childadded', child: null };
// const _childremovedEvent = { type: 'childremoved', child: null };
