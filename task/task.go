package task

import (
 "os";
 "strings"
)

func GetAll() []Task{
 tasks,err := read[Task]("task.json")
 if err != nil {
  if os.IsNotExist(err) {
   os.Create("task.json")
   return []Task{}
  }
  return nil
 }
 return tasks
}

func GetCompleted() []Task{
 task := GetAll()
 var completed []Task
 for _, t := range task {
  if t.Completed {
   completed = append(completed, t)
  }
 }
 return completed
}

func GetPending() []Task{
 task := GetAll()
 var pending []Task
 for _, t := range task {
  if !t.Completed {
   pending = append(pending, t)
  }
 }
 return pending
}

func Add(t string, p int) error{
 path := "task.json"
 task,err := read[Task](path)
 if err != nil {
   if os.IsNotExist(err) || task == nil {
    task = []Task{}
   } else{
    return err
   }
 }
 id := len(task) + 1
 task = append(task, Task{Id: id, Title: t, Completed: false, Priority: p})
 return write(path, task)
}

func RemoveById(id int) error{
 path := "task.json"
 task,err := read[Task](path)
 if err != nil {
   if os.IsNotExist(err) || task == nil {
    return err
   } else{
    return err
   }
 }
 for i, t := range task {
  if t.Id == id {
   task = append(task[:i], task[i+1:]...)
   break
  }
 }
 return write(path, task)
}

func RemoveByTask(taskt string) error{
 path := "task.json"
 task,err := read[Task](path)
 if err != nil {
   if os.IsNotExist(err) || task == nil {
    return err
   } else{
    return err
   }
 }
 taskt = strings.TrimSpace(taskt)
 for i, t := range task {
  if ti := strings.TrimSpace(t.Title); strings.ToLower(ti) == strings.ToLower(taskt) {
   task = append(task[:i], task[i+1:]...)
   break
  }
 }
 return write(path, task)
}

func MarkCompleted(id int) error{
 path := "task.json"
 task,err := read[Task](path)
 if err != nil {
   if os.IsNotExist(err) || task == nil {
    return err
   } else{
    return err
   }
 }
 for i, t := range task {
  if t.Id == id {
   task[i].Completed = true
   break
  }
 }
 return write(path, task)
}

func Undo(id int) error{
 path := "task.json"
 task,err := read[Task](path)
 if err != nil {
   if os.IsNotExist(err) || task == nil {
    return err
   } else{
    return err
   }
 }
 for i, t := range task {
  if t.Id == id {
   task[i].Completed = false
   break
  }
 }
 return write(path, task)
}

func Clear() error{
 path := "task.json"
 if err := delete(path); err != nil {
  return err
 }
 _, err := os.Create(path)
 return err
}