// package main
// import "fmt"

// func main()  {
// 	for i:=0; i<5; i++{
// 		for j:=0; j<i+1; j++{
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
	
// }

package main
import "fmt"

func main(){
    
  for i := 0; i < 5; i++ {
  for j :=5; j<0; j-- {
    fmt.Print("*")
  }
  fmt.Println()
  }
}