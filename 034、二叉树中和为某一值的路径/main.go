/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var result  [][]int

 func pathSum(root *TreeNode, sum int) [][]int {
	 var path []int
 
	 if root != nil {
		 findPath(root, sum, path)
		 // 所有节点
		 // pathSum(root.Left, sum)
		 // pathSum(root.Right, sum)
	 }
 
	 return result
 }
 
 func findPath(root *TreeNode, sum int, path []int) bool {
 
 
	 if root != nil {
		 path = append(path, root.Val)
		 sum = sum - root.Val
		 // fmt.Println(path)
 
		 if sum == 0 {
			 // 到叶节点
			 if root.Left == root.Right && root.Left == nil {
				 result = append(result, path)
				 return true
			 }else {
				 path = path[:len(path) - 1]
				 return false
			 }
			 
		 }else if sum < 0 {
			 path = path[:len(path) - 1]
			 return false
		 }else {
			 left := findPath(root.Left, sum, path)
			 // 不拷贝的话，后面right会修改已经添加到result中的path
			 if left {
				 copy := make([]int,len(path))
				 for i := 0; i < len(copy) ; i++ {
					 copy[i] = path[i]
				 }
				 path = copy
			 }
 
			 right := findPath(root.Right, sum, path)
			 if  left || right {
				 return true
			 }else {
				 path = path[:len(path) - 1]
				 return false
			 }
		 }
	 }
 
	 if sum == 0 {
		 result = append(result, path)
		 return true
	 }else {
		 path = path[:len(path) - 1]
		 return false
	 }
 }