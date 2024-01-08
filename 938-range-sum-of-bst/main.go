package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {

	s := 0

	if (*root).Val >= low && (*root).Val <= high {
		s += (*root).Val
	}

	if root.Left != nil {
		s += rangeSumBST(root.Left, low, high)
	}

	if root.Right != nil {
		s += rangeSumBST(root.Right, low, high)
	}

	return s

}

func main() {

	/*
		Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
		Output: 32
		Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.
	*/

}
