/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	result := 0
	for _, child := range root.Children {
		result = max(result, maxDepth(child))
	}
	return result + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}