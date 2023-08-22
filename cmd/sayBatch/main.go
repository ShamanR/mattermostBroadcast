package main

import (
	"fmt"
	"github.com/ShamanR/mattermostBroadcast/mtClient"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	EnvVarMattermostUrl          = "MATTERMOST_URL"
	DefaultMattermostUrl         = "" // If you want to build binary without ane envs
	EnvVarMattermostAccessToken  = "MATTERMOST_ACCESS_TOKEN"
	DefaultMattermostAccessToken = "" // If you want to build binary without ane envs
)

func main() {
	url := MustReadEnvOrDefault(EnvVarMattermostUrl, DefaultMattermostUrl)
	accessToken := MustReadEnvOrDefault(EnvVarMattermostAccessToken, DefaultMattermostAccessToken)
	client := mtClient.NewClient(url, accessToken)

	message := getMessage()
	destinations := getDestinations()
	for _, destination := range destinations {
		switch mtClient.GetType(destination) {
		case mtClient.DestinationChannel:
			// todo
			log.Printf("[Error] Channel for %s are not supported yet. Skipped", destination)
		case mtClient.DestinationUserEmail:
			// todo
			log.Printf("[Error] Email for %s are not supported yet. Skipped", destination)
		case mtClient.DestinationUserName:
			err := client.PostDirectMessageFromMe(destination, message)
			if err != nil {
				log.Printf("[Error] error occured while sending message to %s: %s", destination, err.Error())
			}
			log.Printf("[OK] Message send to [userName]%s\n", destination)
		default:
			log.Printf("[Error] Unknown destination %s for %s\n", mtClient.GetType(destination), destination)
		}
	}

}
func MustReadEnvOrDefault(envVar, defVal string) string {
	envVal := os.Getenv(envVar)
	if envVal == "" {
		envVal = defVal
	}
	if envVal == "" {
		log.Fatalf("Set %s env for continue", envVar)
	}
	return envVal
}

func getMessage() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)

	data, err := os.ReadFile(filepath.Join(path, "what.txt"))
	if err != nil {
		log.Fatal("cant read what.txt ", err)
	}
	return string(data)
}

func getDestinations() []string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)

	data, err := os.ReadFile(filepath.Join(path, "who.txt"))
	if err != nil {
		log.Fatal("cant read who.txt ", err)
	}
	return strings.Split(string(data), "\n")
}
