package geo

import "github.com/go-gl/mathgl/mgl32"

type Triangle struct {
	A, B, C mgl32.Vec2
}

func (t Triangle) Area() float32 {
	ab := t.B.Sub(t.A)
	ac := t.C.Sub(t.A)

	cp := ab.X()*ac.Y() - ab.Y()*ac.X()
	if cp < 0 {
		cp *= -1
	}
	return cp / 2
}

// Return a new triangle translated by p
func (t Triangle) Add(p mgl32.Vec2) Triangle {
	return Triangle{t.A.Add(p), t.B.Add(p), t.C.Add(p)}
}

func (t Triangle) Contains(p mgl32.Vec2) bool {
	origA := t.Area()
	abp := Triangle{t.A, t.B, p}.Area()
	acp := Triangle{t.A, t.C, p}.Area()
	bcp := Triangle{t.B, t.C, p}.Area()
	return origA == abp+acp+bcp
}
