package settings

import (
	"strconv"
)

type Settings struct {
	Admin           int64
	MaxParticipants int64
	Registration    bool
	MainPrize       string
	GiftCode        string
	Url             string
}

func NewSettings(admin string) *Settings {
	a, err := strconv.ParseInt(admin, 10, 64)
	if err != nil {
		return nil
	}
	return &Settings{
		Admin:           a,
		MaxParticipants: DefaultMaxParticipants,
		Registration:    DefaultRegistrationOpen,
		Url:             WwwUrl,
	}
}
