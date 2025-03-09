package main
//Throught assignment we can update the values.
import "fmt"

func main(){
	var a = 10
	fmt.Println("originla value", a)

	a = 20
	fmt.Println("modified value ", a)

	a, b := 10, 30 //multiple assignment
	fmt.Println(a,b)

	a, b = b , a //swapping in go
	fmt.Println(a,b)

}
