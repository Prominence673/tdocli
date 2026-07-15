package task

import(
 "os";
 "fmt";
 "encoding/json"
)

func read[T any](path string) ([]T, error){
 d,err := os.ReadFile(path)
 if err != nil {
  return nil, fmt.Errorf("Couldn't read json file. %w", err)
 }
 var data []T
 if err := json.Unmarshal(d,&data); err != nil {
  return nil, fmt.Errorf("Invalid json format. %w", err)
 }
 return data, nil
}

func write[T any](path string, data []T) error{
  ser, err := json.MarshalIndent(data,""," ")
  if err != nil {
   return fmt.Errorf("Couldn't serialization data. %w", err)
  }
  return os.WriteFile(path, ser, 0644)
}

func delete(path string) error {
 return os.Remove(path)
}
