# Slack Age Bot
### Requirements:
- Install [Go](https://go.dev/doc/install)
- Slack Token

### Get Slack Token:
Create a new app in [slack api](https://api.slack.com/apps)
Get SOCKET_APP_TOKEN:
- In your app, go to **OAuth & Permissions** => **Socket Mode**
- Turn on **Enable Socket Mode** and receive SOCKET_APP_TOKEN.

Get SOCKET_BOT_TOKEN:
- Go to **Event Subscriptions**, turn on **Enable Events**
- **Add Bot User Event**  in **Subscribe to bot events**, add some basic events like: *app_mention, im_history_changed, message.im, message.channels, message.mpim*... and **Save Changes**
- Go to **OAuth & Permissions**, add some OAuth Scope in **Bot Token Scopes**: *chat:write, chat:write.public, channels:read, im:read, im:write, mpim:read, mpim:write* and **Install to Workspace**

### Run:
```sh
cd slack-age-bot
go build
go run main.go
```
```
Connecting to Slack with Socket Mode.
Connected to Slack with Socket Mode.
Connected as App ID A063V3HA1HD
```