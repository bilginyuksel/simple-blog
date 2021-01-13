package user

import (
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
