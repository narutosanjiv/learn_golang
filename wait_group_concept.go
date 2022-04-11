package main

import (
	"fmt"
	"runtime"
	"sync"
)

func work(wg *sync.WaitGroup) {
	fmt.Println("in do work, no. of goroutines", runtime.NumGoroutine())
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("In main functions")
	arr := []int{
		10, 4, 90,
	}
	for i, i2 := range arr {
		fmt.Printf("index:%d value=%d\n", i, i2)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(&wg)
	}
	wg.Wait()
}
