package fsapi

type FoldServer interface {
	Hello() (string, error)
	MakeServerRequest(string, map[string]interface{}) (string, error)
	SnapshotListRequest() (string, error)
}
