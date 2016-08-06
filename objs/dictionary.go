package objs

import (
	"encoding/json"
	"fmt"
	"time"
)

type M map[string]interface{}

// Snapshot List structure

type SnapshotHash struct {
	SnapshotTime time.Time `json:"snapshot_time"`
	SnapshotHash string    `json:"snapshot_hash"`
}

func (sh *SnapshotHash) UnmarshalJSON(b []byte) error {
	prestr := struct {
		SnapshotTime int64  `json:"snapshot_time"`
		SnapshotHash string `json:"snapshot_hash"`
	}{}
	umerr := json.Unmarshal(b, &prestr)
	if umerr != nil {
		return umerr
	}
	sh.SnapshotHash = prestr.SnapshotHash
	sh.SnapshotTime = time.Unix(prestr.SnapshotTime, 0)
	return nil
}

func (sh *SnapshotHash) MarshalJSON() ([]byte, error) {
	return ([]byte)(fmt.Sprintf(
		`{"snapshot_time":%d,"snapshot_hash":"%s"}`,
		sh.SnapshotTime.Unix(),
		sh.SnapshotHash)), nil
}

func (sh *SnapshotHash) String() string {
	return fmt.Sprintf(`{"snapshot_time":%d,"snapshot_hash":"%s"}`, sh.SnapshotTime.Unix(), sh.SnapshotHash)
}

type SnapshotListResponse struct {
	Ok        bool            `json:"ok"`
	Snapshots []*SnapshotHash `json:"snapshots"`
}

//Snapshot structure

type UserScore struct {
	Username string  `json:"username"`
	Score    float64 `json:"score"`
}

type UserNameplate struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

type SolutionEval struct {
	Resemblance  float64 `json:"resemblance"`
	SolutionSize int     `json:"solution_size"`
}

type ProblemHeader struct {
	Owner           string         `json:"owner"`
	ProblemId       int            `json:"problem_id"`
	ProblemSize     int            `json:"problem_size"`
	ProblemSpecHash string         `json:"problem_spec_hash"`
	PublishTime     int64          `json:"publish_time"`
	Ranking         []SolutionEval `json:"ranking"`
	SolutionSize    int            `json:"solution_size"`
}

type Snapshot struct {
	Leaderboard  []UserScore
	Problems     []ProblemHeader
	SnapshotTime int64
	Users        []UserNameplate
}
