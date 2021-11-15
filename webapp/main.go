package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"

	"github.com/gorilla/securecookie"
)

// go mod init webapp
// go get github.com/gorilla/mux
// go get github.com/joho/godotenv variáveis de ambiente
// go get github.com/gorilla/securecookie

// Essa função será executada apenas uma vez com a finalidade de imprimir o valor da duas variáveis e slavar no arquivo .env
func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16)) // o tamanho da chave que eu quero
	fmt.Println(hashKey)
	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16)) // o tamanho da chave que eu quero
	fmt.Println(blockKey)
}

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()

	r := router.Gerar() // irá me retornar a rota configurada
	fmt.Println(config.Porta)

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
	fmt.Println("Testando")
}
