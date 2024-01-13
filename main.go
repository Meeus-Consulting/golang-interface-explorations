package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
)

type User struct {
	Id   string
	Name string
}

// Define - within the domain - the interaction points that are required

type GetUserById func(id string) User

// Create a factory method that creates a closure around a long-lived dependency
//
//	In this case it's just a message, but it could just as well be a connection pool
//	or an HTTP or AWS client
func makeGetUserAndPrintMessage(message string) GetUserById {
	return func(id string) User {
		var user User
		_ = faker.FakeData(&user)
		fmt.Println(message)
		return user
	}
}

func makeGetUserFromDatabase(db interface{}) GetUserById {
	return func(id string) User {
		// Do some database-y stuff here
		return User{}
	}
}

func main() {
	// This is the composition root, where everything is wired up
	getRandomUser := makeGetUserAndPrintMessage("I expect this message to be printed")

	// And now we invoke our entrypoint, in this case just a simple function call
	_ = getRandomUser("1")
}

// The question is, is this cleaner than full interfaces? I believe it could be
