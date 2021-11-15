package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost) /* meu segundo parametro é o que chama de custo da operação, é o
	numero que você passa que conforme aumenta mais fica dificil de quebrar o hash que você gerou. bctypt.DefaultCost traz meus custo padrão */
}

// VerificaSenha compara uma senha com o hash e retorna se elas são iguais
func VerificarSenha(senhaComHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
