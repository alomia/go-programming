package main

type triangle struct {
	base   float64
	height float64
}

func (t triangle) Area() float64 {
	return t.base * t.height / 2
}

func (t triangle) Name() string {
	return "triangle"
}
