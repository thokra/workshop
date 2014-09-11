package main

import (
	"runtime"
	"sync"

	"fmt"
)

type a struct {
	sync.Mutex
	runtime.BlockProfileRecord
}

//START OMIT

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		a := i
		go func() { // HL
			fmt.Println(a)
			wg.Done()
		}()
	}

	wg.Wait()
}

//END OMIT
