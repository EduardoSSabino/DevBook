package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o MySQL
	StringDeConexaoBanco = ""
	// Porta que a API está rodando
	Porta = 0
	// SecretKey chave que vai ser usada para assinar o token
	SecretKey []byte
)

// Carregar vai inicializar nossas variáveis de ambiente
func Carregar() { /*essa função não recebe parametros e nem retornar nada,
	isso porquê ela vai mexer com algumas variáveis que irão estar fora
	dela, e essas variávies vão ficar disponiveis pra nossa API inteira*/

	// vou usar o pacote godot env pra ler meu arquivo .env
	var erro error

	if erro = godotenv.Load(); erro != nil { // chamando uma função do pacote godotenv que vai ler uma função do nosso arquivo .env
		log.Fatal(erro) // caso eu tiver algum problema na hora de carregar alguma variavel de ambiente, não tem como nossa API rodar
	}

	// no meu arquivo .env, minha porta é igual 5000, so que o valor está no formato string, precisamos mudá-lo para Int, para ser compartivel com minha variavel Porta tipo Int
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT")) // funciona como um ParseInt
	if erro != nil {
		Porta = 9000
	}

	StringDeConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8",
		os.Getenv("DB_USUARIO"), // usuario
		os.Getenv("DB_SENHA"),   // senha do banco
		os.Getenv("DB_NOME"),    // nome do banco
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
