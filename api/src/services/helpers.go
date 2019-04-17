package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

// ParseJSONFromStream : Consume all bytes in the ReadCloser to a buffer and returns it
//							 parsed ad JSON
func ParseJSONFromStream(reader io.ReadCloser) map[string]interface{} {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	data := make(map[string]interface{})
	json.Unmarshal(buf.Bytes(), &data)
	return data
}

// GetDNFieldValue : Read value from distinguished name in X-Client-Subject-DN header
func GetDNFieldValue(context *gin.Context, key string) (string, error) {
	issuer := context.GetHeader("X-Client-Subject-DN")
	entries := strings.Split(issuer, ",")

	for _, elem := range entries {
		kv := strings.Split(elem, "=")
		if string(kv[0]) == key {
			return kv[1], nil
		}
	}

	return "", errors.New(key + " missing in X-Client-Subject-DN header")
}