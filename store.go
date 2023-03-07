package main

import "sync"

type UrlStore struct {
	urls map[string]string // карта где будут хранится наши адреса
	mu   sync.RWMutex      // блокировка к операциям обновления
}

func NewUrlStore() *UrlStore { // функция для создания нашей структуры
	return &UrlStore{urls: make(map[string]string)}
}

func (s *UrlStore) Get(key string) string { // получение url адреса по ключу
	s.mu.RLock()
	url := s.urls[key]
	s.mu.RUnlock()
	return url
}

func (s *UrlStore) Set(key, url string) bool { // добавление нового url
	s.mu.Lock()
	if _, ok := s.urls[key]; ok {
		s.mu.Unlock()
		return false
	}
	s.urls[key] = url
	s.mu.Unlock()
	return true
}

func (s *UrlStore) Count() int { // подсчет url адресов
	s.mu.RLock()
	count := len(s.urls)
	s.mu.RUnlock()
	return count
}

func (s *UrlStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if s.Set(key, url) {
			return key
		}
	}
	return ""
}
