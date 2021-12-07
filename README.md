# rejectChannels
Build from go

Clone this repo
cd rejectChannels

`go mod download &&
go mod tidy &&
go build -o bot cmd/bot/main`

Docker:

`docker build -fDockerfile -t rejectChannels . &&
docker run -d -e BOT_TOKEN=<token> --name rejectChannels \
rejectChannels`

running bot: https://t.me/rejectchannelsbot