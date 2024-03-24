package http

import (
	"crypto/sha256"
	"encoding/base64"
)

/*
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
	fmt.Println(hash(key))
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
*/

func HashURL(url string) string {
	// Calculate SHA-256 hash of the URL
	hash := sha256.New()
	hash.Write([]byte(url))
	hashBytes := hash.Sum(nil)

	// Use base64 encoding to get alphanumeric characters
	encoded := base64.URLEncoding.EncodeToString(hashBytes)

	// Trim padding characters
	trimmed := trimPadding(encoded)

	// Take first 4 characters
	shortHash := trimmed[:4]

	return shortHash
}

func trimPadding(s string) string {
	// Remove padding characters from the end of the string
	for len(s) > 0 && s[len(s)-1] == '=' {
		s = s[:len(s)-1]
	}
	return s
}
