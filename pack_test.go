package main

import (
	"fmt"
	"github.com/cilidm/toolbox/store/instance"
	"mime"
	"path"
	"path/filepath"
	"sync"
	"testing"

	OS "github.com/cilidm/toolbox/Os"
	"github.com/cilidm/toolbox/gomail"
	"github.com/cilidm/toolbox/gozap"
	"github.com/cilidm/toolbox/logging"
	"github.com/cilidm/toolbox/net"
	"github.com/cilidm/toolbox/rand"
	"github.com/cilidm/toolbox/session"
	"github.com/cilidm/toolbox/session/cookie"
	"github.com/gin-gonic/gin"
)

func TestStore(t *testing.T) {
	var conf instance.Config
	conf.CloudType = "qiniu"
	conf.AccessKey = "your accessKey"
	conf.SecretKey = "your secretKey"
	conf.PublicBucket = "your publicBucket"
	conf.PublicBucketDomain = "http://your publicBucketDomain/"
	cloud, err := instance.NewCloudStore(conf, false)
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

func TestMail(t *testing.T) {
	var mailConf gomail.MailConfForm
	mailConf.EmailHost = "xx"
	mailConf.EmailUser = "xx" // 其他全填写完整,此处省略
	var conf gomail.Config
	conf.Config = mailConf
	conf.MailTo = append(conf.MailTo, mailConf.EmailTest)
	conf.Subject = mailConf.EmailTestTitle
	gomail.SendMail(conf)
}

func TestZap(t *testing.T) {
	gozap.InitLogger("./zap.log", "debug")
	gozap.Info("aa", "bb")
}

func TestLogging(t *testing.T) {
	logging.Info("aa", "bb")
}

func TestInitSession(t *testing.T) {
	store := cookie.NewStore([]byte("SESSION_KEY"))
	session.Sessions("_SESSION", store)
}

// session使用
var SessionList sync.Map

func SetSession(c *gin.Context) {
	ses := session.Default(c)
	ses.Set("key", "val")
	ses.Save()
	sessionID := ses.SessionId()
	SessionList.Store(sessionID, c)
}

func TestNet(t *testing.T) {
	localIP, err := net.LocalIP()
	fmt.Printf(localIP, err)
	mac, err := net.LocalMac()
	fmt.Printf(mac, err)
}

func TestOs(t *testing.T) {
	darwin := OS.IsDarwin()
	fmt.Println(darwin)
	pwd := OS.Pwd()
	t.Log(pwd)
}

func TestRand(t *testing.T) {
	r := rand.Int(0, 100)
	t.Log(r)
	rs := rand.String(10)
	t.Log(rs)
}
