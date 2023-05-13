package mtClient

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"log"
	"net/http"
)

func NewClient(url string, mattermostAuthToken string) *MtClient {
	client := model.Client4{
		URL:        url,
		APIURL:     url + model.APIURLSuffix,
		HTTPClient: &http.Client{},
		AuthToken:  mattermostAuthToken,
		AuthType:   model.HeaderBearer,
		HTTPHeader: nil,
	}
	// because we send all messages from ourselves we need me user
	// also its healthcheck for our client
	meUser, _, err := client.GetUser("me", "")
	if err != nil {
		log.Fatalf("error of getting `me` user : %s", err.Error())
	}
	return &MtClient{
		Client: &client,
		Me:     meUser,
	}
}
