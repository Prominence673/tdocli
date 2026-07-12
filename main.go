package main

import (
 "fmt";
 //"sync"
 "tdocli/task"
 //"tdocli/cmd"
)

func main(){
 t := task.GetAll()
 fmt.Println("json ", t[0].Task)
}
