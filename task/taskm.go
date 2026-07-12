package task

import(
 "os";
 "fmt";
 "encoding/json"
)

func Read[T any](path string) ([]T, error){
 d,err := os.ReadFile(path)
 if err != nil {
  return nil, fmt.Errorf("Error : couldn't read json file. %w", err)
 }
 var data []T
 if err := json.Unmarshal(d,&data); err != nil {
  return nil, fmt.Errorf("Error : invalid json format. %w", err)
 }
 return data, nil
}

func Write[T any](path string, data []T) error{
  ser, err := json.MarshalIndent(data,""," ")
  if err != nil {
   return fmt.Errorf("Error : couldn't serialization data. %w", err)
  }

  return os.WriteFile(path, data, 0644)
}
