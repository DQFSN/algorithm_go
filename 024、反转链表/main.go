/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func reverseList(head *ListNode) *ListNode {
    pre := (*ListNode)(nil)
    cru := head
    next := (*ListNode)(nil)
    
    for cru != nil {
        next = cru.Next
        cru.Next = pre
        pre = cru
        cru = next
    } 

    return pre
}


// func reverseList(head *ListNode) *ListNode {
//     guard := &ListNode{}

//     next := (*ListNode)(nil) 

//     for head != nil {
//         next = head.Next
//         head.Next = guard.Next
//         guard.Next = head
//         head = next
//     }

//     return guard.Next
// }