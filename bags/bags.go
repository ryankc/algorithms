package bags

type Stack interface {
	Push(s string)
	Pop() string
	IsEmpty() bool
	Size() int
}

type FixedCapacityStack struct {
	cap int
	curr int
}

func (fcs *FixedCapacityStack) Push(s string) {

}

func (fcs *FixedCapacityStack) Pop() string {
	return ""
}

func (fcs *FixedCapacityStack) IsEmpty() bool {
	return false
}

func (fcs *FixedCapacityStack) Size() int {
	return fcs.curr
}

func NewStack(cap int) *FixedCapacityStack {
	return &FixedCapacityStack{
		cap: cap,
	}
}
