package cache

import "sync"

var cache = sync.Map{}

func GetCachedURL(shortURL string) (string, bool) {
	if url, found := cache.Load(shortURL); found {
		return url.(string), true
	}
	return "", false
}

func SetCachedURL(shortURL, originalURL string) {
	cache.Store(shortURL, originalURL)
}
