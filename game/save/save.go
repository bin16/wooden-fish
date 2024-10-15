package save

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"strings"
)

var (
	wd = "."

	ReadSave  = osReadJSON
	WriteSave = osWriteJSON

	ReadJSON = osReadJSON
	ReadFile = osReadFile

	List = osList
	Find = osFind
)

func init() {
	if d, err := os.Getwd(); err != nil {
		log.Println(err)
		wd = ""
	} else {
		wd = d
	}
}

// List entries of dir,
// {name} relative to {wd}
func osList(name string) ([]string, error) {
	var names = []string{}

	var pathname = path.Join(wd, name)
	entries, err := os.ReadDir(pathname)
	if err != nil {
		return names, err
	}

	for _, d := range entries {
		names = append(names, d.Name())
	}

	return names, err
}

// Read file and return bytes,
// {name} relative to {wd}
func osReadFile(name string) ([]byte, error) {
	var filename = path.Join(wd, name)
	return os.ReadFile(filename)
}

// Read data as JSON,
// {name} relative to {wd}
func osReadJSON(name string, data any) error {
	var filename = path.Join(wd, name)
	raw, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, data)
}

// Find entries has {suffix}, return list of string
// {name} relative to {wd}
func osFind(dir, suffix string) ([]string, error) {
	var (
		results = []string{}
		p       = path.Join(wd, dir)
	)

	entries, err := os.ReadDir(p)
	if err != nil {
		return results, err
	}

	for _, d := range entries {
		if d.IsDir() {
			continue
		}

		if strings.HasSuffix(d.Name(), suffix) {
			results = append(results, d.Name())
		}
	}

	return results, nil
}

// Write JSON data to file,
// {name} relative to {wd}
func osWriteJSON(name string, data any) error {
	var filename = path.Join(wd, name)

	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, raw, 0666)
}
