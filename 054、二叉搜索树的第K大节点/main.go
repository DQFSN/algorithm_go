/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
    if k == 0 {
        return 0
    }
    var result *[]int = new([]int)
    Inorder(root, result)

    if len(*result) <= 0 {
        return 0
    }
    return (*result)[len(*result) - k]
}


func Inorder(root *TreeNode, result *[]int) {

    if root != nil {
        Inorder(root.Left, result)
        *result = append(*result, root.Val)
        Inorder(root.Right, result)
    }

}