/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var q = [](*TreeNode){root}
	var result = [][]int{}
	for len(q) > 0 {
		var curLevel = []int{}
		var sz = len(q)
		for i := 0; i < sz; i++ {
			var current = q[i]
			curLevel = append(curLevel, current.Val)
			if current.Left != nil {
				q = append(q, current.Left)
			}
			if current.Right != nil {
				q = append(q, current.Right)
			}
		}
		q = q[sz:]
		result = append(result, curLevel)
	}
	return result
}