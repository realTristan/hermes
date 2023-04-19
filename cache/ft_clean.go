package cache

// Clean the FullText cache with Mutex Locking
func (ft *FullText) Clean() {
	ft.mutex.Lock()
	defer ft.mutex.Unlock()
	ft.clean()
}

// Clean the FullText cache
func (ft *FullText) clean() {
	ft.wordCache = map[string][]string{}
	ft.data = map[string]map[string]interface{}{}
}
