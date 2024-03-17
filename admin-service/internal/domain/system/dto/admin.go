package dto

type SampleAdmin struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UniqueUsernameQuery struct {
	ExcludeID int64  `json:"id"`
	Username  string `json:"username"`
}

type UniqueEmailQuery struct {
	ExcludeID int64  `json:"id"`
	Email     string `json:"email"`
}

type UniqueResult struct {
	Exist bool  `json:"exist"`
	IsDel uint8 `json:"isDel"`
}
