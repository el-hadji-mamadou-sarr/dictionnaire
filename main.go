package main

import (
	"dictionnaire/dictionnary"
	"fmt"
)


func  run_test(dict dictionnary.Dictionnary){
	ch_add := make(chan dictionnary.Entry)
	ch_remove := make(chan string)

	go dict.Add("paris", "its a city in france", ch_add)
	fmt.Println(<-ch_add)
	go dict.Remove("paris", ch_remove)
	fmt.Println(<-ch_remove)
}

func main() {
	filePath := "dictionary.json"
	dict := dictionnary.NewDictionnary(filePath)
	run_test(dict)
	
}




