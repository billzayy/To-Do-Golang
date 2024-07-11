package ui

import (
	"bufio"
	"fmt"
	"os"

	"github.com/billzayy/to-do-list/pkg"
)

func Content(option int, list *pkg.Todo) {
	switch option {
	case 1:
		AddTodo(list)
	case 2:
		DeleteTodo(list)
	case 3:
		UpdateItem(list)
	case 4:
		DoneTodo(list)
	default:
		break
	}
}

func AddTodo(todo *pkg.Todo) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Print("Add tasks : ")
	reader.Scan()

	item := reader.Text()

	todo.Add(item)
}

func DeleteTodo(todo *pkg.Todo) {
	var index int

	fmt.Print("Choose tasks to delete : ")
	fmt.Scan(&index)

	todo.Delete(index)
}

func UpdateItem(todo *pkg.Todo) {
	reader := bufio.NewScanner(os.Stdin)

	var index int

	fmt.Print("Choose item : ")
	reader.Scan()

	fmt.Print("Type your update : ")

	tasks := reader.Text()

	todo.UpdateTask(index, tasks)
}

func DoneTodo(todo *pkg.Todo) {
	var index int

	fmt.Print("Choose tasks done : ")
	fmt.Scan(&index)

	todo.Done(index)
}
