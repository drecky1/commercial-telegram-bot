#!/bin/bash
APP_NAME="APP_NAME"
ADMIN_CHAT_ID=0
TELEGRAM_API_TOKEN="SOME_BEAUTIFUL_TELEGRAM_TOKEN_FROM_BOT_FATHER"
PARAMS="-admin $ADMIN_CHAT_ID -token $TELEGRAM_API_TOKEN"
GITHUB_URL="git@github.com:drecky1/commercial-telegram-bot.git"

DIR="/path/to/project/on/your/system"
# checking folder
if [ -d "$DIR" ]; then
        echo "Directory $DIR exists."
        git checkout
    else
        echo "Directory $DIR does not exist."
        git clone $GITHUB_URL
    fi
# build binary file
env CGO_ENABLED=0 go build -v -o bot cmd/bot/main.go
# check OS
case "$(uname -s)" in
    Linux)
        COMMAND="./$DIR/$APP_NAME"
        ;;
    Darwin)
        COMMAND="./$DIR/$APP_NAME"
        ;;
    CYGWIN*|MINGW32*|MSYS*|MINGW*)
        COMMAND="./$DIR/$APP_NAME.exe"
        ;;
    *)
        echo "Unsupported OS: $(uname -s)"
        exit 1
        ;;
esac
# start bot
$COMMAND $PARAMS