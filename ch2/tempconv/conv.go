// Copyright 2020. SUNGYUB.COM. All Rights Reserved.
//
// - conv.go
// - author: Sungyub NA <mailto: darkerkorean@gmail.com>

package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
