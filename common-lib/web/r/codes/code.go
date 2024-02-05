package codes

// Code 响应状态
type Code = uint32

const (
	OK                        Code = 0
	BadRequest                Code = 400
	StatusInternalServerError      = 500
	LoginErr                  Code = 1000
	TokenInvalid              Code = 4010
	TokenMission              Code = 4011
	// ....
)
