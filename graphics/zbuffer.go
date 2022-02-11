package graphics

import (
	// "fmt"
	"image"
	"image/color"
	"math"
	"tulip/mymath"
	"tulip/scene"
)

type ZBufferGraphicsEngine struct {
	Cnv     Canvas
	Shadows Canvas

	zbuf  [][]float64
	sbuf  [][]float64
	light scene.Light

	fNear        float64
	fFar         float64
	fFov         float64
	fAspectRatio float64
	fFovRad      float64

	mProjection mymath.Matrix4x4
	mCamera     mymath.Matrix4x4
	mView       mymath.Matrix4x4

	mSun     mymath.Matrix4x4
	mSunView mymath.Matrix4x4
}

func MakeZBuffEngine(cnv Canvas) ZBufferGraphicsEngine {
	var engine ZBufferGraphicsEngine

	engine.Cnv = cnv
	engine.Shadows = MakeImageCanvas(cnv.height(), cnv.width())

	engine.initProjection()

	return engine

}

func (engine ZBufferGraphicsEngine) RenderScene(scn scene.Scene) {
	cnv := engine.Cnv
	engine.light = scn.LightSource

	// vUp := mymath.MakeVec3d(0, 1, 0)
	// t := mymath.MakeVec3d(0, 0, 1)
	// scn.Camera.VLookDir = mymath.MulVectorMatrix(t, mymath.MakeRotationYM(scn.Camera.FYaw))
	// scn.Camera.VTarget = mymath.Vec3dSum(scn.Camera.VCamera, scn.Camera.VLookDir)

	// engine.mCamera = mymath.MakePointAtM(scn.Camera.VCamera, scn.Camera.VTarget, vUp)
	// engine.mView = mymath.InverseTranslationM(engine.mCamera)

	// engine.mSun = mymath.MakePointAtM(scn.LightSource.Pos, mymath.Vec3dSum(scn.LightSource.Pos, scn.LightSource.Direction), vUp)
	// engine.mSunView = mymath.InverseTranslationM(engine.mSun)

	cnv.fill(scn.Background)

	zBack := float64(10000)
	engine.initBuf(cnv.height(), cnv.width(), zBack)

	//engine.ShadowModel(scn.Ground)

	for i := range scn.Objects {
		engine.ShadowModel(scn.Objects[i])
	}

	//fmt.Println(511, 96, engine.sbuf[511][96])

	// for i := range engine.sbuf {
	// 	for j := range engine.sbuf[i] {
	// 		if engine.sbuf[i][j] != zBack {
	// 			fmt.Println(i, j, engine.sbuf[i][j])
	// 		}
	// 	}
	// }

	//engine.RenderModel(scn.Ground)

	for i := range scn.Objects {
		engine.RenderModel(scn.Objects[i])
	}
}

func (engine ZBufferGraphicsEngine) RenderPolygonShadow(v1, v2, v3 scene.Vertex, clr color.NRGBA) {
	t := makeTriangle(v1, v2, v3)
	// world space --> view space
	engine.convertToViewSpace(&t)

	// clip against screen plane
	clipped, count := ClipAgainstPlane(mymath.MakeVec3d(0.0, 0.0, 0.1), mymath.MakeVec3d(0.0, 0.0, 1.0), t)

	for i := 0; i < count; i++ {
		// projection 3D --> 2D
		engine.project(&(clipped[i]))
		// engine.shadowMap(clipped[i], color.NRGBA{0, 0, 0, 255})

		// clip screen edges
		trisToDraw, trisCount := engine.clipScreenEdges(&(clipped[i]))

		engine.convertToSunSpace(&(clipped[i]))

		for j := 0; j < trisCount; j++ {
			// zbuffer rasterization
			engine.shadowMap(trisToDraw[j], color.NRGBA{0, 0, 0, 255})
		}
	}
}

