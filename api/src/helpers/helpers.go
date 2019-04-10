package helpers

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

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
