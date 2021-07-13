package main

import "math/rand"

// arbitrary max "height" for the skip list
const p = 0.5

type skipListNode struct {
	item  Item
	links []*skipListNode // turn this into a slice to indicate level
}

type skipListOC struct {
	maxLevel  int
	currLevel int
	head      *skipListNode
}

func newSkipListOC(maxLevel int) *skipListOC {
	// NB that head node is special, doesn't contain an item and reaches to the top level always
	return &skipListOC{maxLevel, 1, &skipListNode{links: []*skipListNode{nil}}}
}

// FindNext finds the node in the skip list that is equal to or greater than the key provided.
// Returns the next node (or nil) and a slice of nodes indictating the path taken through the skip list.
func (o *skipListOC) FindNext(key string) (*skipListNode, []*skipListNode) {
	prevLinks := make([]*skipListNode, o.maxLevel)
	curr := o.head
	for i := o.currLevel - 1; i >= 0; i-- {
		for curr != nil && (curr.links[i] != nil && curr.links[i].item.Key >= key) {
			curr = curr.links[i]
		}
		if curr != nil {
			prevLinks[i] = curr
		}
	}
	return curr, prevLinks
}

func (o *skipListOC) Get(key string) (string, bool) {
	curr, _ := o.FindNext(key)
	if curr != nil && curr.item.Key == key {
		return curr.item.Value, true
	}
	return "", false
}

func (o *skipListOC) Put(key, value string) bool {
	curr, prevLinks := o.FindNext(key)

	if curr.item.Key == key {
		curr.item.Value = value
		return false
	}

	rl := RandLevel(o.maxLevel)
	if rl > o.currLevel {
		// fill in the blank space between head node and the new node.
		for i := o.currLevel; i < rl; i++ {
			prevLinks[i] = o.head
		}
		o.currLevel = rl
	}

	// Splice slice and update all the links, working from bottom up
	newLinks := make([]*skipListNode, rl)
	node := &skipListNode{Item{key, value}, newLinks}
	for i := 0; i < rl; i++ {
		if prevLinks[i] != nil && len(prevLinks[i].links)-1 >= i {
			newLinks[i] = prevLinks[i].links[i]
			prevLinks[i].links[i] = node
		} else {
			prevLinks[i].links = append(prevLinks[i].links, node)
		}
	}
	node.links = newLinks
	return true
}

func (o *skipListOC) Delete(key string) bool {
	curr, prevLinks := o.FindNext(key)
	if curr != nil && curr.item.Key == key {
		for i := len(prevLinks) - 1; i >= 0; i++ {
			prevLinks[i].links[i] = curr.links[i]
		}
		return true
	}
	return false
}

func (o *skipListOC) RangeScan(startKey, endKey string) Iterator {
	return &skipListOCIterator{}
}

type skipListOCIterator struct {
}

func (iter *skipListOCIterator) Next() {
}

func (iter *skipListOCIterator) Valid() bool {
	return false
}

func (iter *skipListOCIterator) Key() string {
	return ""
}

func (iter *skipListOCIterator) Value() string {
	return ""
}

func RandLevel(maxLevel int) int {
	i := 1
	for p < rand.Float64() && i < maxLevel {
		i++
	}
	return i
}
