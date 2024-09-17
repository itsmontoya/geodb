package geodb

import "fmt"

const (
	branchHasLeaves branchState = iota
	branchChildrenHasLeaves
	branchChildrenHasBranches
)

type branchState uint8

func (b branchState) String() string {
	switch b {
	case branchHasLeaves:
		return "has leaves"
	case branchChildrenHasLeaves:
		return "children has leaves"
	case branchChildrenHasBranches:
		return "children has branches"
	default:
		return fmt.Sprintf("<unsupported (%d)>", b)
	}
}
