//Channels in GO (chan keyword)

// A channel allows goroutines to communicate safely.
// Use make(chan Type) to create a channel.
// Use <- to send and receive data.

package main

import "fmt"

// func main(){
// 	messages := make(chan string)  //creates a channel

// 	go func(){
// 		messages <- "Hello, Goroutine!" //send msg to channel
// 	}()

// 	msg := <- messages //receive data from channel
// 	fmt.Println(msg)

// }
// ✅ Key Takeaways:

// messages := make(chan string) → Creates a channel for strings.
// messages <- "Hello" → Sends data to the channel.
// msg := <-messages → Receives data from the channel.

//Buffered Channels
// A buffered channel allows storing multiple values before blocking.
// make(chan Type, capacity) creates a buffered channel.

// func main(){
//     messages := make(chan string,2)

//     messages <- "message 1"
//     messages <- "message 2"

//     fmt.Println(<-messages)
//     fmt.Println(<-messages)
// }

// ✅ Key Takeaways:
// Buffered channels store values without blocking until full.
// Receiving from a channel removes the value.

//Channel Synchronization (using sync.waitGroup)
// The sync.WaitGroup helps wait for goroutines to finish.

// import "sync"

// func worker(id int, wg *sync.WaitGroup){
//     defer wg.Done()
//     fmt.Println("Worker", id, "is done")
// }

// func main(){
//     var wg sync.WaitGroup

//     for i := 1; i <= 3; i++ {
//         wg.Add(1)
//         go worker(i, &wg)
//     }

//     wg.Wait()
//     fmt.Println("All workers completed")
// }


//Select Statement (Multiplexing Channels)

// The select statement allows waiting for multiple channels.

// import "time"

// func main(){
//     ch1 := make(chan string)
//     ch2 := make(chan string)

//     go func(){
//         time.Sleep(time.Second)
//         ch1 <- "Message from ch1"
//     }()

//     go func(){
//         time.Sleep(time.Second)
//         ch2 <- "Message from ch2"
//     }()

//     select {
//     case msg := <-ch1:
//         fmt.Println(msg)
    
//     case msg := <-ch2:
//         fmt.Println(msg)
//     }

// }

// ✅ Key Takeaways:

// select waits for whichever channel receives first.
// Used for handling multiple channels simultaneously.

// Exercises
// 1. Create a goroutine that prints numbers from 1 to 5.
// 2. Use a channel to send and receive a string message between goroutines.
// 3. Modify worker() to process a list of tasks using channels.


// 1.Que
// func main(){
//     go func() {
//         for i := 1; i <= 5; i++{
//             fmt.Println(i)
//             time.Sleep(time.Second)
//         }
//     }()

//     time.Sleep(6 * time.Second)
// }

// 2.Use a channel to send and receive a string message between goroutines.

// func main(){

//   ch := make(chan string)

//   go func(){
//     message := "Hello Go"
//     ch <- message
//   }()

//   receivedmessage := <-ch
//   fmt.Println(receivedmessage)
// }


//WaitGroup using sync.WaitGroup

import "sync"

func runner1(wg *sync.WaitGroup) {
    defer wg.Done() 
    fmt.Println("\n I am first runner")

}

func runner2(wg *sync.WaitGroup) { 
    defer wg.Done()
    fmt.Println("\n I am second runner")
}

func execute() {
    wg := new(sync.WaitGroup)
    wg.Add(2)
    go runner1(wg)
    go runner2(wg)

    wg.Wait()
}

func main(){
    execute()
}
