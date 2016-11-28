package main

import "math"

var (
	w, h int
)

type Entity struct {
	x, y, vx, vy, r, m float64
}

type Vect struct {
	x, y float64
}

func (v *Vect) normalize() {
	norm := distP(0, 0, v.x, v.y)
	if norm > 0 {
		v.x = v.x / norm
		v.y = v.y / norm
	}
}

func main() {
	w = 16001
	h = 7501

}

func dist2(e1, e2 Entity) float64 {
	return (e1.x-e2.x)*(e1.x-e2.x) + (e1.y-e2.y)*(e1.y-e2.y)
}

func dist(e1, e2 Entity) float64 {
	return math.Sqrt(dist2(e1, e2))
}

func dist2P(x1, y1, x2, y2 float64) float64 {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

func distP(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(dist2P(x1, y1, x2, y2))
}

func dot(v1, v2 Vect) float64 {
	return distP(0, 0, v1.x, v1.y) * distP(0, 0, v2.x, v2.y) * math.Cos(math.Atan2(v2.y-v1.y, v2.x-v1.x))
}

func isCollision(e1, e2 Entity) (isColl bool, result float64) {
	// If speed norm is less than dist less radiuses, then no Collision
	sNorm := distP(e1.vx, e1.vy, e2.vx, e2.vy)
	if sNorm < dist(e1, e2)-float64(e1.r+e2.r) {
		return false, 0
	}
	// Normalize V -> N
	N := Vect{e1.vx + e2.vx, e1.vy + e2.vy}
	N.normalize()
	// Find vector C between centers
	C := Vect{e2.x - e1.x, e2.y - e1.y}
	// Compute product
	lengthC := distP(0, 0, C.x, C.y)
	D := dot(N, C)
	// if direction is opposite, no Collision
	if D <= 0 {
		return false, 0
	}
	//Get F length
	F := lengthC*lengthC - D*D
	sumR2 := (e1.r + e2.r) * (e1.r + e2.r)
	if F >= sumR2 {
		return false, 0
	}
	T := sumR2 - F
	if T < 0 {
		return false, 0
	}
	// Dist to travel
	dist := D - math.Sqrt(T)
	if sNorm < dist {
		return false, 0
	}
	// Compute the ratio
	ratio := dist / sNorm
	return true, ratio
}

func computeMove(e1, e2 Entity, ratio float64) (ne1, ne2 Entity) {
	// First get the things in contact
	e1.x = e1.x + ratio*e1.vx
	e1.y = e1.y + ratio*e1.vy
	e2.x = e2.x + ratio*e2.vx
	e2.y = e2.y + ratio*e2.vy
	// Compute the new speed vectors
	N := Vect{e2.x - e1.x, e2.y - e1.y}
	N.normalize()
	v1 := Vect{e1.vx, e1.vy}
	v2 := Vect{e2.vx, e2.vy}
	a1 := dot(N, v1)
	a2 := dot(N, v2)
	//Compute Optimized param
	optim := 2 * (a1 - a2) / (e1.m + e2.m)
	nv1 := Vect{v1.x - optim*e2.m*N.x, v1.y - optim*e2.m*N.y}
	nv2 := Vect{v2.x - optim*e1.m*N.x, v2.y - optim*e1.m*N.y}

	ne1 = Entity{e1.x + (1-ratio)*nv1.x, e1.y + (1-ratio)*nv1.y, nv1.x, nv1.y, e1.r, e1.m}
	ne2 = Entity{e2.x + (1-ratio)*nv2.x, e2.y + (1-ratio)*nv2.y, nv2.x, nv2.y, e2.r, e2.m}

	return ne1, ne2
}
