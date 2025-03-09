package main

//Go doesnot allow implicit type conversion. 
//we must work on Explicit convert types.
import "fmt"

func main(){
	var a int = 10
	var b float64 = float64(a)
	fmt.Println(b)
}

//Use Explicit conversion like 
	//Float64(i)
	//int(i)
	//bool(i) etc


//In this type conversion Constants are Untyped they will take the Type of the Variables with they work with.
