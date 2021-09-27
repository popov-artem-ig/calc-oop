package structures

import (
	"fmt"
	"strings"
	"sync"
)

type Stack struct {
	// Slice of type string, it holds items in stack.
	items []string
	// rwLock for handling concurrent operations on the stack.
	rwLock sync.RWMutex
}

func (stack *Stack) New() *Stack {

	stack.items = []string{}

	return stack
}

func (stack *Stack) Push(t string) {
	if stack.items == nil {
		stack.items = []string{}
	}
	stack.rwLock.Lock()
	stack.items = append(stack.items, t)
	stack.rwLock.Unlock()
}

func (stack *Stack) Pop() *string {
	if len(stack.items) == 0 {
		return nil
	}
	stack.rwLock.Lock()
	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[0 : len(stack.items)-1]
	stack.rwLock.Unlock()
	return &item
}

func (stack *Stack) Size() int {
	stack.rwLock.RLock()
	defer stack.rwLock.RUnlock()
	return len(stack.items)
}

func (stack *Stack) All() []string {
	stack.rwLock.RLock()
	defer stack.rwLock.RUnlock()
	return stack.items
}

func (stack *Stack) IsEmpty() bool {
	stack.rwLock.RLock()
	defer stack.rwLock.RUnlock()
	return len(stack.items) == 0
}

func (stack *Stack) Peek() *string {
	if len(stack.items) == 0 {
		return nil
	}
	stack.rwLock.Lock()
	item := stack.items[len(stack.items)-1]
	stack.rwLock.Unlock()
	return &item
}

func (stack *Stack) ToString() (string, error) {
	if len(stack.items) > 0 {
		return strings.Join(stack.items, " "), nil
	}
	return "", fmt.Errorf("error. current value empty")
}