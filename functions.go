package main
import "fmt"

func mul(a,b int) int{
	return a*b
}

func cal(a,b int)(int,int){
	return a+b,a-b		//can return two values
}

func div(a,b int)(r1 int){
	r1 = a/b
	return 
}
func main(){
	a:=500
	b:=5
	result:=mul(a,b)
	fmt.Println(result)

	result1,result2 := cal(a,b)
	fmt.Println(result1,result2)

	r1:=div(a,b)
	fmt.Println(r1)

}