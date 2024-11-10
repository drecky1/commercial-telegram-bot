package settings

import (
	"strconv"
)

type Settings struct {
	Admin           int64
	Moderator       int64
	MaxParticipants int64
	Registration    bool
	MainPrize       string
	GiftCode        string
	Url             string
}

func NewSettings(admin, moderator string) *Settings {
	a, err := strconv.ParseInt(admin, 10, 64)
	if err != nil {
		return nil
	}
	m, err := strconv.ParseInt(moderator, 10, 64)
	if err != nil {
		return nil
	}
	return &Settings{
		Admin:           a,
		Moderator:       m,
		MaxParticipants: DefaultMaxParticipants,
		Registration:    DefaultRegistrationOpen,
		Url:             WwwUrl,
	}
}
