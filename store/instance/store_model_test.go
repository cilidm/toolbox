package instance

import (
	"fmt"
	"mime"
	"path"
	"path/filepath"
	"testing"
)

func TestCloudStore_Lists(t *testing.T) {
	var conf Config
	conf.CloudType = "qiniu"
	conf.AccessKey = "your accessKey"
	conf.SecretKey = "your secretKey"
	conf.PublicBucket = "your publicBucket"
	conf.PublicBucketDomain = "http://your publicBucketDomain/"
	cloud, err := NewCloudStore(conf, false)
	if err != nil {
		t.Fatal(err.Error())
	}
	filePath := "./store/README.md"
	storePath := filepath.Join("test", filePath)
	miMe := mime.TypeByExtension(path.Ext(filePath))
	if err := cloud.Upload(filePath, storePath, map[string]string{"Content-Type": miMe}); err != nil {
		fmt.Println(err)
	}

	files, err := cloud.Lists("") // 腾讯云暂时没有lists
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range files {
		fmt.Println(v.Name)
	}
}
