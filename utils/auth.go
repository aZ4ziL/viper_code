package utils

import "golang.org/x/crypto/bcrypt"

// EncryptionPassword password encryption
func EncryptionPassword(pw string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashed)
}

// DecryptionPassword password decryption
func DecryptionPassword(pw, hashed string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw)); err != nil {
		return false
	} else {
		return true
	}
}
