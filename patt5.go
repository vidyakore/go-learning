package main
import "fmt"

func main()  {
  for i:=5; i>0; i--{
    for j:=i; j>0; j--{
      fmt.Print("*")
    }
	fmt.Print("\n")
	
  }
  
}