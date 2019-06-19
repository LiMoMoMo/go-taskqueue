package taskqueue

import (
	"sync"

	"github.com/LiMoMoMo/go-taskqueue/syncqueue"
)

// TaskQueue is a queue of tasks.
type TaskQueue struct {
	conNum    int
	taskCount int
	taskq     *syncq.SyncQueue
	msgChan   chan MessageInterface
	countLck  sync.Mutex
	taskm     map[int]TaskInterface
}

// New return instance of taskQueue
func New(concurrentNum int) *TaskQueue {
	tq := &TaskQueue{
		conNum:    concurrentNum,
		taskCount: 0,
		taskq:     syncq.New(),
		msgChan:   make(chan MessageInterface, 10),
		taskm:     make(map[int]TaskInterface),
	}
	return tq
}

// Start the task management.
func (tq *TaskQueue) Start() {
	go tq.run()
}

func (tq *TaskQueue) genID() int {
	tq.countLck.Lock()
	defer tq.countLck.Unlock()
	tq.taskCount++
	return tq.taskCount
}

// Enqueue insert task into the queue.
func (tq *TaskQueue) Enqueue(t TaskInterface) (taskID int) {
	t.SetChan(tq.msgChan)
	t.SetTaskID(tq.genID())
	tq.taskq.Enqueue(t)
	tq.msgChan <- &Message{Action: DISPATCH}
	return t.GetTaskID()
}

func (tq *TaskQueue) run() {
	for {
		select {
		case message := <-tq.msgChan:
			switch message.GetAction() {
			case DISPATCH:
				if len(tq.taskm) >= tq.conNum {
					continue
				}
				task := tq.taskq.Dequeue()
				tq.taskm[task.(TaskInterface).GetTaskID()] = task.(TaskInterface)
				// go task.(TaskInterface).Run()
				// go task.(TaskInterface).Start()
				go func() {
					task.(TaskInterface).Run()
					task.(TaskInterface).Finish()
				}()
			case FINISH:
				tid := message.(*FinishMessage).Taskid
				delete(tq.taskm, tid)
				tq.msgChan <- &Message{Action: DISPATCH}
			}
		}
	}
}
