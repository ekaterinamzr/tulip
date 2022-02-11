package graphics

import (
	"image/color"
	"tulip/mymath"
	"tulip/scene"
)

type MyGrEngine struct {
	cnv Canvas
	
	projection mymath.Matrix4x4
	viewport   mymath.Matrix4x4

	zBuf  []float64
	zBack float64
}

// Creating new engine
func NewMyGrEngine(cnv Canvas) *MyGrEngine {
	engine := new(MyGrEngine)

	// setting canvas
	engine.cnv = cnv
	
	// setting z-buffer
	engine.zBack = 10000
	engine.zBuf = make([]float64, engine.cnv.height())

	return engine
}

func (engine *MyGrEngine) initZBuf() {
	for i := range engine.zBuf {
		engine.zBuf[i] = engine.zBack
	}
}

// Rendering scene 
func (engine MyGrEngine) RenderScene(scn *scene.Scene) {

	// setting up the engine 

	// rendering
	engine.cnv.fill(scn.Background)

	for i := range scn.Objects {
		engine.renderModel(scn.Objects[i])
	}
}

func (engine MyGrEngine) renderModel(m scene.PolygonialModel) {
	m.IterateOverPolygons(engine.processTriangle)
}

// creating a tringle -> then rendering
func (engine MyGrEngine) processTriangle(v0, v1, v2 scene.Vertex, clr color.NRGBA) {
	t := makeTriangle(v0, v1, v2, clr)

	engine.renderTriangle(t)
}

// applying projection, viewport -> then rasterizing
func (engine MyGrEngine) renderTriangle(t triangle) {
	engine.rasterizeWire(t)
}

func (engine MyGrEngine) rasterizeWire(t triangle) {
	x0, y0 := int(t.v0.Point.X), int(t.v0.Point.Y)
	x1, y1 := int(t.v1.Point.X), int(t.v1.Point.Y)
	x2, y2 := int(t.v2.Point.X), int(t.v2.Point.Y)

	engine.cnv.drawLine(x0, y0, x1, y1, t.clr)
	engine.cnv.drawLine(x1, y1, x2, y2, t.clr)
	engine.cnv.drawLine(x0, y0, x2, y2, t.clr)
}