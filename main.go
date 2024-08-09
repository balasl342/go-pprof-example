package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// Start CPU profiling
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	defer func() {
		pprof.StopCPUProfile()
		cpuFile.Close()
	}()

	pprof.StartCPUProfile(cpuFile)

	// Run the test function
	testUser()

	// Optionally, write a memory profile
	memFile, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("Could not create memory profile:", err)
		return
	}
	defer memFile.Close()

	if err := pprof.WriteHeapProfile(memFile); err != nil {
		fmt.Println("Could not write memory profile:", err)
	}
}

func testUser() {
	time.Sleep(10 * time.Second)
	fmt.Println("Test User")
}
