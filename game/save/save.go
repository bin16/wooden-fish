package save

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

var DIR = "."

var Read = osRead
var Write = osWrite

func init() {
	if d, err := os.Getwd(); err != nil {
		log.Println(err)
		DIR = ""
	} else {
		DIR = d
	}
}

// func Read(name string, data any) error {
// 	return osRead(name, data)
// }

// func Write(name string, data any) error {
// 	return osWrite(name, data)
// }

func osRead(name string, data any) error {
	var filename = path.Join(DIR, name)
	raw, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, data)
}

func osWrite(name string, data any) error {
	var filename = path.Join(DIR, name)

	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, raw, 0666)
}
