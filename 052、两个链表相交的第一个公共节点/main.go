/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
    ha := headA
    hb := headB
    for ha != nil || hb != nil {

        if ha == hb {
            return ha
        }

        if ha == nil {
            ha = headB
        }else {
            ha = ha.Next
        }
        if hb == nil {
            hb = headA
        }else {
            hb = hb.Next
        }
    }

    return nil
}