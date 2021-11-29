package stack

import (
	"testing"
)

func TestEmptyStack(t *testing.T) {
	stk := NewStack[int]()
	if stk.IsEmpty() == false {
		t.Fatal("stack.IsEmpty() not returning correct status")
	}
}

func TestPushPop(t *testing.T) {
	stk := NewStack[int]()
	v := []int{1, 2, 3, 4, 5, 6}
	for _, x := range v {
		stk.Push(x)
	}
	for i := len(v) - 1; i >= 0; i-- {
		if z, alright := stk.Top(); alright {
			if z != v[i] {
				t.Fatal("stack.Top() not returning the right value")
			}
		}
		if y, ok := stk.Pop(); ok {
			if y != v[i] {
				t.Fatal("stack.{Push/Pop}() not operating properly")
			}
		} else {
			t.Fatal("stack.Pop() on non-empty stack returns not ok!")
		}
	}
	_, ok := stk.Pop()
	if ok {
		t.Fatal("stack.Pop() on empty stack return ok!")
	}
}
