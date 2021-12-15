package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	appToken      string
	testChannelId string
	serverPort    int
)

func LoadConfig() {
	var (
		defaultPort = 8212
		// Get current file full path from runtime
		_, b, _, _ = runtime.Caller(0)

		// Root folder of this project
		ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
	)

	env := os.Getenv("ENVIRONMENT")
	if "" == env {
		env = ""
	}

	err := godotenv.Load(ProjectRootPath + "/" + ".env" + env)

	if err != nil {
		log.Fatalln("Cannot load env config")
		return
	}

	appToken = os.Getenv("APP_TOKEN")
	testChannelId = os.Getenv("TEST_CHANNEL_ID")
	serverPort, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		serverPort = defaultPort
	}
}
