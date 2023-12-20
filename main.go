package main

import (
	"dictionnaire/dictionnary"
	"fmt"
)


func  run_test(dict dictionnary.Dictionnary){
	dict.Add("one", "its a number")
	dict.Add("two", "its the result of 1+1")
	fmt.Println("Get 'one':", dict.Get("one"))
	dict.List()
	dict.Remove("two")
	dict.List()
}

func main() {
	filePath := "dictionary.json"
	dict := dictionnary.NewDictionnary(filePath)
	run_test(dict)
	
}




