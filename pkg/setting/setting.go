package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func loadBase() {
	// 由 Must 开头的方法名允许接收一个相同类型的参数来作为默认值，
	// 当键不存在或者转换失败时，则会直接返回该默认值。
	// 但是，MustString 方法必须传递一个默认值。
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		// err === nil 说明没有错误
		log.Fatalf("Fail to parse 'conf/app.ini: %v", err)
	}
	loadBase()
	loadServer()
}
