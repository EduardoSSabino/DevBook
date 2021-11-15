package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// OBS: chamando o método respostas.Erro() pra quando tudo der erro e chamando respostas.JSON() pra quando tudo der certo!

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) { // A Request ou requisição traduzindo diretamente para português, é o pedido que um cliente realiza a nosso servidor.
	// agora irei ler o request.body e jogar isso dentro de um struct
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro) // StatusUnprocessableEntity: entidade não processável
		return
	}

	var usuario modelos.Usuario // criando um usuario que está dentro do nosso pacote de modelos
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro) //StatusBadRequest: é quando a requisição não está atendendo o que esperamos dela. tratando o erro
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro) //StatusBadRequest: é quando a requisição não está atendendo o que esperamos dela. tratando o erro
		return
	}

	// agora irei abrir a conexão com o banco de dados
	db, erro := banco.Conectar() // abri a conexão com o banco
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro) // StatusInternalServerError: erro interno do servidor. É um erro do servidor, não tem nada haver com a requisição
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db) // criei um repositorio e passei o banco pra dentro desse repositório
	usuario.Id, erro = repositorio.Criar(usuario)           // chamando o método criar, passando o usuario que está vindo na requisição
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro) // StatusInternalServerError: erro interno do servidor. É um erro do servidor, não tem nada haver com a requisição
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario) // chamadno repostas.JSON pra quando tudo der certo!
}

// BuscarUsuarios busca varios usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// vamos buscar um usuario filtrando por nome ou por nick
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario")) /* strings.ToLower() padroniza tudo pra letra minúscula.
	r.URL.Query() vai trazer tudo que estiver na Query (consulta, o termos que vem ao fim da URL).
	.Get -> significa que eu estou querendo pegar espicificamente aquele parametro */

	// vou abrir a conexão  com o banco
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro) // StatusInternalServerError : erro interno no servidor
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)

	usuarios, erro := repositorio.Buscar(nomeOuNick) // um método que vai buscar meu usuario
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro) // StatusInternalServerError : erro interno no servidor
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

}

// BuscarUsuario busca apenas um usuario salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// pra começar, eu tenho que ler dois valores, tenho que ler tanto o corpo da requisição quanto o valor do Id que ta vindo la do parametro
	parametros := mux.Vars(r) // pega todos os parametros da nossa rota

	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64) // base 10, 64bits
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//passando desse ponto, irei abrir o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuario, erro := repositorio.BuscarPorId(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)

}

// AtualizarUsuario altera infromações de um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// pra começar, eu tenho que ler dois valores, tenho que ler tanto o corpo da requisição quanto o valor do Id que ta vindo la do parametro
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIdNoToken, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioId != usuarioIdNoToken { // se meu usuarioId for diferente do meu usuarioInNoToken, eu não posso fazer essa operação
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel atualizar um usuario que não é você")) // Forbidden não tem muito haver com autenticação.
		return
	}
	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario

	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// a partir desse ponto vou abrir a conexão com o banco de dados e crir meu repositorio
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	if erro = repositorio.Atualizar(usuarioId, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil) // qunado a gente atualiza um usuario, costumamos nao devolver nada, por isso usei o nil

}

// DeletarUsuario exclui as informações de um usuario no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIdNoToken, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro) // se der erro, nao sera autorizado
		return
	}

	if usuarioId != usuarioIdNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel deletar um usuario que não seja o seu")) // Forbidden não tem muito haver com autenticação.
		return
	}

	// a partir desse ponto vou abrir a conexão com o banco de dados e crir meu repositorio
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	if erro = repositorio.Deletar(usuarioId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil) // qunado a gente atualiza um usuario, costumamos nao devolver nada, por isso usei o nil

}

// SeguiUsuario permite que um usuario siga outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r) // lendo o usuario Id que está nos parametros

	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64) // na base 10, 64bits
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if seguidorId == usuarioId { // garantindo que o usuario não irá seguir a si mesmo
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel seguir você mesmo"))
		return
	}

	// abrinco conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)

	if erro = repositorio.Seguir(usuarioId, seguidorId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario permite um usuario deixar de seguir outro
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// lendo o usuarioId que é o que está na requisição, parametro da requisição
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorId == usuarioId {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel parar de seguir você mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	if erro = repositorio.PararDeSeguir(usuarioId, seguidorId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores traz todos seguidores de um usuário
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64) // na base 10, 64bits
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, seguidores)
}

// BuscarSeguindo traz tdoos os usuarios que um determinado usuario está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64) // na base 10, 64bits
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarios, erro := repositorio.BuscarSeguidores(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

// AtualizarSenha irá atualizar senha de um usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIdNoToken, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioIdNoToken != usuarioId {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não seja o seu"))
		return
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body) // body = corpo

	var senha modelos.Senha
	if erro = json.Unmarshal(corpoDaRequisicao, &senha); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioId)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("Senha atual invalida"))
		return
	}

	// antes de suir a nova conta pro banco de dados, eu tenho que transforma-la em hash
	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarSenha(usuarioId, string(senhaComHash)); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// caso não dê erro

	respostas.JSON(w, http.StatusNoContent, nil)
}
