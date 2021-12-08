package util

import (
	"CLogServer/pkg/setting"
	"log"
	"os"
)

func Setup() {
	path := setting.AppSetting.LogSavePath
	_, err := os.Stat(path)
	if err != nil {
		createDir(path)
	}
	initJWTSecret()
}

func initJWTSecret() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

func createDir(dir string) bool {
	err := os.Mkdir(dir, 0777)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
