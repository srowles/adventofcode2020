package common

// IntStack is an array backed stack of ints
type IntStack []int

// IsEmpty checks if stack is empty
func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *IntStack) Push(val int) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
}

// Pop and return top element of stack. Panics if empty
func (s *IntStack) Pop() int {
	if s.IsEmpty() {
		panic("empty")
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}

// Peek looks at the top element in the stack, panics if empty
func (s *IntStack) Peek() int {
	if s.IsEmpty() {
		panic("peek empty")
	}

	return (*s)[len(*s)-1]
}
