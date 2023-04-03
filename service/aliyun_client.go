package service

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ram20150501 "github.com/alibabacloud-go/ram-20150501/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	go_oss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gosm/config"
	"log"
	"os"
)

var (
	OssClient *go_oss.Client
	RamClient *ram20150501.Client
)

func InitOssClient() {
	client, err := go_oss.New(config.OssEndpoint, config.AliyunAccessKeySecret, config.AliyunAccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	log.Println("init oss success...")
	OssClient = client
}

func InitRamClient() {
	cf := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: tea.String(config.AliyunAccessKeyId),
		// 必填，您的 AccessKey Secret
		AccessKeySecret: tea.String(config.AliyunAccessKeySecret),
	}
	// 访问的域名
	cf.Endpoint = tea.String(config.RamEndpoint)
	_result, err := ram20150501.NewClient(cf)
	if err != nil {
		panic(fmt.Sprintf("err: %s ,ram client connection failed. ", err.Error()))
	}
	RamClient = _result
}
