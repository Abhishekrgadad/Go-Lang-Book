Concurrency

Concurrency allows multiple tasks to run independently.
Go provides goroutines and channels for concurrency.

A goroutine is a lightweight thread managed by Go.
Use go before a function to run it concurrently.

package main

import (
	"fmt"
	"time"
)

func sayhello(){
	fmt.Println("Hello from Goroutine")
}

func main(){
	go sayhello()
	time.Sleep(time.Second)
	fmt.Println("Main function finished")
}

✅ Key Takeaways:
go sayHello() starts a new goroutine.
main() does not wait for goroutines unless you add a delay (time.Sleep).

Go routine Anonymous Goroutine--------------------------------------

func main(){
	go func(){
		fmt.Println("hello from anonymous Goroutine")
	}()

	time.Sleep(time.Second)
	fmt.Println("Main function completed")
}

Key Takeaways
go func() { ... }() runs a function without a name as a goroutine.

