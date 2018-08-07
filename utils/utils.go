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
	return GetBufferBody(buf)
}

// GetBufferBody extracts JSON body from byte stream
func GetBufferBody(body *bytes.Buffer) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal(body.Bytes(), &data)
	if err != nil {
		panic(err)
	}
	return data
}
