Slices in Go:

slice is dynamic collection (more flexible than arrays)
Slices can grow and shrink.


Key points:
*Slices are more flexible than arrays.
*Use append() to add elements dynamically.
*len(slice) returns length.
*cap(slice) returns capacity (internal array size).
*Slices are references → Modifying a slice modifies the original data.
*Sub-slicing is done using [start:end] (half-open range).

// //
package main

import "fmt"

func main(){

	num := []int {1,2,3,4,4,5,6,7,8}
	fmt.Println(num)
	fmt.Println("Length",len(num))
	fmt.Println("capacity",cap(num))
	num = append(num, 1,1,1)
	fmt.Println("Append values:",num)
	fmt.Println("Length:",len(num))
	fmt.Println("Capacity:",cap(num))
}
//

//Slice opeartions (Sub slicing)

func main(){

	num := []int {1,2,3,4,5,6,7,8,9}

	//Create sub-slices
	type1 := num[:2] //first two no.
	type2 := num[3:6] //from index 3 to 6
	type3 := num[4:] //start from 4 to end

	fmt.Println(type1)
	fmt.Println(type2)
	fmt.Println(type3)
}

