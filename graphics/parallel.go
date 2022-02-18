package graphics

import (
	"sync"
	"tulip/scene"
)

func (engine MyGrEngine) processVerticesParallel(vertices []scene.Vertex, indices []int, n int) {
	processed := make([]Vertex, len(vertices))
	k := len(vertices) / n

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		start := i * k
		end := (i + 1) * k
		if (i + 1) == n {
			end = len(vertices)
		}

		wg.Add(1)

		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				processed[j] = engine.shader.vs(vertices[j])
			}
		}(start, end)
	}

	wg.Wait()

	engine.assembleTriangles(processed, indices)
}
