package cache

func Get(key string) (string, bool) {
	globalCache.Mu.RLock()
	defer globalCache.Mu.RUnlock()
	val, ok := globalCache.Data[key]
	return val, ok
}
