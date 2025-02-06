package encrypter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordEncrypter_Ok(t *testing.T) {
	//given
	password := "Password1234"

	//act
	hash, err := PasswordEncrypter(password)

	//assert
	assert.Nil(t, err)
	assert.IsType(t, []byte{}, hash)
}

func TestPasswordDecrypter_Ok(t *testing.T) {
	//given
	password := "Password1234"

	//act
	hash, _ := PasswordEncrypter(password)

	decrypter := PasswordDecrypter(hash, password)

	//assert
	assert.True(t, decrypter)
}

func TestPasswordDecrypter_Err(t *testing.T) {
	//given
	password := "Password1234"
	passwordErr := "Password123"
	//act
	hash, _ := PasswordEncrypter(password)

	decrypter := PasswordDecrypter(hash, passwordErr)

	//assert
	assert.False(t, decrypter)
}
