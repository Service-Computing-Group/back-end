package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Service-Computing-Group/back-end/database"
	"github.com/gorilla/mux"
)

// Handler Function of "/api/people".
func peopleHandler(w http.ResponseWriter, req *http.Request) {
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

func getPeopleById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	data := database.GetValue([]byte("people"), []byte(vars["id"]))
	//判断是否存在此数据，否则404
	if data == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found!"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(data))
	}
}
