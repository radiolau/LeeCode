package main

/**
@link https://leetcode-cn.com/problems/add-two-numbers/
 */

type ListNode struct {
	Var int
	Next *ListNode
}

/**
链表逆序 => 队列；
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList *ListNode = nil

	l1 = linkReverse(l1, nil)
	l2 = linkReverse(l2, nil)

	if l1 != nil || l2 != nil {
		var lastDiv = 0
		var mod = 0
		for ; ; {
			mod, lastDiv = addByLink(l1, l2, lastDiv)
			var node = new(ListNode)
			node.Var = mod
			node.Next = nil
			newList.Next = node

			if l1.Next != nil || l2.Next != nil {
				if nil != l1.Next {
					l1 = l1.Next
				}
				if nil != l2.Next {
					l2 = l2.Next
				}
			} else {
				break
			}
		}
	}

	return newList
}

func addByLink(l1 *ListNode, l2 *ListNode, lastDiv int) (mod int, div int) {
	var sum = lastDiv

	if nil != l1 {
		sum += l1.Var
	}

	if nil != l2 {
		sum += l2.Var
	}

	return sum % 10, sum / 10
}

/**
返回逆序后的链表头节点
*/
func linkReverse(headNode *ListNode, lastNode *ListNode) *ListNode {
	if nil == headNode {
		//println("end")
		return headNode
	}

	var stgPtr = headNode.Next
	headNode.Next = lastNode
	lastNode = headNode

	//printLink(headNode)
	return linkReverse(stgPtr, lastNode)
}

func printLink(head *ListNode)  {
	for ; head != nil;  {
		print( " ", head.Var)
		head = head.Next
	}
	println()
}

func main() {
	var A = ListNode{Var: 3, Next: nil}
	var B = ListNode{Var: 1, Next: nil}
	var C = ListNode{Var: 5, Next: nil}
	var D = ListNode{Var: 3, Next: nil}
	var E = ListNode{Var: 9, Next: nil}
	A.Next = &B
	B.Next = &C
	C.Next = &D
	D.Next = &E

	//printLink(&A)

	printLink(linkReverse(&A, nil))
}
