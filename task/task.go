package task 

type State int 


const (
	Pending State = iota
	Scheudled 
	Running
	Completed 
	Failed
)

import (
	"github.com/google/uuid"
)

type Task Struct {
	ID   uuid.UUID  
	Name string
	State State 
}


import (
	"github.com/google/uuid",
	"github.com/docker/go-connections/nat"
)


type Task struct {
	ID     		   uuid.UUID
	Name   		   string
	State  		   State
	
	Image  		   string
	Memory 		   int
	Disk   		   int 
	ExposedPorts   nat.PortSet
	PortBindings   map[string]string
	RestartPolicy  string
	
	ContainerID    string
	StartTime 	   time.Time 
	FinishTime     time.Time	

}


import (
	"time"
)



type TaskEvent struct {
	ID         uuid.UUID 
	State 	   State 
	Timestamp  time.Time
	Task       Task 
}





	
