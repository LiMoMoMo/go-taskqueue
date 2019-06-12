package taskqueue

type taskQueue struct {
	conNum int
}

func New(concurrentNum int) *taskQueue {
	return &taskQueue {
		conNum: concurrentNum,
	}
}

func (tq *taskQueue) Start() {

}

func (tq *taskQueue) Enqueue () (taskID int) {
	return 0
}

