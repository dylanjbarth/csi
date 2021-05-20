package main

// #include <math.h>
import "C"
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sq := os.Args[1]
	v, err := strconv.Atoi(sq)
	if err != nil {
		fmt.Printf("Can't convert %s to int for squaring.\n", sq)
		return
	}
	fmt.Printf("The square of %d is %d\n", v, square(v))
}

func square(n int) int {
	return int(C.pow(C.double(n), C.double(2)))
}
