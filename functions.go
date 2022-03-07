package main
import "fmt"

func mul(a,b int) int{
	return a*b
}

func cal(a,b int)(int,int){
	return a+b,a-b		//can return two values
}
func main(){
	a:=10
	b:=20
	result:=mul(a,b)
	fmt.Println(result)

	result1,result2 := cal(a,b)
	fmt.Println(result1,result2)
}