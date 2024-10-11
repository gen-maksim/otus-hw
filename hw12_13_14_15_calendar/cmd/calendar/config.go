package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger      LoggerConf
	StorageType string
}

type LoggerConf struct {
	Level string
	// TODO
}

func NewConfig(path string) Config {
	rawConf, err := os.ReadFile(path)
	if err != nil {
		log.Println("bad path")
		log.Println(err)
		return Config{}
	}

	var conf Config
	if _, err := toml.Decode(string(rawConf), &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}
