package main

import (
	"fmt"
	"math"
)

Methods in Go
A method is a function that is associated with a type (usually a struct).
It allows object-oriented-like behavior in Go.

Define a method
Methods are defined using a receiver (a variable before func).
The receiver associates the method with a specific type.
Every methods needs a custom type to perform

type Rectangle struct{
	width, height float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

func main(){
	rect := Rectangle{width:10, height:5}
	fmt.Println("Rectangel:",rect.Area())
}

✅ Key Takeaways:

(r Rectangle) → Receiver binds the method to Rectangle.
rect.Area() → Calls the method like an object-oriented function.

Methods with Pointers

type Counter struct{
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func main(){
	count := Counter{value:10}
	count.Increment();
	fmt.Println(count.value)

}

Methods on Custom Types

type Distance float64

func (d Distance) MtoK() float64{
	return float64(d) / 1000
}

func main(){
	dist := Distance(5000)
	fmt.Println("Distance",dist.MtoK())
}

Methods Chaining

type Counter struct{
	value int
}

func (c *Counter) Add(n int) *Counter {
	c.value += n
	return c
}

func main(){
	count := Counter{value:10}
	count.Add(10).Add(5)
	fmt.Println("Final value:", count.value)
}

Exercise:
1. Create a struct Circle with a method to calculate area.
2. Create a method IncreasePrice() for a Product struct that increases price.
3. Write a method on a custom type Celsius to convert to Fahrenheit.

Que 1 solution.
type Circle struct{
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main(){
	fmt.Println("Area of circle")
	rad := Circle{10}

	fmt.Println("Radius", rad.Area())
}

2. Create a method IncreasePrice() for a Product struct that increases price.

type Product struct{
	price int
}

func (p *Product) IncreasePrice() int{
		  p.price++
		 return p.price

}

func main(){
	pri := Product{100}
	fmt.Println("Price is:",pri.IncreasePrice())
}

3. Write a method on a custom type Celsius to convert to Fahrenheit.

type Celsius float64

func (c Celsius) ctf() float64{
	return float64(c)*9/5 + 32
}

func main(){
     var temp Celsius = 24

	 fahren := temp.ctf()
	 fmt.Println(fahren)
}



