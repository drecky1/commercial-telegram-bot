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
	//pids := []int64{1, 2, 3, 4, 5, 6}
	//ps := map[int64]UserInfo{
	//	1: {
	//		TelegramUsername:  "1",
	//		Title:             "Denis1",
	//		Phone:             "+79999999901",
	//		ParticipantNumber: 1,
	//	},
	//	2: {
	//		TelegramUsername:  "2",
	//		Title:             "Denis2",
	//		Phone:             "+79999999902",
	//		ParticipantNumber: 2,
	//	},
	//	3: {
	//		TelegramUsername:  "3",
	//		Title:             "Denis3",
	//		Phone:             "+79999999903",
	//		ParticipantNumber: 3,
	//	},
	//	4: {
	//		TelegramUsername:  "4",
	//		Title:             "Denis4",
	//		Phone:             "+79999999904",
	//		ParticipantNumber: 4,
	//	},
	//	5: {
	//		TelegramUsername:  "5",
	//		Title:             "Denis5",
	//		Phone:             "+79999999905",
	//		ParticipantNumber: 5,
	//	},
	//	6: {
	//		TelegramUsername:  "6",
	//		Title:             "Denis6",
	//		Phone:             "+79999999906",
	//		ParticipantNumber: 6,
	//	},
	//}
	pids := []int64{}
	ps := map[int64]UserInfo{}
	return &Cache{
		Participants:      ps,
		ParticipantsIDs:   pids,
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
