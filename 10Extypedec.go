package main

import "fmt"

// 1.Define a custom type Temperature based on float64 and write a function that converts Celsius to Fahrenheit.
type Temperature float64

func toFahrenheit(t Temperature) float64{
	return (float64(t) * 9/5) + 32
}
func main(){
	var temp Temperature = 25
	fmt.Println("Tempearture in Fahrenhiet", toFahrenheit(temp))

}

// 2.Create a type alias for string and demonstrate that it behaves exactly like string.

type Temp = string

func main(){
	var a Temp = "hello"
	var b string = "hi"
	fmt.Println(a,b)
}

//3.Try assigning an int value to a Kilometers variable without conversion. What error do you get?

package main

import "fmt"

type Kilometers float64

func main() {
	var distance Kilometers
	distance = 10 // ❌ ERROR: cannot use 10 (type int) as type Kilometers

	fmt.Println("Distance:", distance)
}

