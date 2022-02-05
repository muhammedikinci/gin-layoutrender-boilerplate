package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Mode string
}

func NewConfig(path, fileName string) *Config {
	ex, _ := os.Executable()
	viper.AddConfigPath(filepath.Join(filepath.Dir(ex), path))
	viper.SetConfigName(fileName)
	err := viper.ReadInConfig()

	if err != nil {
		viper.AddConfigPath(path)
		viper.SetConfigName(fileName)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)

	if err != nil {
		panic(err)
	}

	return conf
}
