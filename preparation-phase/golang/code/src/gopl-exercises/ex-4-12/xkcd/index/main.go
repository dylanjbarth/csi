// Creates a comic search index using the raw_data
// Keeping it simple and building a hash table where each unique word links to a comic num and the frequency count.
package index

import (
	"fmt"
	"gopl-exercises/ex-4-12/xkcd/extract"
	"gopl-exercises/ex-4-12/xkcd/types"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const indexfn = "comic_index.json"
const cleaner = "[^a-zA-Z0-9]+"

var pwd string
var cleanreg *regexp.Regexp

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("Failed to fetch pwd %s", err)
	}

	cleanreg, err = regexp.Compile(cleaner)
	if err != nil {
		log.Fatalf("Failed to compile regex %s", err)
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
func clean(s string) string {
	return strings.ToLower(cleanreg.ReplaceAllString(s, ""))
}

// tokenize returns a slice of strings from a comic's transcript and alt tags
func tokenize(c *types.Comic) []string {
	words := strings.Split(c.Title, " ")
	for i, w := range words {
		words[i] = clean(w)
	}
	return words
}

func addToIndex(c *types.Comic, idx *types.ComicIndex) {
	tokens := tokenize(c)
	fmt.Println(tokens)
}
