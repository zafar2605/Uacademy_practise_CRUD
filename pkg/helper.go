package pkg

import (
	"encoding/json"
	"os"
)

func Read(Path string) ([]interface{}, error) {

	body, err := os.ReadFile(Path)
	if err!=nil{
		return nil, err	
	}
	
	var data []interface{}
	err = json.Unmarshal(body, &data)
	if err!=nil{
		return nil, err	
	}

	return data, nil
}

func Write(Path string, data []interface{}) error {

	body, err := json.MarshalIndent(data, "", "    ")
	if err!=nil{
		return err
	}

	err = os.WriteFile(Path, body, os.ModePerm)
	if err!=nil{
		return err
	}

	return nil
}