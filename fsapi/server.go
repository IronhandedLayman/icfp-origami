package fsapi

import (
	"time"

	"github.com/IronhandedLayman/icfp-origami/objs"
)

type FoldServer interface {
	Hello() (string, error)
	GetBlob(string) (string, error)
	MakeServerRequest(string, []string, objs.M, bool) (string, error)
	SnapshotListRequest() (*objs.SnapshotListResponse, error)
	ProblemSubmission(string, time.Time) (string, error)
	SolutionSubmission(int, string) (string, error)
	LatestSnapshot() (*objs.Snapshot, error)
	Scoreboard() ([]objs.UserState, error)
	GetProblemSpec(int) (objs.Problem, error)
	GetRawProblemSpec(int) (string, error)
}
