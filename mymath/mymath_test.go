package mymath

import "testing"

func TestInverse(t *testing.T) {
	m := Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}
	inv := Matrix4x4{{0.25, 0.25, 0.25, -0.25}, {0.25, 0.25, -0.25, 0.25}, {0.25, -0.25, 0.25, 0.25}, {-0.25, 0.25, 0.25, 0.25}}
	got, ok := m.Inverse()

	if !ok {
		t.Error("Not ok")
	}

	if got != inv {
		t.Error("\n got:\t", got, "\n want:\t", inv)
	}
}

func TestDeterminant(t *testing.T) {
	m := Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}
	det := -16.0
	got := m.Determinant()

	if got != det {
		t.Error("got: ", got, "\n want: ", det)
	}
}

type submDetTest struct {
	m      Matrix4x4
	i, j   int
	wanted float64
}

// var testMatrix = Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}
var submDetTests = []submDetTest{
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 0, 0, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 0, 1, 4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 0, 2, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 0, 3, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 1, 0, 4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 1, 1, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 1, 2, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 1, 3, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 2, 0, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 2, 1, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 2, 2, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 2, 3, 4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 3, 0, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 3, 1, -4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 3, 2, 4},
	submDetTest{Matrix4x4{{1, 1, 1, -1}, {1, 1, -1, 1}, {1, -1, 1, 1}, {-1, 1, 1, 1}}, 3, 3, -4},
}

func TestSubmatrixDet(t *testing.T) {
	for _, test := range submDetTests {
		if got := test.m.submatrixDet(test.i, test.j); got != test.wanted {
			t.Error("for ", test.i, test.j, " got: ", got, " want: ", test.wanted)
		}
	}
}
