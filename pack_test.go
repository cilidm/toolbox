package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime"
	"path"
	"path/filepath"
	"toolbox/gomail"
	"toolbox/gozap"
	"toolbox/logging"
	"toolbox/session"
	"toolbox/session/cookie"
	"sync"
	"testing"
	"toolbox/store"
)

func TestStore(t *testing.T)  {
	var conf store.Config
	conf.CloudType = "qiniu"
	conf.AccessKey = ""
	conf.SecretKey = ""
	conf.PublicBucket = ""
	conf.PublicBucketDomain = "http://xxx.xxx.xxx/"
	cloud,err := store.NewCloudStore(conf,false)
	if err != nil{
		t.Fatal(err.Error())
	}
	filePath := "./store/README.md"
	storePath := filepath.Join("test",filePath)
	miMe := mime.TypeByExtension(path.Ext(filePath))
	if err := cloud.Upload(filePath,storePath,map[string]string{"Content-Type": miMe});err != nil{
		fmt.Println(err)
	}

	files,err := cloud.Lists("")		// 腾讯云暂时没有lists
	if err != nil{
		t.Fatal(err.Error())
	}
	for _,v := range files{
		fmt.Println(v.Name)
	}
}

func TestMail(t *testing.T)  {
	var mailConf gomail.MailConfForm
	mailConf.EmailHost = "xx"
	mailConf.EmailUser = "xx" // 其他全填写完整,此处省略
	var conf gomail.Config
	conf.Config = mailConf
	conf.MailTo = append(conf.MailTo, mailConf.EmailTest)
	conf.Subject = mailConf.EmailTestTitle
	gomail.SendMail(conf)
}

func TestZap(t *testing.T)  {
	gozap.InitLogger("./zap.log","debug")
	gozap.Info("aa","bb")
}

func TestLogging(t *testing.T)  {
	logging.Info("aa","bb")
}

func TestInitSession(t *testing.T)  {
	store := cookie.NewStore([]byte("SESSION_KEY"))
	session.Sessions("_SESSION",store)
}
// session使用
var SessionList sync.Map
func SetSession(c *gin.Context)  {
	ses := session.Default(c)
	ses.Set("key","val")
	ses.Save()
	sessionID := ses.SessionId()
	SessionList.Store(sessionID, c)
}
