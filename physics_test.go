package main

import (
	"fmt"
	"testing"
)

func TestDist(t *testing.T) {
	e1 := Entity{0, 0, 0, 0, 0, 0}
	e2 := Entity{100, 0, 0, 0, 0, 0}
	e3 := Entity{0, 100, 0, 0, 0, 0}
	e4 := Entity{100, 100, 100, 100, 0, 0}
	e5 := Entity{125, 317, 100, 100, 0, 0}
	e6 := Entity{412, 23, 100, 100, 0, 0}
	d1 := dist(e1, e2)
	d2 := dist(e1, e3)
	d3 := dist(e1, e4)
	d4 := dist(e1, e5)
	d5 := dist(e1, e6)
	d6 := dist(e5, e6)
	d1P := distP(e1.x, e1.y, e2.x, e2.y)
	d2P := distP(e1.x, e1.y, e3.x, e3.y)
	d3P := distP(e1.x, e1.y, e4.x, e4.y)
	d4P := distP(e1.x, e1.y, e5.x, e5.y)
	d5P := distP(e1.x, e1.y, e6.x, e6.y)
	d6P := distP(e5.x, e5.y, e6.x, e6.y)

	//fmt.Printf("d: %v %v %v %v %v %v\n", d1, d2, d3, d4, d5, d6)
	//fmt.Printf("dP: %v %v %v %v %v %v\n", d1P, d2P, d3P, d4P, d5P, d6P)
	if d1 != d1P || d2 != d2P || d3 != d3P || d4 != d4P || d5 != d5P || d6 != d6P {
		t.Error("Error while comparing dist and distP")
	}
	if d3 != 141.4213562373095 {
		t.Errorf("Error: expected 141.4213562373095 and got %v", d3)
	}
	if d4 != 340.7550439832109 {
		t.Errorf("Error: expected 340.7550439832109 and got %v", d4)
	}
	if d5 != 412.6414908852477 {
		t.Errorf("Error: expected 412.6414908852477 and got %v", d5)
	}
	if d6 != 410.85885654321726 {
		t.Errorf("Error: expected 410.85885654321726 and got %v", d6)
	}
}

func TestDot(t *testing.T) {
	v1 := Vect{0, 1}
	v2 := Vect{0, 100}
	v3 := Vect{100, 0}
	v4 := Vect{100, 100}

	d1 := dot(v1, v2)
	d2 := dot(v1, v3)
	d3 := dot(v1, v4)
	d4 := dot(v4, v1)

	//fmt.Printf("d: %v %v %v %v \n", d1, d2, d3, d4)
	if d1 != 100 {
		t.Errorf("Error: expected 100 and got %v", d1)
	}
	if d2 != 0 {
		t.Errorf("Error: expected 0 and got %v", d2)
	}
	if d3 != 100 {
		t.Errorf("Error: expected 100 and got %v", d3)
	}
	if d3 != d4 {
		t.Errorf("Error: expected 100 and got %v", d4)
	}
}

func TestCollision(t *testing.T) {
	e1 := Entity{1000, 1000, 100, 100, 100, 1}
	e2 := Entity{1050, 1050, 0, 0, 100, 1}
	e3 := Entity{1200, 1200, 0, 0, 100, 1}

	b1, r1 := isCollision(e1, e2)
	b2, r2 := isCollision(e1, e3)

	fmt.Printf("r: %v %v, %v %v \n", b1, r1, b2, r2)
}
