package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mertdogan12/osd-perm/internal/mongo"
	"github.com/mertdogan12/osd/pkg/user"
)

func respond(status int, message string, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(message))
	fmt.Println("users/me:", status, "|", message)
}

func respondErr(err error, w http.ResponseWriter) {
	respond(http.StatusInternalServerError, "Error", w)
	fmt.Fprintln(os.Stderr, err.Error())
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")
	if token[0] != "Bearer" {
		respond(http.StatusUnauthorized, "Authorization token must be Bearer", w)
		return
	}

	user_, err := user.GetUserData(token[1])
	if err != nil {
		if err == user.AuthError {
			respond(http.StatusUnauthorized, "Token in invalid", w)
			return
		}

		respondErr(err, w)
		return
	}

	mongoUser, err := mongo.GetUser(user_.Id)
	if err != nil {
		respondErr(err, w)
		return
	}
	if mongoUser == nil {
		respond(http.StatusNoContent, fmt.Sprintf("User does not exists. Id: %d", user_.Id), w)
		return
	}

	fmt.Println("users/me | Success, id:", user_.Id)
	out, err := json.Marshal(mongoUser)
	if err != nil {
		respondErr(err, w)
		return
	}

	w.Write(out)
}
