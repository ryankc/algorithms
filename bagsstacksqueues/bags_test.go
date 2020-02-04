package bagsstacksqueues

import (
	"github.com/ryankc/algorithms/assert"
	"testing"
)

func TestStackSize(t *testing.T) {
	var capacity = 2
	stack := NewStack(capacity)
	assert.Equal(t, 0, stack.Size())
}

func TestPushAndThenPop(t *testing.T) {
	stack := NewStack(20)
	stack.Push("one")
	stack.Push("two")
	assert.Equal(t, "two", stack.Pop())
}
