package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var (
	configFile = "config.toml"
	infoFile   = "info.toml"
)


type Config struct {
	Redis   RedisConf
	Plat    PlatConf
	Server  ServerConf
}

type RedisConf struct {
	Addr string `toml:"addr"`
}

type PlatConf struct {
	Addr         string `toml:"addr"`
	Stdout       string `toml:"stdout"`
	Stderr       string `toml:"stderr"`
}

type ServerConf struct {
	HttpEnable        bool   `toml:"http_enable"`
	HttpAddr          string `toml:"http_addr"`
	HttpsEnable       bool   `toml:"https_enable"`
	HttpsAddr         string `toml:"https_addr"`
	BasicAuthUser     string `toml:"basic_auth_user"`
	BasicAuthPassword string `toml:"basic_auth_password"`
	Certs             string `toml:"certs"`
	Stdout            string `toml:"stdout"`
	Stderr            string `toml:"stderr"`
}

var Default = Config{
	Redis: RedisConf{
		Addr: "localhost:6379",
	},
	Plat: PlatConf{
		Addr: "localhost:50010",
	},
	Server: ServerConf{
		HttpEnable:        true,
		HttpAddr:          ":8080",
		HttpsEnable:       true,
		HttpsAddr:         ":443",
		BasicAuthUser:     "admin",
		BasicAuthPassword: "irene",
		Certs:             "/etc/certs",
	},
}

func (config *Config) Load() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("%v:\n", err)
	}
	fmt.Printf("dir=%v", dir)

	// Try to find the config in current path
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		// Try to find the config in binary path
		if _, err := toml.DecodeFile(dir+"/"+configFile, &config); err != nil {
			fmt.Printf("%v: Load default config ...\n", err)
			config = &Default
		}
	}
	fmt.Printf("Config: %v\n", config)
}

type Info struct {
	Server ServerInfo
}

type ServerInfo struct {
	Version string `toml:"version"`
	Commit  string `toml:"commit"`
	Date    string `toml:"date"`
}

func (info *Info) Load() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("%v:\n", err)
	}
	fmt.Printf("dir=%v", dir)

	// Try to find the config in current path
	if _, err := toml.DecodeFile(infoFile, &info); err != nil {
		// Try to find the config in binary path
		if _, err := toml.DecodeFile(dir+"/"+infoFile, &info); err != nil {
			fmt.Printf("Not able to find %v\n", infoFile)
			return
		}
	}
	fmt.Printf("Info: %v\n", info)
}
