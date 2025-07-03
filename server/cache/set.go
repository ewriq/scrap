package cache
func Set(key, value string) {
	globalCache.Mu.Lock()
	defer globalCache.Mu.Unlock()
	globalCache.Data[key] = value
}

