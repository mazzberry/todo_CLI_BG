package main

import (
	"fmt"
	"bufio"
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
	fmt.Println("Hello, World!")
	registerUser()
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
		ID: len(userStorage) + 1,
		Name: name,
		Email: email,
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

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			fmt.Println("Login successful")
			return
		}
	}
	fmt.Println("Invalid email or password")
}