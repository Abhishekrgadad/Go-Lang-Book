Composite types in Go 
1.Arrays (Fixed-size collections)
2.Slices (Dynamic-size collections)
3.Maps (Key-value pairs)
4.Structs (Custom data structures)

Key points:
Arrays have Fixed length of collection of same type of elements.
Arrays cannot grow or shrink
Unassigned values are set to 0 by default

//
package main

import "fmt"

func main(){
	var num [5] int //declare
	num [0] = 10
	num [1] = 20
	fmt.Println("Array no.", num)
	fmt.Println("First element", num[0])
}
//

Best efficient way of using Arrays are 

//
func main(){

	num := [5]int {1,2,3,4,5}
	fmt.Println("First element: ", num[0])
	fmt.Println("second no:",num[1])
	fmt.Println(num)
}
//