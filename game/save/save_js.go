package save

import (
	"encoding/json"
	"io"
	"net/http"
	"path"
	"regexp"
	"syscall/js"
)

func init() {
	ReadSave = localStorageReadJSON
	WriteSave = localStorageWriteJSON

	ReadJSON = httpReadJSON
	ReadFile = httpReadBytes

	List = httpList
	Find = httpFind
}

func localStorageWriteJSON(name string, data any) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js.Global().Get("localStorage").Call("setItem", name, string(raw))
	return nil
}

func localStorageReadJSON(name string, data any) error {
	var val = js.Global().Get("localStorage").Call("getItem", name)
	var str = val.String()

	return json.Unmarshal([]byte(str), data)
}

func httpReadJSON(name string, data any) error {
	var u = path.Join("/", name)

	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

func httpList(dir string) ([]string, error) {
	var names = []string{}

	resp, err := http.Get(path.Join("/", dir))
	if err != nil {
		return names, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return names, err
	}

	var rx = regexp.MustCompile(`href="([^\"]+)"`)
	var results = rx.FindAll(raw, -1)
	for _, s := range results {
		names = append(names, string(s[6:len(s)-1]))
	}

	return names, nil
}

func httpFind(dir, suffix string) ([]string, error) {
	var names = []string{}

	resp, err := http.Get(path.Join("/", dir))
	if err != nil {
		return names, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return names, err
	}

	var rx = regexp.MustCompile(`"\w+` + suffix + `"`)
	var results = rx.FindAll(raw, -1)
	for _, s := range results {
		s = s[1 : len(s)-1]
		names = append(names, string(s))
	}

	// fmt.Println(string(raw), names)

	return names, nil
}

func httpReadBytes(name string) ([]byte, error) {
	var u = path.Join("/", name)

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
