package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/leejarvis/swapi"
)

var db *bolt.DB

func OpenDB(str string) {
	existed := false
	var err error
	if _, err := os.Open(str); err == nil {
		existed = true
	}
	db, err = bolt.Open(str, 0666, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !existed {
		err = db.Update(func(tx *bolt.Tx) error {
			tx.CreateBucketIfNotExists([]byte("planets"))
			tx.CreateBucketIfNotExists([]byte("species"))
			tx.CreateBucketIfNotExists([]byte("vehicles"))
			tx.CreateBucketIfNotExists([]byte("starships"))
			tx.CreateBucketIfNotExists([]byte("people"))
			tx.CreateBucketIfNotExists([]byte("films"))
			return nil
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

//1.get planets
func catchPlanets() {
	count := 0
	for i := 70; count <= 61; i-- {
		planet, _ := swapi.GetPlanet(i)
		jsonStr, _ := json.Marshal(planet)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("planets"), []byte(indexStr))) == 0 {
			if AddValue([]byte("planets"), []byte(indexStr), jsonStr) {
				count++
			}
		} else {

		}
	}
}

//2.get species
func catchSpecies() {
	for i := 1; i <= 37; i++ {
		specie, _ := swapi.GetSpecies(i)
		jsonStr, _ := json.Marshal(specie)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("species"), []byte(indexStr))) == 0 {
			if AddValue([]byte("species"), []byte(indexStr), jsonStr) {
			}
		} else {

		}
	}
}

//3.get vehicles
func catchVehicles() {
	for i := 1; i <= 39; i++ {
		vehicle, _ := swapi.GetVehicle(i)
		jsonStr, _ := json.Marshal(vehicle)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("vehicles"), []byte(indexStr))) == 0 {
			if AddValue([]byte("vehicles"), []byte(indexStr), jsonStr) {
			}
		} else {

		}
	}
}

//4.get starships
func catchStarships() {
	count := 0
	for i := 50; count <= 37; i-- {
		starship, _ := swapi.GetStarship(i)
		jsonStr, _ := json.Marshal(starship)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("starships"), []byte(indexStr))) == 0 {
			if AddValue([]byte("starships"), []byte(indexStr), jsonStr) {
				count++
			}
		} else {

		}
	}
}

//5.get people
func catchPeople() {
	count := 0
	for i := 90; count <= 87; i-- {
		person, _ := swapi.GetPerson(i)
		jsonStr, _ := json.Marshal(person)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("people"), []byte(indexStr))) == 0 {
			if AddValue([]byte("people"), []byte(indexStr), jsonStr) {
				count++
			}
		} else {

		}
	}
}

func catchFilms() {
	//电影序号是连续的
	for i := 1; i <= 7; i++ {
		film, _ := swapi.GetFilm(i)
		jsonStr, _ := json.Marshal(film)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("films"), []byte(indexStr))) == 0 {
			if AddValue([]byte("films"), []byte(indexStr), jsonStr) {
			}
		} else {

		}
	}
}

func GetValue(bucketName []byte, key []byte) string {
	var value []byte
	db.View(func(tx *bolt.Tx) error {
		byteLen := len(tx.Bucket([]byte(bucketName)).Get(key))
		value = make([]byte, byteLen)
		copy(value[:], tx.Bucket([]byte(bucketName)).Get(key)[:])
		return nil
	})
	return (string)(value)
}

func AddValue(bucketName []byte, key []byte, value []byte) bool {
	if err := db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket(bucketName).Put(key, value); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func getCount(str string) int {
	count := 0
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(str))
		b.ForEach(func(k, v []byte) error {
			count++
			return nil
		})
		return nil
	})
	return count
}

func main() {
	OpenDB("./test.db")
	catchPeople()
	fmt.Println(getCount("people"))
	// catchSpecies()
	// fmt.Println(getCount("species"))
	//catchPlanets()
	//fmt.Println(getCount("planets"))
	//fmt.Println(GetValue([]byte("planets"), []byte(strconv.Itoa(1))))

	//catchFilms()
	//fmt.Println(getCount("films"))
	db.Close()
}
