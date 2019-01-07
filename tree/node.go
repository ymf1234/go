package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value, "")
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()

}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode

	root = treeNode{value: 10}
	root.left = &treeNode{}
	root.right = &treeNode{8, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	//root.setValue(5)
	root.traverse()

	fmt.Println()


	root.print()
	fmt.Println()

	nodes := []treeNode {
		{value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)
}
