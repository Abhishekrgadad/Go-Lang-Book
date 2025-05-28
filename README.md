# 📘 Go Programming:

Welcome to the **Go-Lang-Book** repository!  
This documentation is inspired by _Go Programming_ by Alan A. A. Donovan & Brian W. Kernighan.

---

## Table of Contents

1. [Introduction to Go](#introduction-to-go)
2. [Setting Up Go](#setting-up-go)
3. [Go Basics](#go-basics)
    - [Hello, World!](#hello-world)
    - [Variables & Types](#variables--types)
    - [Control Structures](#control-structures)
    - [Functions](#functions)
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

## Go Basics

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

### Packages & Imports

- Organize code using packages.
- Import standard or custom packages:
    ```go
    import "fmt"
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

- Define custom types and attach methods.

### Interfaces

- Achieve polymorphism.

### Error Handling

- Idiomatic error handling using `error` type.

---

## Advanced Go

### Concurrency (Goroutines & Channels)

- Lightweight threads and communication.

### Testing in Go

- Write tests using the `testing` package.

### Go Modules & Dependency Management

- Manage dependencies with Go modules.

### Reflection

- Inspect types at runtime.

---

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
