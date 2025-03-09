// 1. The Race Condition Problem
// 2. Avoid using sync.Mutex
// 3. using sync.RWMutex
// 4. using sync.Once
// 5. using sync.Cond

// A race condition occurs when multiple goroutines access and modify a shared variable at the same time.
// This can cause unexpected behavior and data corruption.

package main

import ("fmt"
"time")

var counter int 

func increment(){
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main(){
	go increment()
	go increment()
	time.Sleep(time.Second)
	fmt.Println("final counter value:",counter)
}

✅ Key Takeaways:
Multiple goroutines modify counter at the same time.
Results are unpredictable due to race conditions.

// 2. using sync.Mutex to Prevent Race Condition

import ("fmt"
"sync")

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&wg)
	go increment(&wg)

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}


✅ Key Takeaways:
mu.Lock() → Locks counter so only one goroutine can modify it at a time.
mu.Unlock() → Releases the lock so other goroutines can access counter.
Prevents race conditions and ensures correct results.


// 3. sync.RWMutex for Read/Write Operations

import ("fmt"
"time"
"sync")

var data int
var rwmu sync.RWMutex

func readdata(id int) {
	rwmu.RLock()
	fmt.Println("Reader", id, "reads data:", data)
	time.Sleep(time.Millisecond * 500)
	rwmu.RUnlock()
}

func writedata(){
	rwmu.Lock()
	data++
	fmt.Println("Writer updated data to:", data)
	rwmu.Unlock()
}

func main() {
	for i := 1; i <= 3; i++ {
		go readdata(i)
	}
	go writedata()
	
	time.Sleep(time.Second)
}

(Output may varies)

✅ Key Takeaways:
RLock() allows multiple readers simultaneously.
Lock() ensures only one writer at a time.
Useful when reads are frequent but writes are rare.

// 4. sync.Once for One-Time Initialization

sync.Once ensures a function runs only once, even if called multiple times.
Useful for initializing global resources (e.g., a database connection).

import ("fmt"
    "sync")

var once sync.Once

func initialize() {
    fmt.Println("Initializing system......")
}

func main() {
    for i := 1; i <=3; i++ {
        go once.Do(initialize)
    }
    fmt.Println("Function done")
    fmt.Scanln()
}
✅ Key Takeaways:
Ensures initialize() runs only once even if called multiple times.
Used for one-time setup tasks (e.g., database setup, configuration loading).


// 5. sync.cond for goroutine Coordination

sync.Cond allows one goroutine to signal another when it’s ready.
Useful when one goroutine depends on another to finish a task.

import ( "fmt" 
"sync" 
"time")

var condition = sync.NewCond(&sync.Mutex{}) 

func waiter(id int) {
    condition.L.Lock()
    condition.Wait()
    fmt.Println("Worker", id, "resumed")
    condition.L.Unlock()
}

func main() {
    for i := 1; i <= 3; i++ {
        go waiter(i)
    }

    time.Sleep(time.Second)
    fmt.Println("Signaling workers.....")
    condition.Broadcast()

    time.Sleep(time.Second)
}

✅ Key Takeaways:
Wait() makes goroutines wait until they receive a signal.
Broadcast() wakes up all waiting goroutines and Signal() will wake up for only particular goroutines.
For more inof https://victoriametrics.com/blog/go-sync-cond/

// Exercises
// 1. modify the counter program to use sync.RWMutext for reading and writing safely.

import("fmt"
"sync"
)

type Counter struct {
    mu sync.RWMutex
    value int
}

func (c *Counter) increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Get() int {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.value
}

func main() {
    var wg sync.WaitGroup
    counter := Counter{}

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.increment()
        }()
    }

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println("Counter Value:", counter.Get())
        }()
    }
    wg.Wait()
    fmt.Println("Final Counter Value:", counter.Get())
}

// 2. Implement a one-time database connection setup using sync.Once.

import ("fmt"
"sync")

type Database struct {
    connection string
}

var ( 
    dbinstance *Database
    once sync.Once
)

func dbConnection() *Database {
    once.Do(func() {
        fmt.Println("Initializing database connection...")
        dbinstance = &Database{connection: "Connected to DB"}
    })
    return dbinstance
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            db := dbConnection()
            fmt.Println("Goroutine %d: %v\n", id, db.connection)
        }(i)
    }
    wg.Wait()
}

sync.Once ensures the database connection is initialized only once, even when multiple goroutines attempt to access it simultaneously.
The once.Do() function executes the initialization logic only the first time GetDatabaseInstance() is called.
Multiple goroutines call GetDatabaseInstance(), but the database is connected only once.
This prevents race conditions and ensures a thread-safe singleton pattern.

// 3.Modify the worker() function to wait until a signal using sync.Cond.

import (
    "fmt"
    "sync"
    "time"
)

var (
    mu sync.Mutex
    cond = sync.NewCond(&mu)
)

func worker(id int) {
    mu.Lock()
    cond.Wait()
    mu.Unlock()

    fmt.Println("Worker %d is now running\n", id)
}

func main() {
    var wg sync.WaitGroup
    numofworkers := 5

    for i := 1; i <= numofworkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            worker(id)
        }(i)
    }

    time.Sleep(2 * time.Second)
    fmt.Println("All workers are ready. Sending signal....")

    mu.Lock()
    cond.Broadcast()
    mu.Unlock()

    wg.Wait()
    fmt.Println("All workers have completed execution")
}

Explanation:

sync.Cond is used to allow workers to wait until they receive a signal.
worker(id):
Locks the mutex.
Calls cond.Wait(), which releases the lock and blocks until it's signaled.
Once signaled, it regains the lock and proceeds with execution.
Main function:
Starts multiple worker goroutines.
Sleeps for 2 seconds to simulate setup.
Calls cond.Broadcast() to wake up all waiting workers.
Why Broadcast()?
Ensures all workers receive the signal at once.
If we wanted to wake workers one by one, we could use cond.Signal() instead.