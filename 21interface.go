//Interfaces
// An interface is a type that specifies a set of method signatures.
// A type satisfies an interface if it implements all the methods in that interface.

//Defining and using of interface--------------------

package main

import "fmt"

type Shape interface{
	Area() float64
}

type Circle struct{
	radius float64
}

type Rectangle struct{
	length float64
	width float64

}

type Cube struct{
	side float64
}

type Cuboid struct {
	length , width, height float64
}

func (c Circle) Area() float64{
	return 3.14 * c.radius* c.radius
}

func (r Rectangle) Area() float64{
	return 2 * r.length * r.width
}

func (cu Cube) Area() float64{
	return 6 * cu.side * cu.side
}

func (cd Cuboid) Area() float64{
	return 2 * (cd.length * cd.width * cd.height)
}


func main(){
	var s Shape;
	s = Circle{10}
	fmt.Println("Area of Circle:",s.Area())

	s = Rectangle{length: 13, width:40}
	fmt.Println("Area of Reactangle:",s.Area())

	s = Cube{20}
	fmt.Println("Area of Cube: ",s.Area())

	s = Cuboid{10,20,30}
	fmt.Println("Area of Cuboid:",s.Area())
}


✅ Key Takeaways:
Shape is an interface with a single method Area().
Circle and Rectangle implement Shape by defining Area().
The interface variable s can hold any type that implements Shape.

//Empty Interface -------------------------------------
The empty interface interface{} can hold any type (like any in other languages).

func printValue(value interface{}){
	fmt.Println("Value", value)
}

func main(){
	printValue(34)
	printValue(44)
	printValue(3.14)
	printValue("hello")
}

✅ Key Takeaways:

interface{} allows storing any type.
Useful for generic functions.

Type Assertions ---------------------------------
Type assertions extract a concrete type from an interface{}.

func main(){
	var x interface{} = "Hello, GO!"

	str, ok := x.(string)
	if ok{
		fmt.Println("String:", str)
	} else{
		fmt.Println("Not a string")
	}
}

✅ Key Takeaways:
x.(string) extracts a string from x if x actually holds a string.
The ok value prevents a runtime panic if the assertion fails.

//Type Switches -------------------------------------

func identifyType(x interface{}){
	switch v := x.(type) {
	case int: 
		fmt.Println("Integer",v)
	case string:
		fmt.Println("String:",v)
	case float64:
		fmt.Println("Float:",v)
	default:
		fmt.Println("Unknown Type")
	}
}

func main(){
	identifyType(20)
	identifyType("Hello")
	identifyType(3.23)
	identifyType([]int{12,3,4})
}

✅ Key Takeaways:
switch x.(type) determines the underlying type at runtime.
Useful when handling different data types dynamically.

//Interface Embedding (Composition)--------------------------------

Interfaces can be combined using embedding.
A new interface can inherit methods from multiple interfaces.

type Printer interface{
	Print()
}

type Scanner interface{
	Scan()
}

type Multifuction interface{
	Printer
	Scanner
}

type Machine struct {}

func (m Machine) Print(){
	fmt.Println("Printing.......")
}

func (m Machine) Scan(){
	fmt.Println("Scanning.....")
}

func main(){
	var m Multifuction = Machine{}
	m.Print()
	m.Scan()
}

✅ Key Takeaways:
MultiFunctionDevice embeds Printer and Scanner interfaces.
Machine implements both interfaces, so it satisfies MultiFunctionDevice.

Exercise:
1. Create an interface Animal with a method Speak() and implement it for Dog and Cat.
2. Write a function that takes an interface{} and prints its type using switch x.(type).
3. Modify the Shape interface to include a new method Perimeter(), and implement it for Circle and Rectangle.

// 1. Que

type Animal interface{
	Speak()
}

type Dog struct {}

type Cat struct {}

func (d Dog) Speak(){
	fmt.Println("Dog Barks")
}

func (c Cat) Speak(){
	fmt.Println("Cat Meow")
}

func main(){
	var a Animal
	a = Dog{}
	a.Speak()
	a = Cat{}
	a.Speak()
}

// 2.Write a function that takes an interface{} and prints its type using switch x.(type).

func identifyType(x interface{}){
	switch a := x.(type){
	case int:
		fmt.Println("Integer:",a)
	case float64:
		fmt.Println("Float64:",a)
	case string:
		fmt.Println("String:",a)
	case bool:
		fmt.Println("Boolean",a)
	default:
		fmt.Println("Unknow type")
	}
}

func main(){
	identifyType(22)
	identifyType(2.2)
	identifyType("hello")
	identifyType(true)
	identifyType([]int{12,3})
}

// 3.Modify the Shape interface to include a new method Perimeter(), and implement it for Circle and Rectangle.

type Shape interface{
	Perimeter() float64
}

type Circle struct{
	radius float64
}

type Rectangle struct{
	length float64
	width float64
}

func (c Circle) Perimeter() float64{
	return 2 * 3.14 * c.radius
}

func (r Rectangle) Perimeter() float64{
	return 2 *(r.length + r.width)
}

func main(){
	var s Shape
	s = Circle{20}
	fmt.Printf("Permeter of circle: %.3f\n",s.Perimeter())

	s = Rectangle{10,20}
	fmt.Println("Perimeter of Rectangle:",s.Perimeter())
}
