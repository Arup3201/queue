package queue

import "testing"

func TestEnqueue(t *testing.T) {
	t.Run("enqueue rear is 10 in queue [10]", func(t *testing.T) {
		q := MakeQueue()

		q.Enqueue(10)

		want := 10
		if got := q.GetRear(); got != want {
			t.Errorf("TestEnqueue expected %d, got %d", want, got)
		}
	})
	t.Run("enqueue rear is 30 in queue [10, 20, 30]", func(t *testing.T) {
		q := MakeQueue()

		q.Enqueue(10)
		q.Enqueue(20)
		q.Enqueue(30)

		want := 30
		if got := q.GetRear(); got != want {
			t.Errorf("TestEnqueue expected %d, got %d", want, got)
		}
	})
	t.Run("enqueue front is 10 in queue [10, 20, 30]", func(t *testing.T) {
		q := MakeQueue()

		q.Enqueue(10)
		q.Enqueue(20)
		q.Enqueue(30)

		want := 10
		if got := q.GetFront(); got != want {
			t.Errorf("TestEnqueue expected %d, got %d", want, got)
		}
	})
}

func TestDequeue(t *testing.T) {
	t.Run("dequeue item is 10 in queue [10, 20, 30]", func(t *testing.T) {
		q := MakeQueue()
		q.Enqueue(10)
		q.Enqueue(20)
		q.Enqueue(30)

		item, _ := q.Dequeue()

		want := 10
		if item != want {
			t.Errorf("TestDequeue expected %d, got %d", want, item)
		}
	})
	t.Run("after dequeue front is 20 in queue [10, 20, 30]", func(t *testing.T) {
		q := MakeQueue()
		q.Enqueue(10)
		q.Enqueue(20)
		q.Enqueue(30)

		q.Dequeue()

		want := 20
		if got := q.GetFront(); got != want {
			t.Errorf("TestDequeue expected %d, got %d", want, got)
		}
	})
	t.Run("dequeue empty in queue [10]", func(t *testing.T) {
		q := MakeQueue()
		q.Enqueue(10)

		q.Dequeue()

		if got := q.GetFront(); got != nil {
			t.Errorf("TestDequeue expected nil, got %d", got)
		}
	})
}
