package main

import "fmt"

type LinkList struct {
	a    int
	next *LinkList
}

func (a *LinkList) search(b int) bool {
	temp1 := a
	for temp1 != nil {
		if temp1.a == b {
			return true
		}
		temp1 = temp1.next
	}
	return false
}

func (a *LinkList) insert(b int) bool {

}

func main() {
	arr := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
	}
	link := LinkList{a: arr[0], next: nil}
	temp := &link
	for i := 1; i < len(arr); i++ {
		temp.next = &LinkList{a: arr[i], next: nil}
		temp = temp.next
	}

	if link.search(11) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
