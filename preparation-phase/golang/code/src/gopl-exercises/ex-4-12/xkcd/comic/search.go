package comic

import (
	"log"
	"sort"
)

type kwFreq struct {
	comicID int
	freq    int
}
type kwFreqList []kwFreq

func (k kwFreqList) Len() int           { return len(k) }
func (k kwFreqList) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k kwFreqList) Less(i, j int) bool { return k[i].freq > k[j].freq }

// KWSearch returns a slice of comic ids that match each keyword, ordered by keyword frequency.
func KWSearch(kws []string) []*Comic {
	idx := loadIndex()
	r := search(kws, idx)
	f := make([]*Comic, len(r))
	for i, cid := range r {
		c, err := readByID(cid)
		if err != nil {
			log.Fatalf("Failed to read comic %d. %s", cid, err)
		}
		f[i] = &c
	}
	return f
}

func search(kws []string, i *SearchIndex) []int {
	// comic ID => sum of hits
	freq := make(map[int]int)
	for _, kw := range kws {
		hits, found := (*i)[kw]
		if !found {
			continue
		}
		for cid, ct := range hits {
			_, ok := freq[cid]
			if !ok {
				freq[cid] = ct
			} else {
				freq[cid] += ct
			}
		}
	}

	// TODO Feels clumsy.. storing in map, then moving to slice to sort?
	s := make(kwFreqList, len(freq))
	count := 0
	for k, v := range freq {
		s[count] = kwFreq{comicID: k, freq: v}
		count++
	}
	sort.Sort(s)

	// TODO feels extra clumsy, converting this to a single array now.
	f := make([]int, len(s))
	for ix, v := range s {
		f[ix] = v.comicID
	}
	return f
}
