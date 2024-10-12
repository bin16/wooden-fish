package save

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

var DIR = "."

func init() {
	var d, err = os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	DIR = d

	fmt.Println(DIR)
}

func Read(name string, data any) error {
	var filename = path.Join(DIR, name)
	raw, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, data)
}

func Write(name string, data any) error {
	var filename = path.Join(DIR, name)

	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, raw, 0666)
}
