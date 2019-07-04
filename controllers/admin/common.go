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
	defer f.Close()
	imageType := c.GetString("type")
	ext := path.Ext(h.Filename)
	//uploadDir := "static/upload/" + imageType+ "/" + time.Now().Format("20060102")
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
	c.Data["code"] = 0
	c.Data["msg"] = "上传成功"
	data := make(map[string]interface{})
	data["src"] = fpath
	c.Data["data"] = data
	c.ServeJSON()
}
