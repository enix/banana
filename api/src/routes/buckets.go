package routes

import (
	"errors"
	"net/http"

	"enix.io/banana/src/storage"
	"github.com/gin-gonic/gin"
)

// ServeBucketsList : Fetch and render as JSON the buckets list
func ServeBucketsList(store *storage.ObjectStorage) func(*gin.Context) (int, interface{}) {
	return func(c *gin.Context) (int, interface{}) {
		list, err := store.ListBuckets()

		if err != nil {
			return http.StatusInternalServerError, errors.New("failed to list buckets")
		}

		return http.StatusOK, list
	}
}
