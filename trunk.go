package geodb

type trunk struct {
	branch
}

func (t *trunk) Insert(key string, s Shape) {
	e := makeEntry(key, s)
	t.branch.Insert(&e)
}

func (t *trunk) Split() (out []node) {
	out = t.branch.Split()
	if t.state < branchChildrenHasBranches {
		t.state++
	}
	return
}
