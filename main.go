package main

import "fmt"

func SayHi(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(SayHi(("Diego")))
}
