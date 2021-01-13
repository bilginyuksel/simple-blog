package user

import (
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var emails map[string]string
var usernames map[string]string

func init() {
	emails = make(map[string]string)
	usernames = make(map[string]string)

	emails["noreply@gmail.com"] = "string"
	usernames["bilginyuksel"] = "string"
}

// User ...
type User struct {
	id          int64
	firstname   string
	lastname    string
	email       string
	password    string
	createdTime time.Time
	updatedTime time.Time
	active      bool
	confirmed   bool

	Username string
}

// RegisterRequest Is the struct object when new user is registering the system.
type RegisterRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	// Birthday  time.Time
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

const alreadyExists = "Username or email already exists"
const mandatoryInformationsMissing = "Mandatory informations missing in the RegisterRequest object."

// NewRegisterRequest Create user registeration request with mandatory parameters.
func NewRegisterRequest(username string, email string, password string, password2 string) *RegisterRequest {

	return &RegisterRequest{
		Username:  username,
		Email:     email,
		Password:  password,
		Password2: password2,
	}
}

// Register User registration function.
func (rq *RegisterRequest) Register() (*User, error) {

	if len(rq.Username) == 0 || len(rq.Email) == 0 || len(rq.Password) == 0 || len(rq.Password2) == 0 {
		// rq doesn't created with NewRegisterRequest method so reject registeration.
		return nil, errors.New(mandatoryInformationsMissing)
	}

	// Check if any user signed with the same email or username
	if isEmailExists(rq.Email) || isUsernameExists(rq.Username) {
		return nil, errors.New(alreadyExists)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(rq.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return &User{
		id:        1,
		firstname: rq.Firstname,
		lastname:  rq.Lastname,
		email:     rq.Email,
		Username:  rq.Username,
		password:  string(hash),
	}, nil
}

// AuthenticateWithUsername Basic authentication
func AuthenticateWithUsername(username string, password string) (*User, error) {
	if !isUsernameExists(username) {
		return nil, errors.New("username not exists")
	}
	hashedPassword := usernames[username]
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return nil, nil
}

// AuthenticateWithEmail Basic AuthenticateWithEmail
func AuthenticateWithEmail(email string, password string) (*User, error) {
	if !isEmailExists(email) {
		return nil, errors.New("email not exists")
	}
	hashedPassword := emails[email]
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}
	return nil, nil
}

// AuthenticateWithJWT ...
// func AuthenticateWithJWT(token string) (*User, error) {
// 	token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the algorithm what you expect.
// 		if _, ok := token.Method(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
// 		}
// 		return hmacSampleSecret, nil
// 	})
// 	return nil, nil
// }

// IsEmailExists Check if any user registered with the same email.
func isEmailExists(email string) bool {
	_, ok := emails[email]
	return ok
}

// IsUsernameExists Check if any user registered with the same username
func isUsernameExists(username string) bool {
	_, ok := usernames[username]
	return ok
}
