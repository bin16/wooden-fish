package save

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"syscall/js"
)

var DIR = "."

func init() {
	if d, err := os.Getwd(); err != nil {
		log.Println(err)
		DIR = ""
	} else {
		DIR = d
	}
}

func Read(name string, data any) error {
	if DIR == "" {
		return jsRead(name, data)
	}

	return osRead(name, data)
}

func Write(name string, data any) error {
	if DIR == "" {
		return jsWrite(name, data)
	}

	return osWrite(name, data)
}

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

func jsWrite(name string, data any) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js.Global().Get("localStorage").Call("setItem", name, string(raw))
	return nil
}

func jsRead(name string, data any) error {
	var val = js.Global().Get("localStorage").Call("getItem", name)
	var str = val.String()

	return json.Unmarshal([]byte(str), data)
}
