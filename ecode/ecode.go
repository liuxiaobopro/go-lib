package ecode

type Res struct {
	Code int
	Desc string
}

type BizErr *Res

var (
	SUCCSESS                          = &Res{0, "操作成功"}
	ERROR                             = &Res{99999, "操作失败"}
	ERROR_SERVER                      = &Res{99998, "服务器错误"}
	ERROR_PARAMETER_EXCEPTION         = &Res{99997, "参数异常(缺少必填项或参数格式错误)"}
	ERROR_RESOURCE_DONT_EXISTS        = &Res{99996, "资源不存在"}
	ERROR_RESOURCE_ALREADY_EXISTS     = &Res{99995, "资源已存在"}
	ERROR_TOKEN_INVALID               = &Res{99994, "token无效"}
	ERROR_ID_GREATER_THAN_ZERO        = &Res{99995, "ID值必须大于0"}
	ERROR_NOT_ALLOW_REVERSE_OPERATION = &Res{99994, "不允许逆操作"}
)
