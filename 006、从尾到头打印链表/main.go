package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reversePrint(head *ListNode) []int {
    guard := &ListNode{}

    count := 0
    var next *ListNode = nil
    for head != nil {
        count++
        next = head.Next
        head.Next = guard.Next
        guard.Next = head
        head = next
        
    }

    ret := make([]int, count, count)

    index := 0
    for guard.Next != nil {
        ret[index] = guard.Next.Val
        index++
        guard = guard.Next
    }

    return ret
}

func reversePrintBetter(head *ListNode) []int {

    src := head
    count := 0
    for head != nil {
        count++
        head = head.Next
        
    }

    ret := make([]int, count, count)

    head = src
    for head != nil {
        ret[count - 1] = head.Val
        count--
        head = head.Next
    }

    return ret
}