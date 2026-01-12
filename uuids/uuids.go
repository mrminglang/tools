package uuids

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/mrminglang/tools/times"

	rand2 "math/rand"
)

// GenerateUuid 雪花算法生成19位唯一id
func GenerateUuid() string {
	// Create a new Node with a Node number of 1
	node, _ := snowflake.NewNode(int64(1))

	// Generate a snowflake ID.
	id := node.Generate().String()
	return id
}

// Uuid 返回有连接符的UUID
func Uuid() string {
	return uuid.New().String()
}

// GetRandomUUID 返回无连接符的UUID
func GetRandomUUID() string {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	str := newUUID.String()
	strArray := strings.Split(str, "-")
	str = strings.Join(strArray, "")
	return str
}

// GenerateUniqueNonceStr 生成唯一性随机字符串
func GenerateUniqueNonceStr(length int) (string, error) {
	nonceStr, err := GenerateNonceStr(length)
	if err != nil {
		return "", err
	}
	// 添加时间戳作为唯一标识符
	timestamp := time.Now().UnixNano()
	nonceStr = fmt.Sprintf("%s%d", nonceStr, timestamp)
	return nonceStr, nil
}

// GenerateNonceStr 生成随机字符串
func GenerateNonceStr(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Use base64 encoding to generate the random bytes
	nonceStr := base64.URLEncoding.EncodeToString(randomBytes)

	// Trim the nonceStr to the specified length
	if len(nonceStr) > length {
		nonceStr = nonceStr[:length]
	}

	return nonceStr, nil
}

// 生成随机数去重结构
var genUidDupLock sync.Mutex
var genUidDup = struct {
	TimeKey string
	DupMap  map[string]struct{}
}{DupMap: make(map[string]struct{})}

// GetSimpleRandId 获取简单的uid 使用时间戳
func GetSimpleRandId() string {
	now := time.Now()
	return GetSimpleRandIdWithTime(now)
}

// GetSimpleRandIdWithTime 使用指定时间生成简单的uid
func GetSimpleRandIdWithTime(now time.Time) string {
	genId := func(now time.Time) (id string, timeKey string) {
		date := now.Format(times.YYMMDD)
		daySecond := strconv.Itoa(1e5 + now.Hour()*3600 + now.Minute()*60 + now.Second())[1:]
		randNum := strconv.Itoa(1e5 + rand2.Intn(99999))[1:]
		return date + daySecond + randNum, date + daySecond
	}

	//是否重复 控制循环
	isDup, resId := true, ""
	for isDup {
		id, timeKey := genId(now)
		genUidDupLock.Lock()
		if genUidDup.TimeKey != timeKey {
			//新的一秒 重置
			genUidDup.TimeKey = timeKey
			genUidDup.DupMap = make(map[string]struct{})
		}
		if _, ok := genUidDup.DupMap[id]; !ok {
			//生成的id不存在 未重复
			genUidDup.DupMap[id] = struct{}{}
			isDup = false
			resId = id
		} else {
			//重复 重试
			isDup = true
		}
		genUidDupLock.Unlock()
	}

	return resId
}

// GetTaskId 获取任务id
func GetTaskId() string {
	return GetSimpleRandId()
}

// GetTaskIdWithTime 获取指定时间的任务id
func GetTaskIdWithTime(now time.Time) string {
	return GetSimpleRandIdWithTime(now)
}

// IsTaskID 是否是任务ID
func IsTaskID(taskID string) bool {
	if len(taskID) != 16 {
		return false
	}
	if _, err := time.Parse(times.YYMMDD, taskID[:6]); err != nil {
		return false
	}
	for i := 0; i < len(taskID); i++ {
		if taskID[i] < '0' || taskID[i] > '9' {
			return false
		}
	}

	return true
}
