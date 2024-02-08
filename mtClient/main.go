package mtClient

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
)

type MtClient struct {
	Client *model.Client4
	Me     *model.User
}

// PostDirectMessageFromMe Simple func for sending direct messages
// In MM you can send messages only in channels! And direct - is a channel for two =)
func (m *MtClient) PostDirectMessageFromMe(userName string, message string) error {
	user, _, err := m.Client.GetUserByUsername(userName, "")
	if err != nil {
		return fmt.Errorf("cant get user %s: %w", userName, err)
	}
	channel, _, err := m.Client.CreateDirectChannel(user.Id, m.Me.Id)
	if err != nil {
		return fmt.Errorf("cant create direct channel for user %s: %w", userName, err)
	}
	_, _, err = m.Client.CreatePost(&model.Post{
		Message:   message,
		ChannelId: channel.Id,
	})

	if err != nil {
		return fmt.Errorf("cant send direct message for user %s: %w", userName, err)
	}
	return nil
}
