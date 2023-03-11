package handlers

type Logger interface {
	Printf(format string, v ...any)
}

type MakeNewFriendInfo struct {
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

type DeleteUserInfo struct {
	ID int `json:"target_id"`
}

type UpdateAgeInfo struct {
	Age int `json:"new_age"`
}
