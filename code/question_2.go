package main

/**
@link https://leetcode-cn.com/problems/add-two-numbers/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
链表逆序 => 队列；
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 && nil == l2 {
		return nil
	}

	var sumList *ListNode = nil
	l1 = linkReverse(l1, nil)
	l2 = linkReverse(l2, nil)

	var tmpPtr = sumList
	var node *ListNode = nil
	var lastDiv = 0
	var mod = 0
	for ; ; {
		if (nil == l1) && (nil == l2) && (0 == lastDiv) {
			break
		}

		mod, lastDiv = addByLink(l1, l2, lastDiv)
		node = new(ListNode)
		node.Val = mod
		node.Next = nil

		if nil == tmpPtr {
			sumList = node
			tmpPtr = sumList
		} else {
			tmpPtr.Next = node
			tmpPtr = tmpPtr.Next
		}

		if nil != l1 {
			l1 = l1.Next
		}
		if nil != l2 {
			l2 = l2.Next
		}
	}

	return sumList
}

func addByLink(l1 *ListNode, l2 *ListNode, lastDiv int) (mod int, div int) {
	var sum = lastDiv
	if nil != l1 {
		sum += l1.Val
	}
	if nil != l2 {
		sum += l2.Val
	}
	return sum % 10, sum / 10
}

/**
返回逆序后的链表头节点
*/
func linkReverse(headNode *ListNode, lastNode *ListNode) *ListNode {
	if nil == headNode {
		return headNode
	}

	if nil == headNode.Next {
		headNode.Next = lastNode
		return headNode
	}

	var stgPtr = headNode.Next
	headNode.Next = lastNode
	lastNode = headNode
	return linkReverse(stgPtr, lastNode)
}

func printList(head *ListNode)  {
	for ; head != nil;  {
		print( " ", head.Val)
		head = head.Next
	}
	println()
	//println("end of print link list")
}

func createLinkListByArray(nums []int) *ListNode {
	var list *ListNode = nil
	var tmpNode *ListNode = nil
	var i = 0
	//println("len : ", len(nums))
	for ; i < len(nums); i++ {
		var node = new(ListNode)
		node.Val = nums[i]
		node.Next = nil
		if nil == list {
			list = node
			tmpNode = list
		} else {
			tmpNode.Next = node
			tmpNode = tmpNode.Next
		}
	}

	return list
}

func main() {

	var l1, l2, sumList *ListNode

	l1 = createLinkListByArray([]int {2, 4, 3})
	l2 = createLinkListByArray([]int {5, 6, 4})
	sumList = addTwoNumbers(l1, l2)
	printList(sumList)

	l1 = createLinkListByArray([]int {0})
	l2 = createLinkListByArray([]int {0})
	sumList = addTwoNumbers(l1, l2)
	printList(sumList)

	l1 = createLinkListByArray([]int {9,9,9,9,9,9,9})
	l2 = createLinkListByArray([]int {9,9,9,9})
	sumList = addTwoNumbers(l1, l2)
	printList(sumList)
}
