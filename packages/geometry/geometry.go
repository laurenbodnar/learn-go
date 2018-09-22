// This application will calculate the area and diagonal of a rectangle using different packages. Following tutorial part 7 of (golangbot) [https://golangbot.com/packages/]


package main

import (
   "fmt"
   "learn-go/packages/geometry/rectangle" // importing custom pkg
   "log"
)
/*
 * 1. package variables
*/
var rectLen, rectWidth float64 = -6, 7 

/*
*2. init function to check if length and width are greater than zero
*/
func init() {  
    println("main package initialized")
    if rectLen < 0 {
        log.Fatal("length is less than zero")
    }
    if rectWidth < 0 {
        log.Fatal("width is less than zero")
    }
}

func main() {
   fmt.Println("Geometrical shape properties")
   // Area function of rectangle pkg used
   fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
        /*Diagonal function of rectangle package used
        */
    fmt.Printf("diagonal of the rectangle %.2f", rectangle.Diagonal(rectLen, rectWidth))
}










