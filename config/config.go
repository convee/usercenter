package conf

import (
	"os"
	"io/ioutil"
	"github.com/BurntSushi/toml"
	"usercenter/system"
)

type appConfig  struct {
	DB struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
	}
	Redis struct{
		Host string
		Port string
	}
	Option struct{
		SiteName  string
	}
}


func GetConfig() *appConfig {
	f, err := os.Open(system.RootPath() + "/config/config.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var appConf appConfig
	if err := toml.Unmarshal(buf, &appConf); err != nil {
		panic(err)
	}
	return &appConf
}




