package geodb

import "fmt"

var _ node = &leaf{}

type leaf struct {
	rect    Rect
	entries []*entry
}

func (l *leaf) Append(n node) {
	l.Insert(n)
}

func (l *leaf) Insert(v value) {
	e, ok := v.(*entry)
	if !ok {
		msg := fmt.Sprintf("invalid type, expected %T and received %T", e, v)
		panic(msg)
	}

	l.entries = append(l.entries, e)
	l.rect.SetRect(e.shape.Rect())
}

func (l *leaf) Split() (out []node) {
	var l1, l2 leaf
	l1.rect, l2.rect = l.rect.Split()
	for _, e := range l.entries {
		sr := e.shape.Rect()
		node, ok := sr.GetMatchingNode(&l1, &l2)
		if !ok {
			msg := fmt.Sprintf("leaf.Split(): no split match found for %+v within %v / %v, orig: %v", sr, l1.rect, l2.rect, l.rect)
			panic(msg)
		}

		node.Insert(e)
	}

	out = append(out, &l1)
	out = append(out, &l2)
	return
}

func (l *leaf) Rect() (r Rect) {
	r = l.rect
	return
}

func (l *leaf) Len() int {
	return len(l.entries)
}

func (l *leaf) ContainsCoordinates(c Coordinates) (contain bool) {
	return l.rect.ContainsCoordinates(c)
}

func (l *leaf) IsFullyContained(r *Rect) (contained bool) {
	return l.rect.IsFullyContained(r)
}

func (l *leaf) AppendMatches(in []*entry, c Coordinates) (out []*entry) {
	out = in
	if !l.ContainsCoordinates(c) {
		return
	}

	for _, e := range l.entries {
		if !e.shape.IsWithin(c) {
			continue
		}

		out = append(out, e)
	}

	return
}
