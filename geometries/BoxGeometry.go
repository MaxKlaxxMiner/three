package geometries

import (
	"github.com/MaxKlaxxMiner/three/core"
	"github.com/MaxKlaxxMiner/three/math"
	"github.com/MaxKlaxxMiner/three/utils"
)

type BoxGeometry struct {
	core.BufferGeometry
	parameters boxGeometryParameters
}

type boxGeometryParameters struct {
	width, height, depth                         float64
	widthSegments, heightSegments, depthSegments int
}

func NewBoxGeometry(width, height, depth float64) *BoxGeometry {
	return NewBoxGeometrySegments(width, height, depth, 1, 1, 1)
}

func NewBoxGeometrySegments(width, height, depth float64, widthSegments, heightSegments, depthSegments int) *BoxGeometry {
	this := new(BoxGeometry)
	this.BufferGeometry = *core.NewBufferGeometry()
	this.This = this
	this.Type = "BoxGeometry"

	this.parameters = boxGeometryParameters{
		width:          width,
		height:         height,
		depth:          depth,
		widthSegments:  widthSegments,
		heightSegments: heightSegments,
		depthSegments:  depthSegments,
	}

	// --- buffers ---
	indicesCount := 12 * (widthSegments*heightSegments + heightSegments*depthSegments + depthSegments*widthSegments)
	indices := make([]uint32, 0, indicesCount)
	verticesCount := 6 * (widthSegments*heightSegments + heightSegments*depthSegments + widthSegments*depthSegments + 2*(widthSegments+heightSegments+depthSegments) + 3)
	vertices := make([]float32, 0, verticesCount)
	normals := make([]float32, 0, verticesCount)
	uvs := make([]float32, 0, verticesCount/3*2)

	// --- helper variables ---
	numberOfVertices := 0
	groupStart := 0

	// --- build each side of the box geometry ---
	buildPlane := func(u, v, w int, udir, vdir, width, height, depth float64, gridX, gridY, materialIndex int) {
		segmentWidth := width / float64(gridX)
		segmentHeight := height / float64(gridY)

		widthHalf := width / 2
		heightHalf := height / 2
		depthHalf := depth / 2

		gridX1 := gridX + 1
		gridY1 := gridY + 1

		vertexCounter := 0
		groupCount := 0

		var vector math.Vector3

		// --- generate vertices, normals and uvs ---
		for iy := 0; iy < gridY1; iy++ {
			y := float64(iy)*segmentHeight - heightHalf
			for ix := 0; ix < gridX1; ix++ {
				x := float64(ix)*segmentWidth - widthHalf

				// --- set values to correct vector component ---
				vector.SetComponent(u, x*udir)
				vector.SetComponent(v, y*vdir)
				vector.SetComponent(w, depthHalf)

				// --- now apply vector to vertex buffer ---
				vertices = append(vertices, float32(vector.X), float32(vector.Y), float32(vector.Z))

				// --- set values to correct vector component ---
				vector.SetComponent(u, 0)
				vector.SetComponent(v, 0)
				vector.SetComponent(w, utils.If(depth > 0, 1.0, -1.0))

				// --- now apply vector to normal buffer ---
				normals = append(normals, float32(vector.X), float32(vector.Y), float32(vector.Z))

				// --- uvs ---
				uvs = append(uvs, float32(float64(ix)/float64(gridX)))
				uvs = append(uvs, float32(1.0-float64(iy)/float64(gridY)))

				// --- counters ---
				vertexCounter++
			}
		}

		// --- indices ---

		// 1. you need three indices to draw a single face
		// 2. a single segment consists of two faces
		// 3. so we need to generate six (2*3) indices per segment

		for iy := 0; iy < gridY; iy++ {
			for ix := 0; ix < gridX; ix++ {
				a := numberOfVertices + ix + gridX1*iy
				b := numberOfVertices + ix + gridX1*(iy+1)
				c := numberOfVertices + (ix + 1) + gridX1*(iy+1)
				d := numberOfVertices + (ix + 1) + gridX1*iy

				// --- faces ---
				indices = append(indices, uint32(a), uint32(b), uint32(d))
				indices = append(indices, uint32(b), uint32(c), uint32(d))

				// --- increase counter ---
				groupCount += 6
			}
		}

		// --- add a group to the geometry. this will ensure multi material support ---
		//			this.addGroup( groupStart, groupCount, materialIndex ); todo

		// --- calculate new start value for groups ---
		groupStart += groupCount

		// --- update total number of vertices ---
		numberOfVertices += vertexCounter
	}

	// --- build each side of the box geometry ---
	buildPlane('z', 'y', 'x', -1, -1, depth, height, width, depthSegments, heightSegments, 0)  // px
	buildPlane('z', 'y', 'x', 1, -1, depth, height, -width, depthSegments, heightSegments, 1)  // nx
	buildPlane('x', 'z', 'y', 1, 1, width, depth, height, widthSegments, depthSegments, 2)     // py
	buildPlane('x', 'z', 'y', 1, -1, width, depth, -height, widthSegments, depthSegments, 3)   // ny
	buildPlane('x', 'y', 'z', 1, -1, width, height, depth, widthSegments, heightSegments, 4)   // pz
	buildPlane('x', 'y', 'z', -1, -1, width, height, -depth, widthSegments, heightSegments, 5) // nz

	// --- build geometry ---
	this.SetIndexUint32(indices)
	this.SetAttribute("position", core.NewFloat32BufferAttribute(vertices, 3))
	this.SetAttribute("normal", core.NewFloat32BufferAttribute(normals, 3))
	this.SetAttribute("uv", core.NewFloat32BufferAttribute(uvs, 2))

	return this
}

func (c *BoxGeometry) IsBoxGeometry() bool { return c != nil }

//todo
// 	copy( source ) {
//
// 		super.copy( source );
//
// 		this.parameters = Object.assign( {}, source.parameters );
//
// 		return this;
//
// 	}
//
// 	static fromJSON( data ) {
//
// 		return new BoxGeometry( data.width, data.height, data.depth, data.widthSegments, data.heightSegments, data.depthSegments );
//
// 	}
// import { Float32BufferAttribute } from '../core/BufferAttribute.js';
