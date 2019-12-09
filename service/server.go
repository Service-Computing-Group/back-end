package service

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/leejarvis/swapi"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	//database.LoadData()

	mx.HandleFunc("/api/people", peopleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/people/", peopleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/people/{id}", getPeopleById).Methods("GET")
	mx.HandleFunc("/api/test", testHandler).Methods("GET")
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	people, _ := swapi.GetPerson(1)
	jsonStr, _ := json.MarshalIndent(people, "", "    ")
	//jsonStr, _ := json.Marshal(people)
	w.Write([]byte(jsonStr))
}
