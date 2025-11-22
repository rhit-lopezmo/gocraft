package terrain

const (
	ChunkSizeX = 16
	ChunkSizeY = 64
	ChunkSizeZ = 16
)

type Chunk struct {
	BlockIds [][][]BlockType
	X, Z     int32
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
					chunk.SetBlock(x, y, z, Grass)
				} else {
					// TODO: add other cases once we handle more blocks
					chunk.SetBlock(x, y, z, Grass)
				}
			}
		}
	}
}
