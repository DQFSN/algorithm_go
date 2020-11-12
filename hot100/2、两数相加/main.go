
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    carry := 0

    res := &ListNode{}
    head := res
    for l1 != nil || l2 != nil {
        sum := 0
        if l1 == nil {
            sum = l2.Val + carry
            l2 = l2.Next
        }else if l2 == nil {
            sum = l1.Val + carry
            l1 = l1.Next
        }else {
            sum = l2.Val + l1.Val + carry
            l1 = l1.Next
            l2 = l2.Next
        }
        
        carry = sum / 10
        a := sum % 10
        node := &ListNode{a,nil}
        head.Next = node
        head = node
    }

    if carry > 0 {
        node := &ListNode{carry,nil}
        head.Next = node
        head = node
    }

    return res.Next
    
}