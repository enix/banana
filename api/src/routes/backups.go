package routes

import (
	"net/http"

	"enix.io/banana/src/models"
	"enix.io/banana/src/storage"
	"github.com/gin-gonic/gin"
)

// ServeContainerList : Returns the list of containers (buckets)
func ServeContainerList(store *storage.ObjectStorage) RequestHandler {
	return func(context *gin.Context) (int, interface{}) {
		list, err := store.ListBuckets()

		if err != nil {
			return http.StatusInternalServerError, err
		}

		formattedList := make([]models.BackupContainer, len(list.Buckets))
		for index, elem := range list.Buckets {
			formattedList[index] = models.BackupContainer{
				Name: elem.Name,
			}
		}

		return http.StatusOK, formattedList
	}
}

// ServeBackupListFromContainer : Returns the list of available backups
//															in given bucket
func ServeBackupListFromContainer(store *storage.ObjectStorage) RequestHandler {
	return func(context *gin.Context) (int, interface{}) {
		list, err := store.ListTopLevelObjectsInBucket(context.Param("bucketName"))

		if err != nil {
			return http.StatusBadRequest, err
		}

		formattedList := make([]models.Backup, len(list.CommonPrefixes))
		for index, elem := range list.CommonPrefixes {
			formattedList[index] = models.Backup{
				Name: elem.Prefix,
			}
		}

		return http.StatusOK, formattedList
	}
}
