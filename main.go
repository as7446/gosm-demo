//package main
//
//import (
//	"gosm/dao/mysql"
//	"gosm/router"
//)
//
//func main() {
//	mysql.InitMysql()
//	r := router.InitRouter()
//	r.Run(":8000")
//}
package main

import (
	"fmt"
	"gosm/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID           int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string
	CreatedAt    *time.Time `gorm:"column:create_at" json:"createAt"`
	Email        string     `gorm:"uniqueIndex;type:varchar(100)"`
	Role         string     `gorm:"size:255"`
	MemberNumber string     `gorm:"unique;not null"`
	Num          int
	Address      string `gorm:"index:addr"`
	IgnoreMe     string `gorm:"-"`
}

var DB *gorm.DB

func Create() {
	bucket := Bucket{
		BucketName: "bj-oss-shujiajia-m",
		Area:       "华北-2",
		Labels: []Label{
			{Key: "env", Valule: "product"},
		},
	}
	ram := Ram{
		User:            "fuxuhao",
		AccessKeyId:     "DJFKAJLEJAWLKTEA",
		AccessKeySecret: "FHNKLDSHGENHATLAEJRLKIJAELWJRELAWRA",
		Buckets:         []Bucket{bucket},
	}
	DB.Create(&ram)
}

func CreateTest() {
	//user := model.OssUserInfo{
	//	User:       "fuxuhao",
	//	Username:   "fuxuhao",
	//	Email:      "fuxuhao@datatang.com",
	//	Phone:      "15120027019",
	//	Department: "运维中心",
	//	OssStorageInfos: []model.OssStorageInfo{
	//		{
	//			BucketName:      "bj-oss-shujiajia-m",
	//			AccountName:     "fuxuhao",
	//			AccessKeyId:     "DJFKAJLEJAWLKTEA",
	//			AccessKeySecret: "FHNKLDSHGENHATLAEJRLKIJAELWJRELAWRA",
	//			Region:          "华北-2",
	//			ProjectNumber:   "US-DKJGLJLKAJKLER5646465",
	//			ProjectUsage:    "测试用的bucket",
	//			Labels: []*model.OssLabel{
	//				{
	//					Key:   "env",
	//					Value: "product",
	//				},
	//			},
	//			Status: "1",
	//		},
	//	},
	//	AzureStorageInfos: nil,
	//	CreatedAt:         time.Time{},
	//	UpdatedAt:         time.Time{},
	//	Status:            "1",
	//}
	//azure := model.AzureStorageInfo{
	//	Id:               0,
	//	AccountName:      "fuxuhao",
	//	ResourceGroup:    "fuxuhao",
	//	SecretKey:        "dskjaglejwrlaew",
	//	ConnectionString: "dsagjlewajtlllllllllllllllllllllllinkldsahgkelwhltiewuahlaegt",
	//	Region:           "美国西部",
	//	ProjectNumber:    "US-dsjagk-JDLSAFJDS-12545",
	//	ProjectUsage:     "jgakjddsajglkewjalrtaewr",
	//	Labels:           nil,
	//	CreatedAt:        time.Time{},
	//	UpdatedAt:        time.Time{},
	//	Status:           "1",
	//}
	u := model.OssUserInfo{Id: 2}
	DB.First(&u)
	//DB.Model(&u).Association("AzureStorageInfos").Append(&azure)
	fmt.Println(u)
	o := model.OssStorageInfo{}
	DB.Model(&u).Association("OssStorageInfos").Find(&o)
	fmt.Println(o)

}
func Update() {

}

func Select() {

}

func Delete() {

}

func (User) TableName() string {
	return "user"
}

type Bucket struct {
	BucketName string `gorm:"unique;not null"`
	Area       string
	Labels     []Label `gorm:"foreignKey:LabelID"`
	BucketID   uint
	gorm.Model
}

type Ram struct {
	gorm.Model
	User            string   `gorm:"not null"`
	AccessKeyId     string   `gorm:"unique"`
	AccessKeySecret string   `gorm:"unique"`
	Buckets         []Bucket `gorm:"foreignKey:BucketID"`
}

type Label struct {
	gorm.Model
	Key     string
	Valule  string
	LabelID uint
}

func (Bucket) TableName() string {
	return "bucket"
}
func (Ram) TableName() string {
	return "ram"
}
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接错误：" + err.Error())
	}
	DB = db
	DB.AutoMigrate(model.OssLabel{}, model.OssStorageInfo{}, model.OssUserInfo{})
	DB.AutoMigrate(model.AzureLabel{}, model.AzureStorageInfo{})
	DB.AutoMigrate(model.OssUserInfo{})
	fmt.Println("初始化完成.")
	CreateTest()
}
