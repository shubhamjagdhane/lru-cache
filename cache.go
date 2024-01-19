package main

import "fmt"

type entry struct {
	value string
	data  *node
}

type cache struct {
	records map[int]entry
	size    int
	dll     *doublyLinkedList
}

func newCache(size int) *cache {
	return &cache{
		records: make(map[int]entry, 0),
		dll:     newDoublyLinkedList(),
		size:    size,
	}
}

func (c *cache) Get(key int) (string, error) {
	e, ok := c.records[key]
	if !ok {
		return "", fmt.Errorf("KEY DOES NOT EXISTS")
	}

	c.dll.AdjustNode(e.data)

	return e.value, nil
}

func (c *cache) Set(key int, value string) error {
	e := entry{
		value: value,
		data:  newNode(key),
	}

	if len(c.records) >= c.size {
		delete(c.records, c.dll.tail.data)
		c.dll.RemoveTail()
	}

	c.dll.InsertFirst(e.data)
	c.records[key] = e

	return nil
}
