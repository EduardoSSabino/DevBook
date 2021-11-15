package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuario utilizando a rede social
type Usuario struct {
	Id       uint64    `json:"id,omitempty"` // omitempty :significa que, quando fomos passar esse  usuario para um json e o campo id, estiver em branco, ele tira o campo id do json
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

/* Criando alguns métodosdetrod o nosso struct de usuarios que vão ajudar
a gente na questão de validação dos campos. Irei criar três métodos, dois
deles serão privados, só vão ser visíveis dentro desse próprio pacote
modelos, e o outro será publíco pra poder ser usado em outros pacotes */

// Preparar irá chamar meus outros dois métodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error { // eu passo pra funçõ um parametro que vai dizer em qual etapa estou, etapa de cadastro o etapa de edicao
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil { // irmos colocar o hash na senha
		return erro
	}
	return nil
}

// validar : irá averiguar se todos os campos foram preenchidos
func (usuario *Usuario) validar(etapa string) error { // essa função vai verificar se todos os campos do nosso struct de usuario estão preenchidos
	if usuario.Nome == "" { // se o campo nome do meu usuario estiver em branco
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" { // se o campo nick do meu usuario estiver em branco
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" { // se o campo email do meu usuario estiver em branco
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	// validação do email
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" { // se o campo senha do meu usuario estiver em branco. A senha só sera obrigatória na etapa de cadastro
		return errors.New("A senha é obrigatório e não pode estar em branco")
	}
	return nil
}

// formatar : irá tirar os espaços que estão sobranco nas extremidades da string
func (usuario *Usuario) formatar(etapa string) error {
	// irei usar metodos do pacote string
	usuario.Nome = strings.TrimSpace(usuario.Nome)   // retirando os espaços na extremidades do campo nome do meu usuario
	usuario.Nick = strings.TrimSpace(usuario.Nick)   // retirando os espaços na extremidades do campo nick do meu usuario
	usuario.Email = strings.TrimSpace(usuario.Email) // retirando os espaços na extremidades do campo email do meu usuario

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}
	return nil
}
