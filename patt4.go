package main
import "fmt"

func main()  {
	for i:=0; i<5; i++{
		for j:=i; j<5; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
	
}


// package main
// import "fmt"

// func main()  {
//   for i:=5; i>0; i--{
//     for j:=i; j>0; j--{
//       fmt.Print("*")
//     }
//     fmt.Println()
//   }
  
// }