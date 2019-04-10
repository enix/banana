package routes

import (
	"net/http"

	"enix.io/banana/src/models"
	"enix.io/banana/src/storage"
	"github.com/gin-gonic/gin"
)

// ServeBackupContainerList : Returns the list of containers (buckets)
func ServeBackupContainerList(store *storage.ObjectStorage) RequestHandler {
	return func(context *gin.Context, _ *RequestIssuer) (int, interface{}) {
		list, err := store.ListBuckets()

		if err != nil {
			return http.StatusBadRequest, err
		}

		return http.StatusOK, models.NewBackupContainerList(list)
	}
}

// ServeBackupContainer : Returns a container filled with his top level objects
func ServeBackupContainer(store *storage.ObjectStorage) RequestHandler {
	return func(context *gin.Context, _ *RequestIssuer) (int, interface{}) {
		name := context.Param("containerName")
		list, err := store.ListTopLevelObjectsInBucket(&name)

		if err != nil {
			return http.StatusBadRequest, err
		}

		return http.StatusOK, models.NewFilledBackupContainer(&name, models.NewBackupTreeList(list))
	}
}

// ServeBackupTree : Returns a backup tree filled with his backups
func ServeBackupTree(store *storage.ObjectStorage) RequestHandler {
	return func(context *gin.Context, _ *RequestIssuer) (int, interface{}) {
		containerName := context.Param("containerName")
		treeName := context.Param("treeName")

		list, err := store.ListObjectsWithPrefixInBucket(&containerName, &treeName)
		if err != nil {
			return http.StatusBadRequest, err
		}

		backups, err := models.NewBackupList(list)
		if err != nil {
			return http.StatusBadRequest, err
		}

		return http.StatusOK, models.NewFilledBackupTree(&treeName, backups)
	}
}
