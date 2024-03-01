package status

import (
	"fmt"
	"github.com/gogoclouds/gogo-services/common-lib/web/r/codes"
)

// Status https://cloud.google.com/apis/design/errors
// It is immutable and should be created with New, Newf, or FromProto.
type Status struct {

	// The status code, which should be an enum value of [codes.Code][codes.Code].
	Code codes.Code `json:"code"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message string `json:"message"`
	// A list of messages that carry the error details.  There is a common set of
	// message types for APIs to use.
	Details []any `json:"details"`
}

// New returns a Status representing c and msg.
func New(c uint32, msg string) *Status {
	return &Status{Code: c, Message: msg}
}

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c uint32, format string, a ...any) *Status {
	return New(c, fmt.Sprintf(format, a...))
}

// Err returns an error representing c and msg.  If c is OK, returns nil.
func Err(c uint32, msg string) error {
	return New(c, msg).Err()
}

// Errf returns Err(c, fmt.Sprintf(format, a...)).
func Errf(c uint32, format string, a ...interface{}) error {
	return Err(c, fmt.Sprintf(format, a...))
}

// GetCode returns the status code contained in s.
func (s *Status) GetCode() codes.Code {
	if s == nil {
		return codes.OK
	}
	return s.Code
}

// GetMessage returns the message contained in s.
func (s *Status) GetMessage() string {
	if s == nil {
		return ""
	}
	return s.Message
}

// Err returns an immutable error representing s; returns nil if s.Code() is OK.
func (s *Status) Err() error {
	if s.GetCode() == codes.OK {
		return nil
	}
	return s
}

func (s *Status) WithDetails(details ...any) *Status {
	if s.GetCode() == codes.OK {
		return s
	}
	cp := *s
	cp.Details = append(cp.Details, details...)
	return &cp
}

func (s *Status) String() string {
	return fmt.Sprintf("rpc error: code = %d desc = %s", s.GetCode(), s.GetMessage())
}

func (s *Status) Error() string {
	return s.String()
}
