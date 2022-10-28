package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"

	"gitee.com/liuxiaobopro/golib/console"
	utilStr "gitee.com/liuxiaobopro/golib/utils/string"
	utilTime "gitee.com/liuxiaobopro/golib/utils/time"
)

func CreateFolder(file multipart.File, fileHeader *multipart.FileHeader, filepath string) (string, string) {
	// 生成文件名
	fileExt := path.Ext(fileHeader.Filename)
	randFileName := utilStr.RandNumberAndLetter(16)
	randFileName = utilTime.GetNowTimeUnsigned() + randFileName + fileExt

	// 生成文件夹
	uploadPath := filepath
	// 创建文件夹
	pwd, _ := os.Getwd()
	folder := pwd + uploadPath
	if err := os.MkdirAll(folder, 0744); err != nil {
		console.Console.Error(err.Error(), "创建文件夹失败")
	}
	// 生成文件路径
	fPath := fmt.Sprintf("%s/%s", folder, randFileName)
	fW, err := os.Create(fPath)
	if err != nil {
		console.Console.Error(err.Error(), "创建文件失败")
	}
	defer fW.Close()
	// 复制文件，保存到本地
	_, err = io.Copy(fW, file)
	if err != nil {
		console.Console.Error(err.Error(), "复制文件失败")
	}

	return fPath, randFileName
}
