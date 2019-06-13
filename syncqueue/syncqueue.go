package syncq

import (
	"fmt"
	"sync"
)

const basesize = 500

type SyncQueue struct {
	cond  *sync.Cond
	items []interface{}
}

// New return instance of syncqueue.
func New() *SyncQueue {
	q := &SyncQueue{
		cond:  sync.NewCond(&sync.Mutex{}),
		items: make([]interface{}, 0, basesize),
	}
	return q
}

// Enqueue into this queue.
func (q *SyncQueue) Enqueue(t interface{}) {
	withLock(q.cond.L, func() {
		if t == nil {
			fmt.Printf("t is nill\n")
		}
		q.items = append(q.items, t)
		q.cond.Signal()
	})
}

// Dequeue remove the first item of the queue.
func (q *SyncQueue) Dequeue() (t interface{}) {
	withLock(q.cond.L, func() {
		for len(q.items) <= 0 {
			q.cond.Wait()
		}
		t = q.items[0]
		q.items = q.items[1:]
	})
	return
}

// Front get the first item of the queue.
func (q *SyncQueue) Front() (t interface{}) {
	withLock(q.cond.L, func() {
		for len(q.items) <= 0 {
			q.cond.Wait()
		}
		t = q.items[0]
	})
	return
}

// IsEmpty whether the queue is empty.
func (q *SyncQueue) IsEmpty() (b bool) {
	withLock(q.cond.L, func() {
		if len(q.items) == 0 {
			b = true
		} else {
			b = false
		}
	})
	return
}

// Size get the length of the queue.
func (q *SyncQueue) Size() (length int) {
	withLock(q.cond.L, func() {
		length = len(q.items)
	})
	return
}

func withLock(lk sync.Locker, fn func()) {
	lk.Lock()
	defer lk.Unlock() // in case fn panics
	fn()
}
