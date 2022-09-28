package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

const (
	UserBucket = "user"
	PostBucket = "post"
)

func checkErr(err error) {
	log.Panicf("Damn %v", err)
}

func pl(v ...interface{}) {
	log.Println(v...)
}

func main() {

	path := "my.db"
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	pl("Start tx")
	// Start a writable transaction.
	tx, err := db.Begin(true)
	if err != nil {
		checkErr(err)
	}
	defer tx.Rollback()

	// Use the transaction...
	pl("creating user bucket")
	_, err = tx.CreateBucket([]byte(UserBucket))
	if err != nil && err != bolt.ErrBucketExists {
		checkErr(err)
	}

	pl("get user bucket")
	b := tx.Bucket([]byte(UserBucket))

	pl("write to user bucket, name and age")
	err = b.Put([]byte("name"), []byte("fakih"))
	if err != nil {
		checkErr(err)
	}

	err = b.Put([]byte("age"), []byte("17"))
	if err != nil {
		checkErr(err)
	}

	v := b.Get([]byte("name"))
	pl(fmt.Sprintf("The name is: %s and the sequence is %v\n", v, b.Sequence()))

	// Commit the transaction and check for error.
	if err = tx.Commit(); err != nil {
		checkErr(err)
	}

}
