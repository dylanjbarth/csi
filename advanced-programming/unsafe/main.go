package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("oh hey")
	fmt.Println(unsafe.Sizeof(float32(0)))
	fmt.Println(unsafe.Sizeof(float64(0)))
	fmt.Println(unsafe.Sizeof(make([]int, 100)))
	fmt.Println(unsafe.Sizeof(make(map[string]int)))
}
