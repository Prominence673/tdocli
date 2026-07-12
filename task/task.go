package task

import (
 //"fmt";
 //"os";
 //"encoding/json";	
)

func GetAll() []Task{
 tasks,err := Read[Task]("task.json")
 if err != nil {
  return nil
 }
 return tasks
}

func Add(t string) error{
 return nil
}

func Save() error{
 return nil
}

func Remove(t string) error{
 return nil
}

func MarkCompleted(t string) error{
 return nil
}
