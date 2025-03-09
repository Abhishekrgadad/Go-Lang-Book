Scope defines where a variable, function, or type can be accessed in a Go program.
Go follows a lexical scoping rule, meaning that a variable is accessible only within the block where it is declared.

//Variable Scope
package main

import "fmt"

func main() {
	message := "Hello, Go!" // Local variable
	fmt.Println(message)
}

// fmt.Println(message) // ❌ ERROR: message is not accessible here

//Package level scope
package main

import "fmt"

var globalMessage = "Accessible everywhere in this package" // Package-level variable

func printMessage() {
	fmt.Println(globalMessage) // ✅ Accessible here
}

func main() {
	printMessage()
}

//Exported Identifiers

Functions name starts with Uppercase letter are public to use in other files.

//Unexported Identifiers

Functions names starts with small letter are private to the package cant able to access in other files or packages.

//Nested Scopes in Go
Go follows block-level scoping, meaning inner blocks can access outer variables.

package main

import "fmt"

func main() {
	x := 10 // Outer scope

	if x > 5 {
		y := 20 // Inner scope
		fmt.Println("x:", x) // ✅ Accessible
		fmt.Println("y:", y) // ✅ Accessible
	}

	// fmt.Println(y) // ❌ ERROR: y is not accessible here
}

Exercises:
1. Create a package named mathutils with:
	An exported function Square(n int) int
	An unexported function secretCalculation()
2. Try accessing secretCalculation() from main.go. What happens?
3. Modify secretCalculation() to start with an uppercase letter and try again.