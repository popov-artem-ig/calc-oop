package structures

import (
	"fmt"
	"strings"
	"sync"
)

type Queue struct {
	// Slice of type string, it holds items in Queue.
	items []string
	// rwLock for handling concurrent operations on the Queue.
	rwLock sync.RWMutex
}

func (queue *Queue) New() *Queue {

	queue.items = []string{}

	return queue
}

func (queue *Queue) Enqueue(t string) {
	if queue.items == nil {
		queue.items = []string{}
	}
	queue.rwLock.Lock()
	queue.items = append(queue.items, t)
	queue.rwLock.Unlock()
}

func (queue *Queue) Dequeue() *string {
	if len(queue.items) == 0 {
		return nil
	}
	queue.rwLock.Lock()
	item := queue.items[0]
	queue.items = queue.items[1:]
	queue.rwLock.Unlock()
	return &item
}

func (queue *Queue) Size() int {
	queue.rwLock.RLock()
	defer queue.rwLock.RUnlock()
	return len(queue.items)
}

func (queue *Queue) All() []string {
	queue.rwLock.RLock()
	defer queue.rwLock.RUnlock()
	return queue.items
}

func (queue *Queue) IsEmpty() bool {
	queue.rwLock.RLock()
	defer queue.rwLock.RUnlock()
	return len(queue.items) == 0
}

func (queue *Queue) ToString() (string, error) {
	if len(queue.items) > 0 {
		return strings.Join(queue.items, " "), nil
	}
	return "", fmt.Errorf("error. current value empty")
}