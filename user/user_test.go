package user

import (
	"fmt"
	"testing"
)

var rq = NewRegisterRequest("username", "email", "password", "password")

func TestRegister_UserRegisteration(t *testing.T) {
	user, err := rq.Register()
	if err != nil || user == nil {
		t.Fail()
	}

}

func TestIsEmailExists_CheckEmail(t *testing.T) {
	shouldTrue := isEmailExists("noreply@gmail.com")
	shouldFalse := isEmailExists("demo@gmail.com")
	if shouldTrue == false || shouldFalse == true {
		t.Fail()
	}

}

func TestIsUsernameExists_CheckUsername(t *testing.T) {
	shouldTrue := isUsernameExists("bilginyuksel")
	shouldFalse := isUsernameExists("lemoan")
	if shouldTrue == false || shouldFalse == true {
		t.Fail()
	}

}

func TestRegister_PasswordShouldBeDecryped(t *testing.T) {
	user, _ := rq.Register()
	pass := user.password
	if pass == "password" {
		t.Fail()
	}
}

func TestCreateJWT_GetToken(t *testing.T) {
	user, _ := rq.Register()
	token := user.CreateJWT()
	fmt.Println(token)
	if _, ok := validateJWT(token); !ok {
		t.Fail()
	}
}


func TestAuthenticateWithJWT_GetUser(t *testing.T) {
    expectedUser, _ := rq.Register()
    token := expectedUser.CreateJWT()
    fmt.Printf("JWT is created the new JWT is %s\n", token)
    givenUser, err := AuthenticateWithJWT(token)
    if err != nil || givenUser == nil {
        t.FailNow()
    }

    fmt.Println(givenUser)
    isUsernameSame := expectedUser.Username == givenUser.Username
    isEmailSame := expectedUser.email == givenUser.email
    isPasswordSame := expectedUser.password == givenUser.password
    isCreatedTimeSame := expectedUser.createdTime.Equal(givenUser.createdTime)

    isEqual := isUsernameSame && isEmailSame && isPasswordSame && isCreatedTimeSame
    if !isEqual {
        t.Fail()
    }
}


