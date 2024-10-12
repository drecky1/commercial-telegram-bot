package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg_contour_bot/internal/settings"
	"tg_contour_bot/utils"
)

var (
	// Menu texts
	startMenu       = fmt.Sprintf(utils.HelloMessage, utils.CompanyName)
	participateMenu = utils.ParticipateMessage

	// Button texts
	nextButton        = utils.Next
	backButton        = utils.Back
	participateButton = utils.Participate
	changeTitle       = utils.UpdateTitle
	changePhone       = utils.UpdatePhone
	giftButton        = utils.SendGift

	// Keyboard layout for the first menu. One button, one row
	startMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(nextButton, nextButton),
		),
	)

	// Keyboard layout for the second menu. Two buttons, one per row
	participateMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(participateButton, participateButton),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(backButton, backButton),
		),
	)

	updateParticipationDataMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(giftButton, giftButton),
		),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData(changeTitle, changeTitle),
		//),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData(changePhone, changePhone),
		//),
	)

	webUrlDataMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("За подарком ♥️", settings.WwwUrl)),
	)
)
