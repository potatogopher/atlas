package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"atlas-api/config/schema"
	"atlas-api/db"
	"atlas-api/middleware"
)

// AuthenticatePostData will hold the email and password of the request
// that was sent up by the clientf
type AuthenticatePostData struct {
	Email    string
	Password string
}

// Authenticate - POST
// Will accept an Email and Password. Query the database for Email
// and grab salt and hash from user in the database with the same
// Email. It will then hash the requested password with the existing
// salt and compare the two.
func Authenticate(rw http.ResponseWriter, req *http.Request) {
	var data AuthenticatePostData
	var user schema.User

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}

	err = req.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		middleware.JSONHandler(rw, req)
		rw.WriteHeader(http.StatusBadRequest)
		// TODO: Handle error correctly (JSON)
		return
	}

	if err = db.DB.Where("email = ?", data.Email).Find(&user).Error; err != nil {
		middleware.JSONHandler(rw, req)
		rw.WriteHeader(http.StatusNotFound)
		// TODO: Handle error correctly (JSON)
		return
	}

	err = middleware.Compare(data.Password, user)
	if err != nil {
		middleware.JSONHandler(rw, req)
		rw.WriteHeader(http.StatusNotFound)
		// TODO: Handle error correctly (JSON)
		return
	}

	middleware.JSONHandler(rw, req)
	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
