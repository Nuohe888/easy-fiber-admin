package config

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
)

var filepath = "./config.toml"

func Init() {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败,路径(%s)", filepath))
	}

	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		panic(fmt.Sprintf("无法解析配置文件: %s", err))
	}
}
