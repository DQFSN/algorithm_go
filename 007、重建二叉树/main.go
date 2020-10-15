package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {

    if len(preorder) != len(inorder) || len(preorder) == 0 {
        return nil
    }

    rootVal := preorder[0]
    rootIndex := -1

    for i, v := range inorder {
        if v == rootVal {
            rootIndex = i
        }
    }

    // lLen := rootIndex
    // rLen := len(inorder) - 1 - rootIndex

    lPre, rPre := preorder[1:rootIndex+1], preorder[rootIndex+1:]
    lIn, rIn := inorder[:rootIndex], inorder[rootIndex+1:]

    root := &TreeNode{Val:rootVal}
    root.Left = buildTree(lPre, lIn)
    root.Right = buildTree(rPre, rIn)

    return root


}