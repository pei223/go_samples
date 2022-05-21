package error_sample

import (
	"fmt"
)

type ErrorCode = uint16

type ErrorDetail struct {
	Code    ErrorCode
	Message string
}

type ErrorWithErrorCode struct {
	Code    ErrorCode
	Message string
}

type ErrorWithDetail struct {
	Message string
	Details []ErrorDetail
}

func (err *ErrorWithErrorCode) Error() string {
	return err.Message
}

func (err *ErrorWithErrorCode) ErrorCode() ErrorCode {
	return err.Code
}

func (err *ErrorWithDetail) Error() string {
	return err.Message
}

func (err *ErrorWithDetail) ErrorDetails() []ErrorDetail {
	return err.Details
}

func DoErrorSamples() {
	fmt.Println("\n\nエラー型識別テスト")
	err1 := GenError(1)
	IdentifyAndOutputError(err1)
	fmt.Println()
	err2 := GenError(2)
	IdentifyAndOutputError(err2)
}

func GenError(num uint16) error {
	if num%2 == 0 {
		return &ErrorWithErrorCode{
			Code:    num,
			Message: fmt.Sprintf("Error%d", num),
		}
	}
	return &ErrorWithDetail{
		Message: fmt.Sprintf("Error%d", num),
		Details: []ErrorDetail{
			{
				Message: "Detail1",
				Code:    1,
			},
			{
				Message: "Detail2",
				Code:    2,
			},
		},
	}
}

func IdentifyAndOutputError(err error) {
	switch err := err.(type) {
	case *ErrorWithErrorCode:
		fmt.Printf("[ErrorWithErrorCode]\nCode: %d, Message: %s\n", err.Code, err.Message)
	case *ErrorWithDetail:
		fmt.Printf("[ErrorWithDetail]\nMessage: %s, \n", err.Message)
		for i, detail := range err.Details {
			fmt.Printf("\t%d: Code: %d, Message: %s\n", i, detail.Code, detail.Message)
		}
	}
}
