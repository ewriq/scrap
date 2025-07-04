package cache


func Delete(key string) {
	globalCache.Mu.Lock()
	defer globalCache.Mu.Unlock()
	delete(globalCache.Data, key)
}
