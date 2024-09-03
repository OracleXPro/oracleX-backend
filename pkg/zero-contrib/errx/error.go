package errx

import (
	"fmt"
	"oracleX-backend/pkg/zero-contrib/errx/types"
	"runtime"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CodeError struct {
	Code uint32

	// internal error message
	Internal string

	// external error message
	External string
}

func (e CodeError) Error() {

}

func NewCodeError(code uint32, internal, external any, wraps ...error) CodeError {
	e := CodeError{
		Code: code,
	}

	switch internal := internal.(type) {
	case string:
		e.Internal = internal
	case error:
		e.Internal = internal.Error()
	default:
		e.Internal = fmt.Sprintf("unhandle internal: %+v", internal)
	}

	switch external := external.(type) {
	case string:
		e.External = external
	case error:
		e.External = external.Error()
	default:
		e.External = fmt.Sprintf("unhandle external: %+v", external)
	}

	return e
}

func _new(code uint32, message string) error {
	pc, filename, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	detail := fmt.Sprintf("Err<code=%d, err=%v>, Locaction<file=%s, func=%s, line=%d>", code, message, filename, funcName, line)

	st := status.New(codes.Code(code), message)
	ds, err := st.WithDetails(&types.ErrorMessage{Message: detail})
	if err != nil {
		panic(err)
	}
	return ds.Err()
}

// func New(code uint32, message string) error {
// 	return _new(code, message)
// }

// func Code(errCode uint32) error {
// 	return _new(errCode, MapErrMsg(errCode))
// }

// func Message(errMsg string) error {
// 	return _new(Internal, errMsg)
// }

func Error(code uint32, internal, external any, wraps ...error) error {
	e := NewCodeError(code, internal, external, wraps...)

	pc, filename, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	detail := fmt.Sprintf("\nError:\n\tcode=%d\n\tinner err=%v\n\touter err=%v\n Locaction:\n\tfile=%s\n\tfunc=%s\n\tline=%d\n", code, e.Internal, e.External, filename, funcName, line)

	st := status.New(codes.Code(code), e.External)
	ds, err := st.WithDetails(&types.ErrorMessage{Message: detail})
	if err != nil {
		panic(err)
	}
	return ds.Err()
}

func IsCodeErr(code uint32) bool {
	if _, ok := CodeToMsg[code]; ok {
		return true
	} else {
		return false
	}
}
