package bags

import "errors"

type Stack interface {
	Push(s string)
	Pop() string
	IsEmpty() bool
	Size() int
}

type FixedCapacityStack struct {
	cap          int
	curr         int
	backingArray []string
}

func (fcs *FixedCapacityStack) Push(s string) {
	if fcs.curr == fcs.cap-1 {
		panic(errors.New("cannot push to stack as it is already full"))
	}
	fcs.backingArray[fcs.curr] = s
	fcs.curr++
}

func (fcs *FixedCapacityStack) Pop() string {
	if fcs.curr == 0 {
		panic(errors.New("cannot pop from stack with length 0"))
	}
	value := fcs.backingArray[fcs.curr-1]
	fcs.curr--
	return value
}

func (fcs *FixedCapacityStack) IsEmpty() bool {
	return fcs.curr == 0
}

func (fcs *FixedCapacityStack) Size() int {
	return fcs.curr
}

func NewStack(cap int) *FixedCapacityStack {
	arr := make([]string, cap)
	return &FixedCapacityStack{
		cap:          cap,
		backingArray: arr,
	}
}
