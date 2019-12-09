package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Service-Computing-Group/back-end/database"
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
	}

	tmp := database.GetValue([]byte("user"), []byte(name))

	if tmp != "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(os.Stderr, "Error in register!\n")
		w.Write([]byte("Username existed.\n"))
		return
	}

	database.AddValue([]byte("user"), []byte(name), []byte(psw))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
