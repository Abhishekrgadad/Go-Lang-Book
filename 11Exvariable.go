package main

import "fmt"

one := "abhi"
The above statement is not allowed in Package level scope.



func main(){
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func main(){
	const two = "hello"
	fmt.Println(two)
	two = "newhell"
	fmt.Println(two)
}