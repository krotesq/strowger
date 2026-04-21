package password

import "golang.org/x/crypto/bcrypt"

// get hash from password
func Hash(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hash), err
}

// verify password hash
func Compare(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}