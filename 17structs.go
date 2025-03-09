// A struct is a custom data type that groups multiple fields.

// Keys points:
// Structs group multiple fields (like objects in OOP).
// Access fields using . notation (p1.Name).

package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main(){
	//create struct instance
	p1 := Person{Name: "Alice", Age:"25"}
	p2 := Person{"Bob", 30}
	p3 := Person{"Charlie", 44}

	fmt.Println("First Person:",p1)
	fmt.Println("Second Persion:",p2)
	fmt.Println("Third Person:",p3)
}

//Nested Structs

type Address struct {
	City, Country string
}

type Person struct {
	Name string
	Age int
	Address Address //nested struct
}

func main(){
	per := Person{
		Name : "Abhi",
		Age : 20,
		Address: Address{
			City: "banglore",
			Country: "India",
		},
	}

	fmt.Println("Person:", per)
	fmt.Println("Person with name",per.Name)
	fmt.Println("City:", per.Address.City)
}

✅ Key Takeaways:
Structs can be nested inside other structs.
Access fields using dot notation (p.Address.City).