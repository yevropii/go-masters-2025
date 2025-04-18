package main

import (
	"fmt"
	"time"

	"github.com/yevropii/go-masters-2025/homework1-task-manager/scheduler"
)

// Простая структура для демонстрации
type PrintTask struct {
	Message string
}

// Для демонстрации
func (t *PrintTask) Exec() {
	fmt.Printf("%s: %s\n", time.Now().Format(time.RFC3339), t.Message)
}

func main() {
	sched := scheduler.NewScheduler()

	sched.Add(&PrintTask{Message: "Task #1 after 1 сек"}, time.Now().Add(1*time.Second))
	sched.Add(&PrintTask{Message: "Task #2 after 2 сек"}, time.Now().Add(2*time.Second))
	sched.Add(&PrintTask{Message: "Task #3 через 3 сек"}, time.Now().Add(3*time.Second))
	sched.Add(&PrintTask{Message: "Task #4 через 4 сек"}, time.Now().Add(4*time.Second))
	sched.Add(&PrintTask{Message: "Task #5 через 5 сек"}, time.Now().Add(5*time.Second))

	time.Sleep(10 * time.Second)
}
