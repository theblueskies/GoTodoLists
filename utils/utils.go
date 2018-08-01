package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

// GetJSONBody converts the body to a convenient map
func GetJSONBody(body io.ReadCloser) map[string]interface{} {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	var data map[string]interface{}
	_ = json.Unmarshal(buf.Bytes(), &data)
	return data
}
