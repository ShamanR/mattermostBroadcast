package mtClient

import "net/mail"

type DestinationType string

const DestinationUserName = DestinationType("userName")
const DestinationUserEmail = DestinationType("userEmail")
const DestinationChannel = DestinationType("channel")

func GetType(destination string) DestinationType {
	if destination[0] == '~' {
		return DestinationChannel
	}
	if _, err := mail.ParseAddress(destination); err == nil {
		return DestinationUserEmail
	}
	return DestinationUserName
}
