package main

import (
	"fmt"
	"sync"
)

var countWords = make(map[string]int)
var wg sync.WaitGroup
var psuedo_mutex = make(chan int, 1)

func update_map1(key string, a int) {
	<-psuedo_mutex
	countWords[key] = countWords[key] + 1
	wg.Done()
	psuedo_mutex <- 1
}

func main() {
	psuedo_mutex <- 1
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go update_map1("update1", 1)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go update_map1("update2", 1)
	}

	wg.Wait()

	for s, i := range countWords {
		fmt.Printf("Key=%s, Value=%d\n", s, i)
	}
}
