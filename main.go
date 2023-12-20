package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entry struct {
	Name  string `json:"name"`
	Definition string    `json:"definition"`
}

type dictionnaire struct {
	filePath string
}

func (dict *dictionnaire) run_test(){
	// dict.add("one", "its a number")
	// dict.add("two", "its the result of 1+1")
	fmt.Println("Get 'one':", dict.get("one"))
	// dict.list()
	// dict.remove("two")
	dict.list()
}

func main() {
	filePath := "dictionary.json"
	dict := newDictionary(filePath)
	dict.run_test()
	
}


func (dict *dictionnaire) saveToFile(entries []Entry) error {
	jsonData, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(dict.filePath, jsonData, 644)
	if err != nil {
		return err
	}

	return nil	
}

func (dict *dictionnaire) loadFromFile() ([]Entry, error) {
	jsonData, err := os.ReadFile(dict.filePath)
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

func (dict *dictionnaire) add(name string, definition string) {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return
	}

	// check if the name already exists in the dictionnary
	for i, entry := range entries {
		if entry.Name == name {
			fmt.Printf("Updating existing entry '%s': %s to %s\n", name, entry.Definition, definition)
			entries[i].Definition = definition
			dict.saving(entries)
			return
		}
	}

	entry := Entry{Name: name, Definition: definition}
	entries = append(entries, entry)
	dict.saving(entries)
}

func (dict *dictionnaire) get(name string) string {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return ""
	}

	for _, entry := range entries {
		fmt.Println("Found:", entry)
		if entry.Name == name {
			return entry.Definition
		}
	}

	fmt.Println("Key not found:", name)
	return ""
}

func (dict *dictionnaire) remove(name string) {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return
	}

	for i, entry := range entries {
		if entry.Name == name {
			fmt.Println("Removing:", entry)
			// Remove the entry from the slice
			entries = append(entries[:i], entries[i+1:]...)
			dict.saving(entries)
			return
		}
	}

	fmt.Println("Key not found:", name)
}

func (dict *dictionnaire) list() {
	entries, err := dict.loadFromFile()
	if err != nil {
		fmt.Println("Error loading from file:", err)
		return
	}

	for _, entry := range entries {
		fmt.Println( entry.Name, ":", entry.Definition)
	}
}
