package main

import "fmt"

const enPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return enPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
