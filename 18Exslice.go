
Exercises consist of Slices, map, Struct and Nested struct
1.Create a slice of integers, append values, and print its length & capacity.
2.Create a map of student names and grades, update values, and check if a student exists.
3.Define a struct Book with fields Title, Author, Price, and create a few book instances.
4.Create a nested struct Employee with a field Department (which is another struct).

1....//
package main

import "fmt"

func main(){

	 number := []int{1,2,3,34,4,5,6,7,5}
	 fmt.Println(number)

	number = append(number,3,4,5,6,7)
	fmt.Println("Length:",len(number))
	fmt.Println("Capacity:",cap(number))
}
//

Explanation
✅ number := []int{...} → Creates a slice with initial values.
✅ append(number, 3, 4, 5, 6, 7) → Adds new elements.
✅ len(number) → Returns the number of elements.
✅ cap(number) → Returns the capacity (which may increase dynamically).

2.Create a map of student names and grades, update values, and check if a student exists.

func main(){

	student := map[string]string{
		"Abhi":"A",
		"Raj":"B",
	}

	student["Raj"]="C"
	fmt.Println(student["Raj"])

	name,exists := student["Raj"]
	fmt.Println("Student", name, "Exists?",exists)
}

✅ Key Fixes:

student["Raj"] = "C" (Correct way to update a map key).
student["Raj"] (Correct way to access map values).
value, exists := student["Raj"] (Correct way to check if a key exists).

//

3.Define a struct Book with fields Title, Author, Price, and create a few book instances.

type Book struct{
	Title string
	Author string
	Price int
}
func main(){

	bookone := Book{
		Title:"Aliceborder",
		Author:"James",
		Price: 344,
	}

	booktwo := Book{
		Title:"Jungle book",
		Author:"Bond",
		Price: 3434,
	}

	fmt.Println(bookone)
	fmt.Println(booktwo)

		fmt.Println("Title of first book:", bookone.Title)
	fmt.Println("Price of second book:", booktwo.Price)
}

//
4.Create a nested struct Employee with a field Department (which is another struct).
type Employee struct{
	Department string
}

type Company struct{
	Name string
	Role string
	salary int
	Department Employee
}

func main(){

	emp1 := Company{
		Name : "Abhishek",
		Role : "software developer",
		salary: 12000,
		Department:Employee{
			Department : "IT",
		},
	}
	fmt.Println(emp1)
	fmt.Println(emp1.Name)
	fmt.Println(emp1.Department)
	fmt.Println(emp1.Department.Department)
	
}