package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail_Ok(t *testing.T) {
	//given
	email := "johndoe@example.com"

	//act
	isValid := ValidateEmail(email)

	//assert
	assert.True(t, isValid)

}

func TestPasswordValidator_Ok(t *testing.T) {
	//given
	password := "Password1234"

	//act
	isValid := ValidatePassword(password)

	//assert
	assert.True(t, isValid)
}

func TestPasswordValidator_Err(t *testing.T) {
	//given
	password := "..."

	//act
	isValid := ValidatePassword(password)

	//assert
	assert.False(t, isValid)
}

func TestPasswordValidator_LongPwd(t *testing.T) {
	//given
	password := ".................."

	//act
	isValid := ValidatePassword(password)

	//assert
	assert.False(t, isValid)
}
