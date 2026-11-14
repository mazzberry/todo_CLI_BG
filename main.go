package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var userStorage []User

func main() {
	fmt.Println("Hi it's a todo list CLI app")
	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	if *command != "register-user" || *command != "exit" {
		fmt.Println("you must login first")

		for {
			runCommand(*command)
			scn := bufio.NewScanner(os.Stdin)
			fmt.Println("enter another command")
			*command = scn.Text()

		}
	}
}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		loginUser()
	case "exit":
		os.Exit(1)
	default:
		fmt.Println("command is not valid", command)
	}
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, name, email, password string

	fmt.Println("Enter your name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Enter your email: ")
	scanner.Scan()
	email = scanner.Text()

	id = email

	fmt.Println("Enter your password: ")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println(id, name, email, password)

	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)
}

func loginUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("Enter your email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Enter your password: ")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println(email, password)

		// check if user exists
		var userFlag bool
		for _, user := range userStorage {
			if user.Email == email {
				if user.Password == password {
					userFlag = true
					fmt.Println("you're successfully logged in")

				} else {
					fmt.Println("Invalid password")

				}
			}
		}

		if userFlag {
			fmt.Println("the email or password is incorrect")
			return
		}

}

func createTask() {
	scn := bufio.NewScanner(os.Stdin)
	var name, duedate, category string

	fmt.Println("enter the task title")
	scn.Scan()
	name = scn.Text()

	fmt.Println("enter the task category")
	scn.Scan()
	category = scn.Text()

	fmt.Println("enter the task due date")
	scn.Scan()
	duedate = scn.Text()

	fmt.Println("task :", name, duedate, category)

}

func createCategory() {
	scn := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("enter the category title")
	scn.Scan()
	title = scn.Text()

	fmt.Println("enter the category color")
	scn.Scan()
	color = scn.Text()

	fmt.Println("category", title, color)
}
