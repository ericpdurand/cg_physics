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
	//return distP(0, 0, v1.x, v1.y) * distP(0, 0, v2.x, v2.y) * math.Cos(math.Atan2(v2.y-v1.y, v2.x-v1.x))
	return v1.x*v2.x + v1.y*v2.y
}

func isCollision(e1, e2 Entity) (isColl bool, result float64) {
	// If speed norm is less than dist less radiuses, then no Collision
	sNorm := distP(e1.vx, e1.vy, e2.vx, e2.vy)
	//fmt.Printf("sNorm: %v \n", sNorm)
	if sNorm < dist(e1, e2)-float64(e1.r+e2.r) {
		return false, 0
	}
	// Normalize V -> N
	N := Vect{e1.vx + e2.vx, e1.vy + e2.vy}
	N.normalize()
	//fmt.Printf("N: %v \n", N)
	// Find vector C between centers
	C := Vect{e2.x - e1.x, e2.y - e1.y}
	// Compute product
	lengthC := distP(0, 0, C.x, C.y)
	D := dot(N, C)
	//fmt.Printf("D, lengthC: %v %v \n", D, lengthC)
	// if direction is opposite, no Collision
	if D <= 0 {
		return false, 0
	}
	//Get F length
	F := lengthC*lengthC - D*D
	sumR2 := (e1.r + e2.r) * (e1.r + e2.r)
	//fmt.Printf("F, sumR2: %v %v \n", F, sumR2)
	if F >= sumR2 {
		return false, 0
	}
	T := sumR2 - F
	if T < 0 {
		return false, 0
	}
	// Dist to travel
	dist := D - math.Sqrt(T)
	//fmt.Printf("dist: %v \n", dist)
	if sNorm < dist {
		return false, 0
	}
	// Compute the ratio
	ratio := dist / sNorm
	//fmt.Printf("ratio: %v \n", ratio)
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
	//fmt.Printf("N: %v \n", N)
	N.normalize()
	v1 := Vect{e1.vx, e1.vy}
	v2 := Vect{e2.vx, e2.vy}
	a1 := dot(v1, N)
	a2 := dot(v2, N)
	//fmt.Printf("a1: %v \n", a1)
	//fmt.Printf("a2: %v \n", a2)
	//Compute Optimized param
	optim := 2 * (a1 - a2) / (e1.m + e2.m)
	nv1 := Vect{v1.x - optim*e2.m*N.x, v1.y - optim*e2.m*N.y}
	nv2 := Vect{v2.x + optim*e1.m*N.x, v2.y + optim*e1.m*N.y}
	//Check the impulse
	N1 := distP(0, 0, nv1.x, nv1.y)
	N2 := distP(0, 0, nv2.x, nv2.y)
	if N1 < 100 && N1 > 0 {
		nv1.x = nv1.x / N1 * 100
		nv1.y = nv1.y / N1 * 100
	}
	if N1 == 0 {
		nv1.x = -N.x * 100
		nv1.y = -N.y * 100
	}

	if N2 < 100 && N2 > 0 {
		nv2.x = nv2.x / N2 * 100
		nv2.y = nv2.y / N2 * 100
	}
	if N2 == 0 {
		nv2.x = N.x * 100
		nv2.y = N.y * 100
	}

	ne1 = Entity{e1.x + (1-ratio)*nv1.x, e1.y + (1-ratio)*nv1.y, nv1.x, nv1.y, e1.r, e1.m}
	ne2 = Entity{e2.x + (1-ratio)*nv2.x, e2.y + (1-ratio)*nv2.y, nv2.x, nv2.y, e2.r, e2.m}

	return ne1, ne2
}

func isWallCollision(e Entity, w, h int) (bool, float64) {
	result := false
	ratio := float64(1)
	//First simulate move
	nx := e.x + e.vx
	ny := e.y + e.vy
	//Did we cross a wall?
	if nx-e.r <0 {
		r := nx-e.r / e.vx 
		if r < ratio {
			ratio = r
			result = true
		}  		
	}
	if ny-e.r <0 {
                r := ny-e.r / e.vy 
                if r < ratio {
                        ratio = r
                        result = true
                }
        }
	if nx+e.r >= float64(w) {
                r := (nx+e.r-float64(w)) / e.vx 
                if r < ratio {
                        ratio = r
                        result = true
                }
        }
        if ny+e.r >= float64(h) {
                r := (ny+e.r-float64(h)) / e.vy 
                if r < ratio {
                        ratio = r
                        result = true
                }
        }
	return result,ratio
}

