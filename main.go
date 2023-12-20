package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Entry struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type dictionnaire struct {
	filePath string
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
func (dict *dictionnaire) saveToFile(entries []Entry) error {
	jsonData, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dict.filePath, jsonData, 644)
	if err != nil {
		return err
	}

	return nil	
}

func (dict *dictionnaire) loadFromFile() ([]Entry, error) {
	jsonData, err := ioutil.ReadFile(dict.filePath)
	if err != nil {
		return nil, err
	}

	// Check if the JSON data is empty
	if len(jsonData) == 0 {
		return nil, nil
	}

	var entries []Entry
	err = json.Unmarshal(jsonData, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}


func (dict *dictionnaire) saving(entries []Entry) {
	if err := dict.saveToFile(entries); err != nil {
		fmt.Println("Error saving to file:", err)
	}
}

func newDictionary(filePath string) dictionnaire {
	return dictionnaire{
		filePath: filePath,
	}
}

func (dict *dictionnaire) add(key string, value int) {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return
	}

	entry := Entry{Key: key, Value: value}
	entries = append(entries, entry)
	dict.saving(entries)
}

func (dict *dictionnaire) get(key string) int {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return 0
	}

	for _, entry := range entries {
		if entry.Key == key {
			fmt.Println("Found:", entry)
			return entry.Value
		}
	}

	fmt.Println("Key not found:", key)
	return 0
}

func (dict *dictionnaire) remove(key string) {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return
	}

	for i, entry := range entries {
		if entry.Key == key {
			fmt.Println("Removing:", entry)
			// Remove the entry from the slice
			entries = append(entries[:i], entries[i+1:]...)
			dict.saving(entries)
			return
		}
	}

	fmt.Println("Key not found:", key)
}

func (dict *dictionnaire) list() {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}
}
