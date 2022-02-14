package graphics

type triangle struct {
	v0, v1, v2 Vertex
}

func makeTriangle(v0, v1, v2 Vertex) triangle{
	var t triangle

	t.v0, t.v1, t.v2 = v0, v1, v2
	
	return t
}