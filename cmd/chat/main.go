package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"simple_websocket/internal/chat"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "path", "configs/conf.toml", "provide path to config file")
}


func main() {

	flag.Parse()

	config := chat.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Не найден файл конфигурации. Будут использованы настройки по умолчанию")
	}

	server := chat.New(config)

	log.Fatal(server.Start())

}