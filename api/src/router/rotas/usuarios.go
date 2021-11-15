package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{ // Pra iniciar irei fazer as 5 primeiras rotas padrões que teremos no nosso usuario
	{ // rota que vai criar um usuário
		URI:                "/usuarios", // o que vem depoiis da rota é chamado de parametro
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{ // rota que buscar todos os usuarios
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{ // rota que vai buscar um unico usuario pelo o ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{ // rota que vai atualizar os dados do usuário
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{ // rota que vai deletar um usuário
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{ // rota vai seguir um usuario
		URI:                "/usuarios/{usuarioId}/seguir", // nesse caso o {usuarioId} vai ser do usuario que está sendo seguido
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{ // rota vai parar de seguir um usuario
		URI:                "/usuarios/{usuarioId}/parar-de-seguir", // nesse caso o {usuarioId} vai ser do usuario que está sendo seguido
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{ // rota vai buscar os seguidores
		URI:                "/usuarios/{usuarioId}/seguidores", // nesse caso o {usuarioId} vai ser do usuario que está sendo seguido
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{ // rota vai trazer todos os usuarios que um determinado usuario está seguindo
		URI:                "/usuarios/{usuarioId}/seguindo", // nesse caso o {usuarioId} vai ser do usuario que está sendo seguido
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{ // rota vai atualizar senha do usuario
		URI:                "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
