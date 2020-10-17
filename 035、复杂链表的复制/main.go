/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

    src := head

    guard := &Node{}
    next := guard
    randMap := make(map[*Node]*Node)
    for head != nil {
        newNode := &Node{Val:head.Val}
        randMap[head] = newNode
        next.Next = newNode
        next = next.Next
        head = head.Next
    }

    head = src
    next = guard.Next
    for next != nil {
        next.Random = randMap[head.Random]
        head = head.Next
        next = next.Next
    }

    return guard.Next
}