/*
 * @Author: Xudong0722 
 * @Date: 2025-07-21 17:41:00 
 * @Last Modified by: Xudong0722
 * @Last Modified time: 2025-07-21 22:30:20
 */

package db

import (
	"sync"

	"github.com/Xudong0722/Leveldb-go/utils"
)

const (
	// MaxLevel is the maximum level of the skip list
	MaxLevel = 12
	// P is the probability of promoting a node to the next level
	P = 0.25
)

type Node struct {
	// Key-value data
	key interface{}

	// Pointers to next nodes at each level
	next []*Node

	// Current level of the node
	level int     
}

func NewNode(key interface{}, level int) *Node {
	return &Node {
		key: key,
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

	// Comprator
	cmp utils.Comprator
}

func NewSkipList(cp utils.Comprator) *SkipList {
	return &SkipList {
		head: NewNode(nil, MaxLevel),
		mutex: new(sync.RWMutex),
		maxLevel: MaxLevel,
		cmp: cp,
	}
}

func (sl *SkipList) GetCurrentHeight() int {
	return sl.maxLevel
}

// KeyIsAfterNode returns the given key is greater than 
// Node's key, return true means we need to keep searching in this list
func (sl *SkipList)KeyIsAfterNode(key interface{}, nd *Node) bool {
	if nd == nil {
		return false  //search in lower level
	}
	if sl.cmp(key, nd.key) > 0{
		return true
	}
	return false
}

// GetGreaterOrEqual returns the first node whose key is >= given key
// if prevs is not nil,  it also sets the first m.height elements of prev to the
// preceding node at each height.
func (sl *SkipList) GetGreaterOrEqual(key interface{}, prevs *[]*Node) *Node {
	cur := sl.head
	level := sl.GetCurrentHeight() - 1
	for  {
		next := cur.Next(level)
		if sl.KeyIsAfterNode(key, next) {
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
				if sl.cmp(key, cur.key) == 0 {
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

func (sl *SkipList) Insert(key interface{}) {
	return 
}

func (sl *SkipList) Contains(key interface{}) bool {
	return false
}

