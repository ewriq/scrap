package model

import "sync"

type Cache struct {
	Data map[string]string
	Mu   sync.RWMutex
}