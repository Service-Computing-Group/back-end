package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/leejarvis/swapi"
)

var db *bolt.DB

func CloseDB() {
	db.Close()
}

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

			tx.CreateBucketIfNotExists([]byte("users"))
			return nil
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func LoadData() {
	OpenDB("./database/test2.db")
	catchStarships()
	fmt.Println("starships = ", GetCount("starships"))
	time.Sleep(200 * time.Millisecond)
	catchVehicles()
	fmt.Println("vehicles = ", GetCount("vehicles"))
	time.Sleep(200 * time.Millisecond)
	catchPlanets()
	fmt.Println("planets = ", GetCount("planets"))
	time.Sleep(200 * time.Millisecond)
	catchSpecies()
	fmt.Println("species = ", GetCount("species"))
	time.Sleep(100 * time.Millisecond)
	catchFilms()
	fmt.Println("films = ", GetCount("films"))
	time.Sleep(200 * time.Millisecond)
	catchPeople()
	fmt.Println("people = ", GetCount("people"))
	db.Close()
}

func Test() {
	people, _ := swapi.GetPerson(1)
	jsonStr, _ := json.MarshalIndent(people, "", "\t")
	fmt.Println(jsonStr)
	var p swapi.Person
	json.Unmarshal(jsonStr, &p)
	fmt.Println(p)
}

//1.get planets 1~61 共 61
func catchPlanets() {
	//是连续的
	for i := 1; i <= 61; i++ {
		planet, _ := swapi.GetPlanet(i)
		// if planet.Name == "" {
		// 	continue
		// }
		jsonStr, _ := json.MarshalIndent(planet, "", "    ")
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
		jsonStr, _ := json.MarshalIndent(specie, "", "    ")
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
		jsonStr, _ := json.MarshalIndent(vehicle, "", "    ")
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
		jsonStr, _ := json.MarshalIndent(starship, "", "    ")
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
		if i%20 == 0 {
			time.Sleep(200 * time.Millisecond)
		}
		person, _ := swapi.GetPerson(i)
		if person.Name == "" {
			continue
		}
		jsonStr, _ := json.MarshalIndent(person, "", "    ")
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
		jsonStr, _ := json.MarshalIndent(film, "", "    ")
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

func GetCount(str string) int {
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

//func main() {
//	OpenDB("./test.db")
//	fmt.Println("starships = ", GetCount("starships"))
//	fmt.Println("vehicles = ", GetCount("vehicles"))
//	fmt.Println("planets = ", GetCount("planets"))
//	fmt.Println("species = ", GetCount("species"))
//	fmt.Println("films = ", GetCount("films"))
//	fmt.Println("people = ", GetCount("people"))
//
//	//已完成的
//	// catchStarships()
//	// fmt.Println("starships = ", GetCount("starships"))
//	//catchVehicles()
//	//fmt.Println("vehicles = ", GetCount("vehicles"))
//	//catchPlanets()
//	//fmt.Println("planets = ", GetCount("planets"))
//	//catchSpecies()
//	//fmt.Println("species = ", GetCount("species"))
//	//catchFilms()
//	//fmt.Println("films = ", GetCount("films"))
//	//catchPeople()
//	//fmt.Println("people = ", GetCount("people"))
//	db.Close()
//}
