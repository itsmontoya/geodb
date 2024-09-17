package geodb

import "fmt"

var _ node = &branch{}

type branch struct {
	rect  Rect
	nodes []node

	state branchState
}

func (b *branch) Append(n node) {
	b.nodes = append(b.nodes, n)
}

func (b *branch) Insert(v value) {
	rect := v.Rect()
	match, i := b.getMatch(&rect)
	match.Insert(v)
	b.rect.SetRect(rect)
	if match.Len() <= 8 {
		return
	}

	ns := match.Split()
	ns = append(ns, b.nodes[i+1:]...)
	b.nodes = append(b.nodes[:i], ns...)
}

func (b *branch) Len() int {
	return len(b.nodes)
}

func (b *branch) Split() (out []node) {
	var b1, b2 branch
	b1.rect, b2.rect = b.rect.Split()
	b1.state = b.state
	b2.state = b.state
	for _, n := range b.nodes {
		nr := n.Rect()
		node, ok := nr.GetMatchingNode(&b1, &b2)
		if !ok {
			msg := fmt.Sprintf("no split match found for %+v", nr)
			panic(msg)
		}

		node.Append(n)
	}

	out = append(out, &b1)
	out = append(out, &b2)
	return
}

func (b *branch) AppendMatches(in []*entry, c Coordinates) (out []*entry) {
	out = in
	for _, n := range b.nodes {
		out = n.AppendMatches(out, c)
	}

	return
}

func (b *branch) getMatch(rect *Rect) (match node, index int) {
	var (
		distance Meter
		isSet    bool
	)

	if len(b.nodes) == 0 {
		var l leaf
		b.nodes = append(b.nodes, &l)
	}

	c := rect.Center()
	for i, n := range b.nodes {
		r := n.Rect()
		rc := r.Center()
		m := c.Distance(rc)
		if !isSet || m < distance {
			match = n
			index = i
			distance = m
			isSet = true
		}
	}

	return
}

func (b *branch) Rect() (r Rect) {
	return b.rect
}

func (b *branch) IsFullyContained(r *Rect) (contained bool) {
	return b.rect.IsFullyContained(r)
}

func (b *branch) DoesNotOverlap(r *Rect) (notOverlapping bool) {
	return b.rect.DoesNotOverlap(r)
}
