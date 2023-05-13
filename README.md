# MattermostBroadcast
An easy way to send a message to many people and channels in mattermost

## How to use
1. insert your message in what.txt
2. insert в файле who.txt напишите адресатов с разделителем \n
3. Run Binary\Docker

Supports userName, user@Email, ~channel in who.txt

## How to build binary
You will need Golang 1.20 installed
```go build -o ./sayBatch ./cmd/sayBatch/main.go```

## Configuration
Is made by ENV

`MATTERMOST_URL` - for url of your Mattermost Server

`MATTERMOST_ACCESS_TOKEN` - for your personal mattermost access token

### Where you can find your MATTERMOST_ACCESS_TOKEN
The `MATTERMOST_ACCESS_TOKEN` is your personal token from Web-Mattermost.
Just use value of `MMAUTHTOKEN` cookie in the browser console.
