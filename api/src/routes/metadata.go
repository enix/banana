package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"enix.io/banana/src/services"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func receiveBackupMetadata(context *gin.Context, issuer *requestIssuer) (int, interface{}) {
	data := services.ReadBytesFromStream(context.Request.Body)
	messageID, _ := strconv.Atoi(context.Param("id"))
	err := services.Db.ZAdd(fmt.Sprintf("artifacts:%s:%s", issuer.Organization, issuer.CommonName), redis.Z{
		Score:  float64(messageID),
		Member: data,
	}).Err()

	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, "ok"
}

func serveBackupMetadata(context *gin.Context, issuer *requestIssuer) (int, interface{}) {
	key := fmt.Sprintf("artifacts:%s", context.Param("id"))
	_ = context.Param("artifactID")

	elems, err := services.Db.ZRevRange(key, 0, 100).Result()
	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, elems
}