func (engine ZBufferGraphicsEngine) inShadowOld(p mymath.Vec3d) bool {
	p.MulMatrix(engine.mSunView)

	// z := p.Z
	// p.MulMatrix(engine.mProjection)

	// p.DivW()

	// p.Add(mymath.MakeVec3d(1.0, 1.0, 0.0))

	// p.X *= 0.5 * float64(engine.Cnv.width())
	// p.Y *= 0.5 * float64(engine.Cnv.height())

	// p.Z = z

	//fmt.Println(p.X, p.Y, p.Z)

	if p.X >= 0 && p.Y >= 0 && p.X < float64(engine.Cnv.width()) && p.Y < float64(engine.Cnv.height()) {
		// fmt.Println("x= ", int(p.X), "y= ", int(p.Y), "buf= ", engine.sbuf[int(p.X)][int(p.Y)], "z= ", p.Z)
		if engine.sbuf[int(p.X)][int(p.Y)] < p.Z+5 {
			return true
		}
	}
	return false
}

func (engine *ZBufferGraphicsEngine) initProjection() {
	engine.fNear = 0.1
	engine.fFar = 1000.0
	engine.fFov = 90.0
	engine.fAspectRatio = float64(engine.Cnv.height()) / float64(engine.Cnv.width())
	engine.fFovRad = 1.0 / math.Tan((engine.fFov*0.5)/180.0*math.Pi)

	engine.mProjection[0][0] = engine.fAspectRatio * engine.fFovRad
	engine.mProjection[1][1] = engine.fFovRad
	engine.mProjection[2][2] = engine.fFar / (engine.fFar - engine.fNear)
	engine.mProjection[3][2] = (-engine.fFar * engine.fNear) / (engine.fFar - engine.fNear)
	engine.mProjection[2][3] = 1.0
	engine.mProjection[3][3] = 0.0
}

func (engine *ZBufferGraphicsEngine) initBuf(h, w int, value float64) {
	engine.zbuf = make([][]float64, h)
	engine.sbuf = make([][]float64, h)
	for i := range engine.zbuf {
		engine.zbuf[i] = make([]float64, w)
		engine.sbuf[i] = make([]float64, w)
		for j := range engine.zbuf[i] {
			engine.zbuf[i][j] = value
			engine.sbuf[i][j] = value
		}
	}
}

func projection(h int, v mymath.Vec3d) (int, int) {
	return int(v.X), h - int(v.Y)
}

func (engine ZBufferGraphicsEngine) ShadowModel(obj scene.PolygonialModel) {
	obj.IterateOverPolygons(engine.RenderPolygonShadow)
}

func (engine ZBufferGraphicsEngine) RenderModel(obj scene.PolygonialModel) {
	obj.IterateOverPolygons(engine.RenderPolygon)
}

func vectorIntersectPlane(planeP, planeN, lineStart, lineEnd mymath.Vec3d) mymath.Vec3d {
	planeN.Normalize()
	planeD := -planeN.DotProduct(planeP)
	ad := lineStart.DotProduct(planeN)
	bd := lineEnd.DotProduct(planeN)
	t := (-planeD - ad) / (bd - ad)
	lineStartToEnd := mymath.Vec3dDiff(lineEnd, lineStart)
	lineToIntersect := mymath.Vec3dMul(lineStartToEnd, t)

	return mymath.Vec3dSum(lineStart, lineToIntersect)
}

type triangle [3]scene.Vertex

func makeTriangle(v1, v2, v3 scene.Vertex) triangle {
	var t triangle
	t[0], t[1], t[2] = v1, v2, v3
	return t
}

func dist(p, planeP, planeN mymath.Vec3d) float64 {
	planeN.Normalize()

	return planeN.X*p.X + planeN.Y*p.Y + planeN.Z*p.Z - planeN.DotProduct(planeP)
}

