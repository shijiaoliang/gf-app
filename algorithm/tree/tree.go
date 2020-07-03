package tree

import "fmt"

// 二叉树
type BinaryTree struct {
	Data  int
	Left  *BinaryTree
	Right *BinaryTree
}

func PreOrder(root *BinaryTree) {
	if root == nil {
		return
	}

	fmt.Println(root.Data)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func MidOrder(root *BinaryTree) {
	if root == nil {
		return
	}

	MidOrder(root.Left)
	fmt.Println(root.Data)
	MidOrder(root.Right)
}

func PostOrder(root *BinaryTree) {
	if root == nil {
		return
	}

	PostOrder(root.Right)
	PostOrder(root.Left)
	fmt.Println(root.Data)
}
