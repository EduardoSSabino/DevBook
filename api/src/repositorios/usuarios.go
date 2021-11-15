package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// um struct que vai receber o banco!

// Usuarios representa um repositório
type Usuarios struct {
	db *sql.DB
}

/* O controllers que vai chamar essa função NovoRepositorioUsuario,
e depois a função irá pegar o banco e jogar dentro do struct usuarios,
 ou seja, vamos criar uma instancia desse struct cm o banco que foi aberto.
  Dentro do struct teremos a informações que faremos a ligação direta com o banco de dados */

// NovoRepositioUsuarios cria um repositório de usuários
func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// criando o primeiro método desse nosso repositório de usuario

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) { // esse metodo esta dentro do meu repositio de usuario. u usuario é o meu repositorio
	fmt.Println("lala")
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)") // statement = demonstração
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha) // Exec executa uma instrução preparada com os argumentos fornecidos e retorna um Resultado resumindo o efeito da instrução.
	if erro != nil {
		return 0, erro
	}

	// passando dessa linha, significa que o usuario ja foi inserido no banco
	ultimoIDInserido, erro := resultado.LastInsertId() // LastInsertId é um uint64
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar trás todos os usuario que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // minha string está assim %nomeOuNick%

	linhas, erro := repositorio.db.Query( // Query é uma consulta que vem no fim da minha URL
		"select id, nome, nick, email from usuarios where nome LIKE ? or nick LIKE ?", nomeOuNick, nomeOuNick) /* Selecionar id, nome, nick, email, criadoEm,
	dos usuarios, onde nome ou o nick é igual/de acordo com o que está vindo em "nomeOuNick". Funciona como  barra de pesquisas do facebook
	onde se digitaar-mos "joão", ira trazer varios usuarios cujo tenha joão em alguma parte do nome */

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() { // iterando com as linhas que chegaram do banco de dados
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
		//	&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario) // adiionando usuarioCriado ao meu slice usuarios

	}

	return usuarios, nil
}

// BuscarPorId busca um usuario por ID
func (repositorio Usuarios) BuscarPorId(Id uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email from usuarios where id = ?", Id)
	if erro != nil {
		return modelos.Usuario{}, erro // estou passadno um usuario vazio
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
		//	&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Atualizar altera as informações de um usuario no banco de dados
func (repositorio Usuarios) Atualizar(Id uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuario set nome = ?, nick = ?, email = ? where id = ?") // iremos passar os campos que iremos alterar

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, Id); erro != nil {
		return erro
	}

	return nil
}

// Deletar excliu as informações de um usuario no banco de dados
func (repositorio Usuarios) Deletar(Id uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(Id); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuario por email e retorna um Id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.Id, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

//
func (repositorio Usuarios) Seguir(usuarioId, seguidorId uint64) error {
	statement, erro := repositorio.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)") // "ignore" não insere um dado na tabela caso ele ja esteja la
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorId); erro != nil { // executando nosso statement
		return erro
	}

	return nil
}

// PararDeSeguir faz com que um usuario pare de seguir outro
func (repositorio Usuarios) PararDeSeguir(usuarioId, seguidorId uint64) error {
	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

// BuscarSeguidores traz todos os seguidores de um usuario
func (repositorio Usuarios) BuscarSeguidores(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email
	from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ? 
	`, usuarioId) /* esse query (consulta) irá fazer, basicaemnte, um join (juntar), entre a tabela
	de usuário e a tabela de seguidores, eu tenho a informação de quem segue quem na tabela de seguidores
	so que as informações desses usuarios estão na tabela de usuarios, então, na tabela de seguidores eu tenho a informação
	o usuario 1 segue o usuario 2, mas eu so sei o Id deles, eu não sei a principio que seria o usuario 1, pra saber quem é
	eu tenho que buscar na tabela de usuarios, então é por isso que eu vou juntas as tabelas. */

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarSeguindo traz todos os usuarios que um determinado usuario está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email
	from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ? 
	`, usuarioId) /* esse query (consulta) irá fazer, basicaemnte, um join (juntar), entre a tabela
	de usuário e a tabela de seguidores, eu tenho a informação de quem segue quem na tabela de seguidores
	so que as informações desses usuarios estão na tabela de usuarios, então, na tabela de seguidores eu tenho a informação
	o usuario 1 segue o usuario 2, mas eu so sei o Id deles, eu não sei a principio que seria o usuario 1, pra saber quem é
	eu tenho que buscar na tabela de usuarios, então é por isso que eu vou juntas as tabelas. */

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario // pra cada iteração vou criar um usuario

		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil // retorno meu slice e um erro sendo nil
}

// BuscarSenha traz a senha de um usuário pelo o Id
func (repositorio Usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linhas, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioId) // diferente do .Prepare, eu tenho que passar um valor depois das aspas
	if erro != nil {
		return "", erro
	}

	defer linhas.Close()

	// criando um usuario pra conseguir dar o Scan pra ele
	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil
}

// AtualizarSenha atualiza a senha de um usuário no banco de dados
func (repositorio Usuarios) AtualizarSenha(usuarioId uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?") // diferente do .Query, eu não tenho que passar um valor depois da aspas
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(senha, usuarioId); erro != nil {
		return erro
	}

	return nil
}
