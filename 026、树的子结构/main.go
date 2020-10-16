/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  感觉leetcode测试用例有问题，[1,0,1,-4,-3]，[1,-4]  ------>????
 func isSubStructure(a *TreeNode, b *TreeNode) bool {

    if a == nil || b == nil {
        return false 
    }
    

    return hasSubTree(a, b) || hasSubTree(a.Left,b) || hasSubTree(a.Right, b)

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
    }else {
        return hasSubTree(a.Left,b) || hasSubTree(a.Right, b)
    }

}