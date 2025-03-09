// Consider an Example of converting Boiling water from Fahrenheit to Celsius.
package main

import "fmt"

const bolingF = 212
func main(){
 
	var f = bolingF
	var c = (f - 32)* 5/9
	fmt.Printf("boling point = %dF or %dC\n",f,c)
}

//Some values are declared in Package level using Const 
//Some values are declared under function level scope using var 
//or you can use short variable declaration as c := (something)