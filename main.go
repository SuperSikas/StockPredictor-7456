```go
// Package main provides a program to demonstrate basic data processing
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// User struct to hold user data
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Age       int
}

func main() {
	// Create slice to hold users
	users := make([]User, 0)

	// Open the data file
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Loop over all lines in the file and parse them
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line on comma
		parts := strings.Split(line, ",")

		// Parse the ID and Age as integers
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		age, err := strconv.Atoi(parts[4])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Create a new User object and append it to the slice
		users = append(users, User{
			ID:        id,
			FirstName: parts[1],
			LastName:  parts[2],
			Email:     parts[3],
			Age:       age,
		})
	}

	// Print all users
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	// Example of data processing: find average age
	var totalAge int
	for _, user := range users {
		totalAge += user.Age
	}
	avgAge := float64(totalAge) / float64(len(users))
	fmt.Printf("Average age: %.2f\n", avgAge)
}

// Note: The code assumes that the data file "data.txt" exists in the 
// same directory with each line in the format:
// ID,FirstName,LastName,Email,Age
```

Цей код виконує обробку даних, зчитуючи інформацію про користувачів з файлу, розділяючи її на окремі частини і зберігаючи її в структурі "User". Після цього він знаходить середній вік користувачів. 

Примітка: цей код припускає, що файл даних "data.txt" існує в тій же директорії і кожен рядок має формат: ID,FirstName,LastName,Email,Age