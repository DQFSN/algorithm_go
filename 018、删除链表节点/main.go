/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteNode(head *ListNode, val int) *ListNode {

    pre := (*ListNode)(nil)
    cur := head
    for cur != nil {
        
        if cur.Val != val {
            pre = cur
            cur = cur.Next
        }else {
            if pre == nil {
                return cur.Next
            }
            pre.Next = cur.Next
            return head
        }
    }

    return head
}