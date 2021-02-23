// Creates a comic search index using the raw_data
// Keeping it simple and building a hash table where each unique word links to a comic num and the frequency count.
package main

import (
	"encoding/json"
	"gopl-exercises/ex-4-12/xkcd/extract"
	"gopl-exercises/ex-4-12/xkcd/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const indexfn = "comic_index.json"
const cleaner = "[^a-zA-Z0-9]+"
const space = `[\n\r\s]+`

var pwd string
var cleanreg *regexp.Regexp
var spacereg *regexp.Regexp

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("Failed to fetch pwd %s", err)
	}
	cleanreg = regexp.MustCompile(cleaner)
	spacereg = regexp.MustCompile(space)
}

func main() {
	comicGen := extract.ReadAll()
	index := make(types.ComicIndex)
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
	write(&index)
}

// Fp returns the os safe absolute path to the index file
func Fp() string {
	return filepath.Join(pwd, indexfn)
}

func write(i *types.ComicIndex) {
	b, err := json.Marshal(&i)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Writing comic index to disk => %s\n", Fp())
	if err = ioutil.WriteFile(Fp(), b, 0600); err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}
}

// clean removes punctuation, lower cases everything, etc from the string
func clean(s string) string {
	return strings.ToLower(cleanreg.ReplaceAllString(s, ""))
}

// tokenize returns a slice of strings from a comic's transcript and alt tags
func tokenize(c *types.Comic) []string {
	// Cleaner way to express this in Go? Eg iterate through each attribute we care about?
	words := append(spacereg.Split(c.Title, -1), spacereg.Split(c.Transcript, -1)...)
	words = append(words, spacereg.Split(c.Alt, -1)...)
	words = append(words, spacereg.Split(c.Year, -1)...)
	var cleanedWords []string
	for _, w := range words {
		cleaned := clean(w)
		if cleaned != "" {
			cleanedWords = append(cleanedWords, cleaned)
		}
	}
	return cleanedWords
}

func addToIndex(c *types.Comic, idx *types.ComicIndex) {
	tokens := tokenize(c)
	for _, t := range tokens {
		// We already have a base entry for this comic and word pair
		if _, ok := (*idx)[t]; !ok {
			(*idx)[t] = map[int]int{c.Num: 1}
		} else {
			// check if found already in this comic
			if _, exists := (*idx)[t][c.Num]; !exists {
				(*idx)[t][c.Num] = 1
			} else {
				(*idx)[t][c.Num]++
			}
		}
	}
}
