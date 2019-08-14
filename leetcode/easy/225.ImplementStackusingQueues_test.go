package easy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type MyStack struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.data) == 0 {
		return 0
	} else {
		temp := this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]
		return temp
	}
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.data) == 0 {
		return 0
	} else {
		return this.data[len(this.data)-1]
	}
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

func TestMyStack(t *testing.T) {
	Convey("test", t, func() {
		stack := Constructor()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		So(stack.Empty(), ShouldBeFalse)
		So(stack.Top(), ShouldEqual, 3)
		So(stack.Pop(), ShouldEqual, 3)
		So(stack.Top(), ShouldEqual, 2)
		stack.Pop()
		stack.Pop()
		So(stack.Top(), ShouldEqual, 0)
	})
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
