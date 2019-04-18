package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"
)

// ParseJSONFromStream : Consume all bytes in the ReadCloser to a buffer and returns it parsed as JSON
func ParseJSONFromStream(reader io.ReadCloser) map[string]interface{} {
	bytes := ReadBytesFromStream(reader)
	data := make(map[string]interface{})
	json.Unmarshal(bytes, &data)
	return data
}

// ReadBytesFromStream : Consume all bytes in the ReadCloser to a buffer
func ReadBytesFromStream(reader io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.Bytes()
}

// GetDNFieldValue : Read value from distinguished name in X-Client-Subject-DN header
func GetDNFieldValue(dn, key string) (string, error) {
	entries := strings.Split(dn, ",")

	for _, elem := range entries {
		kv := strings.Split(elem, "=")
		if string(kv[0]) == key {
			return kv[1], nil
		}
	}

	return "", errors.New(key + " missing in certificate DN")
}
