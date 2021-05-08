package main

import "fmt"

func main() {
	ch := make(chan struct{})
	// go func() {
	// 	fmt.Println("Before send.")
	fmt.Println("Zero.")
	ch <- struct{}{}
	fmt.Println("First one.")
	ch <- struct{}{}
	fmt.Println("Second one.")
	ch <- struct{}{}
	// 	fmt.Println("After send.")
	// }()
}