func ClipAgainstPlane(planeP, planeN mymath.Vec3d, t triangle) ([2]triangle, int) {
	planeN.Normalize()

	var (
		inPoints, outPoints           [3]scene.Vertex
		inPointsCount, outPointsCount int
		res                           [2]triangle
	)

	d0 := dist(t[0].Point, planeP, planeN)
	d1 := dist(t[1].Point, planeP, planeN)
	d2 := dist(t[2].Point, planeP, planeN)

	if d0 >= 0 {
		inPoints[inPointsCount] = t[0]
		inPointsCount++
	} else {
		outPoints[outPointsCount] = t[0]
		outPointsCount++
	}

	if d1 >= 0 {
		inPoints[inPointsCount] = t[1]
		inPointsCount++
	} else {
		outPoints[outPointsCount] = t[1]
		outPointsCount++
	}

	if d2 >= 0 {
		inPoints[inPointsCount] = t[2]
		inPointsCount++
	} else {
		outPoints[outPointsCount] = t[2]
		outPointsCount++
	}

	if inPointsCount == 0 {
		return res, 0
	}

	if inPointsCount == 3 {
		res[0] = t
		return res, 1
	}

	if inPointsCount == 1 && outPointsCount == 2 {
		res[0][0] = inPoints[0]

		res[0][1].Point = vectorIntersectPlane(planeP, planeN, inPoints[0].Point, outPoints[0].Point)
		res[0][1].Normal = outPoints[0].Normal

		res[0][2].Point = vectorIntersectPlane(planeP, planeN, inPoints[0].Point, outPoints[1].Point)
		res[0][2].Normal = outPoints[1].Normal

		return res, 1
	}

	if inPointsCount == 2 && outPointsCount == 1 {
		res[0][0] = inPoints[0]

		res[0][1] = inPoints[1]

		res[0][2].Point = vectorIntersectPlane(planeP, planeN, inPoints[0].Point, outPoints[0].Point)
		res[0][2].Normal = outPoints[0].Normal

		res[0][0] = inPoints[1]

		res[0][1] = res[0][2]

		res[0][2].Point = vectorIntersectPlane(planeP, planeN, inPoints[1].Point, outPoints[0].Point)
		res[0][2].Normal = outPoints[0].Normal

		return res, 2
	}

	return res, 0
}

func (engine ZBufferGraphicsEngine) convertToViewSpace(t *triangle) {
	t[0].Point.MulMatrix(engine.mView)
	t[1].Point.MulMatrix(engine.mView)
	t[2].Point.MulMatrix(engine.mView)
}

func (engine ZBufferGraphicsEngine) convertToSunSpace(t *triangle) {
	t[0].Point.MulMatrix(engine.mSunView)
	t[1].Point.MulMatrix(engine.mSunView)
	t[2].Point.MulMatrix(engine.mSunView)
}

func (engine ZBufferGraphicsEngine) project(t *triangle) {
	z0 := t[0].Point.Z
	z1 := t[1].Point.Z
	z2 := t[2].Point.Z

	t[0].Point.MulMatrix(engine.mProjection)
	t[1].Point.MulMatrix(engine.mProjection)
	t[2].Point.MulMatrix(engine.mProjection)

	t[0].Point.DivW()
	t[1].Point.DivW()
	t[2].Point.DivW()

	// t[0].Point.Y *= -1
	// t[1].Point.Y *= -1
	// t[2].Point.Y *= -1

	t[0].Point.Add(mymath.MakeVec3d(1.0, 1.0, 0.0))
	t[1].Point.Add(mymath.MakeVec3d(1.0, 1.0, 0.0))
	t[2].Point.Add(mymath.MakeVec3d(1.0, 1.0, 0.0))

	t[0].Point.X *= 0.5 * float64(engine.Cnv.width())
	t[0].Point.Y *= 0.5 * float64(engine.Cnv.height())
	t[1].Point.X *= 0.5 * float64(engine.Cnv.width())
	t[1].Point.Y *= 0.5 * float64(engine.Cnv.height())
	t[2].Point.X *= 0.5 * float64(engine.Cnv.width())
	t[2].Point.Y *= 0.5 * float64(engine.Cnv.height())

	t[0].Point.Z = z0
	t[1].Point.Z = z1
	t[2].Point.Z = z2
}

