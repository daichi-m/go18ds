// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktreeextended

import (
	"fmt"

	rbt "github.com/rahul1534/go18ds/trees/redblacktree"
)

// RedBlackTreeExtended to demonstrate how to extend a RedBlackTree to include new functions
type RedBlackTreeExtended struct {
	*rbt.Tree[int, string]
}

// GetMin gets the min value and flag if found
func (tree *RedBlackTreeExtended) GetMin() (value string, found bool) {
	node, found := tree.getMinFromNode(tree.Root)
	if node != nil {
		return node.Value, found
	}
	return "", false
}

// GetMax gets the max value and flag if found
func (tree *RedBlackTreeExtended) GetMax() (value string, found bool) {
	node, found := tree.getMaxFromNode(tree.Root)
	if node != nil {
		return node.Value, found
	}
	return "", false
}

// RemoveMin removes the min value and flag if found
func (tree *RedBlackTreeExtended) RemoveMin() (value string, deleted bool) {
	node, found := tree.getMinFromNode(tree.Root)
	if found {
		tree.Remove(node.Key)
		return node.Value, found
	}
	return "", false
}

// RemoveMax removes the max value and flag if found
func (tree *RedBlackTreeExtended) RemoveMax() (value string, deleted bool) {
	node, found := tree.getMaxFromNode(tree.Root)
	if found {
		tree.Remove(node.Key)
		return node.Value, found
	}
	return "", false
}

func (tree *RedBlackTreeExtended) getMinFromNode(node *rbt.Node[int, string]) (
	foundNode *rbt.Node[int, string], found bool) {
	if node == nil {
		return nil, false
	}
	if node.Left == nil {
		return node, true
	}
	return tree.getMinFromNode(node.Left)
}

func (tree *RedBlackTreeExtended) getMaxFromNode(node *rbt.Node[int, string]) (
	foundNode *rbt.Node[int, string], found bool) {
	if node == nil {
		return nil, false
	}
	if node.Right == nil {
		return node, true
	}
	return tree.getMaxFromNode(node.Right)
}

func print(tree *RedBlackTreeExtended) {
	max, _ := tree.GetMax()
	min, _ := tree.GetMin()
	fmt.Printf("Value for max key: %v \n", max)
	fmt.Printf("Value for min key: %v \n", min)
	fmt.Println(tree)
}

// RedBlackTreeExtendedExample main method on how to use the custom red-black tree above
func main() {
	tree := RedBlackTreeExtended{rbt.NewWithIntComparator[string]()}

	tree.Put(1, "a") // 1->x (in order)
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(3, "c") // 1->x, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->x, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->x, 2->b, 3->c, 4->d, 5->e (in order)

	print(&tree)
	// Value for max key: e
	// Value for min key: a
	// RedBlackTree
	// │       ┌── 5
	// │   ┌── 4
	// │   │   └── 3
	// └── 2
	//     └── 1

	tree.RemoveMin() // 2->b, 3->c, 4->d, 5->e (in order)
	tree.RemoveMax() // 2->b, 3->c, 4->d (in order)
	tree.RemoveMin() // 3->c, 4->d (in order)
	tree.RemoveMax() // 3->c (in order)

	print(&tree)
	// Value for max key: c
	// Value for min key: c
	// RedBlackTree
	// └── 3
}
