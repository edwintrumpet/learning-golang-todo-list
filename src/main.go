package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

type taskList struct {
	tasks []*task
}

func (l *taskList) addToList(t *task) {
	l.tasks = append(l.tasks, t)
}

func (l *taskList) removeFromList(index int) {
	l.tasks = append(l.tasks[:index], l.tasks[index+1:]...)
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
	list := taskList{}

	list.addToList(&task{
		name:        "Hacer el curso de algoritmos y pensamiento lógico",
		description: "Este curso lo tomé recientemente",
	})

	list.addToList(&task{
		name:        "Hacer el curso gratis de programación básica",
		description: "Este curso lo tomé hace tiempo",
	})

	list.addToList(&task{
		name:        "Hacer el curso de introducción al desarrollo backend",
		description: "Este curso lo tomé hace tiempo",
	})

	list.addToList(&task{
		name:        "Hacer el curso de introducción a la terminal y línea de comandos",
		description: "Este curso lo tomé hace poco, pero es interesante",
	})

	fmt.Println("amount of tasks", len(list.tasks))

	for _, currentTask := range list.tasks {
		fmt.Printf("%+v\n", *currentTask)
	}

	list.removeFromList(1)

	fmt.Println("amount of tasks", len(list.tasks))

	for _, currentTask := range list.tasks {
		fmt.Printf("%+v\n", *currentTask)
	}
}