func (engine ZBufferGraphicsEngine) clipScreenEdges(t *triangle) ([]triangle, int) {
	height := engine.Cnv.height()
	width := engine.Cnv.width()

	listTriangles := make([]triangle, 10)
	listTriangles[0] = makeTriangle(t[0], t[1], t[2])
	count := 1
	var clipped [2]triangle

	for p := 0; p < 4; p++ {
		trisToAdd := 0
		for count > 0 {
			test, listTriangles := listTriangles[0], listTriangles[1:] // pop from queue
			count--

			switch p {
			case 0:
				clipped, trisToAdd = ClipAgainstPlane(mymath.MakeVec3d(0.0, 0.0, 0.0), mymath.MakeVec3d(0.0, -1.0, 0.0), test)
			case 1:
				clipped, trisToAdd = ClipAgainstPlane(mymath.MakeVec3d(0.0, float64(height-1), 0.0), mymath.MakeVec3d(0.0, 1.0, 0.0), test)
			case 2:
				clipped, trisToAdd = ClipAgainstPlane(mymath.MakeVec3d(0.0, 0.0, 0.0), mymath.MakeVec3d(1.0, 0.0, 0.0), test)
			case 3:
				clipped, trisToAdd = ClipAgainstPlane(mymath.MakeVec3d(float64(width-1), 0.0, 0.0), mymath.MakeVec3d(-1.0, 0.0, 0.0), test)
			}

			for i := 0; i < trisToAdd; i++ {
				listTriangles = append(listTriangles, clipped[i])
			}
		}
		count = len(listTriangles)
	}

	return listTriangles, count
}

func (engine ZBufferGraphicsEngine) rasterize(t triangle, clr color.NRGBA) {
	if t[0].Point.Y > t[1].Point.Y {
		t[0], t[1] = t[1], t[0]
	}
	if t[0].Point.Y > t[2].Point.Y {
		t[0], t[2] = t[2], t[0]
	}
	if t[1].Point.Y > t[2].Point.Y {
		t[1], t[2] = t[2], t[1]
	}

	dyTotal := t[2].Point.Y - t[0].Point.Y

	light := engine.light

	IntensityV0 := scene.VertexIntensity(t[0], light)
	IntensityV1 := scene.VertexIntensity(t[1], light)
	IntensityV2 := scene.VertexIntensity(t[2], light)

	for y := t[0].Point.Y; y <= t[1].Point.Y; y++ {
		dySegment := t[1].Point.Y - t[0].Point.Y + 1
		alpha := float64((y - t[0].Point.Y) / dyTotal)
		beta := float64((y - t[0].Point.Y) / dySegment)

		var a, b mymath.Vec3d

		a = mymath.Vec3dDiff(t[2].Point, t[0].Point)
		a.Mul(alpha)
		a.Add(t[0].Point)

		b = mymath.Vec3dDiff(t[1].Point, t[0].Point)
		b.Mul(beta)
		b.Add(t[0].Point)

		var i1, i2 float64

		i1 = IntensityV0 + (IntensityV2-IntensityV0)*alpha
		i2 = IntensityV0 + (IntensityV1-IntensityV0)*beta

		if a.X > b.X {
			a, b = b, a
			//i1, i2 = i2, i1
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   mymath.Vec3d
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			intensity := i1 + (i2-i1)*phi

			// transform coordinates
			if x >= 0 && y >= 0 && x < float64(engine.Cnv.width()) && y < float64(engine.Cnv.height()) {
				if engine.zbuf[int(x)][int(y)] > p.Z {
					engine.zbuf[int(x)][int(y)] = p.Z
					pixel_x, pixel_y := projection(engine.Cnv.height(), p)
					// pixel_x, pixel_y := int(p.X), int(p.Y)

					if engine.inShadowOld(p) {
						// fmt.Println("in shadow")
						engine.Cnv.setPixel(pixel_x, pixel_y, color.NRGBA{255, 255, 255, 0})
					} else {
						//fmt.Println("out of shadow")
						engine.Cnv.setPixel(pixel_x, pixel_y, scene.Lightness(clr, intensity))
					}
				}
			}
		}
	}

	for y := t[1].Point.Y; y <= t[2].Point.Y; y++ {
		dySegment := t[2].Point.Y - t[1].Point.Y + 1
		alpha := float64((y - t[0].Point.Y) / dyTotal)
		beta := float64((y - t[1].Point.Y) / dySegment)

		var a, b mymath.Vec3d

		a = mymath.Vec3dDiff(t[2].Point, t[0].Point)
		a.Mul(alpha)
		a.Add(t[0].Point)

		b = mymath.Vec3dDiff(t[2].Point, t[1].Point)
		b.Mul(beta)
		b.Add(t[1].Point)

		var i1, i2 float64

		i1 = IntensityV0 + (IntensityV2-IntensityV0)*alpha
		i2 = IntensityV0 + (IntensityV1-IntensityV0)*beta

		if a.X > b.X {
			a, b = b, a
			i1, i2 = i2, i1
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   mymath.Vec3d
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			intensity := i1 + (i2-i1)*phi

			if x >= 0 && y >= 0 && x < float64(engine.Cnv.width()) && y < float64(engine.Cnv.height()) {
				if engine.zbuf[int(x)][int(y)] > p.Z {
					engine.zbuf[int(x)][int(y)] = p.Z
					pixel_x, pixel_y := projection(engine.Cnv.height(), p)
					// pixel_x, pixel_y := int(p.X), int(p.Y)
					if engine.inShadowOld(p) {
						engine.Cnv.setPixel(pixel_x, pixel_y, color.NRGBA{255, 255, 255, 0})
					} else {
						engine.Cnv.setPixel(pixel_x, pixel_y, scene.Lightness(clr, intensity))
					}

				}
			}
		}
	}
}

