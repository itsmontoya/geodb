package geodb

type Transaction interface {
	Insert(key string, shape Shape)
}
