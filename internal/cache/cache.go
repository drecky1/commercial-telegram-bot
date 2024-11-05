package cache

import "sync"

const (
	WaitingForTitle = "waiting_for_title"
	WaitingForPhone = "waiting_for_phone"
	UpdatingTitle   = "updating_title"
	UpdatingPhone   = "updating_phone"
)

type UserInfo struct {
	TelegramUsername  string
	Title             string
	Phone             string
	ParticipantNumber int
}

type Cache struct {
	Participants      map[int64]UserInfo
	ParticipantsIDs   []int64
	ParticipateStates map[int64]string
	Mutex             sync.RWMutex
}

func NewCache() *Cache {
	var p []int64
	return &Cache{
		Participants:      make(map[int64]UserInfo),
		ParticipantsIDs:   p,
		ParticipateStates: make(map[int64]string),
	}
}

func (c *Cache) Get(key int64) (*UserInfo, bool) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	value, ok := c.Participants[key]
	if !ok {
		return nil, false
	}
	return &value, true
}

func (c *Cache) Set(key int64, value UserInfo) bool {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Participants[key] = value
	c.ParticipantsIDs = append(c.ParticipantsIDs, key)
	return true
}

func (c *Cache) UpdateTitle(key int64, title string) bool {
	ui := c.Participants[key]
	ui.Title = title
	if ok := c.Set(key, ui); !ok {
		return false
	}
	return true
}

func (c *Cache) UpdatePhone(key int64, phone string) bool {
	ui := c.Participants[key]
	ui.Phone = phone
	if ok := c.Set(key, ui); !ok {
		return false
	}
	return true
}
