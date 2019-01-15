package tree

import "fmt"

//结构体
type Node struct {
	Value int
	Left, Right *Node
}

//打印
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//修改结构体的值
func (node *Node) SetValue(value int) {
	node.Value = value
}

//遍历（中序遍历）
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()

}


func CreateNode(value int) *Node {
	return &Node{Value: value}
}