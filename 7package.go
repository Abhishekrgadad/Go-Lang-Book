package main //Package name

import (
	"fmt"
	"app/pack"
) //importing other packages

func main(){

	fmt.Println("Hello,Go")
}

// //Separate import Statements
// import "fmt"
// import "math"

// //Grouped Import Statements
// import (
// 	"fmt"
// 	"math"
// )

//For Example i created a package named pack with pack.go file with hello function.
//Then calling that function in this main file.

//import related custom package with root name.

func main(){
	fmt.Println("main started")
	pack.Hello("Abhishek")
}

Each package have an init() function that runs automatically before main()
this is useful for initialize settings before the program starts.

Exercise:
//1.  Create a custom package named mathutils with a function that calculates the square of a number.
//2.Modify mathutils to have an init() function that prints "Math package initialized".
//3. Import and use mathutils in main.go.

I have done these Exercise in file 
Expack with Expack.go file and main function is in this file.
