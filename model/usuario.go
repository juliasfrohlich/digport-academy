package model

import "golang.org/x/crypto/bcrypt"

type Usuario struct {
	Id    string
	Nome  string
	Email string
	senha string
}

// GeraHashDaSenha aplica hash à senha e a armazena no struct do usuário.
func (u *Usuario) GeraHashDaSenha(senha string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.senha = string(hash)
	return nil
}

// ValidaSenha compara a senha fornecida com o hash armazenado.
func (u *Usuario) ValidaSenha(senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.senha), []byte(senha))
	return err == nil
}
