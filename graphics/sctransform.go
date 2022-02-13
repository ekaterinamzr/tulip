package graphics

type psTransformer struct {
	xFactor, yFactor float64
}

// Perspective screen transformation
func makePST(w, h int) psTransformer {
	var pst psTransformer
	pst.xFactor = float64(w) / 2
	pst.yFactor = float64(h) / 2 
	return pst
}

func (pst psTransformer) transform(v *Vertex) {
	wInv := 1.0 / v.Point.W 
	v.Point.Mul(wInv)

	v.Point.X = (v.Point.X + 1.0) * pst.xFactor
	v.Point.Y = (-v.Point.Y + 1.0) * pst.yFactor

	//v.Point.W = wInv
}