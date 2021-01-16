package user

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var emails map[string]string
var usernames map[string]string
var mockUsers map[string]*User

func init() {
	emails = make(map[string]string)
	usernames = make(map[string]string)
    mockUsers = make(map[string]*User)

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

	username string
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

const alreadyExists = "username or email already exists"
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
		username:  rq.Username,
		password:  string(hash),
	}, nil
}

// AuthenticateWithusername Basic authentication
func AuthenticateWithusername(username string, password string) (*User, error) {
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

const secretKey = "my-secret-key"

func validateJWT(jwtToken string) (*jwt.Token, bool) {
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the algorithm what you expect.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	return token, token != nil && token.Valid
}

// AuthenticateWithJWT ...
func AuthenticateWithJWT(jwtToken string) (*User, error) {

	token, ok := validateJWT(jwtToken)
	if ok {
		fmt.Println(token)
	}

    user, ok := mockUsers[jwtToken]

    if !ok {
        return nil, errors.New("No users found. System error JWT is correct!")
    }

	return user, nil
}


// Login ...
func (u *User) Login() string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Now().Add(3).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secretKey))

	fmt.Println(tokenString, err)

    // TODO: don't do this... This is just a mock implementation to test 
    // its behavior.
    mockUsers[tokenString] = u

	return tokenString
}

// IsEmailExists Check if any user registered with the same email.
func isEmailExists(email string) bool {
	_, ok := emails[email]
	return ok
}

// IsusernameExists Check if any user registered with the same username
func isUsernameExists(username string) bool {
	_, ok := usernames[username]
	return ok
}
