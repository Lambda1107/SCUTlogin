package utils

import (
	"io/ioutil"
	"log"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

func March(s string, perfix string, suffix string) string {
	re := regexp.MustCompile(perfix + "(.*)" + suffix)
	result := re.FindStringSubmatch(s)
	return result[1]
}

type Config struct {
	Scode    string `yaml:"Scode"`
	Password string `yaml:"Password"`
}

func GetConfig() Config {
	var Aconfig Config
	File, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("配置信息文件不存在")
	}
	if err := yaml.Unmarshal(File, &Aconfig); err != nil {
		log.Fatal("配置信息反序列化失败")
	}
	return Aconfig
}
