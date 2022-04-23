package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mertdogan12/osd-perm/internal/conf"
	"github.com/mertdogan12/osd-perm/internal/mongo"
)

func main() {
	// .env
	godotenv.Load()

	conf.Parse(os.Args[1:])

	mongo.Connect()
	defer mongo.Disconnect()

	fmt.Println("Server started on port:", conf.Port)
	http.ListenAndServe(":"+fmt.Sprint(conf.Port), nil)
}
