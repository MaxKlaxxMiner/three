package three

import (
	"github.com/MaxKlaxxMiner/three/cameras"
	"github.com/MaxKlaxxMiner/three/core"
	"github.com/MaxKlaxxMiner/three/geometries"
	"github.com/MaxKlaxxMiner/three/materials"
	"github.com/MaxKlaxxMiner/three/math"
	"github.com/MaxKlaxxMiner/three/objects"
	"github.com/MaxKlaxxMiner/three/renderers"
	"github.com/MaxKlaxxMiner/three/renderers/webgl"
	"github.com/MaxKlaxxMiner/three/scenes"
)

var (
	NewWebGLRendererDefaults = renderers.NewWebGLRendererDefaults
	NewWebGLRenderer         = renderers.NewWebGLRenderer
	// export { ShaderLib } from './renderers/shaders/ShaderLib.js';
	// export { UniformsLib } from './renderers/shaders/UniformsLib.js';
	// export { UniformsUtils } from './renderers/shaders/UniformsUtils.js';
	// export { ShaderChunk } from './renderers/shaders/ShaderChunk.js';
	// export { PMREMGenerator } from './extras/PMREMGenerator.js';
	NewWebGLUtils = webgl.NewWebGLUtils
	// export { WebGLArrayRenderTarget } from './renderers/WebGLArrayRenderTarget.js';
	// export { WebGL3DRenderTarget } from './renderers/WebGL3DRenderTarget.js';
	// export { WebGLCubeRenderTarget } from './renderers/WebGLCubeRenderTarget.js';
	// export { WebGLRenderTarget } from './renderers/WebGLRenderTarget.js';
	// export { FogExp2 } from './scenes/FogExp2.js';
	// export { Fog } from './scenes/Fog.js';
	NewScene = scenes.NewScene
	// export { Sprite } from './objects/Sprite.js';
	// export { LOD } from './objects/LOD.js';
	// export { SkinnedMesh } from './objects/SkinnedMesh.js';
	// export { Skeleton } from './objects/Skeleton.js';
	// export { Bone } from './objects/Bone.js';
	NewMesh = objects.NewMesh
	// export { InstancedMesh } from './objects/InstancedMesh.js';
	// export { BatchedMesh } from './objects/BatchedMesh.js';
	// export { LineSegments } from './objects/LineSegments.js';
	// export { LineLoop } from './objects/LineLoop.js';
	// export { Line } from './objects/Line.js';
	// export { Points } from './objects/Points.js';
	// export { Group } from './objects/Group.js';
	// export { VideoTexture } from './textures/VideoTexture.js';
	// export { FramebufferTexture } from './textures/FramebufferTexture.js';
	// export { Source } from './textures/Source.js';
	// export { DataTexture } from './textures/DataTexture.js';
	// export { DataArrayTexture } from './textures/DataArrayTexture.js';
	// export { Data3DTexture } from './textures/Data3DTexture.js';
	// export { CompressedTexture } from './textures/CompressedTexture.js';
	// export { CompressedArrayTexture } from './textures/CompressedArrayTexture.js';
	// export { CompressedCubeTexture } from './textures/CompressedCubeTexture.js';
	// export { CubeTexture } from './textures/CubeTexture.js';
	// export { CanvasTexture } from './textures/CanvasTexture.js';
	// export { DepthTexture } from './textures/DepthTexture.js';
	// export { Texture } from './textures/Texture.js';
	NewBoxGeometry         = geometries.NewBoxGeometry
	NewBoxGeometrySegments = geometries.NewBoxGeometrySegments
	// export * from './geometries/CapsuleGeometry.js';
	// export * from './geometries/CircleGeometry.js';
	// export * from './geometries/ConeGeometry.js';
	// export * from './geometries/CylinderGeometry.js';
	// export * from './geometries/DodecahedronGeometry.js';
	// export * from './geometries/EdgesGeometry.js';
	// export * from './geometries/ExtrudeGeometry.js';
	// export * from './geometries/IcosahedronGeometry.js';
	// export * from './geometries/LatheGeometry.js';
	// export * from './geometries/OctahedronGeometry.js';
	// export * from './geometries/PlaneGeometry.js';
	// export * from './geometries/PolyhedronGeometry.js';
	// export * from './geometries/RingGeometry.js';
	// export * from './geometries/ShapeGeometry.js';
	// export * from './geometries/SphereGeometry.js';
	// export * from './geometries/TetrahedronGeometry.js';
	// export * from './geometries/TorusGeometry.js';
	// export * from './geometries/TorusKnotGeometry.js';
	// export * from './geometries/TubeGeometry.js';
	// export * from './geometries/WireframeGeometry.js';
	// export ShadowMaterial,
	// export SpriteMaterial,
	// export RawShaderMaterial,
	// export ShaderMaterial,
	// export PointsMaterial,
	// export MeshPhysicalMaterial,
	// export MeshStandardMaterial,
	// export MeshPhongMaterial,
	// export MeshToonMaterial,
	// export MeshNormalMaterial,
	// export MeshLambertMaterial,
	// export MeshDepthMaterial,
	// export MeshDistanceMaterial,
	NewMeshBasicMaterial = materials.NewMeshBasicMaterial
	// export MeshMatcapMaterial,
	// export LineDashedMaterial,
	// export LineBasicMaterial,
	NewMaterial = materials.NewMaterial
	// export { AnimationLoader } from './loaders/AnimationLoader.js';
	// export { CompressedTextureLoader } from './loaders/CompressedTextureLoader.js';
	// export { CubeTextureLoader } from './loaders/CubeTextureLoader.js';
	// export { DataTextureLoader } from './loaders/DataTextureLoader.js';
	// export { TextureLoader } from './loaders/TextureLoader.js';
	// export { ObjectLoader } from './loaders/ObjectLoader.js';
	// export { MaterialLoader } from './loaders/MaterialLoader.js';
	// export { BufferGeometryLoader } from './loaders/BufferGeometryLoader.js';
	// export { DefaultLoadingManager, LoadingManager } from './loaders/LoadingManager.js';
	// export { ImageLoader } from './loaders/ImageLoader.js';
	// export { ImageBitmapLoader } from './loaders/ImageBitmapLoader.js';
	// export { FileLoader } from './loaders/FileLoader.js';
	// export { Loader } from './loaders/Loader.js';
	// export { LoaderUtils } from './loaders/LoaderUtils.js';
	// export { Cache } from './loaders/Cache.js';
	// export { AudioLoader } from './loaders/AudioLoader.js';
	// export { SpotLight } from './lights/SpotLight.js';
	// export { PointLight } from './lights/PointLight.js';
	// export { RectAreaLight } from './lights/RectAreaLight.js';
	// export { HemisphereLight } from './lights/HemisphereLight.js';
	// export { DirectionalLight } from './lights/DirectionalLight.js';
	// export { AmbientLight } from './lights/AmbientLight.js';
	// export { Light } from './lights/Light.js';
	// export { LightProbe } from './lights/LightProbe.js';
	// export { StereoCamera } from './cameras/StereoCamera.js';
	NewPerspectiveCameraDefaults = cameras.NewPerspectiveCameraDefaults
	NewPerspectiveCamera         = cameras.NewPerspectiveCamera
	// export { OrthographicCamera } from './cameras/OrthographicCamera.js';
	// export { CubeCamera } from './cameras/CubeCamera.js';
	// export { ArrayCamera } from './cameras/ArrayCamera.js';
	NewCamera = cameras.NewCamera
	// export { AudioListener } from './audio/AudioListener.js';
	// export { PositionalAudio } from './audio/PositionalAudio.js';
	// export { AudioContext } from './audio/AudioContext.js';
	// export { AudioAnalyser } from './audio/AudioAnalyser.js';
	// export { Audio } from './audio/Audio.js';
	// export { VectorKeyframeTrack } from './animation/tracks/VectorKeyframeTrack.js';
	// export { StringKeyframeTrack } from './animation/tracks/StringKeyframeTrack.js';
	// export { QuaternionKeyframeTrack } from './animation/tracks/QuaternionKeyframeTrack.js';
	// export { NumberKeyframeTrack } from './animation/tracks/NumberKeyframeTrack.js';
	// export { ColorKeyframeTrack } from './animation/tracks/ColorKeyframeTrack.js';
	// export { BooleanKeyframeTrack } from './animation/tracks/BooleanKeyframeTrack.js';
	// export { PropertyMixer } from './animation/PropertyMixer.js';
	// export { PropertyBinding } from './animation/PropertyBinding.js';
	// export { KeyframeTrack } from './animation/KeyframeTrack.js';
	// export { AnimationUtils } from './animation/AnimationUtils.js';
	// export { AnimationObjectGroup } from './animation/AnimationObjectGroup.js';
	// export { AnimationMixer } from './animation/AnimationMixer.js';
	// export { AnimationClip } from './animation/AnimationClip.js';
	// export { AnimationAction } from './animation/AnimationAction.js';
	// export { RenderTarget } from './core/RenderTarget.js';
	// export { Uniform } from './core/Uniform.js';
	// export { UniformsGroup } from './core/UniformsGroup.js';
	// export { InstancedBufferGeometry } from './core/InstancedBufferGeometry.js';
	NewBufferGeometry = core.NewBufferGeometry
	// export { InterleavedBufferAttribute } from './core/InterleavedBufferAttribute.js';
	// export { InstancedInterleavedBuffer } from './core/InstancedInterleavedBuffer.js';
	// export { InterleavedBuffer } from './core/InterleavedBuffer.js';
	// export { InstancedBufferAttribute } from './core/InstancedBufferAttribute.js';
	// export { GLBufferAttribute } from './core/GLBufferAttribute.js';
	// export * from './core/BufferAttribute.js';
	NewObject3D = core.NewObject3D
	// export { Raycaster } from './core/Raycaster.js';
	// export { Layers } from './core/Layers.js';
	NewEventDispatcher = core.NewEventDispatcher
	// export { Clock } from './core/Clock.js';
	// export { QuaternionLinearInterpolant } from './math/interpolants/QuaternionLinearInterpolant.js';
	// export { LinearInterpolant } from './math/interpolants/LinearInterpolant.js';
	// export { DiscreteInterpolant } from './math/interpolants/DiscreteInterpolant.js';
	// export { CubicInterpolant } from './math/interpolants/CubicInterpolant.js';
	// export { Interpolant } from './math/Interpolant.js';
	// export { Triangle } from './math/Triangle.js';
	// export { MathUtils } from './math/MathUtils.js';
	// export { Spherical } from './math/Spherical.js';
	// export { Cylindrical } from './math/Cylindrical.js';
	// export { Plane } from './math/Plane.js';
	// export { Frustum } from './math/Frustum.js';
	// export { Sphere } from './math/Sphere.js';
	// export { Ray } from './math/Ray.js';
	// export { Matrix4 } from './math/Matrix4.js';
	// export { Matrix3 } from './math/Matrix3.js';
	// export { Matrix2 } from './math/Matrix2.js';
	// export { Box3 } from './math/Box3.js';
	// export { Box2 } from './math/Box2.js';
	// export { Line3 } from './math/Line3.js';
	NewEuler           = math.NewEuler
	NewEulerWithOrder  = math.NewEulerWithOrder
	NewVector4Defaults = math.NewVector4Defaults
	NewVector4         = math.NewVector4
	NewVector3Defaults = math.NewVector3Defaults
	NewVector3         = math.NewVector3
	// export { Vector2 } from './math/Vector2.js';
	NewQuaternionDefaults = math.NewQuaternionDefaults
	NewQuaternion         = math.NewQuaternion
	NewColor              = math.NewColor
	NewColorHex           = math.NewColorHex
	// export { ColorManagement } from './math/ColorManagement.js';
	// export { SphericalHarmonics3 } from './math/SphericalHarmonics3.js';
	// export { SpotLightHelper } from './helpers/SpotLightHelper.js';
	// export { SkeletonHelper } from './helpers/SkeletonHelper.js';
	// export { PointLightHelper } from './helpers/PointLightHelper.js';
	// export { HemisphereLightHelper } from './helpers/HemisphereLightHelper.js';
	// export { GridHelper } from './helpers/GridHelper.js';
	// export { PolarGridHelper } from './helpers/PolarGridHelper.js';
	// export { DirectionalLightHelper } from './helpers/DirectionalLightHelper.js';
	// export { CameraHelper } from './helpers/CameraHelper.js';
	// export { BoxHelper } from './helpers/BoxHelper.js';
	// export { Box3Helper } from './helpers/Box3Helper.js';
	// export { PlaneHelper } from './helpers/PlaneHelper.js';
	// export { ArrowHelper } from './helpers/ArrowHelper.js';
	// export { AxesHelper } from './helpers/AxesHelper.js';
	// export * from './extras/curves/Curves.js';
	// export { Shape } from './extras/core/Shape.js';
	// export { Path } from './extras/core/Path.js';
	// export { ShapePath } from './extras/core/ShapePath.js';
	// export { CurvePath } from './extras/core/CurvePath.js';
	// export { Curve } from './extras/core/Curve.js';
	// export { Controls } from './extras/Controls.js';
	// export { DataUtils } from './extras/DataUtils.js';
	// export { ImageUtils } from './extras/ImageUtils.js';
	// export { ShapeUtils } from './extras/ShapeUtils.js';
	// export { TextureUtils } from './extras/TextureUtils.js';
	// export { createCanvasElement } from './utils.js';
)
