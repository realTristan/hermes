package cache

func (c *Cache) Import(data map[string]map[string]interface{}, schema map[string]bool, fullText bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, value := range data {
		if err := c.set(key, value, fullText); err != nil {
			return err
		}
	}
	return nil
}

func (c *Cache) ImportJson(file string, schema map[string]bool, fullText bool) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if data, err := readJson(file); err != nil {
		return err
	} else {
		for key, value := range data {
			if err := c.set(key, value, fullText); err != nil {
				return err
			}
		}
	}
	return nil
}
