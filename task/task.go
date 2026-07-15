package task

import (
 "os";
 "strings"
 "fmt"
 "path/filepath"
)

var filename = "task.json"

func SetFilename(path string) error {
	path = strings.TrimSpace(path)
	if path == ""{
		return fmt.Errorf("path cannot be empty")
	}

	if filepath.Ext(path) == "" {
		filename = path + ".json"
	} else {
		filename = path
	}
	return nil
}

func getTasks() ([]Task, error) {
	return read[Task](filename)
}
func saveTask(task []Task) error {
	return write(filename, task)
}
func fixid(task []Task) []Task {
	if len(task) == 0 {
		return task
	}
	for i := range task {
		task[i].Id = i + 1
	}
	return task
}

func GetAll() []Task{
 tasks,err := getTasks()
 if err != nil {
  if os.IsNotExist(err) {
   if err := saveTask([]Task{}); err == nil {
    return []Task{}
   } else {
    return []Task{}
   }
  }
  return []Task{}
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
 task, err := getTasks()
 if err != nil {
   if os.IsNotExist(err) || task == nil{
    task = []Task{}
    if err := saveTask(task); err != nil {
       return err
    }
   } else{
    return err
   }
 }
 maxId := 0
 for _, t := range task {
  if t.Id > maxId {
   maxId = t.Id
  }
 }
 id := maxId + 1
 task = append(task, Task{Id: id, Title: t, Completed: false, Priority: p})
 return saveTask(task)
}

func Export(path string) error{
 back := filename
 var err error
 if _, err := os.Stat(back); err != nil {
  if os.IsNotExist(err) {
   return err
  } 
 }
 path = strings.TrimSpace(path)
 if path == "" {
  return fmt.Errorf("Path is required")
 }
 if filepath.Ext(path) == "" {
 	filename = path + ".json"
 } else {
 	filename = path
 }
 err = os.Rename(back, path)
 if err != nil {
  return err
 }
 if err := saveTask([]Task{}); err != nil {
	return err
 }
 return nil
}

func Edit(title string, id int) error{
 task, err := getTasks()
 if err != nil {
     return err
 }
 found := false
 for i, t := range task {
  if t.Id == id {
   task[i].Title = title
   found = true
   break
  }
 }
 if !found {
  return fmt.Errorf("Task not found")
 }
 return saveTask(task)
}

func RemoveById(id int) error{
 task, err := getTasks()
 if err != nil {
     return err
 }
 found := false
 for i, t := range task {
  if t.Id == id {
   task = append(task[:i], task[i+1:]...)
   found = true
   break
  }
 }
 if !found {
  return fmt.Errorf("Task not found")
 }
 task = fixid(task)
 return saveTask(task)
}

func RemoveByTask(taskt string) error{
 task, err := getTasks()
 if err != nil {
     return err
 }
 taskt = strings.TrimSpace(taskt)
 found := false
 for i, t := range task {
  if ti := strings.TrimSpace(t.Title); strings.ToLower(ti) == strings.ToLower(taskt) {
   task = append(task[:i], task[i+1:]...)
   found = true
   break
  }
 }
 if !found {
  return fmt.Errorf("Task not found")
 }
 task = fixid(task)
 return saveTask(task)
}

func MarkCompleted(id int) error{
 task, err := getTasks()
 if err != nil {
     return err
 }
 found := false
 for i, t := range task {
  if t.Id == id {
   task[i].Completed = true
   found = true
   break
  }
 }
 if !found {
  return fmt.Errorf("Task not found")
 }
 return saveTask(task)
}

func Undo(id int) error{
 task, err := getTasks()
 if err != nil {
     return err
 }
 found := false
 for i, t := range task {
  if t.Id == id {
   task[i].Completed = false
   found = true
   break
  }
 }
 if !found {
  return fmt.Errorf("Task not found")
 }
 return saveTask(task)
}

func Clear() error {
	return saveTask([]Task{})
}