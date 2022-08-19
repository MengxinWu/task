package ecode

import "fmt"

type TaskError struct {
	Code int
	Msg  string
}

func (t TaskError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", t.Code, t.Msg)
}

func New(code int, msg string) TaskError {
	return TaskError{
		Code: code,
		Msg:  msg,
	}
}

var (
	// common 1000~1999
	EngineError = New(1001, "数据库错误")
	RPCError    = New(1002, "RPC错误")

	// resource 2000 ~ 2999
	ResourceNotFound = New(2001, "资源不存在")

	// dag & processor 3000 ~ 3999
	ProcessorNotFound = New(3001, "处理单元不存在")
	ProcessorNotInDag = New(3002, "处理单元不在DAG")

	// process 4000 ~ 4999
	ExecutorHandlerNotFound = New(4001, "执行处理器不存在")
	ProcessStateNotFound    = New(4002, "处理状态不存在")
	ProcessStateWrong       = New(4003, "处理状态错误")
	ProcessRetryCntOver     = New(4004, "处理次数超过最大值")
)
