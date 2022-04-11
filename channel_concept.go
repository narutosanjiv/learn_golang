package main

func work1(ch chan int) {
	ch <- 1
}

func main() {
	ch := make(chan int)
	go work1(ch)
	<-ch
}
