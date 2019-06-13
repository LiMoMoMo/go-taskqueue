package taskqueue

import (
	"fmt"
	"testing"
	"time"
)

type TestTask struct {
	Task
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
	Task
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

func Test_New(t *testing.T) {
	tq := New(2)
	tq.Start()
	fmt.Printf("%v\r\n", tq)
	task1 := TestTask{}
	task2 := DemoTask{}
	tq.Enqueue(&task1)
	tq.Enqueue(&task2)
	time.Sleep(30 * time.Second)
}
