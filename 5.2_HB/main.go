package main

import (
	"bytes"
	"fmt"
	"hash/fnv"
)

const MAP_SIZE = 10

type Node struct {
	key   string
	value string
	next  *Node
}

type HashMap struct {
	Data []*Node
}

func NewDict() *HashMap {
	return &HashMap{Data: make([]*Node, MAP_SIZE)}
}

func (n *Node) String() string {
	return fmt.Sprintf("<Key: %s, Value: %s>\n", n.key, n.value)
}

func (h *HashMap) String() string {
	var output bytes.Buffer
	fmt.Fprintln(&output, "{")
	for _, n := range h.Data {
		if n != nil {
			fmt.Fprintf(&output, "\t%s: %s\n", n.key, n.value)
			for node := n.next; node != nil; node = node.next {
				fmt.Fprintf(&output, "\t%s: %s\n", node.key, node.value)
			}
		}
	}

	fmt.Fprintln(&output, "}")

	return output.String()
}

func (h *HashMap) Insert(key string, value string) {
	i := getIndex(key)
	if h.Data[i] == nil {
		h.Data[i] = &Node{key: key, value: value}
	} else {
		newNode := h.Data[i]
		for ; newNode.next != nil; newNode = newNode.next {
			if newNode.key == key {
				newNode.value = value
				return
			}
		}
		newNode.next = &Node{key: key, value: value}
	}
}

func (h *HashMap) Get(key string) (string, bool) {
	i := getIndex(key)
	if h.Data[i] != nil {
		node := h.Data[i]
		for ; ; node = node.next {
			if node.key == key {
				return node.value, true
			}
			if node.next == nil {
				break
			}
		}
	}
	return "", false
}

func hash(key string) (hash uint32) {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func getIndex(key string) (index int) {
	return int(hash(key)) % MAP_SIZE
}

func (h *HashMap) delete(key string) bool {
	i := getIndex(key)
	if h.Data[i] != nil {
		node := h.Data[i]
		for ; ; node = node.next {
			if node.key == key {
				h.Data[i] = nil
				return true
			}
			if node.next == nil {
				break
			}
		}
	}
	return false
}

func main() {
	a := NewDict()
	a.Insert("name", "ishan")
	a.Insert("gender", "male")
	a.Insert("city", "mumbai")
	a.Insert("lastname", "khare")
	if value, ok := a.Get("name"); ok {
		fmt.Println(value)
	} else {
		fmt.Println("Value did not match!")
	}

	fmt.Println(a)
	// a.Get("city")
	// a.delete("city")
	fmt.Println(a.delete("city"))
	fmt.Println(a)
}
