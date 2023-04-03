package config

import (
	"fmt"
	"time"
)

func GetRegionName(regionId string) string {
	regionMap := map[string]string{"oss-cn-beijing": "bj-oss"}
	return regionMap[regionId]
}

func GetBucketName(username string, region string) string {
	if username == "" || region == "" {
		return ""
	}
	timeFormat := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s-%s-%s", GetRegionName(region), username, timeFormat)
}
