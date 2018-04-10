package main

type Node struct {
	data int
	left *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree)InsertItem(i int)  {
	if tree.root == nil {
		tree.root = &Node{data: i}
		return
	}
	currentNode := tree.root
	for {
		if i > currentNode.data
	}
}
