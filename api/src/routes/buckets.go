package routes

import (
	"net/http"

	"enix.io/banana/src/storage"
	"github.com/gin-gonic/gin"
)

// ServeBucketsList : Fetch and render as JSON the buckets list
func ServeBucketsList(store *storage.ObjectStorage) func(*gin.Context) (int, interface{}) {
	return func(c *gin.Context) (int, interface{}) {
		list, _ := store.ListBuckets()
		return http.StatusOK, list
	}
}
