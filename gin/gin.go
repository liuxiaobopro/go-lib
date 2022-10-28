package gin

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"gitee.com/liuxiaobopro/golib/ecode"
	"gitee.com/liuxiaobopro/golib/response"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

// Send GetDefault 默认
func (th *Handler) Send(c *gin.Context, data interface{}, r ecode.BizErr) {
	c.JSON(http.StatusOK, &response.Result{
		ErrCode: r.Code,
		ErrMsg:  r.Desc,
		Data:    data,
	})
}

// SendSucc GetSuccRes 成功返回
func (th *Handler) SendSucc(c *gin.Context, data interface{}, err ecode.BizErr) {
	if err == nil {
		err = ecode.SUCCSESS
	}
	th.Send(c, data, err)
}

// SendErr GetErrRes 失败返回
func (th *Handler) SendErr(c *gin.Context, data interface{}, err ecode.BizErr) {
	th.Send(c, nil, err)
}

func (th *Handler) SendError(c *gin.Context, err error, bizerr ecode.BizErr) {
	desc := fmt.Sprintf("【%s】%s", bizerr.Desc, err.Error())
	th.SendErr(c, nil, &ecode.Res{Code: ecode.ERROR_PARAMETER_EXCEPTION.Code, Desc: desc})
}

// GetQueryDefault 获取参数(带默认)
func (th *Handler) GetQueryDefault(c *gin.Context, key string, defaultVal string) string {
	value := c.DefaultQuery(key, "")
	if value == "" {
		return defaultVal
	} else {
		return value
	}
}

// GetParam 获取参数
func (th *Handler) GetParam(c *gin.Context, key string) string {
	value := c.Param(key)
	return value
}

// GetFormFile 获取file
func (th *Handler) GetFormFile(c *gin.Context, key string) (multipart.File, *multipart.FileHeader, error) {
	file, fileHeader, err := c.Request.FormFile(key)
	return file, fileHeader, err
}
