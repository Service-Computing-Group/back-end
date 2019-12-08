package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/leejarvis/swapi"
	"log"
	"os"
	"strconv"
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

//1.get planets 1~60 共 60
func catchPlanets() {
	//是连续的
	for i := 1; i <= 60; i++ {
		planet, _ := swapi.GetPlanet(i)
		if planet.Name == "" {
			continue
		}
		jsonStr, _ := json.Marshal(planet)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("planets"), []byte(indexStr))) == 0 {
			if AddValue([]byte("planets"), []byte(indexStr), jsonStr) {

			}
		} else {

		}
	}
}

//2.get species 1~37 共37
func catchSpecies() {
	//是连续的
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

//3.get vehicles 1~80 共 39
func catchVehicles() {
	for i := 1; i <= 80; i++ {
		vehicle, _ := swapi.GetVehicle(i)
		if vehicle.Name == "" {
			continue
		}
		jsonStr, _ := json.Marshal(vehicle)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("vehicles"), []byte(indexStr))) == 0 {
			if AddValue([]byte("vehicles"), []byte(indexStr), jsonStr) {
			}
		} else {

		}
	}
}

//4.get starships 1~80 共 37
func catchStarships() {
	for i := 1; i <= 80; i++ {
		starship, _ := swapi.GetStarship(i)
		if starship.Name == "" {
			continue
		}
		jsonStr, _ := json.Marshal(starship)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("starships"), []byte(indexStr))) == 0 {
			if AddValue([]byte("starships"), []byte(indexStr), jsonStr) {

			}
		} else {

		}
	}
}

//5.get people	1~90 共 87
func catchPeople() {
	count := 0
	for i := 90; count <= 87; i-- {
		person, _ := swapi.GetPerson(i)
		if person.Name == "" {
			continue
		}
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

//6.get films 1~7 共 7
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

	fmt.Println("starships = ", getCount("starships"))
	fmt.Println("vehicles = ", getCount("vehicles"))	
	fmt.Println("planets = ", getCount("planets"))
	fmt.Println("species = ", getCount("species"))
	fmt.Println("films = ", getCount("films"))
	fmt.Println("people = ", getCount("people"))

	//已完成的		
	// catchStarships()
	// fmt.Println("starships = ", getCount("starships"))		
	//catchVehicles()
	//fmt.Println("vehicles = ", getCount("vehicles"))
	//catchPlanets()
	//fmt.Println("planets = ", getCount("planets"))
	//catchSpecies()
	//fmt.Println("species = ", getCount("species"))
	//catchFilms()
	//fmt.Println("films = ", getCount("films"))
	//catchPeople()
	//fmt.Println("people = ", getCount("people"))
	db.Close()
}
