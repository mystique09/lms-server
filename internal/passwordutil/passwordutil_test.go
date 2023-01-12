package passwordutil

import "testing"

func TestPasswordEncryption(t *testing.T) {
	var plain_password string = "testpassword"
	hashed_password, err := Encrypt(plain_password)

	if err != nil {
		t.Errorf("Password %v not encrypted, got %v", hashed_password, plain_password)
	}

}
