/**
 * Created by lock
 * Date: 2019-08-18
 * Time: 18:03
 */
package tools

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"time"

	"github.com/bwmarrin/snowflake"
)

const SessionPrefix = "sess_"

func GetSnowflakeId(nodeIDs ...int64) string {
	//default node id eq 1,this can modify to different serverId node
	var nodeID int64 = 1
	if len(nodeIDs) > 0 {
		nodeID = nodeIDs[0]
	}
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		panic(err)
	}
	// Generate a snowflake ID.
	id := node.Generate().String()
	return id
}

func GetRandomToken(length int) string {
	r := make([]byte, length)
	io.ReadFull(rand.Reader, r)
	return base64.URLEncoding.EncodeToString(r)
}

func CreateSessionId(sessionId string) string {
	return SessionPrefix + sessionId
}

func GetSessionIdByUserId(userId int) string {
	return fmt.Sprintf("sess_map_%d", userId)
}

func GetSessionName(sessionId string) string {
	return SessionPrefix + sessionId
}

func Sha1(s string) (str string) {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func GetNowDateTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}
