package helper

import (
	"fmt"
	"net/http"
	"os"
)

func ApiRespond(status int, message string, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(message))
	fmt.Println("users/me:", status, "|", message)
}

func ApiRespondErr(err error, w http.ResponseWriter) {
	ApiRespond(http.StatusInternalServerError, "Error", w)
	fmt.Fprintln(os.Stderr, err.Error())
}
