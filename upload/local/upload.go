package local

import (
	"mime/multipart"

	"github.com/liuxiaobopro/go-lib/upload"
)

func Upload(file multipart.File, fileHeader *multipart.FileHeader, filepath string) (string, string) {
	return upload.CreateFolder(file, fileHeader, filepath)
}
