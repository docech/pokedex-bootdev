package http

import (
	"io"
	"net/http"
)

type FetchFunc = func(string) ([]byte, error)

func fetch(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CachedFetch(config CacheConfig) FetchFunc {
	cache := newByteCache(config.MaxAge)

	return func(url string) ([]byte, error) {
		if data, ok := cache.Get(url); ok {
			return data, nil
		}

		data, err := fetch(url)
		if err != nil {
			return nil, err
		}

		cache.Set(url, data)

		return data, nil
	}
}