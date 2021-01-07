package collection

import "errors"

type Stack []interface{}

// Init stack
func InitStack() Stack {
	var s Stack
	return s
}

// Check if stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new element to stack.
func (s *Stack) Push(e interface{}) {
	*s = append(*s, e)
}

// Remove and return top element of stack.
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, nil
	}
}

// How to Use it
//stack:=collection.InitStack()
//stack.Push("element")
//item,err:=stack.Pop()
//if err!=nil{
//	fmt.Println(err)
//}
//str:=item.(string)
