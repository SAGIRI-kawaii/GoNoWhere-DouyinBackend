package upload

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

//上传文件，本项目主要包含视频和图片
// 参考项目和文档
// https://gitee.com/pixelmax/gin-vue-admin/blob/main/server/utils/upload/aliyun_oss.go
// https://help.aliyun.com/product/31815.html?spm=5176.7933691.J_5253785160.6.272f4c59KogXWZ
// https://developer.qiniu.com/kodo

// ToQiNiu 上传文件到七牛云对象存储
func UploadVideo(file *[]byte, videoID int64) (string, error) {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "mini-douyin"
	accessKey := "rh21CFIGeHdD0OAW0Cr-618hZ1SEdXhCR5RicAxQ"
	secretKey := "elpJGvJh_PJQgdgZxxSkytBGxeIO_POO8VTuW_an"
	QiNiuServer := "rqbxff5oo.hn-bkt.clouddn.com"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	fileSize := len(*file)
	key := "video/" + strconv.FormatInt(videoID, 10) + ".mp4"
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(*file), int64(fileSize), &putExtra)
	if err != nil {
		return "", err
	}
	url := "http://" + QiNiuServer + "/" + ret.Key
	return url, nil
}
func UploadCover(file *[]byte, videoID int64) (string, error) {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "mini-douyin"
	accessKey := "rh21CFIGeHdD0OAW0Cr-618hZ1SEdXhCR5RicAxQ"
	secretKey := "elpJGvJh_PJQgdgZxxSkytBGxeIO_POO8VTuW_an"
	QiNiuServer := "rqbxff5oo.hn-bkt.clouddn.com"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	key := "cover/" + strconv.FormatInt(videoID, 10) + ".jpg"
	fileSize := len(*file)
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(*file), int64(fileSize), &putExtra)
	if err != nil {
		return "", err
	}
	url := "http://" + QiNiuServer + "/" + ret.Key
	return url, nil
}
func DeleteVideo(videoID int64) (string, error) {
	bucket := "mini-douyin"
	accessKey := "rh21CFIGeHdD0OAW0Cr-618hZ1SEdXhCR5RicAxQ"
	secretKey := "elpJGvJh_PJQgdgZxxSkytBGxeIO_POO8VTuW_an"
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	key := "video/" + strconv.FormatInt(videoID, 10) + ".mp4"
	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return "Delete Error", err
	}
	return "Success", nil
}
