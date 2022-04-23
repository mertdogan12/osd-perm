package api

import (
	"strings"

	"github.com/mertdogan12/osd-perm/internal/mongo"
	"github.com/mertdogan12/osd-perm/pkg/helper"
)

func hasPermission(permission string, user mongo.User) bool {
	if helper.StringArrayConatins(user.Permissions, "*") {
		return true
	}

	perm := strings.Split(permission, ".")
	if helper.StringArrayConatins(user.Permissions, perm[0]+".*") || helper.StringArrayConatins(user.Permissions, permission) {
		return true
	}

	return false
}
