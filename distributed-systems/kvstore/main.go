package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const PROMPT = "kv> "

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		os.Stdout.Write([]byte(PROMPT))
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Unable to read from stdin. %s", err)
		}
		fmt.Println(line)
	}
}
