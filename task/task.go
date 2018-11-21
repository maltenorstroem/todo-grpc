package task

import (
	"../protos"
	"upper.io/db.v3"
)

// Interface zur DB
var dbCollectionTask db.Collection
var paginationDefault uint

func ConnectDatabase(database db.Database) {
	dbCollectionTask = database.Collection("tasks")
	paginationDefault = 23
}
func createTaskItem(data *task.Task) (task.Task, error) {
	var item task.Task
	//todo ulid typ in protobuf bauen

	item.Id = GenerateULID().String()
	item.Title = data.Title
	item.Description = data.Description
	if data.Completed != 0 {
		item.Completed = data.Completed
	} else {
		item.Completed = 1
	}
	// id interface not needed, we create the ids ourself
	_, err := dbCollectionTask.Insert(&item)
	//fire("item.generated",&item)
	return item, err
}

func listTaskItems(options QueryOptions) ([]task.Task, DBMeta, error) {
	res := dbCollectionTask.Find()
	var meta DBMeta
	res, meta = ApplyRequestOptionsToQuery(res, options)
	var items []task.Task
	err := res.All(&items)

	return items, meta, err
}

func completeTaskItem(id string) (task.Task, error) {
	var item task.Task
	item.Completed = 2
	return updateTaskItem(id, &item)
}

func deleteTaskItem(id string) error {
	var item task.Task
	res := dbCollectionTask.Find(db.Cond{"id": id})
	if err := res.One(&item); err != nil {
		return err
	}
	err := res.Delete()
	return err
}
func getTaskItem(id string) (task.Task, error) {
	var item task.Task
	res := dbCollectionTask.Find(db.Cond{"id": id})
	err := res.One(&item)
	return item, err
}

func updateTaskItem(id string, data *task.Task) (task.Task, error) {
	var item task.Task
	res := dbCollectionTask.Find(db.Cond{"id": id})
	// fields to update
	item.Title = data.Title
	item.Description = data.Description
	item.Completed = data.Completed

	if err := res.Update(&item); err != nil {
		return task.Task{}, err
	}
	// read your write
	err := res.One(&item)
	return item, err
}
