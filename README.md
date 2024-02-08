# MattermostBroadcast
An easy way to send direct message to many people and channels in mattermost

## How to use
1. insert your message in what.txt
2. insert в файле who.txt напишите адресатов с разделителем \n
3. Run Binary\Docker

Supports userName, user@Email, ~channel in who.txt

## Specific mode
You can use `--spec` or `-s` option when you need to send specific messages to certain people.
In this case it will read receipents and appropriate messages form `spec.csv` file.
It should follow this schema, without empty or unexepected fields:
```
who,what,
mattermostName,message for @mattermostName,
~channel1,message for ~channel1,
mattermostEmail,message for user associated with this email,
```

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
