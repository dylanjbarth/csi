package main

import (
	"flag"
	"fmt"
	"gopl-exercises/ex-4-12/xkcd/comic"
	"log"
)

const maxDefault = 3

var extract = flag.Bool("extract", false, "Pass the --extract switch to collect the raw data from xkcd.")
var force = flag.Bool("force", false, "Pass the --force switch with extract to force collection of raw data, even if it's already cached.")
var index = flag.Bool("index", false, "Pass the --index switch to rebuild the search index.")
var max = flag.Int("max", maxDefault, fmt.Sprintf("Max search results."))

func main() {
	flag.Parse()
	if *extract {
		log.Printf("Fetching comics. Force: %t\n", *force)
		comic.FetchAll(*force)
	} else if *index {
		log.Printf("Creating search index.\n")
		comic.CreateIndex()
	} else {
		// default mode of operation is search
		cs := comic.KWSearch(flag.Args())
		if len(cs) == 0 {
			fmt.Println("No results found.")
		} else {
			for i, c := range cs {
				if i >= *max {
					break
				}
				fmt.Printf("%v\n", c)
			}
			fmt.Printf("Total Results: %d.\n", len(cs))
			if *max < len(cs) {
				fmt.Printf("%d more result(s) available (use the -max flag).", len(cs)-*max)
			}
		}
	}
}
