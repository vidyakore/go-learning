package main
import (
	"fmt"
	"math"
)

func main()  {
	var a float64 = 63
	var result float64 = math.Sqrt(a)
	fmt.Println(result)

	fmt.Printf("the precised output is %.2g", result)	//for formatting output use printf
}