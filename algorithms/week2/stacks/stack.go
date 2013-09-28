package stack

type Node struct {
	Item string
	Next *Node
}

type StackOfStrings struct {
	Head *Node
}

func NewStackOfStrings() *StackOfStrings {
	sos := StackOfStrings{Head: nil}
	return &sos
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.Head == nil {
		// stack underflow
		return "", false
	}
	r := s.Head.Item
	s.Head = s.Head.Next
	return r, true
}

func (s *StackOfStrings) Push(content string) {
	n := Node{Item: content, Next: s.Head}
	s.Head = &n
}
