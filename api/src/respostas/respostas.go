package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON vai retornar uma resposta em Json para requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json") //
	w.WriteHeader(statusCode)                          // WriteHeader, passa o statusCode na resposta

	/* Minha função para uma resposta generica, simplesmente irá receber o statusCode que vai ser
	passado, vai colocar o statusCode no Header  e dpois pegar os dados que são genéricos e
	transformar pra JSON, e pra isso eu preciso de um ResponseWriter pq é ele que vai dar a resposta */

	if dados != nil {

		if erro := json.NewEncoder(w).Encode(dados); erro != nil { // convertendo para um json
			log.Fatal(erro)
		}
	}

}

// Erro vai me retornar uma resposta em JSON, mas so vai ser acionada quando tiver um  erro, substituindo o log.Fatal()
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct { // chamando  função JSON
		Erro string `json:"erro"`
	}{ // Vou preencher os dados do Erro
		Erro: erro.Error()}) // Método que vai retornar a mesnagem de erro. "Método Erro() dentro do tipo erro"

}
