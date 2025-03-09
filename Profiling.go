// Go Profiling with go tool pprof
// Profiling helps analyze CPU usage, memory allocation, and performance bottlenecks in Go programs. The go tool pprof command is used for this purpose.

// CPU usage
// Memory allocation
// Performance 
// bottlenecks

// 1. What is pprof?
// pprof (Performance Profiler) collects and visualizes profiling data.
// It helps find slow functions, high CPU usage, and memory leaks.
// Works with net/http/pprof for web-based profiling.

//CPU PROFILING

// package main

// import (
// 	"fmt"
// 	"os"
// 	"runtime/pprof"
// 	// "time"
// )

// // Simulate CPU-intensive work
// func doWork() {
// 	for i := 0; i < 1e7; i++ {
// 		_ = i * i
// 	}
// }

// func main() {
// 	// Create a file to save the CPU profile
// 	f, err := os.Create("cpu.prof")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer f.Close()

// 	// Start CPU profiling
// 	pprof.StartCPUProfile(f)
// 	defer pprof.StopCPUProfile()

// 	// Run the function we want to profile
// 	doWork()

// 	fmt.Println("CPU profiling complete. Run 'go tool pprof cpu.prof'")
// }

//Memory Profiling

// package main

// import (
// 	"fmt"
// 	"os"
// 	"runtime/pprof"
// )

// // Simulate memory allocation
// func allocateMemory() {
// 	data := make([]int, 1e6) // Allocate 1 million integers
// 	_ = data
// }

// func main() {
// 	// Create a file to save memory profile
// 	f, err := os.Create("mem.prof")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer f.Close()

// 	// Run function to analyze
// 	allocateMemory()

// 	// Write memory profile to file
// 	pprof.WriteHeapProfile(f)

// 	fmt.Println("Memory profiling complete. Run 'go tool pprof mem.prof'")
// }


//Profiling Go code with Go Code with go-torch and Flame Graph
// go-torch is a tool for generating flame graphs from pprof data.
// It helps visualize CPU usage and identify bottlenecks.

