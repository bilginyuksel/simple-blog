package user

import (
	"fmt"
	"testing"
)

var rq = NewRegisterRequest("username", "email", "password", "password")

func TestRegister_UserRegisteration(t *testing.T) {
	user, err := rq.Register()
	if err != nil || user == nil {
		t.FailNow()
	}

}

func TestIsEmailExists_CheckEmail(t *testing.T) {
	shouldTrue := isEmailExists("noreply@gmail.com")
	shouldFalse := isEmailExists("demo@gmail.com")
	if shouldTrue == false || shouldFalse == true {
		t.FailNow()
	}

}

func TestIsusernameExists_CheckUsername(t *testing.T) {
	shouldTrue := isUsernameExists("bilginyuksel")
	shouldFalse := isUsernameExists("lemoan")
	if shouldTrue == false || shouldFalse == true {
		t.FailNow()
	}

}

func TestRegister_PasswordShouldBeDecryped(t *testing.T) {
	user, _ := rq.Register()
	pass := user.password
	if pass == "password" {
		t.FailNow()
	}
}

func TestCreateJWT_GetToken(t *testing.T) {
	user, _ := rq.Register()
	token := user.Login()
	fmt.Println(token)
	if _, ok := validateJWT(token); !ok {
		t.FailNow()
	}
}

func TestAuthenticateWithJWT_GetUser(t *testing.T) {
	expectedUser, _ := rq.Register()
	token := expectedUser.Login()
	fmt.Printf("JWT is created the new JWT is %s\n", token)
	givenUser, err := AuthenticateWithJWT(token)
	if err != nil || givenUser == nil {
		t.FailNow()
	}

	fmt.Println(givenUser)
	isUsernameSame := expectedUser.username == givenUser.username
	isEmailSame := expectedUser.email == givenUser.email
	isPasswordSame := expectedUser.password == givenUser.password
	isCreatedTimeSame := expectedUser.createdTime.Equal(givenUser.createdTime)

	isEqual := isUsernameSame && isEmailSame && isPasswordSame && isCreatedTimeSame
	if !isEqual {
		t.FailNow()
	}
}

func TestAuthWithJWT_CheckClaims(t *testing.T) {
	user, _ := rq.Register()
	token := user.Login()
	claims, ok := validateJWT(token)
	if !ok {
		t.FailNow()
	}
	fmt.Println(claims)
	if claims["iss"] != "username" {
		t.FailNow()
	}
}
