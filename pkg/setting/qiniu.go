package setting

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
)

func SetupQiniu() {
	accessKey := "FNl_CsHmKkJr0YCITpDbVRMEh4TB96FyxwLf_azd"
	secretKey := "8e3FGDCRpbrIj2rG1e7SyMXTtpx__J2T30xcuD4W"
	mac := qbox.NewMac(accessKey, secretKey)

	bucket := "pso"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	log.Println("upToken:" + upToken)
	QiniuAccessKey = upToken
}
