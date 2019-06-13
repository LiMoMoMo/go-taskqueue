package taskqueue

// ActionType for Message
type ActionType int32

const (
	// DISPATCH task
	DISPATCH ActionType = 0
	// DELETE task
	DELETE ActionType = 1
	// FINISH task
	FINISH ActionType = 2
)

type MessageInterface interface {
	GetAction() ActionType
}

// Message connmunicate with TaskQueue.
type Message struct {
	Action ActionType
}

func (m *Message) GetAction() ActionType {
	return m.Action
}

// FinishMessage means the task has finished.
type FinishMessage struct {
	Message
	Taskid int
}
