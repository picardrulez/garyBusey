package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func dbPost(key string, value string, bucket string) int {
	byteBucket := []byte(bucket)
	db, err := bolt.Open("garybusey.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer db.Close()
	byteKey := []byte(key)
	byteValue := []byte(value)

	err = db.Update(func(tx *bolt.Tx) error {
		mybucket, err := tx.CreateBucketIfNotExists(byteBucket)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = mybucket.Put(byteKey, byteValue)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func dbGet(key string, bucket string) string {
	byteBucket := []byte(bucket)
	db, err := bolt.Open("garybusey.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var value string
	byteKey := []byte(key)

	err = db.View(func(tx *bolt.Tx) error {
		mybucket := tx.Bucket(byteBucket)
		if mybucket == nil {
			return fmt.Errorf("Bucket %q not found!", byteKey)
		}

		val := mybucket.Get(byteKey)
		value = string(val)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func dbRemove(bucket string, key string) int {
	byteBucket := []byte(bucket)

	db, err := bolt.Open("garybusey.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	byteKey := []byte(key)

	err = db.Update(func(tx *bolt.Tx) error {
		mybucket, err := tx.CreateBucketIfNotExists(byteBucket)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = mybucket.Delete(byteKey)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func listBucket(bucket string) map[string]string {
	db, err := bolt.Open("garybusey.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	m := make(map[string]string)
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			m[string(k)] = string(v)
		}
		return nil
	})
	return m
}
