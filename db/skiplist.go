/*
 * @Author: Xudong0722 
 * @Date: 2025-07-21 17:41:00 
 * @Last Modified by: Xudong0722
 * @Last Modified time: 2025-07-21 22:30:20
 */

package db

import (
	"sync"
	"bytes"
)

const (
	// MaxLevel is the maximum level of the skip list
	MaxLevel = 12
	// P is the probability of promoting a node to the next level
	P = 0.25
)

type Node struct {
	// Key-value data
	key []byte

	// Pointers to next nodes at each level
	next []*Node

	// Current level of the node
	level int     
}

func NewNode(kvData []byte, level int) *Node {
	return &Node {
		key: kvData,
		next: make([]*Node, level + 1),
		level: level,
	}
}

func (nd *Node) Next(n int) *Node {
	if n >= nd.level {
		return nil
	}
	return nd.next[n]
}

func (nd *Node) SetNext(n int, x *Node) bool {
	if n >= nd.level {
		return false
	}
	nd.next[n] = x
	return true
}

type SkipList struct {
	// Skip list head node
	head *Node   

	// Mutex for thread safety
	mutex *sync.RWMutex 

	// Current maximum level of the skip list
	maxLevel int
}

func NewSkipList() *SkipList {
	return &SkipList {
		head: NewNode(nil, MaxLevel),
		mutex: new(sync.RWMutex),
		maxLevel: MaxLevel,
	}
}

func (sl *SkipList) GetCurrentHeight() int {
	return sl.maxLevel
}

// KeyIsAfterNode returns the given key is greater than 
// Node's key, return true means we need to keep searching in this list
func KeyIsAfterNode(key []byte, nd *Node) bool {
	if nd == nil {
		return false  //search in lower level
	}
	if bytes.Compare(key, nd.key) > 0{
		return true
	}
	return false
}

// GetGreaterOrEqual returns the first node whose key is >= given key
// if prevs is not nil,  it also sets the first m.height elements of prev to the
// preceding node at each height.
func (sl *SkipList) GetGreaterOrEqual(key []byte, prevs *[]*Node) *Node {
	cur := sl.head
	level := sl.GetCurrentHeight() - 1
	for  {
		next := cur.Next(level)
		if KeyIsAfterNode(key, next) {
			//keep searching in this list
			cur = next
		}else {
			// The key greater than cur but less than or equal to next
			// cur ------key ------next
			// maybe key equal to next or less than next
			if nil != prevs {
				//Before jump to low level, we need to store current level value
				(*prevs)[level] = cur
			}
			if level == 0{
				//if we already at the level0
				if bytes.Equal(key, next.key) {
					return next
				}else {
					return nil
				}
			}else {
				// Switch to next list
				level = level - 1
			}
			
		}
	}
}

func (sl *SkipList) Insert(kvData []byte) {
	return 
}

func (sl *SkipList) Contains(kvData []byte) bool {
	return false
}

