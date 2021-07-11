package main

type skipListNode struct {
	item Item
	next *skipListNode // index in slice indicates level
	// next []*skipListNode // index in slice indicates level
}

type skipListOC struct {
	head *skipListNode
}

func newSkipListOC() *skipListOC {
	return &skipListOC{}
}

func (o *skipListOC) Get(key string) (string, bool) {
	curr := o.head
	for curr != nil {
		if curr.item.Key == key {
			return curr.item.Value, true
		}
		curr = curr.next
	}
	return "", false
}

func (o *skipListOC) Put(key, value string) bool {
	var prev *skipListNode
	curr := o.head
	for curr != nil && curr.item.Key < key {
		prev = curr
		curr = curr.next
	}
	if prev == nil {
		o.head = &skipListNode{Item{key, value}, curr}
		return true
	}
	if curr != nil && curr.item.Key == key {
		curr.item.Value = value
		return false
	}
	prev.next = &skipListNode{Item{key, value}, curr}
	return true
}

func (o *skipListOC) Delete(key string) bool {
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
