package service

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/leejarvis/swapi"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

const (
	pagelen = 10
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
	mx.HandleFunc("/api", apiHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/", apiHandler(formatter)).Methods("GET")

	mx.HandleFunc("/api/people", peopleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/people/", peopleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/people/{id}", getPeopleById).Methods("GET")
	mx.HandleFunc("/api/test", testHandler).Methods("GET")

	mx.HandleFunc("/api/films", filmsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/films/", filmsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/films/{id}", getFilmById).Methods("GET")

	mx.HandleFunc("/api/planets", planetsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/planets/", planetsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/planets/{id}", getPlanetById).Methods("GET")

	mx.HandleFunc("/api/starships", starshipsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/starships/", starshipsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/starships/{id}", getStarshipById).Methods("GET")

	mx.HandleFunc("/api/species", speciesHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/species/", speciesHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/species/{id}", getSpecieById).Methods("GET")

	mx.HandleFunc("/api/vehicles", vehiclesHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/vehicles/", vehiclesHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/vehicles/{id}", getVehicleById).Methods("GET")
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	people, _ := swapi.GetPerson(1)
	jsonStr, _ := json.MarshalIndent(people, "", "    ")
	//jsonStr, _ := json.Marshal(people)
	w.Write([]byte(jsonStr))
}
