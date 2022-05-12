package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mertdogan12/osd-perm/internal/mongo"
	"github.com/mertdogan12/osd-perm/pkg/helper"
	"github.com/mertdogan12/osd/pkg/user"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")
	if token[0] != "Bearer" {
		helper.ApiRespond(http.StatusUnauthorized, "Authorization token must be Bearer", w)
		return
	}

	user_, err := user.GetUserData(token[1])
	if err != nil {
		if err == user.AuthError {
			helper.ApiRespond(http.StatusUnauthorized, "Token in invalid", w)
			return
		}

		helper.ApiRespondErr(err, w)
		return
	}

	mongoUser, err := mongo.GetUser(user_.Id)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}
	if mongoUser == nil {
		helper.ApiRespond(http.StatusNoContent, fmt.Sprintf("User does not exists. Id: %d", user_.Id), w)
		return
	}

	fmt.Println("users/me | Success, id:", user_.Id)
	out, err := json.Marshal(mongoUser)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	w.Write(out)
}
