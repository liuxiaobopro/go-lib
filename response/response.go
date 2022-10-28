package response

import (
	"github.com/liuxiaobopro/go-lib/ecode"
)

type Result struct {
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	Data    interface{} `json:"data"`
}

// GetSuccRes 成功返回
func GetSuccRes(data interface{}) *Result {
	return GetCustomRes(ecode.SUCCSESS, data)
}

// GetErrRes 失败返回
func GetErrRes(err ecode.BizErr) *Result {
	return GetCustomRes(err, nil)
}

// GetCustomRes 自定义返回
func GetCustomRes(err ecode.BizErr, data interface{}) *Result {
	return &Result{
		ErrCode: err.Code,
		ErrMsg:  err.Desc,
		Data:    data,
	}
}
