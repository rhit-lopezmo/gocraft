package terrain

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ChunkSizeX = 16
	ChunkSizeY = 64
	ChunkSizeZ = 16
)

type Chunk struct {
	BlockIds [][][]BlockType
	X, Z     int32
	Vertices []float32
	Normals  []float32
	UVs      []float32
	Indices  []uint16
}

func NewChunk(chunkX, chunkZ int32) *Chunk {
	blockIds := make([][][]BlockType, ChunkSizeX)

	for x := range blockIds {
		blockIds[x] = make([][]BlockType, ChunkSizeY)
		for y := range blockIds[x] {
			blockIds[x][y] = make([]BlockType, ChunkSizeZ)
		}
	}

	return &Chunk{
		BlockIds: blockIds,
		X:        chunkX,
		Z:        chunkZ,
	}
}

func (chunk *Chunk) SetBlock(x, y, z int32, blockType BlockType) {
	chunk.BlockIds[x][y][z] = blockType
}

func (chunk *Chunk) GetBlock(x, y, z int32) BlockType {
	return chunk.BlockIds[x][y][z]
}

func (chunk *Chunk) GenerateFlat(height int32) {
	for x := range int32(ChunkSizeX) {
		for z := range int32(ChunkSizeZ) {
			for y := range height {
				if y == height-1 {
					chunk.SetBlock(x, y, z, BlockGrass)
				} else {
					// TODO: add other cases once we handle more blocks
					chunk.SetBlock(x, y, z, BlockGrass)
				}
			}
		}
	}
}

func (chunk *Chunk) GenerateMesh() rl.Mesh {
	mesh := rl.Mesh{}

	// On the very first call, allocate backing arrays
	if chunk.Vertices == nil {
		chunk.Vertices = make([]float32, 0, 50000)
		chunk.Normals = make([]float32, 0, 50000)
		chunk.UVs = make([]float32, 0, 50000)
		chunk.Indices = make([]uint16, 0, 50000)
	} else {
		// Reuse same backing arrays – no reallocation
		chunk.Vertices = chunk.Vertices[:0]
		chunk.Normals = chunk.Normals[:0]
		chunk.UVs = chunk.UVs[:0]
		chunk.Indices = chunk.Indices[:0]
	}
	vertexOffset := uint16(0)

	for x := range chunk.BlockIds {
		for y := range chunk.BlockIds[x] {
			for z := range chunk.BlockIds[x][y] {
				if chunk.BlockIds[x][y][z] != BlockAir {
					chunk.gatherExposedFaces(x, y, z, &vertexOffset)
				}
			}
		}
	}

	// update mesh info
	if len(chunk.Vertices) > 0 {
		mesh.Vertices = &chunk.Vertices[0]
	}

	if len(chunk.Normals) > 0 {
		mesh.Normals = &chunk.Normals[0]
	}

	if len(chunk.UVs) > 0 {
		mesh.Texcoords = &chunk.UVs[0]
	}

	if len(chunk.Indices) > 0 {
		mesh.Indices = &chunk.Indices[0]
	}

	mesh.VertexCount = int32(len(chunk.Vertices) / 3)
	mesh.TriangleCount = int32(len(chunk.Indices) / 3)
	mesh.VaoID = 0
	mesh.VboID = nil

	rl.UploadMesh(&mesh, false)
	return mesh
}

func (chunk *Chunk) gatherExposedFaces(x, y, z int, vertexOffset *uint16) {
	worldX := float32(x)
	worldY := float32(y)
	worldZ := float32(z)

	for _, face := range FaceDirs {
		// if face != FaceLeft {
		// 	continue
		// }

		if chunk.isFaceExposed(x, y, z, face) {
			template := face.Template()

			// vertices
			for _, vertex := range template.Vertices {
				localX := vertex[0]
				localY := vertex[1]
				localZ := vertex[2]

				chunk.Vertices = append(chunk.Vertices, localX+worldX, localY+worldY, localZ+worldZ)
			}

			// normals
			for range 4 {
				chunk.Normals = append(
					chunk.Normals,
					template.Normals[0],
					template.Normals[1],
					template.Normals[2],
				)
			}

			// UVs
			for _, uv := range template.UVs {
				chunk.UVs = append(chunk.UVs, uv[0], uv[1])
			}

			// Indices
			for _, index := range template.Indices {
				chunk.Indices = append(chunk.Indices, index+*vertexOffset)
			}

			// increase vertex offset so future indices line up
			*vertexOffset += 4
		}
	}

}

func (chunk *Chunk) isFaceExposed(x, y, z int, faceDir FaceDir) bool {
	offsetX, offsetY, offsetZ := faceDir.Offset()

	neighborX, neighborY, neighborZ := x+offsetX, y+offsetY, z+offsetZ

	// check if neighbor is out of bounds
	if neighborX < 0 || neighborX >= ChunkSizeX {
		return true
	}

	if neighborY < 0 || neighborY >= ChunkSizeY {
		return true
	}

	if neighborZ < 0 || neighborZ >= ChunkSizeZ {
		return true
	}

	// check if neighbor block is air
	if chunk.BlockIds[neighborX][neighborY][neighborZ] == BlockAir {
		return true
	}

	// not exposed
	return false
}
