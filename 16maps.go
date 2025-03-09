// //Maps in Go 
// A map is a collection of key-value pairs, similar to dictionaries in Python.

//Key points:
*Maps store key-value pairs.
*delete(map, key) removes a key.
*Checking existence: val, exists := map[key].

//
package main

import "fmt"

func main(){
	//create a map
    num := map[string]int{
        "Alice":20,
        "bob":40,
    }

    //adding new values
    num["charlie"] = 45

    //delete a key
    delete(num,"bob")

    //checking if key exists
    age,exists := num["bob"]
    fmt.Println("bob age:", age, "Exists?", exists)
}
//