//A Tour of Go: Exercise Loops ad Functions
// Daniel's experiment with it, explaining calls and returns showing fibonacci




package main

import (
	"fmt"
)


func sum(x int, y int) int {
	sumResult := x + y
	return sumResult
}

func Sqrt(x float64) float64 {
	z := 1.0  // must define the var before the for loop
	
    for i := float64(1); i < 10; i++ {
		z -= (z*z - x) / (2*z) 
		fmt.Println(z)
	}
	return z
}

func main() {
	
	//result := Sqrt(10)
	result := 1
	result2 := 2
	
	for i :=1; i<10; i++ {
		result3 := sum(result, result2)
		fmt.Printf("loop number %d: \n %d + %d = %d \n", i, result, result2, result3)
		result = result2
		result2 = result3
	}
	
	fmt.Println(result)
}









