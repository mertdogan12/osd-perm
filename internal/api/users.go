package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mertdogan12/osd-perm/internal/mongo"
	"github.com/mertdogan12/osd/pkg/user"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")
	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authorization token must be Bearer"))
		fmt.Println("users/me | No token")
		return
	}

	user := user.GetUserData(token[1])
	mongoUser := mongo.GetUser(user.Id)
	if mongoUser == nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(fmt.Sprintf("User does not exists. Id: %d", user.Id)))
		fmt.Println("users/me | User doesn't exists:", user.Id)
		return
	}

	fmt.Println("users/me | Success, id:", user.Id)
	out, err := json.Marshal(mongoUser)
	if err != nil {
		panic(err)
	}

	w.Write(out)
}
