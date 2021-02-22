package types

import "fmt"

// Comic represents a comic metadata entry for an XKCD comic fetched via https://xkcd.com/Num/info.0.json, eg https://xkcd.com/2427/info.0.json when Num = 2427
type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func (c *Comic) String() string {
	return fmt.Sprintf("Comic Num %d", c.Num)
}

// ComicIndex is a hash table mapping keywords to comics containing them (and the frequency count)
// Example { "river": { 1: 4, 5: 1}} => river appears in comic 1 four times and comic 5 one time.
type ComicIndex map[string]map[int]int
