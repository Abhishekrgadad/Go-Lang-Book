package main

import "fmt"

Exercies :

1. Write a program that declares an int, float64, and string, then tries assigning an int to a float64 without conversion. What error do you get?
func main(){
	var a int = 10
	var c float64 = 3.14

	a = int(c)
	fmt.Println(a)

}

2.Modify the program to correctly convert int to float64 before assignment.
Exact same code as above

3. Try swapping two numbers without using a temporary variable.

func main(){
	var a int = 10
	var b int = 20 
	fmt.Println("Swap of two no.")
	a, b = b , a 
	fmt.Println(a,b)
}