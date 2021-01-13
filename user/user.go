package user

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var emails map[string]bool
var usernames map[string]bool

func init() {
	emails = make(map[string]bool)
	usernames = make(map[string]bool)

	emails["noreply@gmail.com"] = true
	usernames["bilginyuksel"] = true
}

// User ...
type User struct {
	id        int64
	firstname string
	lastname  string
	email     string
	password  string
	Username  string
}

const alreadyExists = "Username or email already exists"

// Register User registration function.
func Register(username string, email string, password string) (*User, error) {
	// Check if any user signed with the same email or username
	if IsEmailExists(email) || IsUsernameExists(username) {
		return nil, errors.New(alreadyExists)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return &User{
		id:        1,
		firstname: "",
		lastname:  "",
		email:     email,
		Username:  username,
		password:  string(hash),
	}, nil
}

// IsEmailExists Check if any user registered with the same email.
func IsEmailExists(email string) bool {
	_, ok := emails[email]
	return ok
}

// IsUsernameExists Check if any user registered with the same username
func IsUsernameExists(username string) bool {
	_, ok := usernames[username]
	return ok
}

// Password Get user decoded password
func (u *User) Password() string {
	return u.password
}
