func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head:=&ListNode{}
    resp:=head
    carry:=0
    sum:=0
    for l1!=nil || l2!=nil || carry>0 {
        sum=carry
        if l1!=nil {
            sum+=l1.Val
            l1=l1.Next
        }
        if l2!=nil {
            sum+=l2.Val
            l2=l2.Next
        }
        carry=sum/10
        sum=sum%10
        node:=&ListNode{sum,nil}
        head.Next=node
        head=head.Next
    }
    return resp.Next
}