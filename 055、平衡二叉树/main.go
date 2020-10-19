/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }else {
        if depth(root.Left) - depth(root.Right) > 1 || depth(root.Left) - depth(root.Right) < -1 {
            return false
        }else {
            if isBalanced(root.Left) && isBalanced(root.Right) {
                return true
            }else {
                return false
            }
        }
    }
}

func depth(root *TreeNode) int {

    if root == nil {
        return 0
    }

    l := depth(root.Left)
    r := depth(root.Right)

    if l > r {
        return l + 1
    }else {
        return r + 1
    }
}