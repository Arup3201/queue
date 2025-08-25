package queue

type Queue struct {
	Front int
	Rear int
	Max int
	Items []int
}

func NewQueue(m int) *Queue {
	return &Queue{
		Front: -1,
		Rear: -1, 
		Max: m, 
		Items: make([]int, m),
	}
}

func (q *Queue) IsEmpty() bool {
	return q.Front == -1
}

func (q *Queue) GetRear() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}

	return q.Items[q.Rear], true
}

func (q *Queue) Enqueue(item int) {
	if q.IsEmpty() {
		q.Front=0
		q.Rear=0
	} else {
		q.Rear += 1
	}

	if q.Rear==q.Max {
		panic("reached the maximum capacity of the queue")
	}

	q.Items[q.Rear] = item
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}

	item := q.Items[q.Front]
	if q.Front==q.Rear {
		q.Front=-1
		q.Rear=-1
	} else {
		q.Front += 1
	}

	return item, true
}
