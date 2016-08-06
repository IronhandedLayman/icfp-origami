package objs

import (
	"encoding/json"
	"fmt"
	"time"
)

type M map[string]interface{}

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
