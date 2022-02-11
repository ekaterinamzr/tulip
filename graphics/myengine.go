package graphics 

import (
	"tulip/mymath"
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

// Rendering scene 
func (engine MyGrEngine) RenderScene() {

}