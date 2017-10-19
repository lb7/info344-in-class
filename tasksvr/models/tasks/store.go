package tasks

import (
	"gopkg.in/mgo.v2/bson"
)

//AllTasksLimit is the maximum number of tasks
//a store should fetch, so that we don't explode
//when there are millions of tasks in the DBMS.
//We can add result set paging later.
const AllTasksLimit = 1000

//Store represents a data store for Tasks
type Store interface {
	Insert(nt *NewTask) (*Task, error)
	GetAll(completed bool) ([]*Task, error)
	Update(id bson.ObjectId, tu *TaskUpdates) (*Task, error)
}
