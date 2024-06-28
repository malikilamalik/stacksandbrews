package model

import "sync"

type Menu struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MenuData struct {
	Data map[string]Menu
	*sync.RWMutex
}
