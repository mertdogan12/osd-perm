package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mertdogan12/osd-perm/internal/mongo"
)

func main() {
	// .env
	godotenv.Load()

	mongo.Connect()
	defer mongo.Disconnect()

	id, err := strconv.Atoi(os.Getenv("OSD_ID"))
	if err != nil {
		panic(err)
	}

	user := mongo.GetUser(id)
	fmt.Println(*user)
}
