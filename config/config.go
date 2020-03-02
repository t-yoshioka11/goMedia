package config

import (
	"log"
	"os"

	ini "gopkg.in/ini.v1"
)

type ConfigList struct {
	LogFile string
}

var Config ConfigList

// Config 取得
func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println("Failed to read file : %v", err)
		os.Exit(1)
	}
	Config = ConfigList{
		LogFile: cfg.Section("logfile").Key("log_file").String(),
	}
}
