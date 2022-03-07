package main
import "fmt"
var a=20 //package level variable
func demo(){
	var a=10		//local variable or function level variable
	fmt.Println("in demo",a)
}
func main(){
	
	fmt.Println("in main",a)
	demo()
}