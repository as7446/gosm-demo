package service

import (
	"errors"
	"gosm/config"
	"gosm/model"
	"log"
	"time"
)

type ossCell model.OssStorageInfo

func (o *ossCell) GetName(filterType string) string {
	switch filterType {
	case "bucket_name":
		return o.BucketName
	case "project_usage":
		return o.ProjectUsage
	case "project_number":
		return o.ProjectNumber
	case "region":
		return o.Region
	default:
		return ""
	}
}

func (o *ossCell) GetCreation() time.Time {
	return o.CreatedAt
}

var Oss oss

type oss struct {
}

func (o *oss) CreateStorage(user model.OssUserInfo, bucketName, region, project_number, project_usage string, labels map[string]string) (err error) {
	ossLabelList := []*model.OssLabel{}
	if len(labels) != 0 {
		for key, value := range labels {
			ossLabelList = append(ossLabelList, &model.OssLabel{Key: key, Value: value})
		}
	}
	if bucketName == "" {
		bucketName = config.GetBucketName(user.Username, region)
	}
	if project_number == "" || project_usage == "" {
		return errors.New("项目编号和项目用途不能为空.")
	}
	bucket, err := OssClient.GetBucketInfo(bucketName)
	if err == nil {
		return errors.New("Bucket 已存在")
	}
	log.Println("bucketGet:", bucket.BucketInfo)
	log.Println("name:", bucketName, bucket)

	//ossInfo := model.OssStorageInfo{
	//	Id:              0,
	//	BucketName:      bucketName,
	//	AccountName:     "",
	//	AccessKeyId:     "",
	//	AccessKeySecret: "",
	//	Region:          "",
	//	ProjectNumber:   "",
	//	ProjectUsage:    "",
	//	Labels:          nil,
	//	CreatedAt:       time.Time{},
	//	UpdatedAt:       time.Time{},
	//	StorageID:       "",
	//	Status:          "",
	//}
	//mysql.DB.Model(&user).Association("OssStorageInfos").Append(&ossInfo)
	return
}

func (o *oss) DeleteStorage() (err error) {
	return
}

func (o *oss) UpdateStorage() (err error) {
	return
}

func (o *oss) GetStorage() {
	return
}

func (o *oss) GetStorageDetail() {

}
