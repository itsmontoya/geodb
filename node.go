package geodb

type node interface {
	Insert(value)
	Append(node)
	Rect() Rect
	Len() int
	Split() []node
	IsFullyContained(*Rect) bool
	DoesNotOverlap(*Rect) bool
	AppendMatches([]*entry, Coordinates) []*entry
}

type value interface {
	Rect() Rect
}
