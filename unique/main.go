package main

import (
	"fmt"
	"runtime"

	"unique"
)

func getUsedMem() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func main() {
	var strs []string

	// unique better if we have a lot of same strings, but not if we have a lot of different strings
	for i := 0; i < 1_000_000; i++ {
		strs = append(strs, "same string")
	}
	
	beforeMem := getUsedMem()
	fmt.Printf("Before memory: %d bytes\n", beforeMem)
	

	var handles []unique.Handle[string]
	for _, str := range strs {
		handles = append(handles, unique.Make(str))
	}
	afterMem := getUsedMem()
	fmt.Printf("After memory: %d bytes\n", afterMem)

}
