package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue(5)

	if q.Front!=-1 || q.Rear!=-1 || len(q.Items)!=5{
		t.Errorf("create queue failed")
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue(5)

	q.Enqueue(10)
	q.Enqueue(12)
	q.Enqueue(14)
	q.Enqueue(16)

	want := 16
	if got, _ := q.GetRear(); got!=want {
		t.Errorf("q.GetRear()=%d, expected %d", got, want)
	}
}

func TestEmptyInsert(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue(10)

	if q.Rear!=q.Front {
		t.Errorf("queue front and rear should be same for the first insert")
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue(10)
	q.Enqueue(12)
	q.Enqueue(14)

	want := 10
	if got, _:=q.Dequeue(); got!=want {
		t.Errorf("q.Dequeue()=%d, expected %d", got, want)
	}
}

func TestDequeueError(t *testing.T) {
	q := NewQueue(5)
	q.Enqueue(10)
	q.Enqueue(12)
	q.Enqueue(14)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	want := false
	if _, got:=q.Dequeue(); got!=want {
		t.Errorf("q.Dequeue()=_, %t, expected %t", got, want)
	}
}
