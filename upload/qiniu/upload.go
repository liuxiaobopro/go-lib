package qiniu

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/liuxiaobopro/go-lib/console"
	"github.com/liuxiaobopro/go-lib/upload"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

/**
文档地址: https://developer.qiniu.com/kodo/1238/go#form-upload-stream
*/

type QiniuConfigType struct {
	AccessKey  string // 七牛云accessKey
	SecretKey  string // 七牛云secretKey
	Bucket     string // 七牛云存储空间
	Domain     string // 七牛云图片域名
	FilePath   string // 上传到服务器的文件路径
	IsDelLocal bool   // 上传之后是否删除本地文件
}

var QiniuConfig = new(QiniuConfigType)

func Upload(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	filePath, randFileName := upload.CreateFolder(file, fileHeader, QiniuConfig.FilePath)
	// 简单上传的凭证
	putPolicy := storage.PutPolicy{
		Scope: QiniuConfig.Bucket,
	}
	mac := qbox.NewMac(QiniuConfig.AccessKey, QiniuConfig.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	// 空间对应机房
	// 其中关于Zone对象和机房的关系如下：
	//    华东    storage.ZoneHuadong
	//    华北    storage.ZoneHuabei
	//    华南    storage.ZoneHuanan
	//    北美    storage.ZoneBeimei
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 七牛云存储空间设置首页有存储区域
		UseCdnDomains: false,               // 不启用HTTPS域名
		UseHTTPS:      false,               // 不使用CND加速
	}

	// 构建上传表单对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选
	putExtra := storage.PutExtra{
		// Params: map[string]string{
		// 	"x:name": "github logo",
		// },
	}

	// err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	err := formUploader.PutFile(context.Background(), &ret, upToken, randFileName, filePath, &putExtra)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s", QiniuConfig.Domain, ret.Key)

	// 删除本地文件
	if QiniuConfig.IsDelLocal {
		defer func(name string) {
			if err := os.Remove(name); err != nil {
				console.Console.Error(err.Error(), "")
			}
		}(filePath)
	}
	return url, nil
}
