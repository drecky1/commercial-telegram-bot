package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"tg_contour_bot/internal/cache"
	"tg_contour_bot/internal/settings"
	"tg_contour_bot/utils"
)

func ChangeParticipants(msg *tgbotapi.Message, s *settings.Settings) *tgbotapi.Message {
	var newMax int64
	_, err := fmt.Sscanf(msg.Text, "/change %d", &newMax)

	if err == nil {
		maxParticipants := newMax
		s.MaxParticipants = maxParticipants
		msg.Text = fmt.Sprintf(utils.UpdateCounfOfParticipants, s.MaxParticipants)
	} else {
		msg.Text = utils.WrongChangeCommand
	}
	return msg
}

func PickWinner(msg *tgbotapi.Message, c *cache.Cache) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if len(c.Participants) == 0 {
		msg.Text = utils.NoParticipantsMessage
		return
	}
	randomValue := rand.Intn(len(c.ParticipantsIDs))
	winnerInfo := c.Participants[c.ParticipantsIDs[randomValue]]

	msg.Text = fmt.Sprintf(utils.WinnerMessage, winnerInfo.ParticipantNumber, winnerInfo.Title, winnerInfo.Phone, winnerInfo.TelegramUsername)
}

func DeleteParticipants(msg *tgbotapi.Message, s *settings.Settings, c *cache.Cache) {
	if !s.Registration {
		c.Mutex.Lock()
		defer c.Mutex.Unlock()
		var p []int64
		newMap := make(map[int64]cache.UserInfo)
		c.Participants = newMap
		c.ParticipantsIDs = p
		msg.Text = utils.ListOfParticipantsDeleted
	} else {
		msg.Text = utils.WrongDeleteParticipantsNotClosed
	}
}

func CheckSettings(msg *tgbotapi.Message, s *settings.Settings, c *cache.Cache) {
	var text string
	if s.Registration {
		text = "ДА"
	} else {
		text = "НЕТ"
	}
	msg.Text = fmt.Sprintf(utils.NowSettingsMessage, s.MaxParticipants, text, len(c.ParticipantsIDs), s.GiftCode, s.MainPrize)
}

func ShowParticipants(msg *tgbotapi.Message, c *cache.Cache) {
	var text string
	if len(c.ParticipantsIDs) == 0 {
		msg.Text = utils.NoParticipantsMessage
		return
	}
	for _, participantID := range c.ParticipantsIDs {
		v, ok := c.Participants[participantID]
		if !ok {
			continue
		}
		text += fmt.Sprintf("%d %s %s %s\n", v.ParticipantNumber, v.Title, v.Phone, v.TelegramUsername)
	}
	msg.Text = text
}