func (engine ZBufferGraphicsEngine) shadowMap(t triangle, clr color.NRGBA) {
	if t[0].Point.Y > t[1].Point.Y {
		t[0], t[1] = t[1], t[0]
	}
	if t[0].Point.Y > t[2].Point.Y {
		t[0], t[2] = t[2], t[0]
	}
	if t[1].Point.Y > t[2].Point.Y {
		t[1], t[2] = t[2], t[1]
	}

	dyTotal := t[2].Point.Y - t[0].Point.Y

	//light := engine.light

	// IntensityV0 := scene.VertexIntensity(t[0], light)
	// IntensityV1 := scene.VertexIntensity(t[1], light)
	// IntensityV2 := scene.VertexIntensity(t[2], light)

	for y := t[0].Point.Y; y <= t[1].Point.Y; y++ {
		dySegment := t[1].Point.Y - t[0].Point.Y + 1
		alpha := float64((y - t[0].Point.Y) / dyTotal)
		beta := float64((y - t[0].Point.Y) / dySegment)

		var a, b mymath.Vec3d

		a = mymath.Vec3dDiff(t[2].Point, t[0].Point)
		a.Mul(alpha)
		a.Add(t[0].Point)

		b = mymath.Vec3dDiff(t[1].Point, t[0].Point)
		b.Mul(beta)
		b.Add(t[0].Point)

		// var i1, i2 float64

		// i1 = IntensityV0 + (IntensityV2-IntensityV0)*alpha
		// i2 = IntensityV0 + (IntensityV1-IntensityV0)*beta

		if a.X > b.X {
			a, b = b, a
			//i1, i2 = i2, i1
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   mymath.Vec3d
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			//intensity := i1 + (i2-i1)*phi

			// transform coordinates
			if x >= 0 && y >= 0 && x < float64(engine.Cnv.width()) && y < float64(engine.Cnv.height()) {
				if engine.sbuf[int(x)][int(y)] > p.Z {
					engine.sbuf[int(x)][int(y)] = p.Z
					//pixel_x, pixel_y := projection(engine.Cnv.height(), p)
					// pixel_x, pixel_y := int(p.X), int(p.Y)

					//engine.Cnv.setPixel(pixel_x, pixel_y, scene.Lightness(clr, intensity))
				}
			}
		}
	}

	for y := t[1].Point.Y; y <= t[2].Point.Y; y++ {
		dySegment := t[2].Point.Y - t[1].Point.Y + 1
		alpha := float64((y - t[0].Point.Y) / dyTotal)
		beta := float64((y - t[1].Point.Y) / dySegment)

		var a, b mymath.Vec3d

		a = mymath.Vec3dDiff(t[2].Point, t[0].Point)
		a.Mul(alpha)
		a.Add(t[0].Point)

		b = mymath.Vec3dDiff(t[2].Point, t[1].Point)
		b.Mul(beta)
		b.Add(t[1].Point)

		//var i1, i2 float64

		// i1 = IntensityV0 + (IntensityV2-IntensityV0)*alpha
		// i2 = IntensityV0 + (IntensityV1-IntensityV0)*beta

		// if a.X > b.X {
		// 	a, b = b, a
		// 	i1, i2 = i2, i1
		// }

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   mymath.Vec3d
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			//intensity := i1 + (i2-i1)*phi

			if x >= 0 && y >= 0 && x < float64(engine.Cnv.width()) && y < float64(engine.Cnv.height()) {
				if engine.sbuf[int(x)][int(y)] > p.Z {
					engine.sbuf[int(x)][int(y)] = p.Z
					// pixel_x, pixel_y := projection(engine.Cnv.height(), p)
					// // pixel_x, pixel_y := int(p.X), int(p.Y)
					// engine.Cnv.setPixel(pixel_x, pixel_y, scene.Lightness(clr, intensity))
				}
			}
		}
	}
}

