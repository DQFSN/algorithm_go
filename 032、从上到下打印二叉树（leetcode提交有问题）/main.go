/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var queue []*TreeNode
 var result = make([]int, 0)
 
 func levelOrder(root *TreeNode) []int {
 
	 if root != nil {
		 queue = append(queue, root)
		 result = append(result, queue[0].Val)
		 queue = queue[1:]
		 levelOrder(root.Left)
		 levelOrder(root.Right)
	 }
 
	 return result
 }