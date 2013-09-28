package stack

import (
	"fmt"
	"testing"
)

func Test_Stack(t *testing.T) {
	input := []string{
		"to",
		"be",
		"or",
		"not",
		"to", "-",
		"be", "-", "-",
		"that", "-", "-", "-",
		"is",
	}

	stack := NewStackOfStrings()
	for i := range input {
		s := input[i]
		if s == "-" {
			p, ok := stack.Pop()
			if ok {
				fmt.Println("popped", p)
			} else {
				t.Error("stack underflow")
			}
		} else {
			fmt.Println("pushing", s)
			stack.Push(s)
		}
	}
	// expecting two elements on the stack now
	p, ok := stack.Pop()
	if !ok {
		t.Error("unexepected stack underflow")
	}
	if p != "is" {
		t.Error("wrong element on top of stack", p)
	}
	p, ok = stack.Pop()
	if !ok {
		t.Error("unexepected stack underflow")
	}
	if p != "to" {
		t.Error("wrong element on bottom of stack", p)
	}
	p, ok = stack.Pop()
	if ok {
		t.Error("exepected stack underflow didn't happen")
	}
	if p != "" {
		t.Error("should have been a stack underflow", p)
	}
}
