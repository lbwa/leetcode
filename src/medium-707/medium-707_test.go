package medium707

import (
	"leetcode-solutions/src/shared"
	"testing"
)

func TestSolution(t *testing.T) {
	list := CreateLinkedList()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	shared.Expect(t, list.Get(1), 2)
	list.DeleteAtIndex(1)
	shared.Expect(t, list.Get(1), 3)
}
