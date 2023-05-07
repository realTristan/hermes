package cache

// Set a value in the cache and in the full-text cache.
type WFT struct {
	value string
}

// Function to set a value in the cache and in the full-text cache.
func WithFT(value string) WFT {
	return WFT{value}
}

// Get the full-text value from a map
func fullTextMap(value interface{}) string {
	if _, ok := value.(map[string]interface{}); !ok {
		return ""
	}

	// Verify that the map has the correct length
	var v map[string]interface{} = value.(map[string]interface{})
	if len(v) == 0 || len(v) > 2 {
		return ""
	}

	// Verify that the map has the correct keys
	if _, ok := v["$hermes.full_text"]; !ok {
		return ""
	}
	if _, ok := v["value"]; !ok {
		return ""
	}

	// Verify that the map has the correct value types
	if _, ok := v["$hermes.full_text"].(bool); !ok {
		return ""
	}
	if _, ok := v["value"].(string); !ok {
		return ""
	}

	// Check if the full-text value is true
	if !v["$hermes.full_text"].(bool) {
		return ""
	}

	// Return the value
	return v["value"].(string)
}
