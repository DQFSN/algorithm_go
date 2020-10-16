/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getKthFromEnd(head *ListNode, k int) *ListNode {

    src := head
    count := 0
    for head != nil {
        head = head.Next
        count++
        if count == k {
            break
        }
    }

    if count < k {
        return nil
    }

    for head != nil {
        head = head.Next
        src = src.Next
    }

    return src

}