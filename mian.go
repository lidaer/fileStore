package main

import (
	cfg "fileStore/config"
	"fileStore/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func startApiService() {
	router := router.Router()
	router.Run(cfg.UploadServiceHost)
}

func main() {

	// api 服务
	go startApiService()

}
