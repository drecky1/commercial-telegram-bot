package main

import (
	"bufio"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
	"tg_contour_bot/internal/handler"
	"tg_contour_bot/internal/telegram"
	"tg_contour_bot/utils"
)

func main() {
	admin := os.Getenv("ADMIN")
	if admin == "" {
		log.Println("Сделайте: export ADMIN=<CHAT_ID>")
		return
	}

	token := os.Getenv("TELEGRAM_APITOKEN")
	if token == "" {
		log.Println("Сделайте: export TELEGRAM_APITOKEN=<TOKEN>")
		return
	}

	service := telegram.NewService(admin, token)
	adminInt64, err := strconv.ParseInt(admin, 10, 64)
	if err != nil {
		log.Fatalf("Введенный админ не является числовым значение его чата в телеграмм. ")
	}
	go func() { _ = service.SendMessage(adminInt64, utils.WhatCanBotDo) }()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := service.Bot.GetUpdatesChan(u)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	receiveUpdates(ctx, updates, service)

	// Wait for a newline symbol, then cancel handling updates
	_, err = bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		return
	}
}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel, service *telegram.Service) {
	for {
		select {
		case <-ctx.Done():
			return
		case update := <-updates:
			go handler.HandleUpdate(update, service)
		}
	}
}
