package model

import "time"

/*
	bucketName 桶名
	accountName ram用户名
	accessKeyId ak
	accessKeySecret sk
	area 区域
	projectNumber 项目编号
	projectUsage 项目用途
	label 标签
	createUptime 创建时间
	updateTime 更新时间
	status 状态 1 启用 2 禁用 3 删除
*/
type OssLabel struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Key   string `gorm:"type:varchar(255)"`
	Value string `gorm:"type:varchar(255)"`
}

func (OssLabel) TableName() string {
	return "oss_label"
}

type OssStorageInfoOssLabel struct {
	OssStorageInfoId int `gorm:"primaryKey"`
	LabelId          int `gorm:"primaryKey"`
}

func (OssStorageInfoOssLabel) TableName() string {
	return "oss_storage_info_labels"
}

type OssStorageInfo struct {
	Id              int64       `gorm:"primaryKey" json:"id"`
	BucketName      string      `gorm:"type:varchar(255);unique;"`
	AccountName     string      `gorm:"type:varchar(255);unique;"`
	AccessKeyId     string      `gorm:"type:varchar(255)"`
	AccessKeySecret string      `gorm:"type:varchar(255)"`
	Region          string      `gorm:"type:varchar(255)"`
	ProjectNumber   string      `gorm:"type:varchar(255)"`
	ProjectUsage    string      `gorm:"type:text"`
	Labels          []*OssLabel `gorm:"many2many:oss_storage_info_labels;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	StorageID       string `gorm:"type:varchar(255)"`
	Status          string `gorm:"type:varchar(1)"`
}

func (OssStorageInfo) TableName() string {
	return "oss_storage_info"
}

/*
	uid 用户id
	user 用户
	username 用户名
	email 邮箱
	department 部门
	buckets oss账号
	createTime 开通时间
	updateTime 更新时间
	status 状态 1 使用中 2 删除
*/
type AzureLabel struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Key   string `gorm:"type:varchar(255)"`
	Value string `gorm:"type:varchar(255)"`
}
type AzureStorageInfo struct {
	Id               int64         `gorm:"primaryKey" json:"id"`
	AccountName      string        `gorm:"type:varchar(255);unique;"`
	ResourceGroup    string        `gorm:"type:varchar(255)"`
	SecretKey        string        `gorm:"type:varchar(255)"`
	ConnectionString string        `gorm:"type:varchar(255)"`
	Region           string        `gorm:"type:varchar(255)"`
	ProjectNumber    string        `gorm:"type:varchar(255)"`
	ProjectUsage     string        `gorm:"type:text"`
	Labels           []*AzureLabel `gorm:"many2many:azure_storage_info_labels;"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	StorageID        string `gorm:"type:varchar(255)"`
	Status           string `gorm:"type:varchar(1)"`
}

func (AzureStorageInfo) TableName() string {
	return "azure_storage_info"
}

type AzureStorageInfoAzureLabel struct {
	AzureStorageInfoId int `gorm:"primaryKey"`
	LabelId            int `gorm:"primaryKey"`
}

func (AzureStorageInfoAzureLabel) TableName() string {
	return "azure_storage_info_labels"
}

type OssUserInfo struct {
	Id                int64              `gorm:"primary_key" json:"id"`
	User              string             `gorm:"type:varchar(255)"`
	Username          string             `gorm:"type:varchar(255);unique;"`
	Email             string             `gorm:"type:varchar(255);"`
	Phone             string             `gorm:"type:varchar(128);"`
	Department        string             `gorm:"type:varchar(255);"`
	OssStorageInfos   []OssStorageInfo   `gorm:"foreignKey:StorageID"`
	AzureStorageInfos []AzureStorageInfo `gorm:"foreignKey:StorageID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Status            string `gorm:"type:varchar(1)"`
}

func (OssUserInfo) TableName() string {
	return "oss_user_info"
}
