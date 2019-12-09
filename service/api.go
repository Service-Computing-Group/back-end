package service

import (
	"net/http"

	"github.com/unrolled/render"
)

//handle a request with method GET and path "/api/".
func apiHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Films     string `json:"films"`
			People    string `json:"people"`
			Planets   string `json:"planets"`
			Species   string `json:"species"`
			Starships string `json:"starships"`
			Vehicles  string `json:"vehicles"`
		}{
			Films:     "https://localhost:8081/api/films/",
			People:    "https://localhost:8081/api/people/",
			Planets:   "https://localhost:8081/api/planets/",
			Species:   "https://localhost:8081/api/species/",
			Starships: "https://localhost:8081/api/starships/",
			Vehicles:  "https://localhost:8081/api/vehicles/"})
	}
}
