package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func NewList(elems []int) *ListNode {
	var result *ListNode = nil

	var current *ListNode
	for i, e := range elems {
		if i == 0 {
			result = &ListNode{
				Val: e,
			}
			current = result
		} else {
			current.Next = &ListNode{
				Val: e,
			}
			current = current.Next
		}

	}
	return result
}

func PrintList(list *ListNode) {
	fmt.Print("[")
	for list != nil {
		fmt.Print(list.Val, ",")
		list = list.Next
	}

	fmt.Println("]")
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var current, result, next *ListNode
	for !(list1 == nil && list2 == nil) {
		if list1 == nil {
			next = list2
			list2 = list2.Next
		} else if list2 == nil {
			next = list1
			list1 = list1.Next
		} else if list1.Val <= list2.Val {
			next = list1
			list1 = list1.Next
		} else {
			next = list2
			list2 = list2.Next
		}
		if result == nil {
			result = next
			current = result
		} else {
			current.Next = next
			current = current.Next
		}

	}
	return result
}

func main() {
	test_cases := [][]*ListNode{
		{NewList([]int{1,2,4}),NewList([]int{1,3,4})},
		{NewList([]int{}),NewList([]int{})},
		{NewList([]int{}),NewList([]int{0})},
	}
	for _, test_case := range test_cases {
		PrintList(mergeTwoLists(test_case[0], test_case[1]))
	}
}
