package nocache

// Get the full-text value from a map
func ftFromMap(value interface{}) string {
	if _, ok := value.(map[string]interface{}); !ok {
		return ""
	}

	// Verify that the map has the correct length
	var v map[string]interface{} = value.(map[string]interface{})
	if len(v) != 2 {
		return ""
	}

	// Verify that the map has the correct keys
	if ft, ok := v["$hermes.full_text"]; ok {
		if ft, ok := ft.(bool); ok {
			if !ft {
				return ""
			} else if v, ok := v["$hermes.value"]; ok {
				if v, ok := v.(string); ok {
					return v
				}
			}
		}
	}
	return ""
}
