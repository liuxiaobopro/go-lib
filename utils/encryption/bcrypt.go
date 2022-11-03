package encryption

import "golang.org/x/crypto/bcrypt"

// BcryptEncrypt Bcrypt加密
func BcryptEncrypt(str string) (string, error) {
	pwd := []byte(str)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// BcryptCheck Bcrypt校验
func BcryptCheck(hashStr string, readyStr string) bool {
	byteHash := []byte(hashStr)
	byteReady := []byte(readyStr)
	err := bcrypt.CompareHashAndPassword(byteHash, byteReady)
	return err == nil
}
