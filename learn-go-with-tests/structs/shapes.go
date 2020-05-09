// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - shapes.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
