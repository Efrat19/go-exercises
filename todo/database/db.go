package database

import (
	"bytes"
	"encoding/gob"
	"github.com/Efrat19/gophercises/todo/types"
	bolt "go.etcd.io/bbolt"
	"time"
)

var db *bolt.DB

func init() {
	createDBIfNotExist()

}

func ListTasks() ([]types.Task, error) {
	tasks := []types.Task{}
	err := db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("tasks")).ForEach(func(k, v []byte) error {
			var task types.Task
			dec := gob.NewDecoder(bytes.NewBuffer(v))
			dec.Decode(&task)
			tasks = append(tasks, task)
			return nil
		})
	})
	return tasks, err
}

func AddTask(name string) (types.Task, error) {
	id, err := getID()
	if err != nil {
		panic(err)
	}
	task := types.Task{id, name, time.Now()}
	return writeTask(task)
}

func CheckTask(id int) (int, error) {
	defer closeDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return err
		}
		var idBytes bytes.Buffer
		idenc := gob.NewEncoder(&idBytes)
		idenc.Encode(id)
		return b.Delete([]byte(idBytes.Bytes()))
	})
	return id, err
}

func createDBIfNotExist() {
	var err error
	db, err = bolt.Open("/Users/efrat/.boltDB", 0666, nil)
	if err != nil {
		panic(err)
	}
}

func getID() (int, error) {
	max := 0
	err := db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("tasks")).ForEach(func(k, v []byte) error {
			var id int
			dec := gob.NewDecoder(bytes.NewBuffer(k))
			dec.Decode(&id)
			if id > max {
				max = id
			}
			return nil
		})
	})
	return max + 1, err
}

func writeTask(task types.Task) (types.Task, error) {
	defer closeDB()
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return err
		}
		var idBytes bytes.Buffer
		idenc := gob.NewEncoder(&idBytes)
		idenc.Encode(task.Id)
		var taskBytes bytes.Buffer
		taskenc := gob.NewEncoder(&taskBytes)
		taskenc.Encode(task)
		return b.Put([]byte(idBytes.Bytes()), []byte(taskBytes.Bytes()))
	})
	return task, err
}

func closeDB() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
