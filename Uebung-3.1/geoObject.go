package main

import "fmt"

type Position struct {
	X float64
	Y float64
}

type Circle struct {
	GeoObject
	Radius float64
}

type Rectangle struct {
	GeoObject
	Width float64
	Height float64
}

type Triangle struct {
	GeoObject
	B Position
	C Position
}

type GeoObject struct {
	Position
	Color int
}

type Painter interface {
	Paint()
}

func (obj GeoObject) Paint() {
	 fmt.Printf("X=%v, Y=%v, Color=%v", obj.X, obj.Y, obj.Color)
}

func (circle Circle) Paint() {
	fmt.Printf("%v, Radius=%v", circle.GeoObject, circle.Radius)
}

func (rectangle Rectangle) Paint() {
	fmt.Printf("%v, Width=%v, Height=%v", rectangle.GeoObject, rectangle.Width, rectangle.Height)
}

func (triangle Triangle) Paint() {
	fmt.Printf("PointA=%v, PointB=%v, PointC=%v", triangle.GeoObject, triangle.B, triangle.C)
}
