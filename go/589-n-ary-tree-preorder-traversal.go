/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// var result = []int{};
// func preorder(root *Node) []int {

//	        if root==nil {
//	            return result;
//	        }
//	        result = append(result, root.Val);
//	        for _,node := range root.Children {
//	            preorder(node);
//	        }
//	        return result;
//	}
func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, r.Val)
		tmp := []*Node{}
		for _, v := range r.Children {
			tmp = append([]*Node{v}, tmp...)
		}
		stack = append(stack, tmp...)
	}
	return res
}

func preorderdfs(root *Node, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		for i := 0; i < len(root.Children); i++ {
			preorderdfs(root.Children[i], res)
		}
	}
}