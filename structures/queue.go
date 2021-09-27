package structures

import "sync"

type Queue struct {
	// Slice of type string, it holds items in stack.
	items []string
	// rwLock for handling concurrent operations on the stack.
	rwLock sync.RWMutex
}

func (queue *Queue) New() *Queue {

	queue.items = []string{}

	return queue
}

func (queue *Queue) Enqueue(t string) {
	//Initialize items slice if not initialized
	if queue.items == nil {
		queue.items = []string{}
	}
	// Acquire read, write lock before inserting a new item in the stack.
	queue.rwLock.Lock()
	// Performs append operation.
	queue.items = append(queue.items, t)
	// This will release read, write lock
	queue.rwLock.Unlock()
}

func (queue *Queue) Dequeue() *string {
	// Checking if stack is empty before performing pop operation
	if len(queue.items) == 0 {
		return nil
	}
	// Acquire read, write lock as items are going to modify.
	queue.rwLock.Lock()
	// Popping item from items slice.
	item := queue.items[0]
	//Adjusting the item's length accordingly
	queue.items = queue.items[1:]
	// Release read write lock.
	queue.rwLock.Unlock()
	// Return last popped item
	return &item
}

func (queue *Queue) Size() int {
	// Acquire read lock
	queue.rwLock.RLock()
	// defer operation of unlock.
	defer queue.rwLock.RUnlock()
	// Return length of items slice.
	return len(queue.items)
}

func (queue *Queue) All() []string {
	// Acquire read lock
	queue.rwLock.RLock()
	// defer operation of unlock.
	defer queue.rwLock.RUnlock()
	// Return items slice to caller.
	return queue.items
}

func (queue *Queue) IsEmpty() bool {
	// Acquire read lock
	queue.rwLock.RLock()
	// defer operation of unlock.
	defer queue.rwLock.RUnlock()
	return len(queue.items) == 0
}