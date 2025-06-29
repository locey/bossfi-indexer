package main

import (
	"bossfi-indexer/src/core"
	_ "bossfi-indexer/src/docs"
)

const (
	// ConfigFile 配置文件路径
	ConfigFile = "config.toml"
)

func main() {
	core.Start(ConfigFile)
}
