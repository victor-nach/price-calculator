package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	defaultPort = "8080"
	defaultUrl  = "https://api.coindesk.com/v1/bpi/currentprice.json"
)

// Secrets contain all the config that this application needs
type Secrets struct {
	Port        string `json:"port"`
	CoindeskURL string `json:"coindesk_url"`
}

// LoadSecrets loads secrets from the environment and returns it
// if a .env file is present, it would be loaded first
// default values are also set
func LoadSecrets(filename ...string) *Secrets {
	f := ".env"
	if len(filename) > 0 {
		f = filename[0]
	}
	_ = godotenv.Load(f)
	secrets := &Secrets{}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}
	secrets.Port = port

	coindeskUrl, ok := os.LookupEnv("COINDESK_URL")
	if !ok {
		coindeskUrl = defaultUrl
	}
	secrets.CoindeskURL = coindeskUrl

	return secrets
}
