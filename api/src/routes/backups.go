package routes

import (
	"net/http"

	"enix.io/banana/src/models"
	"enix.io/banana/src/services"
	"github.com/gin-gonic/gin"
)

// ServeBackupContainerList : Returns the list of containers (buckets)
func ServeBackupContainerList(context *gin.Context, _ *RequestIssuer) (int, interface{}) {
	list, err := services.Storage.ListBuckets()

	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, models.NewBackupContainerList(list)
}

// ServeBackupContainer : Returns a container filled with his top level objects
func ServeBackupContainer(context *gin.Context, _ *RequestIssuer) (int, interface{}) {
	name := context.Param("containerName")
	list, err := services.Storage.ListTopLevelObjectsInBucket(&name)

	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, models.NewFilledBackupContainer(&name, models.NewBackupTreeList(list))
}

// ServeBackupTree : Returns a backup tree filled with his backups
func ServeBackupTree(context *gin.Context, _ *RequestIssuer) (int, interface{}) {
	containerName := context.Param("containerName")
	treeName := context.Param("treeName")

	list, err := services.Storage.ListObjectsWithPrefixInBucket(&containerName, &treeName)
	if err != nil {
		return http.StatusBadRequest, err
	}

	backups, err := models.NewBackupList(list)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, models.NewFilledBackupTree(&treeName, backups)
}
