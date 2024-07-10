package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Add tasks : ")

	item, _ := reader.ReadString('\n')

	todo.Add(strings.Split(item, "\n")[0])
}

func DeleteTodo(todo *pkg.Todo) {
	var index int

	fmt.Print("Choose tasks to delete : ")
	fmt.Scan(&index)

	todo.Delete(index)
}

func UpdateItem(todo *pkg.Todo) {
	reader := bufio.NewReader(os.Stdin)

	var index int

	fmt.Print("Choose item : ")
	fmt.Scan(&index)

	fmt.Print("Type your update : ")

	tasks, _ := reader.ReadString('\n')

	todo.UpdateTask(index, strings.Split(tasks, "\n")[0])
}

func DoneTodo(todo *pkg.Todo) {
	var index int

	fmt.Print("Choose tasks done : ")
	fmt.Scan(&index)

	todo.Done(index)
}
