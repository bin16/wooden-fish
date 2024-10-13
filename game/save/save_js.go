package save

import (
	"encoding/json"
	"syscall/js"
)

func init() {
	Read = jsRead
	Write = jsWrite
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
