/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSubStructure(a *TreeNode, b *TreeNode) bool {

    if a == nil || b == nil {
        return false 
    }

    return hasSubTree(a, b) || isSubStructure(a.Left,b) || isSubStructure(a.Right, b)

}

func hasSubTree(a *TreeNode, b *TreeNode) bool {

    if a == nil && b == nil {
        return true
    }
    
    if a == nil {
        return false
    }

    if b == nil {
        return true
    }
    
    if a.Val == b.Val {
        return hasSubTree(a.Left,b.Left) && hasSubTree(a.Right, b.Right)
    }

    return false

}