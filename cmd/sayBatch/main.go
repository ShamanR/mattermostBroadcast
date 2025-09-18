package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ShamanR/mattermostBroadcast/mtClient"
	"github.com/gocarina/gocsv"
)

const (
	EnvVarMattermostUrl          = "MATTERMOST_URL"
	DefaultMattermostUrl         = "" // If you want to build binary without ane envs
	EnvVarMattermostAccessToken  = "MATTERMOST_ACCESS_TOKEN"
	DefaultMattermostAccessToken = "" // If you want to build binary without ane envs
)

type Message struct {
	Recipients  string `csv:"who"`
	MessageText string `csv:"what"`
}

func main() {
	// url should not contain trailing slash
	url := strings.TrimRight(MustReadEnvOrDefault(EnvVarMattermostUrl, DefaultMattermostUrl), "/")
	accessToken := MustReadEnvOrDefault(EnvVarMattermostAccessToken, DefaultMattermostAccessToken)
	client := mtClient.NewClient(url, accessToken)
	args := os.Args[1:]

	if len(args) > 0 && (args[0] == "--spec" || args[0] == "-s") {
		log.Printf("Sending messages from spec.csv\n")
		messages := getSpecList()
		for _, msg := range messages {
			if msg.Recipients == "" || msg.MessageText == "" {
				log.Fatal("invalid csv format, see README.md")
			}
		}
		for _, msg := range messages {
			sendMessage(client, strings.Split(msg.Recipients, ","), msg.MessageText)
		}
		os.Exit(0)
	}

	message := getMessage()
	destinations := getDestinations()
	sendMessage(client, destinations, message)
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

	data, err := os.ReadFile(filepath.Join(path, "what.txt"))
	if err != nil {
		log.Fatal("cant read what.txt ", err)
	}
	return strings.TrimSpace(string(data))
}

func getDestinations() []string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(filepath.Join(path, "who.txt"))
	if err != nil {
		log.Fatal("cant read who.txt ", err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func sendMessage(client *mtClient.MtClient, destinations []string, message string) {
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
			} else {
				log.Printf("[OK] Message send to [userName]%s\n", destination)
			}
		default:
			log.Printf("[Error] Unknown destination %s for %s\n", mtClient.GetType(destination), destination)
		}
	}
}

func getSpecList() []*Message {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	in, err := os.Open(filepath.Join(path, "spec.csv"))
	if err != nil {
		log.Fatal("can't read spec.csv ", err)
	}
	defer in.Close()

	messages := []*Message{}
	if err := gocsv.UnmarshalFile(in, &messages); err != nil {
		log.Fatal("can't parse spec.csv ", err)
	}

	return messages
}
