package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, val := range os.Args[1:] {
		fmt.Println(idx, val)
	}
}
