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

type UserReq struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// CreateUser will accept a POST request and add a user to the database
func CreateUser(rw http.ResponseWriter, req *http.Request) {
	var user schema.User
	var userReq UserReq
	var credentials middleware.Credentials

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	if err := req.Body.Close(); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &userReq); err != nil {
		middleware.JSONHandler(rw, req)
		rw.WriteHeader(422)
	}

	credentials, err = middleware.CreateCredentials(userReq.Password)
	if err != nil {
		log.Fatal(err)
	}

	user.FirstName = userReq.FirstName
	user.LastName = userReq.LastName
	user.Email = userReq.Email
	user.PasswordHash = credentials.Hash
	user.PasswordSalt = credentials.Salt

	db.DB.Create(&user)

	middleware.JSONHandler(rw, req)
	rw.WriteHeader(200)

}