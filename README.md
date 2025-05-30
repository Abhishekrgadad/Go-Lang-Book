# 📘 Go Programming:

Welcome to the **Go-Lang-Book** repository!  
This documentation is inspired by _Go Programming_ by Alan A. A. Donovan & Brian W. Kernighan.

---

## Table of Contents

1. [Introduction to Go](#introduction-to-go)
2. [Setting Up Go](#setting-up-go)
3. [Go Basics](#go-basics)
    - [Go Modules and Dependency Management](#go-modules--dependency-management)
        - [Initialize a Module](#initializing-a-go-module)
    - [Hello, World!](#hello-world)
    - [Variables & Types](#variables--types)
    - [Control Structures](#control-structures)
    - [Functions](#functions)
    - [Methods](#methods)
    - [Packages & Imports](#packages--imports)
4. [Intermediate Go](#intermediate-go)
    - [Arrays, Slices & Maps](#arrays-slices--maps)
    - [Structs & Methods](#structs--methods)
    - [Interfaces](#interfaces)
    - [Error Handling](#error-handling)
5. [Advanced Go](#advanced-go)
    - [Concurrency (Goroutines & Channels)](#concurrency-goroutines--channels)
    - [Testing in Go](#testing-in-go)
    - [Go Modules & Dependency Management](#go-modules--dependency-management)
    - [Reflection](#reflection)
6. [Best Practices](#best-practices)
7. [Resources](#resources)
8. [Contact](#contact)

---

## Introduction to Go

Go (or Golang) is an open-source programming language designed for simplicity, reliability, and efficiency.  
It is ideal for building scalable and high-performance applications.

---

## Setting Up Go

- Download and install Go from the [official website](https://golang.org/dl/).
- Set up your `GOPATH` and workspace.
- Verify installation:
    ```sh
    go version
    ```

---

## Why Choose Go?

Go (Golang) stands out among modern programming languages for several reasons:

### 🚀 Key Features of Go
- **Simplicity:** Clean and concise syntax makes code easy to read and maintain.
- **Performance:** Compiled to machine code, Go programs run fast and efficiently.
- **Concurrency:** Built-in support for concurrent programming with goroutines and channels.
- **Strong Standard Library:** Rich set of libraries for networking, web servers, testing, and more.
- **Cross-Platform:** Easily build binaries for multiple operating systems from a single codebase.
- **Garbage Collection:** Automatic memory management for safer and more reliable code.
- **Static Typing:** Catches many errors at compile time, improving code quality.
- **Fast Compilation:** Extremely quick build times, even for large projects.

### 🌟 Why Go is Outstanding Compared to Other Languages
- **Concurrency Model:** Go's goroutines and channels make concurrent programming simpler and safer than traditional threading models in languages like Java or C++.
- **Deployment:** Go produces single, statically-linked binaries, making deployment and distribution straightforward—no dependency hell.
- **Developer Productivity:** Minimalistic design and tooling (like `gofmt`) encourage best practices and reduce time spent on configuration.
- **Scalability:** Widely used for building scalable cloud services, microservices, and distributed systems (e.g., Docker, Kubernetes are written in Go).
- **Community & Ecosystem:** Strong, active community and growing ecosystem of libraries and frameworks.

Go is an excellent choice for backend development, cloud-native applications, DevOps tools, and any scenario where performance, reliability, and simplicity are priorities.

---

## Go Basics

### Go Modules & Dependency Management

Go modules are the standard way to manage dependencies in modern Go projects. Modules make it easy to track and control the external packages your project depends on.

### Initializing a Go Module

To initialize a new module in your project directory, run:

```sh
go mod init <module-name>
```

This command creates two important files:

- **go.mod:**
  - Declares the module path (your project name or repository URL).
  - Lists the Go version and all direct dependencies required by your project.
  - Example content:
    ```
    module github.com/username/projectname

    go 1.21
    ```

- **go.sum:**
  - Contains cryptographic checksums for each dependency to ensure integrity and security.
  - Automatically updated when you add, update, or remove dependencies.

With modules, you can easily add dependencies using:

```sh
go get <package-path>
```

To remove the unnecessary imported packages use:
```sh
go mod tidy
```
And build reproducible, portable Go projects with confidence.

---

### To run Go files,

```go
go run <filename>.go
```

### Hello, World!

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Variables & Types

- Declaring variables:
    ```go
    var age int = 30
    name := "Alice"
    ```

- Basic types: `int`, `float64`, `string`, `bool`

### Control Structures

- **If/Else:**
    ```go
    if age > 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }
    ```
- **For Loop:**
    ```go
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    ```
- **Range Loop:**
    ```go
    nums := []int{1, 2, 3}
    for idx, val := range nums {
        fmt.Printf("Index: %d, Value: %d\n", idx, val)
    }
    ```

### Functions

A function in Go is defined using the `func` keyword. The general syntax is:

```go
func functionName(parameters) returnType {
    // function body
}
```

#### Function Signatures & Types

| Type                              | Signature Example                                 | Description                                 |
|-----------------------------------|---------------------------------------------------|---------------------------------------------|
| Simple Function                   | `func greet() { ... }`                            | No parameters, no return value              |
| Function with Parameters          | `func add(a int, b int) { ... }`                  | Takes parameters, no return value           |
| Function with Return Value        | `func getAge() int { ... }`                       | Returns a value                             |
| Function with Parameters & Return | `func sum(a int, b int) int { ... }`              | Takes parameters and returns a value        |

**Examples:**

- Simple Function:
    ```go
    func greet() {
        fmt.Println("Hello!")
    }
    ```
- Function with Parameters:
    ```go
    func add(a int, b int) {
        fmt.Println(a + b)
    }
    ```
- Function with Return Value:
    ```go
    func getAge() int {
        return 30
    }
    ```
- Function with Parameters & Return Value:
    ```go
    func sum(a int, b int) int {
        return a + b
    }
    ```
- Function with Parameters and Named returns
    ```go
    func product(a,b int) (result int) {
        return result
    }
    ```
---
### Methods
Method is a function with Special receiver argument. Methods allow us to associate function with specific types.
Receiver: The variable to which the method is bound.It can be a value or pointer.
Type: The type the method is associated with (usually a struct).

**Syntax**
```go
func (receiver Type) MethodName(params) returnType {
    // method body
}
```

**Examples**
```go
package main
import "fmt"

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    fmt.Println("Original Rectangle:", rect)
    fmt.Println("Area:", rect.Area())

    rect.Scale(2)
    fmt.Println("Scaled Rectangle:", rect)
    fmt.Println("New Area:", rect.Area())
}
```

**Pass by Value:**
The function argumnets are copied and passed to other methods. if any changes in the other functions won't affect the original values
```go
func increment(x int) {
    x = x + 1
}

func main() {
    a := 10 
    increment(a)
    fmt.Println(a) // Output: 10
}
```

**Pass by Reference:**
The function which pass the address or memory location for the other functions using pointers. if there any change in other functions will affect the original values.
```go
func incrementPtr(x *int) {
    *x = *x + 1
}

func main() {
    a := 10
    incrementPtr(&a)
    fmt.Println(a) // Output: 11
}
```

**Function Chaining:**
Function chaining is a technique where multiple method calls are linked together in a single statement.in other sense the output of the one function is taken as input to next function.
```go
package main

import (
    "fmt"
)

type Counter struct {
    value int
}

func (c *Counter) Add(n int) *Counter {
    c.value += n
    return c
}

func (c *Counter) Multiply(n int) *Counter {
    c.value *= n
    return c
}

func main() {
    c := &Counter{}
    c.Add(5).Multiply(3).Add(2)
    fmt.Println(c.value) // Output: 17
}
```
**Defer Function:**
The defer statement in go is used to schedule a function call to be run after the function completes,just before it returns. Deferred functions are often used for cleanup tasks, such as closing files or releasing resources.
```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Start")
    defer fmt.Println("Deferred: This runs last")
    fmt.Println("End")
}
```
Also in such cases:
```go
file, err := os.Open("data.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```
---

## Intermediate Go

### Arrays, Slices & Maps

| Type   | Description                                      |
|--------|--------------------------------------------------|
| Array  | Fixed-size collection of elements of same type.   |
| Slice  | Dynamic-size, flexible view into arrays.          |
| Map    | Unordered collection of key-value pairs.          |

### Structs & Methods
A structure in Go is a composite data type that groups together zero or more fields with varying data types under a single name.
- Define a struct 
```go
type Person struct {
    Name string
    Age  int
}
```
- creating and using structs
```go
func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Name) // Output: Alice
    fmt.Println(p.Age)  // Output: 30

    p.Age = 31
    fmt.Println(p.Age) // Output: 31
}
```
**Custom types:**
You can define your own custom types using the type keyword.Custom types help make your code more readable,type-safe and organized.
```go
type Age int
var myAge Age = 30

type Person struct {
    Name string
    Age  int
}

var p Person = Person{Name: "Alice", Age: 25}
```

### Interfaces

An interface in Go is a type that specifies set of method signatures(behavior).but does not provide implementations. Any type that implements those methods satisfies the interface no explicit declaration is needed.
```go
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    rect := Rectangle{Width: 4, Height: 3}
    printArea(rect)
}
```
**Empty Interface:**
The empty interface (inteface{}) in Go can hold values of any type because it has zero methods. it is used when you need a variable,function parameter,or return value that can accept any type.
```go
func describe(i interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", i, i)
}

describe(42)        
describe("hello")    
describe([]int{1,2}) 
```

### Error Handling

- Idiomatic error handling using `error` type.

---

## Advanced Go

### Concurrency (Goroutines & Channels)

- Lightweight threads and communication.

### Testing in Go

- Write tests using the `testing` package.

## Best Practices

- Use `gofmt` for formatting.
- Write clear, concise code.
- Handle errors gracefully.

---

## Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)

---

## Contact

For questions or suggestions, feel free to email me!

---
