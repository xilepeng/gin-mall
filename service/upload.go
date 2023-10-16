package service

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/xilepeng/gin-mall/conf"
)

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
