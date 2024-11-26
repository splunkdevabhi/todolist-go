package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.add("Buy milk")
	todos.add("Buy bread")
	todos.toggle(0)
	todos.print()
	//fmt.Printf("%+v\n\n", todos)
	//todos.delete(0)
	//fmt.Printf("%+v", todos)
}
