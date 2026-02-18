package utils

import (
	"fmt"
	"os"
	"server/config"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const configFileName = "config.yaml"

func LoadYaml() ([]byte, error) {
	return os.ReadFile(configFileName)
}

func SaveYAML() error {
	viper.AddConfigPath(".")
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func YamlUnmarshal(c *config.Config) error {
	viper.SetConfigFile(configFileName)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(c)
	if err != nil {
		panic(err)
	}
	//fmt.Println(c)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Watching config file for changes...")
		err = viper.Unmarshal(c)
		if err != nil {
			panic(err)
		}
	})
	return nil
}
