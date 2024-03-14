package http

import (
	"bytes"
	"fmt"
)

type Node struct {
	Key   string
	Value string
	Next  *Node
}

type HashMap struct {
	Data []*Node
}

var mapSize int = 50

func NewMap() *HashMap {
	return &HashMap{
		Data: make([]*Node, mapSize),
	}
}

func hash(key string) (hash uint32) {
	hash = 0
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return
}

func getIndex(key string) (index int) {
	return int(hash(key)) % mapSize
}

func (h *HashMap) Insert(key string, value string) {
	index := getIndex(key)

	if h.Data[index] == nil {
		//index is empty, then you insert the new one
		h.Data[index] = &Node{
			Key:   key,
			Value: value,
		}
	} else {
		// there is a collision, get into linked-list mode
		starting_node := h.Data[index]
		for ; starting_node.Next != nil; starting_node = starting_node.Next {
			if starting_node.Key == key {
				//the key exist and is a modifying operation
				starting_node.Value = value
				return
			}
		}
		starting_node.Next = &Node{
			Key:   key,
			Value: value,
		}
	}
}

func (h *HashMap) Get(key string) (string, bool) {
	index := getIndex(key)
	if h.Data[index] != nil {
		//check if another value is chained
		starting_node := h.Data[index]
		for ; ; starting_node = starting_node.Next {
			if starting_node.Key == key {
				//key matched
				return starting_node.Value, true
			}
			if starting_node.Next == nil {
				break
			}
		}
	}

	return "", false
}

func (h *HashMap) String() string {
	var output bytes.Buffer
	fmt.Println(&output, "{")
	for _, n := range h.Data {
		if n != nil {
			fmt.Fprintf(&output, "\t%s: %s\n", n.Key, n.Value)
			for node := n.Next; node != nil; node = node.Next {
				fmt.Fprintf(&output, "\t%s: %s\n", node.Key, node.Value)
			}
		}
	}
	fmt.Fprintln(&output, "}")

	return output.String()
}
