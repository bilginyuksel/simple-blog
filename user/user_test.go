package user

import (
	"testing"
)

func TestRegister_UserRegisteration(t *testing.T) {
	user, err := Register("username", "email", "password")
	if err != nil || user == nil {
		t.Fail()
	}

}

func TestIsEmailExists_CheckEmail(t *testing.T) {
	shouldTrue := IsEmailExists("noreply@gmail.com")
	shouldFalse := IsEmailExists("demo@gmail.com")
	if shouldTrue == false || shouldFalse == true {
		t.Fail()
	}

}

func TestIsUsernameExists_CheckUsername(t *testing.T) {
	shouldTrue := IsUsernameExists("bilginyuksel")
	shouldFalse := IsUsernameExists("lemoan")
	if shouldTrue == false || shouldFalse == true {
		t.Fail()
	}

}

func TestRegister_PasswordShouldBeDecryped(t *testing.T) {
	user, _ := Register("username", "email", "password")
	if user.Password() == "password" {
		t.Fail()
	}
}
