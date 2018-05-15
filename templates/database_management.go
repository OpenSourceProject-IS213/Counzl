package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
)

const (
	bucket = "./database/file.db"
)

type Something struct {
	Value1 string
	Value2   string
}

func Store_Something(key, value1, value2 string) {
	something := &Something{value1, value2}

	s := string(something.serialize())
	db, err := buntdb.Open(bucket)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, s, nil)
		return err
	})
	fmt.Println("New item saved to db with key: '" + key + "'")
	return true
}

func FetchSomething(key string) (string, string) {
	var value1 string
	var value2 string
	db, err := buntdb.Open(bucket) // First you have to open the file
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // You always have to close the db-file after you are done writing/reading to/from it

	err = db.View(func(tx *buntdb.Tx) error { // View only has read-access
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		tmp := deserialize([]byte(val)) 
		value1 = tmp.Value1
		value2 = tmp.Value2
		return nil
	})
	return value1, value2
}

func PrintAllObjectsFromBucket() { // Iterates the whole db-file
	db, err := buntdb.Open(bucket)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	erro := db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			if key != "" {
				val := deserialize([]byte(value))
				fmt.Printf("Key: %s, value1: %s, value2: %s\n", key, val.Value1, val.Value2)
				return true // Return true if key has a value (!= "")
			} else {
				return false // Return false when there is no more objects in the DB
			}
		})

		return err // Return error when View is done
	})
	fmt.Print(erro)
}

func (s *Something) serialize() []byte { // Serializes Something: writes value1 and value2 to a byte slice ([]byte) so it is possible to save multiple values in the DB
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(u)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserialize(d []byte) *Something { // Returns a pointer to one Something variable
	var s Something

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&s)
	if err != nil {
		log.Panic(err)
	}
	return &s
}
