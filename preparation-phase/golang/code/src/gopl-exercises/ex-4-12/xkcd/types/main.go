package types

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
