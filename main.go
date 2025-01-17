package main


import (
	"cube/node"
	"cube/task"
	"fmt"
	"time"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"cube/manager"
	"cube/worker"
)


func main(){
	
	t := task.Task{
		ID : uuid.New(),
		Name : "Task-1",
		State : task.Pending,
		Image : "Image-1",
		Memory : 1024,
		Disk : 1,
	} 

	te := task.TaskEvent {
		ID: uuid.New(),
		State : task.Pending,
		Timestamp : time.Now(),
		Task : t,
	}


	fmt.Printf("task: %v \n ", t)
	fmt.Printf("task event : %v\n", te)
	

}
