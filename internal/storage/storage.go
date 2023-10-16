package storage

import (
	"errors"
	// "log"
	"sync"
)

type storage struct {
	m *sync.Map
}
type Repository interface {
	AddURL(id, url string) error
	GetURL(id string) (string, error)
}

func New() *storage {
	return &storage{
		m: &sync.Map{},
	}
}

func (s *storage) AddURL(id, url string) error {
	s.m.Store(id, url)
	return nil
}

func (s *storage) GetURL(id string) (string, error) {
	url, ok := s.m.Load(id)
	if ok {
		return url.(string), nil
	}
	return "", errors.New("url not found")
}
