package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"gitee.com/liuxiaobopro/golib/ecode"
	"gitee.com/liuxiaobopro/golib/response"

	"github.com/gin-gonic/gin"
)

type logJson struct {
	Level       string `json:"level"`       // 日志级别
	RequestTime string `json:"requestTime"` // 请求时间
	Err         string `json:"err"`         // 错误信息
	SpendTime   string `json:"spendTime"`   // 耗时
	HostName    string `json:"hostName"`    // 主机名
	StatusCode  int    `json:"statusCode"`  // 状态码
	ClientIp    string `json:"clientIp"`    // 客户端ip
	UserAgent   string `json:"userAgent"`   // 客户端浏览器信息
	DataSize    int    `json:"dataSize"`    // 返回数据大小
	Method      string `json:"method"`      // 请求方法
	Path        string `json:"path"`        // 请求路径
	Req         string `json:"req"`         // 请求参数
	Rep         string `json:"rep"`         // 返回参数
}

type Config struct {
	Path        string // 日志路径
	IsAbort     bool   // 是否终止后续接口调用(false 终止，true 不终止)
	IsRetuenErr bool   // 是否返回错误信息
}

var Configs = new(Config)

func init() {
	if Configs.Path == "" {
		Configs.Path = "runtime/log"
	}
}

func Log(c *gin.Context) {
	startTime := time.Now()

	blw := &ResponseWriterWrapper{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	var parms []byte
	if c.Request.Body != nil {
		// 获取请求体
		parms, _ = ioutil.ReadAll(c.Request.Body)
		// 将请求体写回请求
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(parms))
	}

	defer func() {
		if r := recover(); r != nil {
			// 打印错误堆栈信息
			// log.Printf("logger panic: %v\n", r)
			// fmt.Printf("Stack: %s", string(debug.Stack()))

			errMsg := errorToString(r)
			_, file, line, _ := runtime.Caller(3)
			errMsg1 := fmt.Sprintf("%s【%s:%d】", errMsg, file, line)

			writeErrLog(c, errMsg1, startTime, blw, parms)
			if Configs.IsRetuenErr {
				c.JSON(http.StatusInternalServerError, response.GetCustomRes(ecode.ERROR_SERVER, errMsg))
			}
			// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			if !Configs.IsAbort {
				c.Abort()
			}
		} else {
			writeErrLog(c, "", startTime, blw, parms)
		}
	}()
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func writeErrLog(c *gin.Context, err string, startTime time.Time, blw *ResponseWriterWrapper, parms []byte) {
	now := time.Now()
	e := os.MkdirAll(Configs.Path, 0744)
	if e != nil {
		panic("创建日志目录失败")
	}
	filePath := fmt.Sprintf("%s/%v.log", Configs.Path, now.Format("20060102"))
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	defer file.Close()

	var log logJson
	if err != "" {
		log.Level = "error"
	} else {
		log.Level = "info"
	}
	log.RequestTime = now.Format("2006-01-02 15:04:05")
	log.Err = err

	stopTime := time.Since(startTime)
	log.SpendTime = fmt.Sprintf("%v秒", float64(stopTime.Nanoseconds())/1000000000.0)

	log.HostName, _ = os.Hostname()
	log.StatusCode = c.Writer.Status()
	log.ClientIp = c.ClientIP()
	log.UserAgent = url.QueryEscape(c.Request.UserAgent())
	log.DataSize = c.Writer.Size()
	log.Method = c.Request.Method

	// 获取请求地址参数
	var reqstr string
	q := c.Request.URL.Query()
	if q != nil {
		for k, v := range q {
			reqstr += fmt.Sprintf("%s=%s&", k, v[0])
		}
		if len(reqstr) > 0 {
			reqstr = reqstr[:len(reqstr)-1]
		}
	}
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	log.Path = fmt.Sprintf("%s://%s%s?%s", scheme, c.Request.Host, c.Request.URL.Path, reqstr)
	// 问号结尾的url，去掉问号
	if strings.HasSuffix(log.Path, "?") {
		log.Path = strings.Replace(log.Path, "?", "", -1)
	}

	// 获取请求参数
	log.Req = strings.ReplaceAll(string(parms), "\r\n", "")

	// 获取返回信息
	log.Rep = blw.Body.String()

	all, _ := json.Marshal(log)
	write := bufio.NewWriter(file)
	_, _ = write.WriteString(string(all) + "\r")
	// Flush将缓存的文件真正写入到文件中
	_ = write.Flush()
}

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body *bytes.Buffer // 缓存
}

func (w ResponseWriterWrapper) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriterWrapper) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
