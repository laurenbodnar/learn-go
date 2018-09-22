// Actual answer to https://tour.golang.org/flowcontrol/8



package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i < 100; i++ {
		fmt.Println("---------- ", i, " ----------")
		// Adjust z based on how close z^2 is to x
		calc := z - (z*z - x) / (2*z)
		fmt.Println(z)
		// Stop loop once calculation has not
        // changed since last iteration
		fmt.Println(z, calc, (calc - z), (z - calc))
		if (calc - z > 0 && calc - z < 0.0001) || (z - calc > 0 && z - calc < 0.0001)  {
			fmt.Println("breaking:")
            break
	    }
		
		z = calc
		
    }

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}





