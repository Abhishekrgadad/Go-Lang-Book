package main

import "fmt"


// Basic function
func Hello(){
    fmt.Println("Hi Hello how are you")
}

func main(){
    Hello()
}

✅ Key Takeaways:

Functions in Go start with the func keyword.
greet() is a function without parameters and without return values.

// Function with Parameters

func Add(a int , b int){
    sum := a + b
    fmt.Println(sum)
}

func main(){
    Add(3,4)
}

// Function with Return Values

func mult(a int, b int) int {
    return a * b
}

func main(){
    product := mult(3,4)
    fmt.Println("Product is ", product)
}
✅ Key Takeaways:

The function returns an int, so the return type is specified after parameters (int).
The return statement returns a value from the function.

// Function with Multiple Return Values

func Divide(a int , b int ) (int,int){
    quotient := a / b 
    remainder := a % b 
    return quotient,remainder
}
func main(){
    q,r := Divide(10,3)
    fmt.Println("Quotient",q, "Remainder",r)
}

✅ Key Takeaways:

Return multiple values by specifying multiple types in parentheses ((int, int)).
Store returned values in multiple variables (q, r := divide(10, 3)).


// Named Return Values

func Rect(length, width int) (area,perimeter int){
    area = length * width
    perimeter = 2 * (length + width)
    return
}
func main(){
    a,p := Rect(3,4)
    fmt.Println("Area:",a , "Perimeter:", p)

}

✅ Key Takeaways:

Return variables are declared in the function signature ((area int, perimeter int)).
The return statement does not need explicit values

// Variadic Function (Multiple Arguments)

func sum(num ...int) int {
    total := 0
    for _, num := range num {
        total += num 
    }
    return total
}
func main(){
    fmt.Println(sum(1,2,3,4))
    fmt.Println(sum(10,20,30,40))
}

✅ Key Takeaways:

Use ... to accept multiple arguments (numbers ...int).
Inside the function, numbers behaves like a slice.

// Deferred Function 
The defer keyword delays function execution until the surrounding function returns.

func main(){
    fmt.Println("Function starts here")
    defer fmt.Println("Function ends here")
    fmt.Println("Function is still executing")
}

✅ Key Takeaways:

defer executes after all other function statements.
Often used for cleaning up resources (e.g., closing files).

// Error Handling in Go

panic() → Stops normal execution immediately.
recover() → Catches and handles a panic.

func safediv(a,b int){
    defer func(){
        if r := recover(); r != nil{
            fmt.Println("Recovered from panic",r)
        }
    }()

    if b == 0{
        panic("Division by zero")
    }
}

func main(){
    safediv(10,2)
    safediv(5,0)
}

✅ Key Takeaways:
panic("error message") stops execution.
recover() catches the panic to prevent crashing.


Exercises
1. Write a function subtract(a, b int) int that returns the difference of two numbers.
2. Modify subtract() to return both the difference and absolute difference.
3. Write a variadic function that takes multiple numbers and returns their average.
4. Write a function that returns an error if division by zero occurs, instead of using panic().

1. subtract(a , b int)int
2. Modify subtract() to return both the difference and absolute difference.

import "math"
func subtract(a,b int) (int, float64){
    diff := a - b
    absdiff := math.Abs(float64(a-b))
    return diff,absdiff
}
func main(){
    d, ad := subtract(10,12)
    fmt.Println(d,ad)
    
}

3.Write a variadic function that takes multiple numbers and returns their average.

func avg(num ...float64) float64 {
    if len(num) == 0 {
        return 0
    }

    total := 0.0
    for _, num := range num {
        total += num
    }
    return total / float64(len(num))

}

func main(){
    fmt.Println("Average:", avg(10,28,30))
    fmt.Println("Average", avg(5,10,15,30))
}
✅ Explanation
...float64 → Allows passing multiple floating-point numbers.
Loop through numbers → Calculate the total sum.
Divide sum by length → Compute the average.
Handles empty input (average() returns 0 instead of dividing by zero).

4. Write a function that returns an error if division by zero occurs, instead of using panic().

package main

import (
	"errors"
	
)

// Function to safely divide two numbers
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed") // ✅ Return an error instead of panic
	}
	return a / b, nil // ✅ No error, return result
}

func main() {
	// Test cases
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = safeDivide(5, 0) // Division by zero
	if err != nil {
		fmt.Println("Error:", err) // ✅ Handles the error properly
	} else {
		fmt.Println("Result:", result)
	}
}

✅ Explanation
safeDivide(a, b float64) (float64, error)

Returns both the result (float64) and an error (error).
errors.New("division by zero is not allowed")

If b == 0, it returns an error instead of crashing the program.
Checking for errors in main()

If err != nil, print the error message.
Otherwise, print the result.
