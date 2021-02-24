package main

import (
	"flag"
	"gopl-exercises/ex-4-12/xkcd/comic"
	"log"
)

var extract = flag.Bool("extract", false, "Pass the --extract switch to collect the raw data from xkcd.")
var force = flag.Bool("force", false, "Pass the --force switch with extract to force collection of raw data, even if it's already cached.")
var index = flag.Bool("index", false, "Pass the --index switch to rebuild the search index.")

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
		log.Printf("Search time.\n")
	}
}
