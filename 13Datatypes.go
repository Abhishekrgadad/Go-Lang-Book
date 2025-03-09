Integer types (int, uint, byte, rune)
Floating-point types (float32, float64)
Complex numbers (complex64, complex128)
Boolean type (bool)
String type (string)

Type	Size (bits)	Signed?	Range (Approx.)
int8		8			✅ Yes	-128 to 127
uint8		8			❌ No	0 to 255 (alias: byte)
int16		16			✅ Yes	-32,768 to 32,767
uint16		16			❌ No	0 to 65,535
int32		32			✅ Yes	-2,147,483,648 to 2,147,483,647 (alias: rune)
uint32		32			❌ No	0 to 4,294,967,295
int64		64			✅ Yes	Very large numbers
uint64		64			❌ No	Only positive large numbers

package main

import "fmt"

//Integer Variables
func main() {
	var a int = -10
	var b uint = 20
	var c byte = 255 // byte is an alias for uint8

	fmt.Println("Signed int:", a)
	fmt.Println("Unsigned int:", b)
	fmt.Println("Byte:", c)
}

//Floating Point Types
package main

import "fmt"

func main() {
	var pi float64 = 3.141592653589793
	var e float32 = 2.71828

	fmt.Println("Pi:", pi)
	fmt.Println("Euler's Number:", e)
}

//Boolean Types
package main

import "fmt"

func main() {
	var isGoFun bool = true
	isLearning := false // Short declaration

	fmt.Println("Is Go fun?", isGoFun)
	fmt.Println("Am I learning?", isLearning)
}

//String Types

package main

import "fmt"

func main() {
	var name string = "Alice"
	message := "Hello, " + name

	fmt.Println(message)
}

//Multi line strings using Backticks(``)
package main

import "fmt"

func main() {
	text := `This is
a multi-line
string!`
	fmt.Println(text)
}

//String length and Indexing
package main

import "fmt"

func main() {
	word := "Golang"
	fmt.Println("Length:", len(word))
	fmt.Println("First character:", string(word[0])) // Convert to string
}

//Unicode and Runes
package main

import "fmt"

func main() {
	var ch rune = 'G' // Single quotes for runes
	fmt.Println("Rune:", ch) // Prints ASCII/Unicode code
	fmt.Println("Character:", string(ch)) // Converts rune to string
}



