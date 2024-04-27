package codes

// Code 响应状态
type Code = uint32

const (
	OK                  Code = 200
	BadRequest          Code = 400
	InternalServerError Code = 500
	TokenInvalid        Code = 4010
	TokenMission        Code = 4011
	// ....
)
