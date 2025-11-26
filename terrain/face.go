package terrain

type FaceDir int

const (
	FaceRight FaceDir = iota
	FaceLeft
	FaceTop
	FaceBottom
	FaceFront
	FaceBack
)

var FaceDirs = []FaceDir{
	FaceRight,
	FaceLeft,
	FaceTop,
	FaceBottom,
	FaceFront,
	FaceBack,
}

var templateIndices = [6]uint16{
	0, 1, 2,
	2, 3, 0,
}

var templateUVs = [4][2]float32{
	{0, 0},
	{1, 0},
	{1, 1},
	{0, 1},
}

type UVRegion struct {
	MinU, MinV float32
	MaxU, MaxV float32
}

var oneThird = float32(1) / float32(3)
var twoThirds = float32(2) / float32(3)

var grassTopRegion = UVRegion{0, 0, oneThird, 1}
var grassSideRegion = UVRegion{oneThird, 0, twoThirds, 1}
var grassBottomRegion = UVRegion{twoThirds, 0, 1, 1}

var grassSideUVs = [4][2]float32{
	{grassSideRegion.MinU, grassSideRegion.MinV},
	{grassSideRegion.MaxU, grassSideRegion.MinV},
	{grassSideRegion.MaxU, grassSideRegion.MaxV},
	{grassSideRegion.MinU, grassSideRegion.MaxV},
}

var grassTopUVs = [4][2]float32{
	{grassTopRegion.MinU, grassTopRegion.MinV},
	{grassTopRegion.MaxU, grassTopRegion.MinV},
	{grassTopRegion.MaxU, grassTopRegion.MaxV},
	{grassTopRegion.MinU, grassTopRegion.MaxV},
}

var grassBottomUVs = [4][2]float32{
	{grassBottomRegion.MinU, grassBottomRegion.MinV},
	{grassBottomRegion.MaxU, grassBottomRegion.MinV},
	{grassBottomRegion.MaxU, grassBottomRegion.MaxV},
	{grassBottomRegion.MinU, grassBottomRegion.MaxV},
}

type FaceTemplate struct {
	Vertices [4][3]float32
	Normals  [3]float32
	UVs      [4][2]float32
	Indices  [6]uint16
}

var faceRightTemplate = FaceTemplate{
	Vertices: [4][3]float32{
		{0.5, -0.5, 0.5},  // BL
		{0.5, -0.5, -0.5}, // BR
		{0.5, 0.5, -0.5},  // TR
		{0.5, 0.5, 0.5},   // TL
	},
	Normals: [3]float32{1, 0, 0},
	UVs:     grassSideUVs,
	Indices: templateIndices,
}

var faceLeftTemplate = FaceTemplate{
	Vertices: [4][3]float32{
		{-0.5, -0.5, -0.5}, // BL
		{-0.5, -0.5, 0.5},  // BR
		{-0.5, 0.5, 0.5},   // TR
		{-0.5, 0.5, -0.5},  // TL
	},
	Normals: [3]float32{-1, 0, 0},
	UVs:     grassSideUVs,
	Indices: templateIndices,
}

var faceTopTemplate = FaceTemplate{
	Vertices: [4][3]float32{
		{-0.5, 0.5, -0.5},
		{-0.5, 0.5, 0.5},
		{0.5, 0.5, 0.5},
		{0.5, 0.5, -0.5},
	},
	Normals: [3]float32{0, 1, 0},
	UVs:     grassTopUVs,
	Indices: templateIndices,
}

var faceBottomTemplate = FaceTemplate{
	Vertices: [4][3]float32{
		{-0.5, -0.5, -0.5},
		{0.5, -0.5, -0.5},
		{0.5, -0.5, 0.5},
		{-0.5, -0.5, 0.5},
	},
	Normals: [3]float32{0, -1, 0},
	UVs:     grassBottomUVs,
	Indices: templateIndices,
}

var faceFrontTemplate = FaceTemplate{
	Vertices: [4][3]float32{
		{-0.5, -0.5, 0.5},
		{0.5, -0.5, 0.5},
		{0.5, 0.5, 0.5},
		{-0.5, 0.5, 0.5},
	},
	Normals: [3]float32{0, 0, 1},
	UVs:     grassSideUVs,
	Indices: templateIndices,
}

var faceBackTemplate = FaceTemplate{
	Vertices: [4][3]float32{
		{0.5, -0.5, -0.5},
		{-0.5, -0.5, -0.5},
		{-0.5, 0.5, -0.5},
		{0.5, 0.5, -0.5},
	},
	Normals: [3]float32{0, 0, -1},
	UVs:     grassSideUVs,
	Indices: templateIndices,
}

func (f FaceDir) Template() FaceTemplate {
	switch f {
	case FaceRight:
		return faceRightTemplate
	case FaceLeft:
		return faceLeftTemplate
	case FaceTop:
		return faceTopTemplate
	case FaceBottom:
		return faceBottomTemplate
	case FaceFront:
		return faceFrontTemplate
	case FaceBack:
		return faceBackTemplate
	}
	panic("unknown face")
}

func (f FaceDir) Offset() (x, y, z int) {
	switch f {
	case FaceRight:
		return 1, 0, 0
	case FaceLeft:
		return -1, 0, 0
	case FaceTop:
		return 0, 1, 0
	case FaceBottom:
		return 0, -1, 0
	case FaceFront:
		return 0, 0, 1
	case FaceBack:
		return 0, 0, -1
	}
	return 0, 0, 0
}
