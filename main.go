package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Learn Go")
	todos.add("Get call from JPMC")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v", todos)
}