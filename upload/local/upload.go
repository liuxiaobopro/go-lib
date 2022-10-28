package local

import (
	"mime/multipart"

	"gitee.com/liuxiaobopro/golib/upload"
)

func Upload(file multipart.File, fileHeader *multipart.FileHeader, filepath string) (string, string) {
	return upload.CreateFolder(file, fileHeader, filepath)
}
