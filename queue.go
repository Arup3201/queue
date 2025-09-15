package queue

import "sync"

type elementType any

type queue struct {
	front, rear int
	items       []elementType
	rwLock      sync.RWMutex
}

func MakeQueue() *queue {
	return &queue{
		front: -1,
		rear:  -1,
		items: make([]elementType, 0),
	}
}

func (q *queue) Enqueue(item elementType) {
	q.rwLock.Lock()
	defer q.rwLock.Unlock()

	if q.rear == -1 {
		q.front = 0
		q.rear = 0
	} else {
		q.rear += 1
	}

	q.items = append(q.items, item)
}

func (q *queue) Dequeue() (elementType, bool) {
	q.rwLock.Lock()
	defer q.rwLock.Unlock()

	switch q.front {
	case -1:
		return nil, false
	case q.rear:
		item := q.items[q.front]
		q.front = -1
		q.rear = -1
		return item, true
	default:
		item := q.items[q.front]
		q.front++
		return item, true
	}
}

func (q *queue) GetFront() elementType {
	if q.front == -1 {
		return nil
	}
	return q.items[q.front]
}

func (q *queue) GetRear() elementType {
	if q.rear == -1 {
		return nil
	}
	return q.items[q.rear]
}
