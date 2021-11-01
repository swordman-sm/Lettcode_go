package t2AddTwoNums

/**
 * You are given two non-empty linked lists representing two non-negative integers. The digits are stored in
 * reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked
 * list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 * Example:
 * 题⽬⼤意
 * 2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点。
 * <p>
 * Example:
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 * Explanation: 342 + 465 = 807.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNums(first *ListNode, second *ListNode) *ListNode {
	//1.初始化一个临时虚拟节点(0)
	dummy := &ListNode{
		Val: 0,
	}
	//2.定义临时变量
	current := dummy
	//3.定义每一轮的余数(进位)
	carry := 0
	for first != nil || second != nil {
		x := 0
		if first != nil {
			x = first.Val
		}
		y := 0
		if second != nil {
			y = second.Val
		}
		//补上一轮进位的余数
		sum := carry + x + y
		carry = sum / 10
		//下一节点指向本轮各节点和的商值
		current.Next = &ListNode{
			Val: sum % 10,
		}
		//开启下一轮循环(节点位置向下移动)
		current = current.Next
		if first != nil {
			first = first.Next
		}
		if second != nil {
			second = second.Next
		}
	}
	//余数超过0进位末位节点
	if carry > 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}
	return dummy.Next
}
