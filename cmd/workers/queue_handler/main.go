package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("start")
	wg := &sync.WaitGroup{}
	test(wg)
	wg.Wait()
	fmt.Println("end")
}

func test(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("go 1:", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("go 2:", i)
		}
	}()
}
