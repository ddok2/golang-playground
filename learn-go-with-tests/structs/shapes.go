// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - shapes.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0
}
