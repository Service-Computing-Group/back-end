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
func peopleHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		vals := req.URL.Query()
		page := 1

		itemCount := database.GetCount("people")

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
			item := database.GetValue([]byte("people"), []byte(strconv.Itoa(i)))
			if len(item) != 0 {
				count++
				if count > 10*(page-1) {
					w.Write([]byte(item))
					if count >= pagelen*page || count >= database.GetCount("people") {
						break
					}
					w.Write([]byte(", \n"))
				}
			}
		}
		w.Write([]byte("]\n}"))
	}
}

func getPeopleById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	data := database.GetValue([]byte("people"), []byte(vars["id"]))
	w.Write([]byte(data))
}

func peoplePagesHandler(w http.ResponseWriter, req *http.Request) {
	data := database.GetCount("people")
	w.Write([]byte(strconv.Itoa(data)))
}
