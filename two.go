//Program 1

package main
import "fmt"

func main(){
    
  for i := 0; i < 5; i++ {
  for j :=0; j<5; j++ {
    fmt.Print("*")
  }
  fmt.Println()
  }
}

//Program 2

func main()  {
	for i:=0; i<5; i++{
		for j:=i; j<5; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	
}

//Program 3


func main(){
    
  for i := 0; i < 5; i++ {
  for j :=5; j<0; j-- {
    fmt.Print("*")
  }
  fmt.Println()
  }
}