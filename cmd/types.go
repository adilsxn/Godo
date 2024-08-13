package cmd

import "time"

type Task struct {
	id int
	name string
	dateOfCreation time.Time
	// dateOfModification time.Time
	isDone bool
}
func (t Task) newTask(name string) Task{
	newTask := Task{
		id: 0,
		name: name,
		dateOfCreation: time.Now(),
		isDone: false,
	}
	return newTask;
}

func (t* Task)completeTask(){
	t.isDone = true;
}
