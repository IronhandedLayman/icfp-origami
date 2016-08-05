package fsapi

type FoldServer interface {
	Hello() (string, error)
}
