package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// mysql -h 127.0.0.1 -P 3306 -u root -p
// go get github.com/badoux/checkmail usado para validar um email
// go get golang.org/x/crypto/bcrypt sera usado para ocultar minha senha colcoando hast ********
// go get github.com/dgrijalva/jwt-go sera usado pra habitar o uso do token
// no ubuntu -> mysql> select from usuarios; tras a senha em hash de todos funcionarios

// func init será usado apenas uma vez, depois do uso posso excluuí-lo ou comentá-lo
/* func init() {
	chave := make([]byte, 64)
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64) // será usado para assirnamos nosso token

} */

func main() {
	config.Carregar()
	fmt.Println(config.Porta)
	fmt.Println(config.StringDeConexaoBanco) // só pra ver se está com o valor que esperamos, nossa URL de conexão com o banco de dados

	fmt.Println(config.SecretKey)

	fmt.Printf("Escutando na porta %d\n", config.Porta)

	fmt.Println("Rodando API!")
	r := router.Gerar() // irá a fazer a importação do meu pacote router

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r)) // subindo o servidor na porta 5000. Meu segundo parametro será "r", meu roouter que acabei de criar

}
