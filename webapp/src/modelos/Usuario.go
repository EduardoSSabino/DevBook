package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

//
type Usuario struct {
	Id          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUsuarioCompleto faz 4 requisições na API para montar o usuário
func BuscarUsuarioCompleto(usuarioId uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioId, r)
	go BuscarSeguidores(canalSeguidores, usuarioId, r)
	go BuscarSeguindo(canalSeguindo, usuarioId, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioId, r)

	var ( // vou jogar as respostas das goroutines dentro dessas variáveis
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	// vou fazr um for que vai iterar 4x
	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.Id == 0 {
				return Usuario{}, errors.New("Erro ao buscar o usuário")
			}

			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("Erro ao buscar os seguidores")
			}

			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				return Usuario{}, errors.New("Erro ao buscar quem o usuário está seguindo")
			}

			seguindo = seguindoCarregados
		case publicacoesCarregados := <-canalPublicacoes:
			if publicacoesCarregados != nil {
				return Usuario{}, errors.New("Erro ao buscar quem o usuário está seguindo")
			}

			publicacoes = publicacoesCarregados

		}
	}
	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil

}

// vou chmar uma função para cada canal

// BuscarDadsoDoUsuario chama a API para buscar os dados do usuario
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{} // mandando usuario em branco pro canal
		return
	}
	defer response.Body.Close()
	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}
	canal <- usuario

}

// BuscarSeguidores chama a API para buscar os seguidores do usuário
func BuscarSeguidores(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil // mandando usuario em branco pro canal
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

// BuscarSeguindo chama a API para buscar os usuarios seguidor por um usuário
func BuscarSeguindo(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil // mandando usuario em branco pro canal
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}

	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguindo
}

// BuscarPublicacoes chama a API para buscar as publicações de um usuário
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil // mandando usuario em branco pro canal
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
