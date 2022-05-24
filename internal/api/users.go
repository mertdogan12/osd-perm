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

func GetMe(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")
	if token[0] != "Bearer" {
		helper.ApiRespond(http.StatusUnauthorized, "Authorization token must be Bearer", w)
		return
	}

	user_, err := osuapiv2.NewWithToken(token[1]).Me("osu")
	if err != nil {
		if err == osuapiv2.ErrorUnauthorized {
			helper.ApiRespond(http.StatusUnauthorized, err.Error(), w)
			return
		}

		helper.ApiRespondErr(err, w)
		return
	}

	mongoUser, err := mongo.GetUser(user_.ID)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}
	if mongoUser == nil {
		helper.ApiRespond(http.StatusNoContent, fmt.Sprintf("User does not exists. Id: %d", user_.ID), w)
		return
	}

	fmt.Println("users/me | Success, id:", user_.ID)
	out, err := json.Marshal(mongoUser)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	w.Write(out)
}
