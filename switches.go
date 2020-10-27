//Algorithm no 4
package main

import (
	"fmt"
	"math"
)

func main(){
	end := 100
	a := 1
	pa := 3
	pb := 2

	n := math.Floor(math.Sqrt(float64(end))) - 1
	PSn := math.Pow(n, 2) + 2 * n
	PUn := pa + pb * (int(n) - 1)
	pn := ((PUn - pa)/pb) + 1
	b := int(PSn)/pn
	Un := PSn + 1
	rn := ((int(Un) - a)/b) + 1
	fmt.Println(rn)
	// 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
}