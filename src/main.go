package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) complete() {
	t.completed = true
}

func (t *task) setDescription(description string) {
	t.description = description
}

func (t *task) setName(name string) {
	t.name = name
}

func main() {
	t := task{
		name:        "complete my go course",
		description: "I must to complete it tomorrow",
	}

	fmt.Printf("task: %+v\n", t)

	t.setName("completar mi curso de Go")
	t.setDescription("debo completarlo mañana")
	t.complete()

	fmt.Printf("task: %+v\n", t)
}
