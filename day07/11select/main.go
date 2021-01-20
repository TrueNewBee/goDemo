package main

import "fmt"

// selcet  随机的选取
func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
