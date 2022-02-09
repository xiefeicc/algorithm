package problemtwo

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	dummp := &ListNode{}
	tmp := dummp
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		tmp.Next = &ListNode{
			Val: sum % 10,
		}
		tmp = tmp.Next
	}
	if carry > 0 {
		tmp.Next = &ListNode{
			Val: 1,
		}
	}
	return dummp.Next
}
