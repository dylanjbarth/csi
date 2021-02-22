// Creates a comic search index using the raw_data
// Keeping it simple and building a hash table where each unique word links to a comic num and the frequency count.
package main

import (
	"fmt"
	"gopl-exercises/ex-4-12/xkcd/extract"
	"gopl-exercises/ex-4-12/xkcd/types"
	"log"
	"os"
	"path/filepath"
)

const indexfn = "comic_index.json"

var pwd string

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("Failed to fetch pwd %s", err)
	}
}

func main() {
	comicGen := extract.ReadAll()
	var index types.ComicIndex
	for {
		raw, err := comicGen()
		switch err.(type) {
		case *extract.ReadAllExhaustedError:
			break
		case nil:
			addToIndex(&raw, &index)
			continue
		default:
			log.Fatalf("Unexpected error reading commics. %s", err)
		}
		break
	}
	// write it to disk```
	fmt.Println(index)
}

// IndexFp returns the os safe absolute path to the index file
func IndexFp() string {
	return filepath.Join(pwd, indexfn)
}

// clean removes punctuation, lower cases everything, etc from the string
func clean(s string) {

}

// tokenize returns a slice of strings from a comic's transcript and alt tags
func tokenize(c *types.Comic) {

}

func addToIndex(c *types.Comic, idx *types.ComicIndex) {
	fmt.Printf("Adding %s to index\n", c)
}
