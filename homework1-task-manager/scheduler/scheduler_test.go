package scheduler

import (
	"sync"
	"testing"
	"time"
)

type testTask struct {
	id   int
	done chan int
}

func (t *testTask) Exec() {
	t.done <- t.id
}

// Поверяет выполнение задач по их времени
func TestSchedulerExecutesInOrder(t *testing.T) {
	s := NewScheduler()
	done := make(chan int, 2)
	now := time.Now()

	s.Add(&testTask{id: 1, done: done}, now.Add(50*time.Millisecond))
	s.Add(&testTask{id: 2, done: done}, now.Add(10*time.Millisecond))

	var order []int
	for i := 0; i < 2; i++ {
		order = append(order, <-done)
	}

	if order[0] != 2 || order[1] != 1 {
		t.Errorf("ожидался порядок выполнения [2,1], получен %v", order)
	}
}

// Проверяет, что задачи с истекшим временем выполняются без задержки
func TestSchedulerExecutesImmediate(t *testing.T) {
	s := NewScheduler()
	done := make(chan int, 1)

	s.Add(&testTask{id: 99, done: done}, time.Now().Add(-time.Second))

	select {
	case id := <-done:
		if id != 99 {
			t.Errorf("ожидался id 99, получен %d", id)
		}
	case <-time.After(20 * time.Millisecond):
		t.Error("задача с истекшим временем не выполнилась немедленно")
	}
}

// Проверяет добавление задач из нескольких горутин
func TestSchedulerConcurrentAdd(t *testing.T) {
	s := NewScheduler()
	done := make(chan int, 10)

	start := time.Now().Add(20 * time.Millisecond)
	spacing := 20 * time.Millisecond

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			execTime := start.Add(time.Duration(id) * spacing)
			s.Add(&testTask{id: id, done: done}, execTime)
		}(i)
	}
	wg.Wait()

	// Ожидаем id придут в порядке 0,1,...,9
	for expected := 0; expected < 10; expected++ {
		select {
		case id := <-done:
			if id != expected {
				t.Errorf("ожидался id %d, получен %d", expected, id)
			}

		case <-time.After(spacing + 30*time.Millisecond):
			t.Fatalf("таймаут при ожидании задачи %d", expected)
		}
	}
}
