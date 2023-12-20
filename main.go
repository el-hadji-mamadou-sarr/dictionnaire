package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type dictionnaire struct {
	filePath string
	entries  map[string]int
}

func (dict *dictionnaire) checkError(){
	if err := dict.saveToFile(); err != nil {
		fmt.Println("Error saving to file:", err)
	}
}

func main() {
	filePath := "dictionary.json"
	dict := newDictionary(filePath)

	dict.add("one", 1)
	dict.add("two", 2)
	fmt.Println("Get 'one':", dict.get("one"))
	dict.list()
	dict.remove("two")
	dict.list()
}

func newDictionary(filePath string) dictionnaire {
	return dictionnaire{
		filePath: filePath,
		entries:  make(map[string]int),
	}
}

func (dict *dictionnaire) add(key string, value int) {
	dict.entries[key] = value
	dict.checkError()
}

func (dict *dictionnaire) get(key string) int {
	return dict.entries[key]
}

func (dict *dictionnaire) remove(key string) {
	delete(dict.entries, key)
	dict.checkError()
}

func (dict *dictionnaire) list() {
	for key, value := range dict.entries {
		fmt.Println(key, value)
	}
}

func (dict *dictionnaire) saveToFile() error {
	data, err := json.MarshalIndent(dict.entries, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dict.filePath, data, 644)



}