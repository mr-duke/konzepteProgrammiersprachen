package main

import "testing"

func TestPrint (t *testing.T) {

	l:= []Painter {
		Circle {
			GeoObject: GeoObject{
				Color: 10,
				Position: Position {
					X: 2,
					Y: 7,
					},
				},    
			Radius: 1,
		},
		Triangle {
			GeoObject: GeoObject{
				Position: Position{
					X: 2,
					Y: 3,
					},
				Color: 7,
				},
			B: Position {
				X: 8,
				Y: 7,
			},
			C: Position {
				X: 12,
				Y: 20,
			},
		},
	}

	for _, object := range l {
		object.Paint()
	}
}