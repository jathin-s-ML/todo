package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/jathin-s-ML/todo/internal" // Updated import path
)

func main() {
	fmt.Println("TODO Application")
	reader := bufio.NewReader(os.Stdin)
	todoManager := internal.NewTodoList()

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List All Tasks")
		fmt.Println("3. Mark a Task as Completed")
		fmt.Println("4. Delete a Task")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter the task you want to add: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			todoManager.Add(title)

		case 2:
			todoManager.List()

		case 3:
			fmt.Print("Enter the Task ID to mark as complete: ")
			var id int
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid task ID.")
				continue
			}
			if err := todoManager.MarkAsCompleted(id); err != nil {
				fmt.Println(err)
			}

		case 4:
			fmt.Print("Enter the Task ID to delete: ")
			var id int
			_, err := fmt.Scan(&id)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid task ID.")
				continue
			}
			if err := todoManager.DeleteTask(id); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task deleted successfully")
			}

		case 5:
			fmt.Println("Exiting TODO application")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option")
		}
	}
}
