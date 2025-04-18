package scheduler

import (
	"container/heap"
	"time"
)

type Task interface {
	Exec()
}

// Обертка над Task с добавлением времени и индексом для кучи
type scheduledTask struct {
	execTime time.Time
	task     Task
	index    int
}

// Куча задач, отсортированная по execTime.
type taskHeap []*scheduledTask

// Количество элементов в куче
func (h taskHeap) Len() int {
	return len(h)
}

// Определяет приотритет задач (меньшее время - высокий приоритет)
func (h taskHeap) Less(i, j int) bool {
	return h[i].execTime.Before(h[j].execTime)
}

// Меняет местами элементы и их индексы
func (h taskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

// Добавляет новый элемент в кучу
func (h *taskHeap) Push(x interface{}) {
	item := x.(*scheduledTask)
	item.index = len(*h)
	*h = append(*h, item)
}

// Извлекает последний элемент из кучи.
func (h *taskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.index = -1
	*h = old[0 : n-1]
	return item
}

// Хранит очередь задач
// newTaskChan канал для приёма новых элементов
type Scheduler struct {
	newTaskChan chan *scheduledTask
}

// Cоздаёт новый планировщик и запускает его
func NewScheduler() *Scheduler {
	s := &Scheduler{
		newTaskChan: make(chan *scheduledTask, 1024),
	}
	go s.run()
	return s
}

// Добавляет новую задачу в планировщик
func (s *Scheduler) Add(task Task, t time.Time) {
	st := &scheduledTask{execTime: t, task: task}
	s.newTaskChan <- st
}

// Цикл планировщика - управляет кучей с одним таймером
func (s *Scheduler) run() {
	var pq taskHeap
	heap.Init(&pq)
	var timer *time.Timer
	var timerChan <-chan time.Time

	for {
		if pq.Len() == 0 {
			st := <-s.newTaskChan
			heap.Push(&pq, st)
			continue
		}

		now := time.Now()
		next := pq[0]
		if !next.execTime.After(now) {
			heap.Pop(&pq)
			go next.task.Exec()
			continue
		}

		delay := next.execTime.Sub(now)
		if timer == nil {
			timer = time.NewTimer(delay)
			timerChan = timer.C
		} else {
			timer.Reset(delay)
		}

		select {
		case st := <-s.newTaskChan:
			heap.Push(&pq, st)
			if st.execTime.Before(pq[0].execTime) {
				timer.Reset(st.execTime.Sub(time.Now()))
			}

		case <-timerChan:
		}
	}
}
