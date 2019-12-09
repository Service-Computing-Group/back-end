package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Service-Computing-Group/back-end/database"
	"github.com/Service-Computing-Group/back-end/token"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	User     string
	Password string
}

func registerHandler(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("username")
	psw := req.FormValue("password")
	if name == "" || psw == "" {
		fmt.Fprintf(os.Stderr, "Error in register!\n")
		return
	}

	tmp := database.GetValue([]byte("users"), []byte(name))

	if tmp != "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(os.Stderr, "Error in register!\n")
		w.Write([]byte("Username existed.\n"))
		return
	}

	database.AddValue([]byte("users"), []byte(name), []byte(psw))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("username")
	psw := req.FormValue("password")
	if name == "" || psw == "" {
		fmt.Fprintf(os.Stderr, "Error in login!\n")
		return
	}

	tmp := database.GetValue([]byte("users"), []byte(name))

	if tmp == "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(os.Stderr, "Error in register!\n")
		w.Write([]byte("Username does not exist.\n"))
		return
	}
	if tmp != psw {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(os.Stderr, "Error in register!\n")
		w.Write([]byte("Password is not correct.\n"))
		return
	}

	// copy
	mytoken := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	mytoken.Claims = claims
	tokenString, err := mytoken.SignedString([]byte(token.Secretkey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(tokenString))
}
