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
			tx.CreateBucketIfNotExists([]byte("Planets"))
			tx.CreateBucketIfNotExists([]byte("Species"))
			tx.CreateBucketIfNotExists([]byte("Vehicle"))
			tx.CreateBucketIfNotExists([]byte("Starship"))
			tx.CreateBucketIfNotExists([]byte("People"))
			tx.CreateBucketIfNotExists([]byte("Film"))
			return nil
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func catchPeople() {
	count := 0
	for i := 90; count <= 87; i-- {
		people, _ := swapi.GetPerson(i)
		jsonStr, _ := json.Marshal(people)
		indexStr := strconv.Itoa(i)
		if len(GetValue([]byte("People"), []byte(indexStr))) == 0 {
			if AddValue([]byte("People"), []byte(indexStr), jsonStr) {
				count++
			}
		} else {
		}
	}
}

func catchFilm() {

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
	//catchPeople()
	//fmt.Print(GetValue([]byte("People"), []byte(strconv.Itoa(1))))
	fmt.Println(getCount("People"))
	db.Close()
}
