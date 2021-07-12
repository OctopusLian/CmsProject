package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"CmsProject/util"
	"github.com/astaxie/beego/orm"
	"CmsProject/models"
	"strconv"
	"os"
	"encoding/json"
	"time"
)

type FileController struct {
	beego.Controller
}

/**
 * 更新用户头像
 */
func (this *FileController) UpdateAdminAvatar() {
	util.LogInfo("更新用户头像")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	//判断是否登录的权限
	if !this.IsLogin() {
		resp["status"] = util.RECODE_UNLOGIN
		resp["type"] = util.EEROR_UNLOGIN
		resp["message"] = util.Recode2Text(util.EEROR_UNLOGIN)
		return
	}

	//获取文件操作 thisl.GetFike(name)
	file, head, err := this.GetFile("file")
	defer file.Close()

	//beego.Info(head.Filename)
	//beego.Info(head.Size)

	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTUREADD
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
		return
	}

	//filename : default.jpg
	// fileArr : [ default jpg]
	fileArr := strings.Split(head.Filename, ".")
	//文件类型判断
	if (fileArr[1] != "png" && fileArr[1] != "jpg" && fileArr[1] != "jpeg") {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTURETYPE
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTURETYPE)
		return
	}

	//文件大小判断 控制文件在2M以内
	if head.Size > 1024*1024*2 {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTURESIZE
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTURESIZE)
		return
	}

	uploadPath := "./img/"
	//判断upload是否存在，不存在先创建
	if exist, _ := util.IsExists(uploadPath); !exist {
		//err := os.Mkdir(uploadPath, os.ModePerm)
		//if err != nil {
		//	//失败
		//	beego.Info(err.Error())
		//	resp["status"] = util.RECODE_FAIL
		//	resp["type"] = util.RESPMSG_ERROR_PICTUREADD
		//	resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
		//	return
		//}

		//go语言中更常见的形式
		if err := os.Mkdir(uploadPath, os.ModePerm); err != nil {
			//失败
			beego.Info(err.Error())
			resp["status"] = util.RECODE_FAIL
			resp["type"] = util.RESPMSG_ERROR_PICTUREADD
			resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
			return
		}
	}

	//目录创建成功，继续保存文件
	//自己为上传的文件重新命名
	fileArray := strings.Split(head.Filename, ".")
	fileName := "avatar" + strconv.Itoa(int(time.Now().UnixNano())) + "." + fileArray[1]
	path := uploadPath + fileName

	//真正执行保存文件的操作
	//文件切记关闭文件

	if err = this.SaveToFile("file", path); err != nil {
		//失败
		beego.Info(err.Error())
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTUREADD
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
		return
	}

	//文件保存到目录成功，更新数据库
	om := orm.NewOrm()
	adminId, _ := strconv.Atoi(this.Ctx.Input.Param(":adminId"))
	admin := models.Admin{Id: adminId}
	//select * from admin where id = ? value adminId
	if om.Read(&admin) == nil {
		admin.Avatar = fileName
		//update admin set avatar = ? where id = ? value (fileName, adminId)
		if _, err := om.Update(&admin, "avatar"); err == nil {
			//返回正常图片链接
			resp["status"] = util.RECODE_OK
			resp["image_path"] = fileName
			return
		}
	}

	//失败
	resp["status"] = util.RECODE_FAIL
	resp["type"] = util.RESPMSG_ERROR_PICTUREADD
	resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
}

/**
 * 上传图片
 */
func (this *FileController) UploadImg() {

	util.LogInfo("上传图片")

	resp := make(map[string]interface{})
	this.Data["json"] = resp
	defer this.ServeJSON()

	//使用getFile来获取上传的文件
	//file 为文件类型结构体，head可以得到文件的大小和文件的名字
	file, head, err := this.GetFile("file")
	beego.Info("文件:", head.Filename, head.Size)
	if err != nil {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTUREADD
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
		return
	}

	fileArr := strings.Split(head.Filename, ".")
	//文件类型判断
	if (fileArr[1] != "png" && fileArr[1] != "jpg" && fileArr[1] != "jpeg") {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTURETYPE
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTURETYPE)
		return
	}

	//文件大小判断 控制文件在2M以内
	if head.Size > 1024*1024*2 {
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTURESIZE
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTURESIZE)
		return
	}

	//文件切记关闭文件
	path := "./img/" + head.Filename
	file.Close()
	err = this.SaveToFile("file", path)
	if err != nil {
		//失败
		resp["status"] = util.RECODE_FAIL
		resp["type"] = util.RESPMSG_ERROR_PICTUREADD
		resp["failure"] = util.Recode2Text(util.RESPMSG_ERROR_PICTUREADD)
		return
	}

	//返回正常图片链接
	resp["status"] = util.RECODE_OK
	resp["image_path"] = head.Filename
}

/**
 * 判断用户是否已经登陆过：通过session进行判断
 */
func (this *FileController) IsLogin() bool {
	adminByte := this.GetSession(ADMIN)
	if adminByte != nil {
		var admin models.Admin
		json.Unmarshal(adminByte.([]byte), &admin)
		return admin.Id > 0
	}
	return false
}
