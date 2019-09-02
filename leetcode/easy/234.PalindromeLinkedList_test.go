package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome2(head *ListNode) bool {
	var values []int
	tempHead := head
	for tempHead != nil {
		values = append(values, tempHead.Val)
		tempHead = tempHead.Next
	}

	tempHead = head
	for i := len(values) - 1; i >= len(values)/2; {
		if tempHead.Val != values[i] {
			return false
		}

		i--
		tempHead = tempHead.Next
	}

	return true
}

func isPalindrome3(head *ListNode) bool {
	// count the length
	var length int
	tempHead := head
	for tempHead != nil {
		length++
		tempHead = tempHead.Next
	}

	// reverse half of the list
	var reversedHead, next *ListNode
	for i := 0; i < length/2; i++ {
		next = head.Next
		head.Next = reversedHead
		reversedHead = head
		head = next
	}

	// skip odd head
	if length&1 == 1 {
		head = head.Next
	}

	// compare reverse with the other half
	for i := 0; i < length/2; i++ {
		if reversedHead.Val != head.Val {
			return false
		}
		reversedHead = reversedHead.Next
		head = head.Next
	}

	return true
}

func TestIsPalindrome2(t *testing.T) {
	Convey("test", t, func() {
		So(isPalindrome3(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		}), ShouldBeTrue)

		So(isPalindrome3(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		}), ShouldBeFalse)

		So(isPalindrome3(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		}), ShouldBeTrue)
	})
}
