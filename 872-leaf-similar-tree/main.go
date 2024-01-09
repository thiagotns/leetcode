package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	leafs1 := []int{}
	leafs2 := []int{}

	leafsExtract(root1, &leafs1)
	leafsExtract(root2, &leafs2)

	if len(leafs1) != len(leafs2) {
		return false
	}

	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}

	return true

}

func leafsExtract(root *TreeNode, leafs *[]int) {

	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, (*root).Val)
		return
	}

	if root.Left != nil {
		leafsExtract((*root).Left, leafs)
	}

	if root.Right != nil {
		leafsExtract(root.Right, leafs)
	}
}

func main() {

	/*
		Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
		Output: true
	*/

}
