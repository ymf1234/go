package main

import (
	"basic/tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()

	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}

func main() {
	var root tree.Node

	root = tree.Node{Value: 10}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{8, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	//root.setValue(5)
	root.Traverse()
	fmt.Println()

	myRooot := myTreeNode{&root}
	myRooot.postOrder()
	fmt.Println()



	root.Print()
	fmt.Println()

	nodes := []tree.Node {
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)
}


