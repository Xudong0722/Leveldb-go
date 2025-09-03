/*
 * @Author: Xudong0722
 * @Date: 2025-07-21 17:41:00
 * @Last Modified by: Xudong0722
 * @Last Modified time: 2025-07-21 22:30:20
 */

package db

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/Xudong0722/Leveldb-go/utils"
)

const (
	// MaxHeight is the maximum level of the skip list
	MaxHeight = 12

	Branching = 4
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
	return &Node{
		key:   key,
		next:  make([]*Node, level+1),
		level: level,
	}
}

func (nd *Node) SetNext(n int, x *Node) bool {
	if n >= nd.level {
		return false
	}
	nd.next[n] = x
	return true
}

func (nd *Node) GetNext(n int) *Node {
	if n >= nd.level {
		return nil
	}
	return nd.next[n]
}

type SkipList struct {
	// Skip list head node
	head *Node

	// Mutex for thread safety
	mutex *sync.RWMutex

	// Current maximum level of the skip list
	maxHeight int

	// Comprator
	cmp Comprator
}

func NewSkipList(cp Comprator) *SkipList {
	return &SkipList{
		head:      NewNode(nil, MaxHeight),
		mutex:     new(sync.RWMutex),
		maxHeight: 1, //初始高度为1， 但是head的next数组是有MaxHeight层
		cmp:       cp,
	}
}

func (sl *SkipList) GetCurrentHeight() int {
	return sl.maxHeight
}

// KeyIsAfterNode returns the given key is greater than
// Node's key, return true means we need to keep searching in this list
func (sl *SkipList) KeyIsAfterNode(key interface{}, nd *Node) bool {
	if nd == nil {
		return false //search in lower level
	}
	res, _ := sl.cmp(key, nd.key)
	return res > 0
}

// GetGreaterOrEqual returns the first node whose key is >= given key
func (sl *SkipList) GetGreaterOrEqual(key interface{}) (*Node, [MaxHeight]*Node) {
	cur := sl.head
	level := sl.GetCurrentHeight() - 1
	var prevs [MaxHeight]*Node
	for {
		next := cur.GetNext(level)
		if sl.KeyIsAfterNode(key, next) {
			//keep searching in this list
			cur = next
		} else {
			// The key greater than cur but less than or equal to next
			// cur ------key ------next
			// maybe key equal to next or less than next

			//Before jump to low level, we need to store current level value
			prevs[level] = cur

			if level == 0 {
				//if we already at the level0
				// res, err := sl.cmp(key, cur.key)
				// if err == nil && res == 0 {
				return next, prevs
				// } else {
				// 	return nil, prevs
				// }
			} else {
				// Switch to next list
				level = level - 1
			}
		}
	}
}

// Return the latest node with a key < key.
// Return head if there is no such node.
func (sl *SkipList) GetLessThan(key interface{}) (*Node) {
	cur := sl.head
	level := sl.GetCurrentHeight() - 1
	for {
		next := cur.GetNext(level)
		// If next is nil, or next.key is greater than key, we stop at previous node.
		// 8 ----> 13
		// 8 ->9-> 13
		//key = 10, cur = head
		//round1: next = 8, level = 1; 8 > 10? false, cur = 8
		//round2: next = 13, level = 1; 13 > 10? true, level = 0, cur = 8
		//round3: next = 9, level = 0; 9 > 10? false, cur = 9
		//round4: next = 13, level = 0; 13 > 10? true, because level = 0, return cur = 9
		res := -1
		if next != nil {
			res, _ = sl.cmp(next.key, key)
		}
		if next == nil ||  res >= 0  {
			if level == 0{
				return cur
			}else{
				level = level - 1
			}
		} else {
			cur = next
		}
	}
}

func (sl *SkipList) GetLast() *Node {
	cur := sl.head
	level := sl.GetCurrentHeight() - 1
	for {
		next := cur.GetNext(level)
		if next == nil {
			if level == 0 {
				return cur
			} else {
				level = level - 1
			}
		}else{
			cur = next
		}
	}
}

func (sl *SkipList) Insert(key interface{}) {
	node, prevs := sl.GetGreaterOrEqual(key)
	fmt.Println(node)
	new_height := sl.randomHeight()

	if sl.GetCurrentHeight() < new_height {
		for i := sl.GetCurrentHeight(); i < new_height; i++ {
			prevs[i] = sl.head //这个高度之前没有人达到，用head进行补充，因为head的next数组是最高的
		}
		sl.maxHeight = new_height
	}

	new_node := NewNode(key, new_height)
	for i := 0; i < new_height; i++ {
		//  prevs[i]  -> new_node -> prevs[i].next[i]
		new_node.SetNext(i, prevs[i].GetNext(i))
		prevs[i].SetNext(i, new_node)
	}

}

func (sl *SkipList) Contains(key interface{}) bool {
	if key == nil {
		return false
	}
	tar, _ := sl.GetGreaterOrEqual(key)
	if tar == nil {
		return false
	}
	res, err := sl.cmp(key, tar.key)
	if err == nil && res == 0 {
		return true
	}
	return false
}

func (sl *SkipList) Delete(key interface{}) bool {
	if !sl.Contains(key) {
		return false
	}

	target, prevs := sl.GetGreaterOrEqual(key)
	if nil == target {
		return false
	}
	res, err := sl.cmp(key, target.key)
	if err != nil || res != 0 {
		return false
	}

	//这个地方应该使用target的高度去遍历，target level小于跳表当前高度情况下，高于target level的perv的next并不指向target，不能修改
	//wrong: for i := 0; i < sl.GetCurrentHeight(); i++
	for i := 0; i < target.level; i++ {
		prevs[i].SetNext(i, target.GetNext(i))
	}
	target = nil
	return true
}

func (sl *SkipList) randomHeight() int {
	height := 1
	for height < MaxHeight && (rand.Intn(Branching) == 0) { // 1/4 enter higher level.
		height++
	}
	return height
}

type SkipListIterator struct {
	//当前所在的节点
	cur *Node

	//对应的跳表
	sl *SkipList
}

func NewSkipListIterator(sl_ *SkipList) *SkipListIterator {
	return &SkipListIterator{
		cur: nil,
		sl:  sl_,
	}
}

func (sl_iter *SkipListIterator) Valid() bool {
	return sl_iter.cur != nil
}

func (sl_iter *SkipListIterator) Key() interface{} {
	utils.Assert(sl_iter.Valid(), "Current node is nil.")
	return sl_iter.cur.key
}

func (sl_iter *SkipListIterator) Next() {
	utils.Assert(sl_iter.Valid(), "Current node is nil.")
	sl_iter.cur = sl_iter.cur.next[0]
}

func (sl_iter *SkipListIterator) Prev() {
	utils.Assert(sl_iter.Valid(), "Current node is nil.")
	sl_iter.cur = sl_iter.sl.GetLessThan(sl_iter.cur.key)
	// 如果返回的是 head 节点，说明已经到达了跳表的开始，将迭代器设置为无效
	if sl_iter.cur == sl_iter.sl.head {
		sl_iter.cur = nil
	}
}

func (sl_iter *SkipListIterator) Seek(target interface{}) { //TODO, param->interface{}
	utils.Assert(sl_iter.Valid(), "Current node is nil.")
	sl_iter.cur, _ = sl_iter.sl.GetGreaterOrEqual(target)
}

func (sl_iter *SkipListIterator) SeekToFirst() {
	sl_iter.cur = sl_iter.sl.head.next[0]
}

func (sl_iter *SkipListIterator) SeekToLast() {
	sl_iter.cur = sl_iter.sl.GetLast()
}
