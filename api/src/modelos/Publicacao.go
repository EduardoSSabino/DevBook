package modelos

import (
	"errors"
	"strings"
)

// Publicação representa uma publicação feita por um usuário
type Publicacao struct {
	Id        uint64 `json:"id,omitempty"`
	Titulo    string `json:"titulo,omitempty"`
	Conteudo  string `json:"conteudo,omitempty"`
	AutorId   uint64 `json:"autorId,omitempty"`
	AutorNick string `json:"AutorNick,omitempty"`
	Curtidas  uint64 `json:"curtidas,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar a publicação recebida
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

// validar : vai verificar se o titulo e o conteudo tem alguma coisa dentro
func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O titulo é obrigatório e não pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("O conteudo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (publicacao *Publicacao) formatar() error {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)     // tira os espaçoes em branco das extremidades
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo) // tira os espaços em branco das extremidades

	return nil
}
