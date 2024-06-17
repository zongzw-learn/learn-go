package main

import (
	"log"
	"sort"
	"strings"
)

type RadixNode struct {
	Key   string
	Nodes RadixNodes
}

func NewRadixNode(key string) *RadixNode {
	var nodes RadixNodes = nil
	if key != "" {
		nodes = []*RadixNode{{Key: ""}}
	}

	return &RadixNode{
		Key:   key,
		Nodes: nodes,
	}
}

func (rn *RadixNode) Insert(key string) {
	n := maxOverlap(rn.Key, key)
	if n == len(rn.Key) {
		k := rn.Nodes.findKeyPrefix(key[n:])
		if k == -1 {
			rn.Nodes = append(rn.Nodes, NewRadixNode(key[n:]))
		} else {
			rn.Nodes[k].Insert(key[n:])
		}
	} else {
		newKey := rn.Key[0:n]
		newNodes := []*RadixNode{
			NewRadixNode(rn.Key[n:]),
			NewRadixNode(key[n:]),
		}
		newNodes[0].Nodes = rn.Nodes
		rn.Key = newKey
		rn.Nodes = newNodes
	}

	sort.Sort(rn.Nodes)
}

// Delete 删除Radix树中的元素，同时合并需要合并的子项
func (rn *RadixNode) Delete(key string) {
	n := maxOverlap(rn.Key, key)
	if n != len(rn.Key) {
		return // not found
	}
	k := rn.Nodes.findKeyPrefix(key[n:])
	if k == -1 {
		return // not found
	}

	subrn := rn.Nodes[k]
	if !strings.HasPrefix(key[n:], subrn.Key) {
		return // not found
	}
	if subrn.Key == key[n:] {
		if len(subrn.Nodes) == 1 { // 无其他子节点
			rn.Nodes = append(rn.Nodes[0:k], rn.Nodes[k+1:]...)
			if len(rn.Nodes) == 1 && rn.Nodes[0].Key != "" {
				rn.Key = rn.Key + rn.Nodes[0].Key
				rn.Nodes = rn.Nodes[0].Nodes
			}
		} else if len(subrn.Nodes) == 2 { // 有一个子节点
			subrn.Key = subrn.Key + subrn.Nodes[1].Key
			subrn.Nodes = subrn.Nodes[1].Nodes
		} else { // 删除该节点本身
			subrn.Nodes = subrn.Nodes[1:]
		}
		return
	} else {
		rn.Nodes[k].Delete(key[n:])
	}
}

func (rn *RadixNode) Iterate() []string {
	if len(rn.Nodes) <= 1 {
		return []string{rn.Key}
	}
	rlt := []string{}
	for _, k := range rn.Nodes {
		krlt := k.Iterate()
		for _, kr := range krlt {
			rlt = append(rlt, rn.Key+kr)
		}
	}
	return rlt
}

func main() {
	words := []string{"opening", "open", "book", "books", "blue"}
	root := NewRadixNode("")
	for _, word := range words {
		root.Insert(word)
	}
	log.Printf("%v", root.Iterate())

	root.Delete("zongzhaowei")
	root.Delete("op")
	root.Delete("open")
	root.Delete("books")
	root.Delete("blue")

	log.Printf("%v", root.Iterate())

	words = []string{"0", "1", "10", "11", "100", "101", "110", "111", "1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111"}
	root = NewRadixNode("")
	for _, word := range words {
		root.Insert(word)
	}
	log.Printf("%v", root.Iterate())

	root.Delete("1010")
	log.Printf("%v", root.Iterate())

}

func maxOverlap(k1, k2 string) int {
	i := 0
	for ; i < len(k1) && i < len(k2) && k1[i] == k2[i]; i++ {
	}
	return i
}

func (rns RadixNodes) findKeyPrefix(key string) int {
	if key == "" {
		return -1
	}
	i, j := 0, len(rns)-1
	for i < j {
		if len(rns[i].Key) == 0 {
			i++
		}
		m := (i + j) / 2
		if rns[m].Key[0] > key[0] {
			j = m - 1
		} else if rns[m].Key[0] < key[0] {
			i = m + 1
		} else {
			return m
		}
	}
	if i == j && len(rns[i].Key) > 0 && rns[i].Key[0] == key[0] {
		return i
	}
	return -1
}

type RadixNodes []*RadixNode

func (rns RadixNodes) Len() int           { return len(rns) }
func (rns RadixNodes) Swap(i, j int)      { rns[i], rns[j] = rns[j], rns[i] }
func (rns RadixNodes) Less(i, j int) bool { return rns[i].Key < rns[j].Key }
