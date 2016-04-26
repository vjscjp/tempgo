package controllers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/CiscoCloud/shipped-utils/consts"
	f "github.com/MustWin/gomarathon"
)

var (
	url  string
	user string
	pwd  string
)

func SetEnv() {
	url = os.Getenv(consts.MARATHON_URL)
	if len(url) < 1 {
		url = consts.M_url
		os.Setenv(consts.MARATHON_URL, consts.M_url)
	}

	user = os.Getenv(consts.MARATHON_USER)
	if len(user) < 1 {
		user = consts.M_user
		os.Setenv(consts.MARATHON_USER, consts.M_user)
	}

	pwd = os.Getenv(consts.MARATHON_PWD)
	if len(pwd) < 1 {
		pwd = consts.M_pwd
		os.Setenv(consts.MARATHON_PWD, consts.M_pwd)
	}

}
func NewMarathonClient() (client *f.Client, err error) {

	auth := f.HttpBasicAuth{user, pwd}
	client, err = f.NewClient(url, &auth, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("Invalid Details : " + err.Error())
	}
	return
}

func ServeJsonResponseWithCode(w http.ResponseWriter, responseBodyObj interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(responseBodyObj); err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
