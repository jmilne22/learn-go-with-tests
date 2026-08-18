package main

import "fmt"

const englisHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englisHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
