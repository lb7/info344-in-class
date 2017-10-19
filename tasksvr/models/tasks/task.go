package tasks

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Task represents an existing task stored in the database
type Task struct {
	ID         bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Title      string        `json:"title,omitempty"`
	Completed  bool          `json:"completed,omitempty"`
	Tags       []string      `json:"tags,omitempty"`
	CreatedAt  time.Time     `json:"createdAt,omitempty"`
	ModifiedAt *time.Time    `json:"modifiedAt,omitempty"`
}

//NewTask is a new task to be stored in the database.
//Clients may supply only a Title and tags; other fields
//will be set to appropriate defaults
type NewTask struct {
	Title string   `json:"title,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

//Validate validates a NewTask
func (nt *NewTask) Validate() error {
	if len(nt.Title) == 0 {
		return fmt.Errorf("missing task title")
	}
	return nil
}

//ToTask validates and converts a NewTask to a Task
func (nt *NewTask) ToTask() (*Task, error) {
	if err := nt.Validate(); err != nil {
		return nil, err
	}

	now := time.Now()
	task := &Task{
		ID:        bson.NewObjectId(),
		Title:     nt.Title,
		Completed: false,
		Tags:      nt.Tags,
		CreatedAt: now,
	}
	return task, nil
}

//TaskUpdates represents updates to an existing task.
//We currently only allow updating of the Completed field
type TaskUpdates struct {
	Completed bool `json:"completed,omitempty"`
}
