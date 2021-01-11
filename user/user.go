package user

import "fmt"

func init() {
	fmt.Println("Hello world")
}

// User ...
type User struct {
	ID        int64
	Firstname string
	Lastname  string
	Email     string
	Nickname  string
}
