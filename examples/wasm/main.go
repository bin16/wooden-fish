package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall/js"
	"time"
)

func main() {
	LocalStorage.SetItem("a", Data{time.Now()})
	fmt.Println("wasm", LocalStorage.GetItem("a"))
	fmt.Println(os.Getenv("GOOS"), os.Getenv("GOARCH"))
}

type localStorage struct{}

func (localStorage) GetItem(k string) (v any) {
	return js.Global().Get("localStorage").Call("getItem", k)
}

func (localStorage) SetItem(k string, v any) {
	data, _ := json.Marshal(v)

	js.Global().Get("localStorage").Call("setItem", k, string(data))
}

var LocalStorage localStorage

type Data struct {
	Time time.Time
}
