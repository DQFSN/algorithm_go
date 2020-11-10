package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    res := make([][]int ,0,5)
    traversal(root,0,&res)

    return res

}

func traversal(root *TreeNode, l int, res *[][]int) {

    if root != nil {
        if len(*res) <= l {
            *res = append(*res,[]int{root.Val})
        }else {
            (*res)[l] = append((*res)[l],root.Val) 
        }

        traversal(root.Left,l+1,res)
        traversal(root.Right,l+1,res)
    }

}

func main()  {
    
}