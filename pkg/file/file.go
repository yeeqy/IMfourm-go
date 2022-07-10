package file

import (
	"IMfourm-go/pkg/app"
	"IMfourm-go/pkg/auth"
	"IMfourm-go/pkg/helpers"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

//操作辅助函数

//Put 将数据存入文件
func Put(data []byte,to string) error{
	err := ioutil.WriteFile(to,data,0644)
	if err != nil {
		return err
	}
	return nil
}
//Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _,err := os.Stat(fileToCheck); os.IsNotExist(err){
		return false
	}
	return true
}

func FileNameWithoutExtension(fileName string) string  {
	return strings.TrimSuffix(fileName,filepath.Ext(fileName))
}

func SaveUploadAvatar(c *gin.Context,file *multipart.FileHeader)(string,error)  {
	var avatar string
	//确保目录存在，不存在则创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/avatars/%s/%s/",
		app.TimeNowInTimezone().Format("2006/01/02"),auth.CurrentUID(c))
	os.MkdirAll(publicPath+dirName,0755)
	//保存文件
	fileName := randomNameFromUploadFile(file)
	// public/uploads/avatars/2021/12/22/1/nFDacgaWKpWWOmOt.png
	avatarPath := publicPath + dirName + fileName
	//gin提供的c.SaveUploadedFile方法来保存文件
	if err:= c.SaveUploadedFile(file,avatarPath);err!=nil{
		return avatar,err
	}

	//图片裁剪也封装在此
	img,err := imaging.Open(avatarPath,imaging.AutoOrientation(true))
	if err != nil {
		return avatar,err
	}
	resizeAvatar := imaging.Thumbnail(img,256,256,imaging.Lanczos)
	resizeAvatarName := randomNameFromUploadFile(file)
	resizeAvatarPath := publicPath + dirName + resizeAvatarName

	err = imaging.Save(resizeAvatar,resizeAvatarPath)
	if err != nil {
		return avatar,err
	}
	//删除老文件
	err = os.Remove(avatarPath)
	if err !=nil{
		return avatar,err
	}
	return dirName + resizeAvatarName ,nil
}
func randomNameFromUploadFile(file *multipart.FileHeader) string  {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)

}