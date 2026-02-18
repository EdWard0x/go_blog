package core

import (
	"server/config"
	"server/utils"
)

func InitConf() *config.Config {
	c := &config.Config{}
	err := utils.YamlUnmarshal(c)
	if err != nil {
		return nil
	}
	//time.Sleep(60 * time.Second)
	return c
}
