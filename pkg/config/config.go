package config

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
)

var filepath = "./config.toml"

func Init() {
	InitWithPath(filepath)
}

func InitWithPath(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败,路径(%s)", path))
	}

	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		panic(fmt.Sprintf("无法解析配置文件: %s", err))
	}
}
