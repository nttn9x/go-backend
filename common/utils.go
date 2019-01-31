package common

import (
	"encoding/json"
	"log"
	"os"
)

type jwtSetting struct {
	Key    string `json:"Key"`
	Issuer string `json:"Issuer"`
}

type elastic struct {
	URL    string `json:"Url"`
	Prefix string `json:"Prefix"`
}

type configuration struct {
	Port       string     `json:"Port"`
	Elastic    elastic    `json:"Elastic"`
	JWTSetting jwtSetting `json:"JWTSetting"`
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// InitConfig AppConfig
func InitConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("config.json")

	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}

	err = decoder.Decode(&AppConfig)

	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
