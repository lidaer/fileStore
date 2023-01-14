package main

import (
	cfg "fileStore/config"
	"fileStore/router"
)

func main() {
	// gin framework
	router := router.Router()
	router.Run(cfg.UploadServiceHost)
}
