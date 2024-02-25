package main

import (
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

// CriptografaSenha criptografa a senha usando bcrypt
func CriptografaSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(bytes), err
}

// VerificaSenha compara a senha fornecida com a hash armazenada
func VerificaSenha(senha string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	return err == nil
}
