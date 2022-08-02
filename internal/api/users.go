package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mertdogan12/go-osuapiv2"
	"github.com/mertdogan12/osd-perm/internal/mongo"
	"github.com/mertdogan12/osd-perm/pkg/helper"
)

func checkToken(authToken string, w http.ResponseWriter) *osuapiv2.User {
	token := strings.Split(authToken, " ")
	if token[0] != "Bearer" {
		helper.ApiRespond(http.StatusUnauthorized, "Authorization token must be Bearer", w)
		return nil
	}

	user, err := osuapiv2.NewWithToken(token[1]).Me("osu")
	if err != nil {
		if err == osuapiv2.ErrorUnauthorized {
			helper.ApiRespond(http.StatusUnauthorized, err.Error(), w)
			return nil
		}

		helper.ApiRespondErr(err, w)
		return nil
	}

	return &user
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	user := checkToken(r.Header.Get("Authorization"), w)
	if user == nil {
		return
	}

	mongoUser, err := mongo.GetUser(user.ID)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	if mongoUser == nil {
		helper.ApiRespond(http.StatusNoContent, fmt.Sprintf("User does not exists. Id: %d", user.ID), w)
		return
	}

	fmt.Println("users/me | Success, id:", user.ID)
	out, err := json.Marshal(mongoUser)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	w.Write(out)
}

func Register(w http.ResponseWriter, r *http.Request) {

}
