/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isSymmetric(root *TreeNode) bool {
//     newTree := mirror(copy(root))
//     return same(root, newTree)
// }


func isSymmetric(root *TreeNode) bool {

    if root == nil {
        return true
    }

    return mirrorSame(root.Left, root.Right)
}


func mirrorSame(left, right *TreeNode) bool {

    if left == right && left == nil {
        return true
    }else if left == nil || right == nil {
        return false
    }

    if left.Val == right.Val {
        return mirrorSame(left.Left, right.Right) && mirrorSame(left.Right, right.Left)
    }else {
        return false
    }

}


func same(root, newroot *TreeNode) bool {
    if root == nil && newroot == nil {
        return true
    }else if root == nil || newroot == nil {
        return false
    }

    if root.Val != newroot.Val {
        return false
    }

    return same(root.Left, newroot.Left) && same(root.Right, newroot.Right)
}

func mirror(root *TreeNode) *TreeNode {

    if root == nil {
        return nil
    }

    temp := root.Left
    root.Left = root.Right
    root.Right = temp

    mirror(root.Left)
    mirror(root.Right)

    return root

}

func copy(root *TreeNode) *TreeNode {
    if root == nil {
        return nil 
    }

    newRoot := &TreeNode{Val:root.Val, Left:copy(root.Left), Right:copy(root.Right)}

    return newRoot

}