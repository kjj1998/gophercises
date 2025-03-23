package database

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var pathsToUrlsMap = map[string]string{
	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}
var BUCKET_NAME = "PathsToUrlsBucket"

func OpenDB() (*bolt.DB, error) {
	db, err := bolt.Open("paths.db", 0600, nil)

	return db, err
}

func SearchURL(db *bolt.DB, path string) string {
	var url string

	db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket([]byte(BUCKET_NAME)).Get([]byte(path))
		if value == nil {
			fmt.Printf("URL for path: %s cannot be found.\n", path)
			url = ""
		} else {
			url = string(value)
		}

		return nil
	})

	return url
}

func CreateBucket(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME))

		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}

		fmt.Println("Populating bucket ...")
		for k, v := range pathsToUrlsMap {
			addKeyValErr := bucket.Put([]byte(k), []byte(v))
			if addKeyValErr != nil {
				return fmt.Errorf("add key-val: %s", addKeyValErr)
			}
		}

		return nil
	})
}
