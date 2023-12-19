package main

import "fmt" 
func main() { 
	fmt.Println("👋 Hello World.")
	m := make(map[string]int)
	m["k1"] = 7
    m["k2"] = 13
	get(m, "k1")
	remove(m, "k2")
	list(m)

}

func get(m map[string]int, key string) int {
	return m[key]
}

func remove(m map[string]int, key string) {
	delete(m, key)
}

func list(m map[string]int) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}