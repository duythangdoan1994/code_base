package core

import (
	"os"
	"github.com/youtube/vitess/go/vt/log"
	"encoding/json"
)

type configuration struct {
	Port string
}

var AppConfig configuration

func InitConfig()  {
	loadAppConfig()
}

func loadAppConfig()  {
	file, err := os.Open("core/config.json")
	defer file.Close()
	if err != nil {
		log.Fatal("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
