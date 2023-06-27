package main

import "go-horse-race/structs"

func main() {
	m := structs.NewMatch("1", 10, 100)
	m.Start()
}
