package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var Conf *Config

type Config struct {
	App     AppConfig
	Monitor MonitorConfig
	Pgsql   PgsqlConfig
	Redis   RedisConfig
	Kafka   KafkaConfig
	Chains  []ChainConfig
}

type AppConfig struct {
	Name    string `toml:"name" json:"name"`
	Port    string `toml:"port" json:"port"`
	Version string `toml:"version" json:"version"`
	ChainId int    `toml:"chain_id" json:"chainId"` // 当前主链id
}

type MonitorConfig struct {
	PprofEnable bool `toml:"pprof_enable" json:"pprofEnable"`
	PprofPort   int  `toml:"pprof_port" json:"pprofPort"`
}

type PgsqlConfig struct {
	Host     string `toml:"host" json:"host"`
	Port     string `toml:"port" json:"port"`
	Username string `toml:"username" json:"username"`
	Password string `toml:"password" json:"password"`
	Database string `toml:"database" json:"database"`
}

type RedisConfig struct {
	Host        string `toml:"host" json:"host"`
	Port        string `toml:"port" json:"port"`
	Password    string `toml:"password" json:"password"`
	Db          int    `toml:"db" json:"db"`
	MaxIdle     int    `toml:"max_idle" json:"maxIdle"`
	MaxActive   int    `toml:"max_active" json:"maxActive"`
	IdleTimeout int    `toml:"idle_timeout" json:"idleTimeout"`
}

type KafkaConfig struct {
	Brokers []string `toml:"brokers" json:"brokers"`
	Topic   string   `toml:"topic" json:"topic"`
	GroupID string   `toml:"group_id" json:"group_id"`
}

type ChainConfig struct {
	Name               string `toml:"name" json:"name"`
	ChainId            int    `toml:"chain_id" json:"chainId"`
	Endpoint           string `toml:"endpoint" json:"endpoint"`
	BossTokenAddress   string `toml:"boss_token_address" json:"bossTokenAddress"`
	BossStakingAddress string `toml:"boss_staking_address" json:"bossStakingAddress"`
	StartBlockNumber   int64  `toml:"start_block_number" json:"startBlockNumber"`
}

// InitConfig 初始化配置
func InitConfig(configFile string) *Config {
	// 尝试从以下路径顺序加载：
	// 1. 当前工作目录下的 configs/ 子目录
	// 2. 可执行文件所在目录的 configs/ 子目录
	paths := []string{
		filepath.Join("./configs", configFile),
		filepath.Join(filepath.Dir(os.Args[0]), "configs", configFile),
	}
	var tomlFile string
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			tomlFile = path
			break
		}
	}

	if tomlFile == "" {
		// 没有找到配置文件，请检查以下两个目录
		panic("config file not found, please check the following directories:" + strings.Join(paths, ","))
	}

	conf := Config{}
	if _, err := toml.DecodeFile(tomlFile, &conf); err != nil {
		panic("read toml file err: " + err.Error())
	}

	Conf = &conf
	return &conf
}

func getCurrentAbPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}

	// 获取当前文件所在目录
	dir := filepath.Dir(filename)

	// 获取上两级目录
	abPath := filepath.Join(dir, "..", "..", "..")

	// Clean 会清理多余的 ../ 和 . 等符号，确保路径合法
	clean := filepath.Clean(abPath)
	return clean
}

func getConfigAbPath() string {
	return getCurrentAbPath() + "/configs"
}
