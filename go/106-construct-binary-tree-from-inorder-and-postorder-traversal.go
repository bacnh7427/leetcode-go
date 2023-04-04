/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	length := len(postorder)
	root := TreeNode{Val: postorder[length-1]}
	mid := findIndex(inorder, postorder[length-1])

	root.Left = buildTree(inorder[:mid], postorder[:mid])
	root.Right = buildTree(inorder[mid+1:], postorder[mid:length-1])
	return &root
}

func findIndex(arr []int, target int) int {
	for idx, val := range arr {
		if val == target {
			return idx
		}
	}
	return -1
}
