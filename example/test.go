package main

import (
	"fmt"
	"time"

	tq "github.com/LiMoMoMo/go-taskqueue"
)

type TestTask struct {
	tq.Task
}

func (t *TestTask) Run() {
	count := 10
	for count > 0 {
		fmt.Printf("This is TeskTask, count is %d\r\n", count)
		count--
		time.Sleep(1 * time.Second)
	}
	t.Finish()
}

type DemoTask struct {
	tq.Task
}

func (t *DemoTask) Run() {
	count := 5
	for count > 0 {
		fmt.Printf("This is DemoTask, count is %d\r\n", count)
		count--
		time.Sleep(1 * time.Second)
	}
	t.Finish()
}

type QWETask struct {
	tq.Task
}

func (t *QWETask) Run() {
	count := 5
	for count > 0 {
		fmt.Printf("This is QWETask, count is %d\r\n", count)
		count--
		time.Sleep(1 * time.Second)
	}
	t.Finish()
}

func main() {
	tq := tq.New(2)
	tq.Start()
	task1 := TestTask{}
	task2 := DemoTask{}
	task3 := QWETask{}
	tq.Enqueue(&task1)
	tq.Enqueue(&task2)
	tq.Enqueue(&task3)
	time.Sleep(30 * time.Second)
}
