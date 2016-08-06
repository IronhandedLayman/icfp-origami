package objs

import (
	"sort"
	"strconv"
)

type UserState struct {
	UserId   int     `json:"user_id"`
	Username string  `json:"username"`
	Score    float64 `json:"score"`
}

func (us *UserState) FromUserData(np UserNameplate, ns UserScore) {
	us.Username = np.DisplayName
	us.UserId, _ = strconv.Atoi(np.Username)
	us.Score = ns.Score
}

func MergeUserData(nps []UserNameplate, nss []UserScore) []UserState {
	preNP := make(map[string]UserNameplate)
	preNS := make(map[string]UserScore)
	for _, np := range nps {
		preNP[np.Username] = np
	}
	for _, ns := range nss {
		preNS[ns.Username] = ns
	}
	ans := make([]UserState, len(nps))
	i := 0
	for _, np := range nps {
		us := UserState{}
		(&us).FromUserData(np, preNS[np.Username])
		ans[i] = us
	}
	return ans
}

type ByUsers func(u1, u2 *UserState) bool

type userStateSorter struct {
	userstates []UserState
	by         ByUsers
}

func (byu ByUsers) Sort(us []UserState) {
	sort.Sort(&userStateSorter{
		userstates: us,
		by:         byu,
	})
}

// Len is part of sort.Interface.
func (s *userStateSorter) Len() int {
	return len(s.userstates)
}

// Swap is part of sort.Interface.
func (s *userStateSorter) Swap(i, j int) {
	s.userstates[i], s.userstates[j] = s.userstates[j], s.userstates[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *userStateSorter) Less(i, j int) bool {
	return s.by(&s.userstates[i], &s.userstates[j])
}
