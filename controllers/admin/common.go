package admin

import (
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
)

type CommonController struct {
	baseController
}

func (c *CommonController) Upload() {
	f, h, err := c.GetFile("file")
	if err != nil {
		log.Fatal("getfile err", err)
	}
	result := make(map[string]interface{})
	if err == nil {
		imageType := c.GetString("type")
		ext := path.Ext(h.Filename)
		if ext != ".jpg" && ext != ".png" && ext != ".gif" {
			result["code"] = 1
			result["message"] = "只支持jpg,png,gif格式的图片"
		} else {
			uploadDir := path.Join("static/upload/", imageType, time.Now().Format("20060102"))
			err = os.MkdirAll(uploadDir, 777)
			if err != nil {
				log.Fatal("create directory err :", err)
			}
			rand.Seed(time.Now().UnixNano())
			randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
			hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
			fileName := fmt.Sprintf("%x", hashName) + ext
			fpath := path.Join(uploadDir, fileName)
			err = c.SaveToFile("file", fpath)
			if err != nil {
				log.Fatal("save file err", err)
			}
			result["code"] = 0
			result["msg"] = "上传成功"
			data := make(map[string]interface{})
			data["src"] = "/" + fpath
			result["data"] = data
		}

	} else {
		result["code"] = 1
		result["message"] = "上传异常"
	}
	defer f.Close()
	c.Data["json"] = result
	c.ServeJSON()
}
