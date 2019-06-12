package taskqueue
import (
    "fmt"
    "testing"
)
func Test_New(t *testing.T) {
	tq := New(12)
	fmt.Printf("%v\r\n", tq)
}