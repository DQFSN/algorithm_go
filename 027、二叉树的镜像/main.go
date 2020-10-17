/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func mirrorTree(root *TreeNode) *TreeNode {

    if root != nil {
        temp := root.Left
        root.Left = root.Right
        root.Right = temp
        mirrorTree(root.Left)
        mirrorTree(root.Right)
    }

    return root

}