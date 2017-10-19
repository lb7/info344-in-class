package tasks

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type completedFilter struct {
	Completed bool
}

type updateDoc struct {
	Completed  bool
	ModifiedAt time.Time
}

//MongoStore implements Store for MongoDB
type MongoStore struct {
}

//NewMongoStore constructs a new MongoStore
func NewMongoStore(sess *mgo.Session, dbName string, collectionName string) *MongoStore {
	panic("TODO:")
}

//Insert inserts a new task into the store
func (s *MongoStore) Insert(nt *NewTask) (*Task, error) {
	panic("TODO:")
}

//GetAll gets all tasks (up to AllTasksLimit) with a particular `completed` value
func (s *MongoStore) GetAll(completed bool) ([]*Task, error) {
	panic("TODO:")
}

//Update updates the task with `id` based on the values in `tu`
func (s *MongoStore) Update(id bson.ObjectId, tu *TaskUpdates) (*Task, error) {
	panic("TODO:")
}
