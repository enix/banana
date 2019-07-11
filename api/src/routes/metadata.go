package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"enix.io/banana/src/services"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func receiveBackupArtifacts(context *gin.Context, issuer *requestIssuer) (int, interface{}) {
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

func serveBackupArtifacts(context *gin.Context, issuer *requestIssuer) (int, interface{}) {
	key := fmt.Sprintf("artifacts:%s", context.Param("id"))
	messageID := context.Param("messageID")

	elems, err := services.Db.ZRevRangeByScore(key, redis.ZRangeBy{
		Min: messageID,
		Max: messageID,
	}).Result()
	if err != nil || len(elems) == 0 {
		return http.StatusNotFound, err
	}

	return http.StatusOK, []byte(elems[0])
}
