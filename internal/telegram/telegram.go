package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
	"tg_contour_bot/internal/cache"
	"tg_contour_bot/internal/settings"
)

type Service struct {
	Bot      *tgbotapi.BotAPI
	Settings *settings.Settings
	Cache    *cache.Cache
	Mutex    sync.Mutex
}

func NewService(admin, moderator, token string) *Service {
	s := settings.NewSettings(admin, moderator)
	c := cache.NewCache()
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		// in case of smth went wrong
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &Service{
		Bot:      bot,
		Settings: s,
		Cache:    c,
	}
}

func (s *Service) SendMessage(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.DisableWebPagePreview = true
	_, err := s.Bot.Send(msg)
	return err
}

func (s *Service) SendMenu(chatId int64, menu string, markup tgbotapi.InlineKeyboardMarkup) error {
	msg := tgbotapi.NewMessage(chatId, menu)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = markup
	msg.DisableWebPagePreview = true
	_, err := s.Bot.Send(msg)
	return err
}
