package utils

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

var GlobalConfig *Config

func LoadConfig(envPath string) *Config {
	var con Config

	if err := cleanenv.ReadConfig(envPath, &con); err != nil {
		panic(fmt.Sprintf("%s 文件加载失败：%s", envPath, err.Error()))
	}

	return &con
}
