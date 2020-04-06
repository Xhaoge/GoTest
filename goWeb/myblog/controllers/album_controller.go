package controllers

import (
	"goWeb/myblog/models"
	"log"
	"fmt"
	"os"
	"time"
	"github.com/astaxie/beego"
)

type AlbumController struct {
	BaseController
}

type UploadController struct {
	BaseController
}

func (this *AlbumController) Get(){
	albums,err := models.FindAllAlbums()
	if err != nil {
		log.Error(err)
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
}

func (this *UploadController) Post(){
	fmt.Println("file upload...")
	fileData,fileHeader,err := this.GetFile("upload")
	if err != nil {
		this.responseErr(err)
		return 
	}
	fmt.Println("name:",fileHeader.Filename,fileHeader.Size)
	fmt.Println(fileData)
	now := time.Now()
	fmt.Println("exti: ",filepath.Ext(fileHeader.Filename))
	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存到数据库中；
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	// 文件路径
	timeStamp := time.Now().Unix()
	fileName :=fmt.Sprintf("%d-%s".timeStamp,fileHeader.Filename)
	filePathStr := filepath.Join(fileDir,fileName)
	desFile,err := os.Create(filePathStr)
	if err != nil {
		this.responseErr(err)
		return
	}
	//将浏览器客户端上传的文件拷贝到本地路径的文件里面；
	_,err = io.Copy(desFile,fileData)
	if err != nil {
		this.responseErr(err)
		return
	}
	if fileType == "img" {
		album :=  models.Album{0,filePathStr,fileName,0,timeStamp}
		models.InsertAlbum(album)
	}
	this.Data["json"] = map[string]interface{}{"code":1,"message":"上传成功"}
	this.ServeJSON()
}


func (this *UploadController) responseErr(err error){
	this.Data["json"] = map[string]interface{}{"code":0,"message":err}
	this.ServeJSON()
}


