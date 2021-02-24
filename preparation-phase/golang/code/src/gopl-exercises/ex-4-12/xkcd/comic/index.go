package comic

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

const indexfn = "comic_index.json"
const cleaner = "[^a-zA-Z0-9]+"
const space = `[\n\r\s]+`

var cleanreg *regexp.Regexp
var spacereg *regexp.Regexp

func init() {
	cleanreg = regexp.MustCompile(cleaner)
	spacereg = regexp.MustCompile(space)
}

// CreateIndex builds a new search index
func CreateIndex() {
	comicGen := readAll()
	index := make(SearchIndex)
	count := 0
	for {
		raw, err := comicGen()
		switch err.(type) {
		case *readAllExhaustedError:
			break
		case nil:
			index.add(&raw)
			count++
			continue
		default:
			log.Fatalf("Unexpected error reading commics. %s", err)
		}
		break
	}
	log.Printf("Processed %d comics.", count)
	index.write()
}

func loadIndex() *SearchIndex {
	b, err := ioutil.ReadFile(indexPath())
	if err != nil {
		log.Fatalf("Unable to read comic search index from disk. %s", err)
	}
	i := make(SearchIndex)
	jErr := json.Unmarshal(b, &i)
	if jErr != nil {
		log.Fatalf("Unable to unmarshall json of comic search index from disk. Could be corrupted. %s", jErr)
	}
	return &i
}

// indexPath returns the os safe absolute path to the index file
func indexPath() string {
	p, err := filepath.Abs(filepath.Join(".", indexfn))
	if err != nil {
		log.Fatalf("Unable to produce index filepath. %s", err)
	}
	return p
}

func (i *SearchIndex) write() {
	b, err := json.Marshal(&i)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Writing comic index to disk => %s\n", indexPath())
	if err = ioutil.WriteFile(indexPath(), b, 0600); err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}
}

// clean removes punctuation, lower cases everything, etc from the string
func clean(s string) string {
	return strings.ToLower(cleanreg.ReplaceAllString(s, ""))
}

// tokenize returns a slice of strings from a comic's transcript and alt tags
func (c *Comic) tokenize() []string {
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

func (i *SearchIndex) add(c *Comic) {
	tokens := c.tokenize()
	for _, t := range tokens {
		// We already have a base entry for this comic and word pair
		if _, ok := (*i)[t]; !ok {
			(*i)[t] = map[int]int{c.Num: 1}
		} else {
			// check if found already in this comic
			if _, exists := (*i)[t][c.Num]; !exists {
				(*i)[t][c.Num] = 1
			} else {
				(*i)[t][c.Num]++
			}
		}
	}
}
