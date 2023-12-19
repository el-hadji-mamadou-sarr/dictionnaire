package main

import "fmt" 
type dictionnaire struct {
    m map[string]int
}

func main() { 
	
	dict := new()
	dict.add("one", 1)
	dict.add("two", 2)
	dict.get("one")
	dict.list()
	dict.remove("two")
}

func new() dictionnaire {
	return dictionnaire{m: make(map[string]int)}
}

func (dict dictionnaire) add( key string, value int) {
	dict.m[key] = value
	
}

func  (dict dictionnaire) get( key string) int {
	return dict.m[key]
}

func (dict dictionnaire) remove( key string) {
	delete(dict.m, key)
}

func  (dict dictionnaire) list() {
	for key, value := range dict.m {
		fmt.Println(key, value)
	}
}