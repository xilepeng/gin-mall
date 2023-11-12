package util

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/xilepeng/gin-mall/conf"
)

// UploadAvatarToLocalStatic 上传图片到本地
func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId)) // 路径拼接
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

// UploadAvatarToQiNiu 封装上传图片到七牛云然后返回状态和图片的url，单张
func UploadAvatarToQiNiu(file multipart.File, fileSize int64) (filePath string, err error) {
	var AccessKey = conf.AccessKey
	var SerectKey = conf.SerectKey
	var Bucket = conf.Bucket

	var ImgUrl = conf.QiniuServer
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", err
	}
	url := ImgUrl + ret.Key
	return url, nil
}

func UploadProductToLocalStatic(file multipart.File, userId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId)) // 路径拼接
	basePath := "." + conf.ProductPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return "boss" + bId + "/" + productName + ".jpg", nil
}

// DirExistOrNot 判断路径存不存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0777) // 完全控制权限
	return err == nil
}
