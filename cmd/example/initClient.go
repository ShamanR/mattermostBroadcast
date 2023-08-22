package main

import (
	"github.com/ShamanR/mattermostBroadcast/mtClient"
	"log"
	"os"
)

const (
	EnvVarMattermostUrl          = "MATTERMOST_URL"
	DefaultMattermostUrl         = "https://mt.avito.ru" // If you want to build binary without ane envs
	EnvVarMattermostAccessToken  = "MATTERMOST_ACCESS_TOKEN"
	DefaultMattermostAccessToken = "kynmrmmicjdq3rmnwfdhubedao" // If you want to build binary without ane envs
)

func newMTClient() *mtClient.MtClient {
	url := MustReadEnvOrDefault(EnvVarMattermostUrl, DefaultMattermostUrl)
	accessToken := MustReadEnvOrDefault(EnvVarMattermostAccessToken, DefaultMattermostAccessToken)
	return mtClient.NewClient(url, accessToken)
}

func MustReadEnvOrDefault(envVar, defVal string) string {
	envVal := os.Getenv(EnvVarMattermostUrl)
	if envVal == "" {
		envVal = defVal
	}
	if envVal == "" {
		log.Fatalf("Set %s env for continue", envVar)
	}
	return envVal
}
