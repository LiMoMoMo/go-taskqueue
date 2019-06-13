package taskqueue

import "fmt"

type Task struct {
	taskID  int
	msgChan chan MessageInterface
}

type TaskInterface interface {
	Run()
	Stop()

	SetChan(chan MessageInterface)
	SetTaskID(id int)
	GetTaskID() int
}

func (task *Task) SetChan(ch chan MessageInterface) {
	task.msgChan = ch
}

func (task *Task) SetTaskID(id int) {
	task.taskID = id
}

func (task *Task) GetTaskID() int {
	return task.taskID
}

func (task *Task) Run() {
	fmt.Printf("This is Task run function, task id is %d\r\n", task.taskID)
	task.Finish()
}

func (task *Task) Finish() {
	msg := &FinishMessage{Taskid: task.taskID}
	msg.Action = FINISH
	task.msgChan <- msg
}

func (task *Task) Stop() {

}
