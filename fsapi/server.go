package fsapi

import (
	"time"

	"github.com/IronhandedLayman/icfp-origami/objs"
)

type FoldServer interface {
	Hello() (string, error)
	GetBlob(string) (string, error)
	MakeServerRequest(string, []string, objs.M) (string, error)
	SnapshotListRequest() (*objs.SnapshotListResponse, error)
	ProblemSubmission(string, time.Time) (string, error)
	SolutionSubmission(int, string) (string, error)
}
