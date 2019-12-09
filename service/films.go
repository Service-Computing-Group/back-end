package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/Service-Computing-Group/back-end/database"
)

//handle a request with method GET and path "/api/".
func filmsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		vals := req.URL.Query()
		page := 1

		itemCount := database.GetCount("films")

		if vals["page"] != nil {
			var err error
			page, err = strconv.Atoi(vals["page"][0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", vals)
			}
		}
		if page == 0 || page >= (itemCount+pagelen-1)/pagelen+1 {
			fmt.Println((itemCount+pagelen-1)/pagelen + 1)
			w.Write([]byte("404 Not Found!"))
			return
		}
		w.Write([]byte("{\n    \"count\" : "))
		w.Write([]byte(strconv.Itoa(itemCount)))
		w.Write([]byte(",\n    \"result\" : [\n"))

		count := 0
		for i := 1; count < pagelen*page; i++ {
			item := database.GetValue([]byte("films"), []byte(strconv.Itoa(i)))
			if len(item) != 0 {
				count++
				if count > 10*(page-1) {
					w.Write([]byte(item))
					if count >= pagelen*page || count >= database.GetCount("films") {
						break
					}
					w.Write([]byte(", \n"))
				}
			}
		}
		w.Write([]byte("]\n}"))
	}
}

func getFilmById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	data := database.GetValue([]byte("films"), []byte(vars["id"]))
	//判断是否存在此数据，否则404
	if data == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found!"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	}
}

func filmsPagesHandler(w http.ResponseWriter, req *http.Request) {
	data := database.GetCount("films")
	w.Write([]byte(strconv.Itoa(data)))
}
