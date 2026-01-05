package internal

import "Unofficial_API/utils"

var cache *utils.Cache

func init() {
	cache = utils.New(1 << 20)
}

func TryToGetToken(slug, name string) string {
	token, ok := cache.Get(name + slug)
	if ok {
		return token.(string)
	}
	return ""
}

func StoreToken(name, slug, token string) {
	cache.Add(name+slug, token)
}
