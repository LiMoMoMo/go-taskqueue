package taskqueue

type Task struct {
	TaskID int
}

type TaskInterface interface {
	run()
	stop()
}