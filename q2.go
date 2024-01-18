package main

import "fmt"

// creating a tree
type Node struct {
	val   string
	left  *Node
	right *Node
}

func CreateNode(val string) *Node {
	return &Node{val: val, left: nil, right: nil}
}
func preOrder(node *Node) {
	if node != nil {
		fmt.Print(node.val)
		preOrder(node.left)
		preOrder(node.right)
	}
}
func postOrder(node *Node) {
	if node != nil {
		postOrder(node.left)
		postOrder(node.right)
		fmt.Printf(node.val)
	}
}

func main() {
	root := CreateNode("+")
	root.left = CreateNode("a")
	root.right = CreateNode("-")
	root.right.left = CreateNode("b")
	root.right.right = CreateNode("c")
	fmt.Println("Preorder traversal is:")
	preOrder(root)
	fmt.Println()
	fmt.Println("Postorder traversal is:")
	postOrder(root)
	fmt.Println()

}
