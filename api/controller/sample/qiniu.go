package sample

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"

	"pss/pkg/setting"
)

func QiniuUpload(c *gin.Context) {

	//read image from assets folder
	//file, err1 := os.Open("./assets/1.jpg")
	//if err1 != nil {
	//	panic(err1)
	//}

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	log.Println("accessKey:" + setting.QiniuAccessKey)
	err := formUploader.PutFile(context.Background(), &ret, setting.QiniuAccessKey, "sample.jpg", "./assets/1.jpg", &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash, ret.PersistentID)

	c.JSON(200, gin.H{
		"message": "dbinsert",
	})

}
