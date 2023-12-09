package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashPassword(password string, salt []byte) string {
	pwBytes := []byte(password)
	pwBytes = append(pwBytes, salt...)

	var shaHasher = sha512.New()

	shaHasher.Write(pwBytes)

	var hashedPasswordBytes = shaHasher.Sum(nil)

	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func ComparePasswords(hashedPassword, currentPassword string, salt []byte) bool {
	var currentPasswordHashed = HashPassword(currentPassword, salt)

	return hashedPassword == currentPasswordHashed
}
