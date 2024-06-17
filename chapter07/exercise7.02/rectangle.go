package main

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}

func (r rectangle) Name() string {
	return "rectangle"
}
