/*
https://stackoverflow.com/a/19908214
*/
package arange

import "math"

/*
for example:

arange(0., 10., 0.5)
should return
[0 0.5 1 1.5 2 2.5 3 3.5 4 4.5 5 5.5 6 6.5 7 7.5 8 8.5 9 9.5]
*/

func Arange(start, stop, step float64) []float64 {
	N := int(math.Ceil((stop - start) / step))
	rnge := make([]float64, N, N)
	i := 0
	for x := start; x < stop; x += step {
		rnge[i] = x
		i += 1
	}
	return rnge
}
