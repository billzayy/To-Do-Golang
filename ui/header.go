package ui

import "fmt"

func Header() int {
	var options int

	fmt.Println("To Do List :")
	fmt.Println("----------------")
	fmt.Println("1. Add Items")
	fmt.Println("2. Delete Items")
	fmt.Println("3. Update Items")
	fmt.Println("4. Done Items")
	fmt.Println("5. Exist !!!")
	fmt.Println("----------------")
	fmt.Print("Choose your options : ")
	fmt.Scan(&options)

	return options
}
