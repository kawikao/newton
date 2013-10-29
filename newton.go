// Copyright 2013 Kawika Ohumukini.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// General environment variables.

package newton

import "math"

// Sqrt finds the square root of a number using Newton's method.
func Sqrt(x float64) float64 {
	for i := 0; i < 10; i++ {
		newtonX := x - (((x * x) - 2) / (2 * x))
		if (math.Abs(newtonX - x)) < 10e-10 {
			return newtonX
		}
		x = newtonX
	}
	return x
}
