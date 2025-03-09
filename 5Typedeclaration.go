package main

import "fmt"

type name string

//Custom types
func main(){
	var myname name = "Abhishek"
	fmt.Println(myname)
	
}

//Custom types in Functions

type kilometer float64

func distinmiles(km kilometer) float64 {
	return float64(km) * 6.4654
}
func main(){
	var trip kilometer = 43
	distance := distinmiles(trip)
	fmt.Println(distance)
}

//Aliasing Types VS Defining New Types

type Id int // strictly different type

type Myint = int //just another name for int

func main(){

	var a id = 123
	var b Myint = 345
	var c int = 344

	fmt.Println(a,b,c)

	// b = c //allowed as same type
	// a = c //strictly new type
}

