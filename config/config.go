package config

import (
	"flag"

	"github.com/spf13/viper"
)

var JwtSecret []byte

var Wechat struct {
	AppId     string
	AppSecret string
}

var Database struct {
	Uri  string
	Name string
}

func init() {
	confPath := flag.String("conf", "conf/prod.yml", "config file path")
	flag.Parse()

	readConfigFile(confPath)
}

func readConfigFile(confPath *string) {
	v := viper.New()
	v.SetConfigFile(*confPath)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	JwtSecret = []byte(v.Get("jwtSecret").(string))
	Wechat.AppId = v.Get("wechat.appId").(string)
	Wechat.AppSecret = v.Get("wechat.appSecret").(string)
	Database.Uri = v.Get("database.uri").(string)
	Database.Name = v.Get("database.name").(string)
}