func (engine ZBufferGraphicsEngine) RenderPolygon(v1, v2, v3 scene.Vertex, clr color.NRGBA) {
	t := makeTriangle(v1, v2, v3)
	// world space --> view space
	engine.convertToViewSpace(&t)

	// clip against screen plane
	clipped, count := ClipAgainstPlane(mymath.MakeVec3d(0.0, 0.0, 0.1), mymath.MakeVec3d(0.0, 0.0, 1.0), t)

	for i := 0; i < count; i++ {
		// projection 3D --> 2D
		engine.project(&(clipped[i]))

		// clip screen edges
		trisToDraw, trisCount := engine.clipScreenEdges(&(clipped[i]))

		for j := 0; j < trisCount; j++ {
			// zbuffer rasterization
			engine.rasterize(trisToDraw[j], clr)
		}
	}
}

func (engine ZBufferGraphicsEngine) RenderWire(v1, v2, v3 scene.Vertex, clr color.NRGBA) {
	cnv := engine.Cnv
	cnv.drawLine(int(v1.Point.X), cnv.height()-int(v1.Point.X), int(v2.Point.X), cnv.height()-int(v2.Point.X), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v1.Point.X), cnv.height()-int(v1.Point.X), int(v3.Point.X), cnv.height()-int(v3.Point.X), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v2.Point.X), cnv.height()-int(v2.Point.X), int(v3.Point.X), cnv.height()-int(v3.Point.X), color.RGBA{0, 0, 0, 255})
}

func (engine ZBufferGraphicsEngine) RenderNormals(v1, v2, v3 scene.Vertex, clr color.NRGBA) {
	cnv := engine.Cnv
	cnv.drawLine(int(v1.Point.X), cnv.height()-int(v1.Point.X), int(v1.Point.X+v1.Normal.X), cnv.height()-int(v1.Point.X+v1.Normal.X), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v2.Point.X), cnv.height()-int(v2.Point.X), int(v2.Point.X+v2.Normal.X), cnv.height()-int(v2.Point.X+v2.Normal.X), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v3.Point.X), cnv.height()-int(v3.Point.X), int(v3.Point.X+v3.Normal.X), cnv.height()-int(v3.Point.X+v3.Normal.X), color.RGBA{0, 0, 0, 255})
}

func (engine ZBufferGraphicsEngine) Image() image.Image {
	return engine.Cnv.Image()
}
