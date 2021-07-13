package main

import (
	"fmt"
	"math/rand"
)

// arbitrary max "height" for the skip list
const p = 0.5

type skipListNode struct {
	item  Item
	links []*skipListNode // the index of link indicates level, so index 0 = level 1, index 1 = level 2 and so forth.
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

// Search finds the node in the skip list that is equal to or less than the key provided.
// Returns the nearest node and a slice of nodes indictating the path taken through the skip list to get to the returned node.
func (o *skipListOC) Search(key string) (*skipListNode, []*skipListNode) {
	prevLinks := make([]*skipListNode, o.currLevel)
	curr := o.head
	// Work downward
	for i := o.currLevel - 1; i >= 0; i-- {
		// Work across
		for curr != nil && (curr.links[i] != nil && curr.links[i].item.Key <= key) {
			// Store the node we moved laterally from from
			curr = curr.links[i]
		}
		if curr != nil {
			prevLinks[i] = curr
		}
	}
	return curr, prevLinks
}

func (o *skipListOC) prettyPrint() {
	fmt.Printf("Current state of the skiplist: \n")
	fmt.Printf("# Levels: %d/%d\n", o.currLevel, o.maxLevel)
	curr := o.head
	for i := o.currLevel - 1; i >= 0; i-- {
		fmt.Printf("Level %d:", i)
		local := curr
		fmt.Printf(" HEAD -> ")
		for local != nil && (local.links[i] != nil) {
			fmt.Printf(" %s -> ", local.links[i].item.Key)
			local = local.links[i]
		}
		fmt.Printf("\n")
	}
	fmt.Printf("End list\n")
}

func (o *skipListOC) Get(key string) (string, bool) {
	curr, _ := o.Search(key)
	if curr != nil && curr.item.Key == key {
		return curr.item.Value, true
	}
	return "", false
}

func (o *skipListOC) Put(key, value string) bool {
	o.prettyPrint()
	curr, prevLinks := o.Search(key)

	if curr.item.Key == key {
		curr.item.Value = value
		return false
	}

	rl := RandLevel(o.maxLevel)
	if rl > o.currLevel {
		// fill in the blank space between head node and the new node.
		for i := o.currLevel; i < rl; i++ {
			prevLinks = append(prevLinks, o.head)
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
	curr, prevLinks := o.Search(key)
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
